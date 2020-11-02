package entities

type TokenRequest struct {
	Username      string `json:"name"`
	Password      string `json:"password"`
	Grant_Type    string `json:"grant-type"`
	Client_Secret string `json:"client-secret"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	CreateAt string `json:"createAt"`
	UpdateAt string `json:"updateAt"`
}

type Data struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	CreateAt string `json:"createAt"`
	UpdateAt string `json:"updateAt"`
}
