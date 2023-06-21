package route

import (
	"net/http"

	"github.com/Aldikon/API-user/api/handlers"
	"github.com/Aldikon/API-user/bootstrap"
	"github.com/Aldikon/API-user/db"
	"github.com/Aldikon/API-user/redis"
	"github.com/Aldikon/API-user/service"
)

func Setup(app *bootstrap.Application, mux *http.ServeMux) {
	userDB := db.NewUser(app.DB)
	cacheDB := redis.NewCache(app.RDB)
	userService := service.NewUser(userDB, cacheDB)
	userHandlers := handlers.NewUser(userService)

	mux.HandleFunc("/user", userHandlers.Users)
	mux.HandleFunc("/user/", userHandlers.User)

	mux.HandleFunc("/", handlers.Index)
}
