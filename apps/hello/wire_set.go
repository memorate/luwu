package hello

import (
	"github.com/google/wire"
	"luwu/apps/hello/access"
)

var HelloSchemaSet = wire.NewSet(
	access.NewHelloApp,
)

var HelloSet = wire.NewSet(
	HelloSchemaSet,
)
