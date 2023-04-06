package access

import (
	"luwu/libs/router"
)

func (h *HelloApp) URLPatterns() []router.Route {
	var routers []router.Route
	group := router.NewRouterGroup("/luwu")
	{
		route := group.Group("/hello")
		{
			route.GET("/say_hello", h.SayHello)
			route.GET("/say_failed", h.SayFailed)
		}
	}
	routers = append(routers, group.GetRouters()...)
	return routers
}
