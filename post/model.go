package post

type Post struct {
	ID      uint   `gorm:"primaryKey;autoIncrement"`
	ID_User uint   `gorm:"not null"`
	Body    string `gorm:"not null"`
}
