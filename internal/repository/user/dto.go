package user

type CookieAuth struct {
	UserID int
	Login  string
}

type Create struct {
	Login    string `json:"-" form:"login" binding:"required"`
	Password string `json:"-" form:"password" binding:"required"`
	Name     string `json:"-" form:"name" binding:"required"`
	Age      int    `json:"-" form:"age" binding:"required"`
}

type Auth struct {
	Login    string `json:"login" form:"-" binding:"required"`
	Password string `json:"password" form:"-" binding:"required"`
}

type Detail struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
}
