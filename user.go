package todo

type User struct {
	Id       int    `json:"-" db:"user_id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Passowrd string `json:"password" binding:"required"`
}
