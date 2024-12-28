package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type config struct {
	port string
}

type application struct {
	config *config
	logger *slog.Logger
}

func (app *application) test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("RSS House"))
}

// i need a way to pass in cmd line arg to config stuct
func main() {
	var cfg config
	flag.StringVar(&cfg.port, "port", "4040", "API server port")

	flag.Parse()

	app := &application{
		config: &cfg,
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}

	err := app.Serve()
	if err != nil {
		app.logger.Error(err.Error())
		os.Exit(1)
	}
}
