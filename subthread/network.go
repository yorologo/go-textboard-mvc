package subthread

import (
	realip "github.com/Ferluci/fast-realip"
	"github.com/gogearbox/gearbox"
)

type response struct {
	subthreads []Subthread
}

func GetRoutes() []*gearbox.Route {
	gb := gearbox.New()

	routes := []*gearbox.Route{
		gb.Get("/:idthread", func(ctx gearbox.Context) {
			var res response
			res.subthreads, _ = getLastSubthreads(ctx.Param("idthread"))

			ctx.Set("Access-Control-Allow-Origin", "*")
			ctx.SendJSON(res)
		}),
		gb.Post("/new", func(ctx gearbox.Context) {
			client_ip := realip.FromRequest(ctx.Context())

			subthread, _ := newSubthread(client_ip, ctx.Param("thread"), ctx.Param("body"))
			if subthread != 0 {
				ctx.Context().Response.SetStatusCode(301)
				ctx.Context().Redirect("/"+ctx.Param("thread")+"", 301)
			} else {
				ctx.SendString("Error")
			}
		}),
	}

	return routes
}
