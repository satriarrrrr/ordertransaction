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

	"github.com/satriarrrrr/store/helpers"
	"github.com/satriarrrrr/store/products"
	"goji.io/pat"

	"github.com/satriarrrrr/store/infrastructures"
	"github.com/spf13/viper"
	"goji.io"
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
	productRepository := products.NewProductsRepository(db)
	productService := products.NewProductsService(logger, productRepository)
	productController := products.NewProductsController(productService)

	// Routing
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/"), func(w http.ResponseWriter, r *http.Request) {
		helpers.ResponseJSON(w, "Hello, world!", 1000, 200)
	})
	mux.HandleFunc(pat.Get("/healthcheck"), func(w http.ResponseWriter, r *http.Request) {
		helpers.ResponseJSON(w, "Ok!", 1000, 200)
	})
	mux.HandleFunc(pat.Get("/ping"), func(w http.ResponseWriter, r *http.Request) {
		if err = db.Ping(); err != nil {
			helpers.ResponseJSON(w, "Failed!", 1000, 200)
		} else {
			helpers.ResponseJSON(w, "Pong", 1000, 200)
		}
	})

	mux.HandleFunc(pat.Get("/products"), productController.GetProducts)
	mux.HandleFunc(pat.Get("/products/:id"), productController.GetProductByID)

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
