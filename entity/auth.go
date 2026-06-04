package entity

import (
	"time"

	"gorm.io/gorm"
)

type AccessDoor struct {
	ID            string         `gorm:"primary_key;type:varchar(39)" json:"id"`
	FullName      string         `gorm:"full_name;type:varchar(70)" json:"full_name"`
	Gender        string         `gorm:"type:varchar(5)" json:"gender"`
	Age           int            `gorm:"age" json:"age"`
	Whatsapp      string         `gorm:"uniqueIndex;type:varchar(200)" json:"whatsapp"`
	Email         string         `gorm:"uniqueIndex;type:varchar(150)" json:"email"`
	Password      string         `gorm:"password" json:"password"`
	DetailID      string         `gorm:"detail_id" json:"detail_id"`
	GoogleID      string         `gorm:"google_id" json:"google_id"`
	Provider      string         `gorm:"provider" json:"provider"`
	AccessRoleID  string         `gorm:"access_role_id" json:"access_role_id"`
	LoginAttempts int            `gorm:"login_attempts" json:"login_attempts"`
	Suspended     bool           `gorm:"suspended" json:"suspended"`
	LastAttempt   time.Time      `gorm:"last_attempt" json:"last_attempt"`
	Verified      UserVerified   `json:"verified" gorm:"foreignKey:UserID"`
	CreatedBy     string         `gorm:"created_by" json:"created_by"`
	UpdatedBy     string         `gorm:"updated_by" json:"updated_by"`
	DeletedBy     string         `gorm:"deleted_by" json:"deleted_by"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
