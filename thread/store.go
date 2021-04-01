package thread

import "gorm.io/gorm"

var md *gorm.DB = model()

func createThread(Thread *Thread) (uint, error) {
	result := md.Create(&Thread)

	return Thread.ID, result.Error
}

func readThread(id uint) (Thread, error) {
	var Thread Thread
	result := md.First(&Thread, "ID = ?", id)

	return Thread, result.Error
}

func readLastsThreads(n int) ([]Thread, error) {
	var Thread []Thread
	result := md.Limit(n).Find(&Thread)

	return Thread, result.Error
}
