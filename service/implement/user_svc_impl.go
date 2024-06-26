package implement

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	dto "github.com/jackyuan2022/workspace/api/dto"
	app "github.com/jackyuan2022/workspace/app"
	core "github.com/jackyuan2022/workspace/core"
	model "github.com/jackyuan2022/workspace/domain/model"
	oauth "github.com/jackyuan2022/workspace/domain/oauth"
	repo "github.com/jackyuan2022/workspace/domain/repository"
	repoImpl "github.com/jackyuan2022/workspace/domain/repository/implement"
	svc "github.com/jackyuan2022/workspace/service"
	util "github.com/jackyuan2022/workspace/util"
)

type UserServiceImpl struct {
	userRepo    repo.UserRepository
	sessionRepo repo.OAuthSessionRepository
	oauthMaker  oauth.OAuthMaker
}

func NewUserService() svc.UserService {
	userRepo := repoImpl.NewUserRepository(app.AppContext.APP_DbContext)
	sessionRepo := repoImpl.NewOAuthSessionRepository(app.AppContext.APP_DbContext)
	oauthMaker, err := oauth.NewPasetoMaker(app.AppContext.APP_CONFIG.OAuthConfig)

	if err != nil {
		panic("NewUserService error")
	}

	userSvc := &UserServiceImpl{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		oauthMaker:  oauthMaker,
	}

	return userSvc
}

func (s *UserServiceImpl) Login(ctx *gin.Context, u *dto.LoginRequest) (r *dto.LoginResponse, err *core.AppError) {
	user, err := s.userRepo.FindByMobile(ctx, u.Mobile)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, core.NewValidationError("手机号不存在")
	}

	if user.DenyLogin {
		return nil, core.NewValidationError("用户被禁止登录")
	}

	if isPass := util.BcryptCheck(u.Password, user.Password); !isPass {
		return nil, core.NewValidationError("密码错误")
	}

	accessToken, _, oauthErr := s.oauthMaker.GenerateAccessToken(user.Mobile, user.Name)
	if oauthErr != nil {
		return nil, core.NewUnexpectedError("获取access token失败")
	}
	refreshToken, claims, oauthErr := s.oauthMaker.GenerateRefreshToken(user.Mobile, user.Name)
	if oauthErr != nil {
		return nil, core.NewUnexpectedError("获取refresh token失败")
	}

	session := model.OAuthSession{
		UserId:       user.DbBaseModel.Id,
		Mobile:       user.Mobile,
		UserName:     user.Name,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiredAt:    claims.ExpiredAt.UnixMilli(),
		DbBaseModel:  core.NewDbBaseModel(claims.Id),
	}

	_, err = s.sessionRepo.Insert(ctx, &session)
	if err != nil {
		return nil, core.NewUnexpectedError("Insert Session Data Failure")
	}

	r = &dto.LoginResponse{
		User: dto.UserDTO{
			Mobile:   user.Mobile,
			UserId:   user.DbBaseModel.Id,
			UserName: user.Name,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return r, nil
}

func (s *UserServiceImpl) Register(ctx *gin.Context, u *dto.RegisterRequest) *core.AppError {
	if u == nil {
		return core.NewValidationError("参数错误")
	}

	if strings.Trim(u.User.Mobile, " ") == "" {
		return core.NewValidationError("手机号不能为空")
	}

	if strings.Trim(u.User.Password, " ") == "" {
		return core.NewValidationError("密码不能为空")
	}

	user := model.User{
		Mobile:      u.User.Mobile,
		Password:    u.User.Password,
		Name:        u.User.UserName,
		DbBaseModel: core.NewDbBaseModel(util.GenerateId()),
	}

	_, err := s.userRepo.InsertUser(ctx, &user)
	return err
}

func (s *UserServiceImpl) RefreshToken(ctx *gin.Context, u *dto.RefreshTokenRequest) (r *dto.RefreshTokenResponse, err *core.AppError) {
	claims, oauthErr := s.oauthMaker.VerifyToken(u.RefreshToken)
	if oauthErr != nil {
		return nil, core.NewValidationError("Refresh Token is Invalid")
	}

	session, err := s.sessionRepo.FindById(ctx, claims.Id)
	if err != nil {
		return nil, err
	}

	if session.IsBlocked {
		return nil, core.NewAuthenticationError("Session Blocked")
	}

	if session.UserName != claims.UserName {
		return nil, core.NewAuthenticationError("Incorrect Session User")
	}

	if session.RefreshToken != u.RefreshToken {
		return nil, core.NewAuthenticationError("Incorrect Session Refresh Token")
	}

	expiredAt := time.UnixMilli(session.ExpiredAt)
	if time.Now().After(expiredAt) {
		return nil, core.NewAuthenticationError("Session Expired")
	}

	accessToken, _, oauthErr := s.oauthMaker.GenerateAccessToken(session.Mobile, session.UserName)
	if oauthErr != nil {
		return nil, core.NewUnexpectedError("获取access token失败")
	}
	refreshToken, claims, oauthErr := s.oauthMaker.GenerateRefreshToken(session.Mobile, session.UserName)
	if oauthErr != nil {
		return nil, core.NewUnexpectedError("获取refresh token失败")
	}

	_, err = s.sessionRepo.DeleteById(ctx, session.Id)
	if err != nil {
		return nil, core.NewUnexpectedError("Delete Session Data Failure")
	}

	session = &model.OAuthSession{
		UserId:       session.UserId,
		Mobile:       session.Mobile,
		UserName:     session.UserName,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiredAt:    claims.ExpiredAt.UnixMilli(),
		DbBaseModel:  core.NewDbBaseModel(claims.Id),
	}

	_, err = s.sessionRepo.Insert(ctx, session)
	if err != nil {
		return nil, core.NewUnexpectedError("Insert Session Data Failure")
	}

	r = &dto.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return r, nil
}

func (s *UserServiceImpl) GetUserList(ctx *gin.Context, r *dto.GetUserListRequest) (res *dto.DataListResponse[dto.UserDTO], err *core.AppError) {
	wheres := []core.DbQueryWhere{}
	if len(r.UserId) > 0 {
		filters := []core.DbQueryFilter{core.NewDbQueryFilter("user_id", []interface{}{r.UserId}, "EQ", "string")}
		wheres = append(wheres, core.NewDbQueryWhere(filters, "AND"))
	}
	query := &core.DbQuery{
		QueryWheres: wheres,
		PageSize:    r.Pagination.PageSize,
		PageNumber:  r.Pagination.PageNumber,
	}
	result, err := s.userRepo.QueryData(ctx, query)
	if err != nil {
		return nil, core.NewUnexpectedError("Query User Data Failure")
	}
	var catagoreList []dto.UserDTO
	index := len(result)
	if index > r.Pagination.PageSize {
		index = r.Pagination.PageSize
	}
	for i := 0; i < index; i++ {
		item := result[i]
		dto := dto.UserDTO{
			UserId:    item.DbBaseModel.Id,
			UserName:  item.Name,
			Password:  "",
			DenyLogin: item.DenyLogin,
			Mobile:    item.Mobile,
		}
		catagoreList = append(catagoreList, dto)
	}
	res = &dto.DataListResponse[dto.UserDTO]{
		DataList: catagoreList,
		Pagination: dto.PageDTO{
			PageSize:    r.Pagination.PageSize,
			PageNumber:  r.Pagination.PageNumber,
			HasNextPage: len(result) > r.Pagination.PageSize,
		},
	}
	return res, nil
}

func (s *UserServiceImpl) UpdateUser(ctx *gin.Context, r *dto.DataRequest[dto.UserDTO]) (res *dto.DataResponse[dto.UserDTO], err *core.AppError) {
	if r == nil {
		return nil, core.NewValidationError("参数错误")
	}

	if strings.Trim(r.Data.UserName, " ") == "" {
		return nil, core.NewValidationError("名称不能为空")
	}
	category := model.User{
		Name:        r.Data.UserName,
		Password:    r.Data.Password,
		Mobile:      r.Data.Mobile,
		DenyLogin:   r.Data.DenyLogin,
		DbBaseModel: core.NewDbBaseModel(r.Data.UserId),
	}

	result, err := s.userRepo.UpdateUser(ctx, &category)
	if err != nil {
		return nil, core.NewUnexpectedError("Update User Data Failure")
	}
	res = &dto.DataResponse[dto.UserDTO]{
		Data: dto.UserDTO{
			UserId:    result.Id,
			UserName:  result.Name,
			Password:  "",
			DenyLogin: result.DenyLogin,
		},
	}
	return res, nil
}

func (s *UserServiceImpl) DeleteUser(ctx *gin.Context, r *dto.DataRequest[dto.UserDTO]) (res *dto.DataResponse[dto.UserDTO], err *core.AppError) {
	_, err = s.userRepo.DeleteUserById(ctx, r.Data.UserId)
	if err != nil {
		return nil, core.NewUnexpectedError("Delete User Data Failure")
	}
	res = &dto.DataResponse[dto.UserDTO]{
		Data: r.Data,
	}
	return res, nil
}
