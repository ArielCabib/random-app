package routes

import (
	"github.com/ArielCabib/random-app/server/api/todo/routes"
	"github.com/ArielCabib/random-app/server/common/static"
	"github.com/julienschmidt/httprouter"
)

func Init(router *httprouter.Router) {
	todoroutes.Init(router)
	static.Init(router)
}
