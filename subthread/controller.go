package subthread

import (
	"github.com/yorologo/board/post"
	"github.com/yorologo/board/user"
)

func newSubthread(ip string, title string, id_thread uint) (uint, error) {
	var subthread *Subthread
	IdIp, error := user.ArchiveUser(ip)

	if error != nil {
		subthread = &Subthread{
			{
			Post: post.Post{
				Body:    body,
				ID_User: IdIp,
			},
			ID_Thread: id_thread,
		}
		return createSubthread(subthread)

	}
	return 0, error
}

func getSubthread(id_Subthread uint) (Subthread, error) {
	return readSubthread(id_Subthread)
}

func getLastSubthreads() ([]Subthread, error) {
	return readLastsSubthreads(20)
}
