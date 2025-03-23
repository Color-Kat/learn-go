package main

import (
	"demo/http/configs"
	"demo/http/internal/auth"
	"demo/http/internal/link"
	"demo/http/internal/stat"
	"demo/http/internal/user"
	"demo/http/pkg/database"
	"demo/http/pkg/event"
	"demo/http/pkg/middleware"
	"fmt"
	"net/http"
)

func App() http.Handler {
	config := configs.LoadConfig()
	db := database.NewDB(config)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()

	// Repositories
	linkRepository := link.NewLinkRepository(db)
	userRepository := user.NewUserRepository(db)
	statRepository := stat.NewStatRepository(db)

	// Services
	authService := auth.NewAuthService(userRepository)
	statService := stat.NewStatService(&stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepository,
	})

	// --- handlers --- //
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      config,
		AuthService: authService,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		EventBus:       eventBus,
		Config:         config,
	})
	stat.NewStatHandler(router, stat.StatHandlerDeps{
		StatRepository: statRepository,
		Config:         config,
	})
	// --- handlers --- //

	// Middlewares
	middlewares := middleware.Chain(
		middleware.CORS,
		middleware.Logging,
	)

	go statService.AddClick()

	return middlewares(router)
}

func main() {
	app := App()

	port := "8081"
	server := http.Server{
		Addr:    ":" + port,
		Handler: app,
	}

	fmt.Println("Server is listening on port " + port)
	server.ListenAndServe()

}
