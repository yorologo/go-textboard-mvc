package subthread

import "gorm.io/gorm"

var md *gorm.DB = model()

func createSubthread(Subthread *Subthread) (uint, error) {
	result := md.Create(&Subthread)

	return Subthread.ID, result.Error
}

func readLastsSubthreads(id_thread uint, n int) ([]Subthread, error) {
	var Subthread []Subthread
	result := md.Where("ID_Thread = ?", id_thread).Limit(n).Find(&Subthread)

	return Subthread, result.Error
}
