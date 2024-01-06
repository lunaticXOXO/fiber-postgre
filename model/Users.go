package model

type Users struct{

	Username string     `gorm:"type:varchar(255)" json:"username"`
	Password string     `gorm:"type:varchar(255)" json:"password"`
	User_type int    	`gorm:"type:int" json:"user_type"`
	Usertype Usertype	`gorm:"foreignKey:User_type;references:Id"`

}