package thread

import (
	"strconv"

	"github.com/yorologo/board/post"
	"github.com/yorologo/board/user"
)

func newThread(ip string, title string, body string) (uint, error) {
	var thread *Thread
	IdIp, error := user.ArchiveUser(ip)

	if error != nil {
		thread = &Thread{
			Title: title,
			Post: post.Post{
				Body:    body,
				ID_User: IdIp,
			},
		}
		return createThread(thread)
	}
	return 0, error
}

func getThread(id_thread string) (Thread, error) {
	id, _ := strconv.ParseUint(id_thread, 10, 32)
	return readThread(uint(id))
}

func getLastThreads() ([]Thread, error) {
	return readLastsThreads(20)
}
