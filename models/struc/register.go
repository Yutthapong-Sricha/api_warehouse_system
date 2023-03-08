package models

// Binding from JSON
type Register struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
}

type Login struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Id_branch string `json:"fullname" binding:"required"`
}
