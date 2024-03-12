package dto

type CaptchaResponse struct {
	CaptchaId     string `json:"captcha_id"`
	PicPath       string `json:"picture_path"`
	CaptchaLength int    `json:"captcha_length""`
}

type LoginRequest struct {
	Mobile    string `json:"mobile"`
	Password  string `json:"password"`
	CaptchaId string `json:"captcha_id"`
	Captcha   string `json:"captcha"`
}

type UserDTO struct {
	UserId   string `json:"user_id"`
	Mobile   string `json:"mobile"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User         UserDTO `json:"user"`
	AccessToken  string  `json:"access_token"`
	RefreshToken string  `json:"refresh_token"`
}

type RegisterRequest struct {
	User UserDTO `json:"user"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
