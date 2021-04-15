package thread

import (
	"fmt"

	realip "github.com/Ferluci/fast-realip"
	"github.com/gogearbox/gearbox"
)

type response struct {
	threads []Thread
}

func GetRoutes() []*gearbox.Route {
	gb := gearbox.New()

	routes := []*gearbox.Route{
		gb.Get("/getlast", func(ctx gearbox.Context) {
			var res response
			res.threads, _ = getLastThreads()

			ctx.Set("Access-Control-Allow-Origin", "*")
			ctx.SendJSON(res)
		}),
		gb.Get("/:id", func(ctx gearbox.Context) {
			var thread Thread
			thread, _ = getThread(ctx.Param("id"))

			ctx.Set("Access-Control-Allow-Origin", "*")
			ctx.SendJSON(thread)
		}),
		gb.Post("/new", func(ctx gearbox.Context) {
			client_ip := realip.FromRequest(ctx.Context())

			thread, _ := newThread(client_ip, ctx.Param("title"), ctx.Param("body"))

			ctx.Context().Response.SetStatusCode(301)
			ctx.Context().Redirect("/"+fmt.Sprint(thread)+"", 301)
		}),
	}

	return routes
}
