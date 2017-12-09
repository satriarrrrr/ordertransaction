package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/satriarrrrr/store/infrastructures"
	"github.com/spf13/viper"
)

var (
	configPath     *string
	configFilename *string
	location       *time.Location
	logger         *log.Logger
)

func init() {
	configPath = flag.String("config-path", ".", "Path to config file, exclude file name")
	configFilename = flag.String("config-filename", "config", "Config filename without extension")
	flag.Parse()
	location, _ = time.LoadLocation("Asia/Jakarta")
	logger = log.New(os.Stdout, fmt.Sprintf("[%s] ", time.Now().In(location).Format("2006-01-02 15:04:05")), log.Lshortfile)
}

func main() {
	// Read configuration file
	if err := infrastructures.ReadConfig(*configPath, *configFilename); err != nil {
		logger.Fatalf("[%s] Failed to read config. Got error: %v", "ERROR", err)
	}

	// Open database connection
	conn := infrastructures.CreateMySQLConnection("store")
	db, err := conn.Open()
	if err != nil {
		logger.Fatalf("[%s] Failed to open db connection. Got error: %v", "ERROR", err)
	}
	defer db.Close()

	// Subscribe to quit signal
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	// Initialize service

	// Routing
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/text")
		w.Write([]byte("Hello, world!"))
	})

	// Run server
	s := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", viper.GetString("http.addr"), viper.GetString("http.port")),
		Handler:      mux,
		ReadTimeout:  time.Duration(viper.GetInt("http.read_timeout")) * time.Second,
		WriteTimeout: time.Duration(viper.GetInt("http.write_timeout")) * time.Second,
	}

	go func() {
		<-quit
		logger.Printf("[%s] Shutting down server", "INFO")
		if err := s.Shutdown(context.Background()); err != nil {
			logger.Printf("[%s] Failed to shutdown. Got error: %v ", "ERROR", err)
		}
	}()

	logger.Printf("[%s] Listening on %s \n", "INFO", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		logger.Fatal(err)
	}
	logger.Printf("[%s] Server gracefully stopped", "INFO")
}
