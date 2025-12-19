package models

type User struct {
	ID    string `gorm:"primaryKey;type:uuid"`
	Email string `gorm:"-:all"`
}

func (u User) TableName() string {
	return "auth.users"
}

type Editor struct {
	Model
	UserName string `gorm:"uniqueIndex;type:string;size:256;not null"`
	User     User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
