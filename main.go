// GameReview Service Sample。
// hogehoge
package main

import (
	"github.com/bmizerany/pat"
	"github.com/macococo/go-gamereviews/conf"
	"github.com/macococo/go-gamereviews/controllers"
	"github.com/yvasiyarov/go-metrics"
	"github.com/yvasiyarov/gorelic"
	"log"
	"net/http"
	"runtime"
	"strconv"
)

var (
	agent *gorelic.Agent
)

func initGorelic() {
	if !conf.IsDev() {
		return
	}

	agent = gorelic.NewAgent()
	agent.NewrelicName = "go-gamereviews"
	agent.Verbose = true
	agent.NewrelicLicense = conf.Config.NewrelicLicense
	agent.CollectHTTPStat = true
	agent.HTTPTimer = metrics.NewTimer()
	agent.Run()
}

func wrapController(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	if agent == nil {
		return f
	}
	return agent.WrapHTTPHandlerFunc(f)
}

func initRouter() {
	port := conf.Config.Port

	m := pat.New()
	m.Get("/api/user/list", http.HandlerFunc(wrapController(controllers.UserListController)))
	m.Get("/api/user/create", http.HandlerFunc(wrapController(controllers.UserCreateController)))
	m.Get("/api/chatroom/list", http.HandlerFunc(wrapController(controllers.ChatroomListController)))
	m.Get("/api/chatroom/create", http.HandlerFunc(wrapController(controllers.ChatroomCreateController)))
	m.Get("/routing/:name/:id", http.HandlerFunc(wrapController(controllers.RoutingController)))
	http.Handle("/", m)
	// http.HandleFunc("/api/user/list", wrapController(controllers.UserListController))
	// http.HandleFunc("/api/user/create", wrapController(controllers.UserCreateController))
	// http.HandleFunc("/api/chatroom/list", wrapController(controllers.ChatroomListController))
	// http.HandleFunc("/api/chatroom/create", wrapController(controllers.ChatroomCreateController))

	log.Println("HTTP listen port:", port)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func main() {
	log.Println("runmode:", conf.Config.Runmode)
	log.Println("NumCPU:", runtime.NumCPU())
	log.Println("GOMAXPROCS:", runtime.GOMAXPROCS(runtime.NumCPU()))

	initGorelic()
	initRouter()
}
