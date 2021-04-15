package subthread

import (
	"strconv"

	"github.com/yorologo/board/post"
	"github.com/yorologo/board/user"
)

func newSubthread(ip string, id_thread string, body string) (uint, error) {
	var subthread *Subthread
	IdIp, error := user.ArchiveUser(ip)
	idthread, _ := strconv.ParseUint(id_thread, 10, 32)

	if error != nil {
		subthread = &Subthread{
			Post: post.Post{
				Body:    body,
				ID_User: IdIp,
			},
			ID_Thread: uint(idthread),
		}

		return createSubthread(subthread)

	}
	return 0, error
}

func getLastSubthreads(id_thread string) ([]Subthread, error) {
	id, _ := strconv.ParseUint(id_thread, 10, 32)
	return readLastsSubthreads(uint(id), 20)
}
