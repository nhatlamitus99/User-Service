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
