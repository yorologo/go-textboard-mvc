package thread

import (
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

func getThread(id_thread uint) (Thread, error) {
	return readThread(id_thread)
}

func getLastThreads() ([]Thread, error) {
	return readLastsThreads(20)
}
