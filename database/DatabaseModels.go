package database

type ShortLink struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	OriginalLink string `gorm:"size:255" json:"originalLink"`
	ShortedLink  string `gorm:"size:20" json:"shortedLink"`
	UserID       uint   `gorm:"not null"`
	User         User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type User struct {
	ID         uint        `gorm:"primaryKey" json:"id"`
	FirstName  string      `gorm:"size:15" json:"firstName"`
	SecondName string      `gorm:"size:15" json:"secondName"`
	Login      string      `gorm:"size:20;unique" json:"login"`
	Password   string      `gorm:"size:100" json:"password"`
	Links      []ShortLink `gorm:"foreignKey:UserID"`
}
