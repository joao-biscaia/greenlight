package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func (app *application) serve() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		ReadTimeout:  time.Minute,
		ErrorLog:     log.New(app.logger, "", 0),
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	app.logger.PrintInfo("Starting %s server on %v", map[string]string{
		"addr": srv.Addr,
		"env":  app.config.env,
	})
	return srv.ListenAndServe()
}
