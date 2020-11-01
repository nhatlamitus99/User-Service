package entities

type Owner struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
