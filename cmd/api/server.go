package main

import (
	"log/slog"
	"net/http"
	"time"
)

func (app *application) Serve() error {
	srv := http.Server{
		Addr:         ":" + app.config.port,
		Handler:      app.route(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
		ErrorLog:     slog.NewLogLogger(app.logger.Handler(), slog.LevelError),
	}
	app.logger.Info("starting server", "addr", srv.Addr)
	return srv.ListenAndServe()
}
