package routes

import (
	"github.com//random-app/server/api/todo/routes"
	"github.com//random-app/server/common/static"
	"github.com/julienschmidt/httprouter"
)

func Init(router *httprouter.Router) {
	todoroutes.Init(router)
	static.Init(router)
}
