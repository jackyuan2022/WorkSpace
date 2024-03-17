package controller

import (
	"github.com/gin-gonic/gin"
	response "github.com/jackyuan2022/workspace/api/core"
	dto "github.com/jackyuan2022/workspace/api/dto"
	service "github.com/jackyuan2022/workspace/service"
	serviceImpl "github.com/jackyuan2022/workspace/service/implement"
	util "github.com/jackyuan2022/workspace/util"
	"github.com/mojocn/base64Captcha"
)

// var store = base64Captcha.DefaultMemStore

var store = util.NewDefaultCaptchaRedisStore()

type AuthorityController struct {
	userService service.UserService
}

func NewAuthorityController() *AuthorityController {
	userService := serviceImpl.NewUserService()

	return &AuthorityController{
		userService: userService,
	}
}

func (t *AuthorityController) Captcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
	// cp := base64Captcha.NewCaptcha(driver, store)
	cp := base64Captcha.NewCaptcha(driver, store.UseWithContext(c))
	if id, b64s, _, err := cp.Generate(); err != nil {
		response.Fail(c, "验证码获取失败", map[string]interface{}{})
	} else {
		response.Ok(c, "验证码获取成功", dto.CaptchaResponse{
			CaptchaId:     id,
			PicPath:       b64s,
			CaptchaLength: 4,
		})
	}
}

func (t *AuthorityController) Login(c *gin.Context) {
	var l dto.LoginRequest
	if err := c.ShouldBindJSON(&l); err != nil {
		response.BadRequest(c, "Bad Request:Invalid Parameters", map[string]interface{}{})
		return
	}

	if store.Verify(l.CaptchaId, l.Captcha, true) {
		r, err := t.userService.Login(c, &l)
		if err != nil {
			response.Fail(c, err.Message, map[string]interface{}{})
			return
		}

		response.Ok(c, "登录成功", r)
	} else {
		response.Fail(c, "验证码错误", map[string]interface{}{})
	}
}

func (t *AuthorityController) Register(c *gin.Context) {
	var l dto.RegisterRequest
	if err := c.ShouldBindJSON(&l); err != nil {
		response.BadRequest(c, "Bad Request:Invalid Parameters", map[string]interface{}{})
		return
	}

	if store.Verify(l.CaptchaId, l.Captcha, true) {
		err := t.userService.Register(c, &l)

		if err != nil {
			response.Fail(c, err.Message, map[string]interface{}{})
		} else {
			response.Ok(c, "注册成功", map[string]interface{}{})
		}
	} else {
		response.Fail(c, "验证码错误", map[string]interface{}{})
	}

}

func (t *AuthorityController) RefreshToken(c *gin.Context) {
	var l dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&l); err != nil {
		response.BadRequest(c, "Bad Request:Invalid Parameters", map[string]interface{}{})
		return
	}

	r, err := t.userService.RefreshToken(c, &l)
	if err != nil {
		response.Fail(c, err.Message, map[string]interface{}{})
	} else {
		response.Ok(c, "Refresh token success", r)
	}
}
