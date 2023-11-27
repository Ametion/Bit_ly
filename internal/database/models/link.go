package database_models

type ShortLink struct {
	ID           uint   `gorm:"primaryKey" json:id`
	OriginalLink string `gorm:"size:255" json:originalLink`
	ShortLink    string `gorm:"size:20" json:shortLink`
}
