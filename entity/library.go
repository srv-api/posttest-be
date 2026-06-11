package entity

type LibraryRequest struct {
	ID        string `gorm:"primaryKey;type:uuid"`
	UserID    string `gorm:"type:uuid;not null;index"`
	DetailID  string `gorm:"type:varchar(100);not null"`
	CreatedBy string `gorm:"type:varchar(100);not null"`
}
