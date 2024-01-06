package model

type Usertype struct{
	Id int `gorm:"type:int;primaryKey" json:"id"`
	Description string `gorm:"type:varchar(255)" json:"description"`
}

