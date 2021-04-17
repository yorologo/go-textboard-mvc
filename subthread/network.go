package subthread

import (
	"fmt"

	realip "github.com/Ferluci/fast-realip"
	"github.com/gogearbox/gearbox"
)

func GetRoutes(gb gearbox.Gearbox) []*gearbox.Route {
	routes := []*gearbox.Route{
		gb.Get("/", func(ctx gearbox.Context) {
			ctx.SendString("Subthreads")
		}),
		gb.Get("/:idthread", func(ctx gearbox.Context) {
			subthreads, _ := getLastSubthreads(ctx.Param("idthread"))

			if subthreads != nil {
				ctx.Set("Access-Control-Allow-Origin", "*")
				ctx.SendJSON(subthreads)
			} else {
				ctx.Context().Response.SetStatusCode(404)
				ctx.SendString("Empty")
			}
		}),
		gb.Post("/new", func(ctx gearbox.Context) {
			client_ip := realip.FromRequest(ctx.Context())

			subthread, _ := newSubthread(client_ip, string(ctx.Context().FormValue("thread")), string(ctx.Context().FormValue("body")))
			if subthread != 0 {
				// ctx.Context().Response.SetStatusCode(301)
				// ctx.Context().Redirect("/"+string(ctx.Context().FormValue("thread"))+"", 301)
				ctx.SendString(fmt.Sprint(subthread))
			} else {
				ctx.SendString("Error")
			}
		}),
	}

	return routes
}
