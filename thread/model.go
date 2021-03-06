package thread

import (
	"github.com/yorologo/board/post"
	"github.com/yorologo/board/subthread"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Thread struct {
	ID         uint                  `gorm:"primaryKey;autoIncrement"`
	Post       post.Post             `gorm:"embedded"`
	Tittle     string                `gorm:"not null"`
	Subthreads []subthread.Subthread `gorm:"foreignKey:ID_Thread;references:ID"`
}

func model() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Thread{})

	return db
}
