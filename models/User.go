package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Email     string    `gorm:"type:citext;uniqueIndex;not nullj" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"password"`
	Age       *int      `gorm:"type:int" json:"age"`
	Sex       *string   `gorm:"type:varchar(10)" json:"sex"`
	Photo     *string   `gorm:"type:text" json:"photo"`
	Bio       *string   `gorm:"type:text" json:"bio"`
	Instagram *string   `gorm:"type:text" json:"instagram"`
	Linkedin  *string   `gorm:"type:text" json:"linkedin"`
	Github    *string   `gorm:"type:text" json:"github"`
}
