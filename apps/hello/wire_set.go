package hello

import (
	"gin-practice/apps/hello/access"
	"github.com/google/wire"
)

var HelloSchemaSet = wire.NewSet(
	access.NewHelloApp(),
)

var HelloSet = wire.NewSet(
	HelloSchemaSet,
)
