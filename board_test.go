package generator

import (
	"testing"

	"github.com/gogearbox/gearbox"
	"github.com/yorologo/board/subthread"
	"github.com/yorologo/board/thread"
)

func TestMain(t *testing.T) {
	// Setup gearbox
	g := gearbox.New()

	// Group account routes
	g.Get("/", func(ctx gearbox.Context) {
		ctx.SendString("Working")
	})
	g.Group("/thread", thread.GetRoutes(g))
	g.Group("/subthread", subthread.GetRoutes(g))

	// Start service
	g.Start(":3000")
}
