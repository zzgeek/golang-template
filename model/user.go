package model

type User struct {
	UserId     int    `json:"userId"`
	UserPwd    string `json:"userPwd"`
	UserName   string `json:"userName"`
	UserStatus int    `json:"userStatus"`
	UserRole   string `json:"userRole"`
	UserToken  string `json:"userToken"`
}
