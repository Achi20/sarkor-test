package phone

// may include fields such as limit and offset
type Filter struct {
	Phone *string
}

type Create struct {
	Phone       string `json:"phone" binding:"required,max=12"`
	Description string `json:"description"`
	IsFax       bool   `json:"is_fax" binding:"required"` // using CamelCase to write idiomatic go code
	UserID      int    `json:"-"`
}

// does not specify which one is used: json or form
type Update struct {
	PhoneID     int    `json:"phone_id" form:"phone_id" binding:"required,max=12"`
	Phone       string `json:"phone" form:"phone"`
	Description string `json:"description" form:"description"`
	IsFax       bool   `json:"is_fax" form:"is_fax"`
	UserID      int    `json:"-" form:"-"`
}

type List struct {
	UserID      int    `json:"user_id"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
	IsFax       bool   `json:"is_fax"`
}
