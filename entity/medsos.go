package entity

type Medsos struct {
	ID        string `gorm:"primaryKey;type:varchar(100)"`
	UserID    string `gorm:"type:varchar(100);not null"`
	Caption   string `gorm:"type:varchar(255);not null"`
	Status    int    `gorm:"type:int;not null"`
	DetailID  string `gorm:"type:varchar(100);not null"`
	CreatedBy string `gorm:"type:varchar(100);not null"`
	ImageURL  string `gorm:"type:varchar(500)"`
}
