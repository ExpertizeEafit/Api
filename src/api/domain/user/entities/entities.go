package entities

type LoginBasicInfo struct {
	Id             int    `json:"id"`
	Username       string `json:"username"`
	ChangePassword bool   `json:"change_password"`
}

type LoginData struct {
	Id       string `json:"dni"`
	Password string `json:"password"`
}
