package implement

import (
	"strings"

	"github.com/gin-gonic/gin"

	dto "github.com/jackyuan2022/workspace/api/dto"
	app "github.com/jackyuan2022/workspace/app"
	core "github.com/jackyuan2022/workspace/core"
	"github.com/jackyuan2022/workspace/domain/model"
	repo "github.com/jackyuan2022/workspace/domain/repository"
	repoImpl "github.com/jackyuan2022/workspace/domain/repository/implement"
	svc "github.com/jackyuan2022/workspace/service"
	util "github.com/jackyuan2022/workspace/util"
)

type CategoryServiceImpl struct {
	dataRepo repo.CategoryRepository
}

func NewCategoryService() svc.CategoryService {
	dataRepo := repoImpl.NewCategoryRepository(app.AppContext.APP_DbContext)

	categorySvc := &CategoryServiceImpl{
		dataRepo: dataRepo,
	}

	return categorySvc
}

func (s *CategoryServiceImpl) GetCategoryList(ctx *gin.Context, r *dto.GetCategoryListRequest) (res *dto.GetCategoryListResponse, err *core.AppError) {
	values := []interface{}{r.CategoryType}
	filters := []core.DbQueryFilter{core.NewDbQueryFilter("category_type", values, "EQ", "string")}
	wheres := []core.DbQueryWhere{core.NewDbQueryWhere(filters, "AND")}
	query := &core.DbQuery{
		QueryWheres: wheres,
		PageSize:    r.PageSize,
		PageNumber:  r.PageNumber,
	}
	result, err := s.dataRepo.QueryData(ctx, query)
	if err != nil {
		return nil, core.NewUnexpectedError("Query Category Data Failure")
	}
	var catagoreList []dto.CategoryDTO
	for _, item := range result {
		dto := dto.CategoryDTO{
			Id:           item.DbBaseModel.Id,
			Icon:         item.Icon,
			Name:         item.Name,
			CategoryType: item.CategoryType,
			Order:        item.Order,
		}
		catagoreList = append(catagoreList, dto)
	}
	res = &dto.GetCategoryListResponse{
		CategoryList: catagoreList,
		HasNextPage:  len(result) >= r.PageSize+1,
	}
	return res, nil
}

func (s *CategoryServiceImpl) CreateCategory(ctx *gin.Context, r *dto.CategoryRequest) (res *dto.CategoryResponse, err *core.AppError) {
	if r == nil {
		return nil, core.NewValidationError("参数错误")
	}

	if strings.Trim(r.Category.Name, " ") == "" {
		return nil, core.NewValidationError("名称不能为空")
	}
	category := model.Category{
		Name:         r.Category.Name,
		Icon:         r.Category.Icon,
		Order:        r.Category.Order,
		CategoryType: r.Category.CategoryType,
		DbBaseModel:  core.NewDbBaseModel(util.GenerateId()),
	}

	result, err := s.dataRepo.Insert(ctx, &category)
	if err != nil {
		return nil, core.NewUnexpectedError("Insert Category Data Failure")
	}
	res = &dto.CategoryResponse{
		Category: dto.CategoryDTO{
			Id:           result.Id,
			Name:         result.Name,
			Icon:         result.Icon,
			Order:        result.Order,
			CategoryType: result.CategoryType,
		},
	}
	return res, nil
}

func (s *CategoryServiceImpl) UpdateCategory(ctx *gin.Context, r *dto.CategoryRequest) (res *dto.CategoryResponse, err *core.AppError) {
	if r == nil {
		return nil, core.NewValidationError("参数错误")
	}

	if strings.Trim(r.Category.Name, " ") == "" {
		return nil, core.NewValidationError("名称不能为空")
	}
	category := model.Category{
		Name:         r.Category.Name,
		Icon:         r.Category.Icon,
		Order:        r.Category.Order,
		CategoryType: r.Category.CategoryType,
		DbBaseModel:  core.NewDbBaseModel(r.Category.Id),
	}

	result, err := s.dataRepo.Update(ctx, &category)
	if err != nil {
		return nil, core.NewUnexpectedError("Update Category Data Failure")
	}
	res = &dto.CategoryResponse{
		Category: dto.CategoryDTO{
			Id:    result.Id,
			Name:  result.Name,
			Icon:  result.Icon,
			Order: result.Order,
		},
	}
	return res, nil
}

func (s *CategoryServiceImpl) DeleteCategory(ctx *gin.Context, r *dto.CategoryRequest) (res *dto.CategoryResponse, err *core.AppError) {
	_, err = s.dataRepo.DeleteById(ctx, r.Category.Id)
	if err != nil {
		return nil, core.NewUnexpectedError("Delete Category Data Failure")
	}
	res = &dto.CategoryResponse{
		Category: r.Category,
	}
	return res, nil
}
