package models

type User struct {
	Id         string `json:"id" gorm:"PRIMARY_KEY;NOT_NULL"`
	Pw         string `json:"password"`
	Name       string `json:"name"`
	Isstudent  string `json:"isStudent"`
	Department string `json:"department"`
	Cardinal   string `json:"cardinal"`
	Belong     string `json:"belong"`

	//UserBoard []Board `gorm:"foreignKey:UserId;association_foreignkey:Id"`
}

type Board struct {
	Num     int    `gorm:"AUTO_INCREMENT;PRIMARY_KEY;NOT_NULL"`
	Userid  string `json:"userid" gorm:"NOT_NULL"`
	Title   string `json:"title" gorm:"NOT_NULL"`
	Content string `json:"content" gorm:"NOT_NULL"`
}
