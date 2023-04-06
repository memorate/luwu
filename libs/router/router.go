package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"luwu/common/constant"
	errorLib "luwu/libs/error"
	httpLib "luwu/libs/http"
	"net/http"
	"reflect"
	"sync"
)

type Route struct {
	Method       string      //Method is one of the following: GET,PUT,POST,DELETE. required
	Path         string      //Path contains a path pattern. required
	ResourceFunc handlerFunc //the func this API calls.
}

type routerGroup struct {
	root     *routerGroup
	routers  []Route
	basePath string // the common prefix for all routes in this group
}

type Router interface {
	//URLPatterns returns route
	URLPatterns() []Route
}

func (r *Route) RegisterRoute(e *gin.Engine) {
	tempF := r.ResourceFunc
	f := func(ctx *gin.Context) {
		rsp, err := tempF(ctx)
		if err != nil {
			if err.Code == 0 {
				err.Code = constant.RequestFailed
			}
			ctx.JSON(http.StatusOK, httpLib.NewResponse(nil, err))
		} else {
			ctx.JSON(http.StatusOK, httpLib.NewResponse(rsp, nil))
		}
	}
	e.Handle(r.Method, r.Path, f)
}

func NewRouterGroup(basePath string) *routerGroup {
	rg := &routerGroup{basePath: basePath}
	rg.root = rg
	return rg
}

func (rg *routerGroup) GetRouters() []Route {
	return rg.routers
}

func (rg *routerGroup) Group(path string) *routerGroup {
	bp := rg.basePath + path
	if rg.basePath == "/" {
		bp = path
	}
	return &routerGroup{
		basePath: bp,
		root:     rg.root,
	}
}

type handlerFunc func(ctx *gin.Context) (interface{}, *errorLib.Error)

func (rg *routerGroup) POST(relativePath string, handlers handlerFunc) *routerGroup {
	route := Route{
		Method:       http.MethodPost,
		Path:         rg.basePath + relativePath,
		ResourceFunc: handlers,
	}
	rg.combine(route)
	return rg
}

func (rg *routerGroup) GET(relativePath string, handlers handlerFunc) *routerGroup {
	route := Route{
		Method:       http.MethodGet,
		Path:         rg.basePath + relativePath,
		ResourceFunc: handlers,
	}
	rg.combine(route)
	return rg
}

func (rg *routerGroup) combine(route Route) {
	root := rg.root
	if rg.root == nil {
		root = rg
	}
	root.routers = append(root.routers, route)
}

func (rg *routerGroup) GetRouteList() []Route {
	return rg.routers
}

func GetUrlPatterns(schema interface{}) ([]Route, error) {
	v, ok := schema.(Router)
	if !ok {
		return []Route{}, fmt.Errorf("can not register APIs to server: %s", reflect.TypeOf(schema).String())
	}
	return v.URLPatterns(), nil
}

var ginInstance = &GinInstance{}

type Schema struct {
	serverName string
	schema     interface{} // business application ptr instance
}

type GinInstance struct {
	version string
	schemas []*Schema
	mu      sync.Mutex
}

func RegisterSchema(severName string, schema interface{}) {
	ginInstance.registerSchema(severName, schema)
}

func (g *GinInstance) registerSchema(serverName string, structPtr interface{}) {
	schema := &Schema{
		serverName: serverName,
		schema:     structPtr,
	}
	g.mu.Lock()
	g.schemas = append(g.schemas, schema)
	g.mu.Unlock()
}

func RegisterUrlPatterns(engine *gin.Engine) error {
	for _, schema := range ginInstance.schemas {
		if routes, e := GetUrlPatterns(schema.schema); e != nil {
			return e
		} else {
			for _, route := range routes {
				route.RegisterRoute(engine)
			}
		}
	}
	return nil
}
