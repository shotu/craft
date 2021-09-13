package main

import (
	"context"
	"fmt"
	"os"

	"github.com/labstack/gommon/log"

	// "github.com/docker/distribution/registry/handlers"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/random"
	"github.com/shotu/craft/config"
	handlers "github.com/shotu/craft/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	CorrelationID = "X-Correlation-ID"
)

// can be used in integration testing
var (
	c   *mongo.Client
	db  *mongo.Database
	col *mongo.Collection
	cfg config.Properties
)

func init() {

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatal("Could not read dconfig : %v", err)
	}

	connectURI := fmt.Sprintf("mongodb://%s:%s", cfg.DBHost, cfg.DBPort) + "/?connect=direct"

	c, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectURI))
	if err != nil {
		log.Fatal("unabale to connect to db : %v", err)
	}
	fmt.Println("Successfully connected to mongodb")
	db = c.Database(cfg.DBName)
	col = db.Collection(cfg.CollectionName)
	// col.Find()
}

func addCorrelationID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Request().Header.Get(CorrelationID)

		var newID string
		if id == "" {
			// generater a random number
			newID = random.String(12)

		} else {
			newID = id
		}

		c.Request().Header.Set(CorrelationID, newID)
		c.Response().Header().Set(CorrelationID, newID)

		return next(c)
	}

}

func main() {

	port := os.Getenv("CRAFT_APP_PORT")
	if port == "" {
		port = "8080"
	}
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	e.Pre(middleware.RemoveTrailingSlash())
	// adding correlation id[tracing accross microservices ]
	e.Pre(addCorrelationID)
	h := handlers.ProductHandler{Col: col}
	// api := e.Group("/api/v1", serverHeader)

	//TODO add api groups and versions
	e.POST("/sl-game", handlers.CreateBoard)
	e.GET("/sl-game/:id", handlers.GetBoard)

	e.PUT("/sl-game/:board_id/players/:id", handlers.RollDice)

	e.POST("/products", h.CreateProducts, middleware.BodyLimit("1M"))
	e.GET("/products", h.GetProducts)

	e.Logger.Infof("Listening on the port %s:%s", cfg.Host, cfg.Port)
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)))

}
