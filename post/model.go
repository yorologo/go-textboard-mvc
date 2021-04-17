package post

type Post struct {
	ID_User uint   `gorm:"not null"`
	Body    string `gorm:"not null"`
}
