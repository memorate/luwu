package access

import (
	"gin-practice/libs/router"
)

func (h *HelloApp) URLPatterns() []router.Route {
	var routers []router.Route
	group := router.NewRouterGroup("/gin_practice")
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
