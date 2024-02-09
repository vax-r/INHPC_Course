package bootstrap

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Application struct {
	Env    *Env
	Conn   *gorm.DB
	Engine *gin.Engine
}

func App() *Application {
	env := NewEnv()
	db := NewDB(env)
	engine := gin.Default()

	// Set timezone
	tz, err := time.LoadLocation(env.Server.TimeZone)
	if err != nil {
		log.Fatal(err)
	}
	time.Local = tz

	app := &Application{
		Env:    env,
		Conn:   db,
		Engine: engine,
	}

	return app
}

// Run the Application
func (app *Application) Run() {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.Env.Server.Port),
		Handler: app.Engine,
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Printf("Server is running on port %d", app.Env.Server.Port)
		serverErrors <- server.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		log.Fatalf("Error starting server: %v", err)

	case <-shutdown:
		log.Println("Shutting down the server...")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		err := server.Shutdown(ctx)
		if err != nil {
			log.Fatalf("Could not stop server gracefully: %v", err)
			err = server.Close()
			if err != nil {
				log.Fatalf("Could not stop http server: %v", err)
			}
		}
	}

}
