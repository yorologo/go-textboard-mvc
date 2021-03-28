package user

import (
	"github.com/yorologo/board/subthread"
	"github.com/yorologo/board/thread"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID         uint                  `gorm:"primaryKey;autoIncrement"`
	IP         string                `gorm:"not null"`
	Blocked    bool                  `gorm:"default:false"`
	Threads    []thread.Thread       `gorm:"foreignKey:ID_User;references:ID"`
	Subthreads []subthread.Subthread `gorm:"foreignKey:ID_User;references:ID"`
}

func model() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})

	return db
}
