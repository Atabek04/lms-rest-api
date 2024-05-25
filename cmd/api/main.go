package main

import (
	"errors"
	"github.com/rs/zerolog"
	"lms-crud-api/internal/data"
	"log"
	"time"

	_ "database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

type application struct {
	config config
	logger zerolog.Logger
	models data.Models
}

func main() {
	var cfg config
	cfg.port = 4000
	cfg.env = "development"
	cfg.db.dsn = "user=youruser password=yourpassword dbname=yourdbname sslmode=disable"

	db, err := openDB(cfg)
	if err != nil {
		log.Fatalln("Failed to connect to database")
	}
	defer db.Close()

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	migrationsPath := "file://migrations"
	m, err := migrate.New(migrationsPath, cfg.db.dsn)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to prepare database migration")
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.Fatal().Err(err).Msg("Failed to migrate database")
	}

	router := gin.Default()
	router.POST("/v1/courses", app.createCourseHandler)
	router.GET("/v1/courses/:id", app.showCourseHandler)
	router.PUT("/v1/courses/:id", app.updateCourseHandler)
	router.DELETE("/v1/courses/:id", app.deleteCourseHandler)
	router.POST("/v1/modules", app.createModuleHandler)
	router.GET("/v1/modules/:id", app.showModuleHandler)
	router.PUT("/v1/modules/:id", app.updateModuleHandler)
	router.DELETE("/v1/modules/:id", app.deleteModuleHandler)
	router.POST("/v1/lessons", app.createLessonHandler)
	router.GET("/v1/lessons/:id", app.showLessonHandler)
	router.PUT("/v1/lessons/:id", app.updateLessonHandler)
	router.DELETE("/v1/lessons/:id", app.deleteLessonHandler)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Info().Msgf("Starting server on %s", srv.Addr)
	err = srv.ListenAndServe()
	if err != nil {
		logger.Fatal().Err(err).Msg("Could not start server")
	}
}

func openDB(cfg config) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (app *application) badRequestResponse(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func (app *application) serverErrorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func (app *application) notFoundResponse(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
}

func (app *application) writeJSON(c *gin.Context, statusCode int, data gin.H) {
	c.JSON(statusCode, data)
}
