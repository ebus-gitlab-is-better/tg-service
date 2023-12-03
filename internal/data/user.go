package data

type User struct {
	ID     uint32 `gorm:"primaryKey"`
	ChatID int64  `gorm:"unique"`
}
