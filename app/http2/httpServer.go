package http2

import (
	"day1/internal/router"
	util2 "day1/util"
)

func Start() {
	r := router.InitRouter()

	r.Run("localhost:" + util2.AnalyIni())
}
