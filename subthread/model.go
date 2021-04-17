package subthread

import (
	"github.com/yorologo/board/post"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Subthread struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Post      post.Post `gorm:"embedded"`
	ID_Thread uint
}

func model() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Subthread{})

	return db
}
