package main

import (
	"context"
	"fmt"
	"os"

	"github.com/labstack/gommon/log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/random"
	"github.com/shotu/craft/config"
	handlers "github.com/shotu/craft/handlers"
	handlersv1 "github.com/shotu/craft/handlersv1"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	CorrelationID = "X-Correlation-ID"
)

// can be used in integration testing
var (
	c        *mongo.Client
	db       *mongo.Database
	col      *mongo.Collection
	boardCol *mongo.Collection
	cfg      config.Properties
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
	boardCol = db.Collection(cfg.BoardCollectionName)
	// col.UpdateOne()
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
	b := handlers.BoardHandler{Col: boardCol}
	// api := e.Group("/api/v1", serverHeader)

	//APIs without DB(in memory Implementation)
	e.POST("/api/v1/board", handlersv1.CreateBoard)
	e.GET("/api/v1/board/:id", handlersv1.GetBoard)
	e.PUT("/api/v1/board/:board_id/players/:id", handlersv1.RollDice)

	//APIs with DB(mongodb as persistent storage Implementation)
	e.POST("/api/v2/boards", b.CreateBoard)
	e.GET("/api/v2/boards/:id", b.GetBoard)
	e.PUT("/api/v2/boards/:board_id/players/:id", b.RollDice)

	// e.GET("/api/v2/board/:id", handlers.GetBoard)
	// e.PUT("/api/v2/board/:board_id/players/:id", handlers.RollDice)

	e.POST("/products", h.CreateProducts, middleware.BodyLimit("1M"))
	e.GET("/products", h.GetProducts)
	e.PUT("/products/:id", h.UpdateProduct, middleware.BodyLimit("1M"))

	e.Logger.Infof("Listening on the port %s:%s", cfg.Host, cfg.Port)
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)))

}
