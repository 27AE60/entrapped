package entrapped

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Start(addr string) {
	if len(addr) == 0 {
		addr = ":7000"
	}

	// initialize hunt
	go hub.run()

	// initialize router
	router := httprouter.New()

	// make routes
	router.GET("/", home)
	router.GET("/players/:id", addPlayer)

	listenErr := http.ListenAndServe(addr, router)
	if listenErr != nil {
		logger.Println(listenErr)
		return
	}
}

func addPlayer(rw http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	conn, connErr := newConnection(rw, req)
	if connErr != nil {
		error500(rw, connErr)
		return
	}

	hub.add(id, conn)
}

func home(rw http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(rw, req, "./home.html")
}
