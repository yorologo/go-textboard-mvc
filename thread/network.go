package thread

import (
	"fmt"

	realip "github.com/Ferluci/fast-realip"
	"github.com/gogearbox/gearbox"
)

func GetRoutes(gb gearbox.Gearbox) []*gearbox.Route {
	routes := []*gearbox.Route{
		gb.Get("/", func(ctx gearbox.Context) {
			ctx.SendString("Threads")
		}),
		gb.Get("/getlast", func(ctx gearbox.Context) {
			threads, _ := getLastThreads()

			if threads != nil {
				ctx.Set("Access-Control-Allow-Origin", "*")
				ctx.SendJSON(threads)
			} else {
				ctx.Context().Response.SetStatusCode(404)
				ctx.SendString("Empty")
			}
		}),
		gb.Get("/details:idthread", func(ctx gearbox.Context) {
			var thread Thread
			thread, _ = getThread(ctx.Param("idthread"))

			ctx.Set("Access-Control-Allow-Origin", "*")
			ctx.SendJSON(thread)
		}),
		gb.Post("/new", func(ctx gearbox.Context) {
			client_ip := realip.FromRequest(ctx.Context())

			thread, error := newThread(client_ip, string(ctx.Context().FormValue("title")), string(ctx.Context().FormValue("body")))
			if error == nil {
				// ctx.Context().Response.SetStatusCode(301)
				// ctx.Context().Redirect("/"+fmt.Sprint(thread)+"", 301)
				ctx.SendString(fmt.Sprint(thread))
			} else {
				panic(error)
			}
		}),
	}

	return routes
}
