package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/apex/log"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func Start() *echo.Echo {

	e := echo.New()

	go func() {
		log.Warnf("[server] Starting the apps on port %s \n", "8080")
		if err := e.Start(fmt.Sprintf(":%s", "8080")); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("[server] Shutting the apps... [%s]", err)
		}
	}()

	return e

}

func Stop(e *echo.Echo) {

	log.Info("stopping HTTP server")
	log.Info("wait 30 second before calling server shutdown")
	time.Sleep(time.Second * 1)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
