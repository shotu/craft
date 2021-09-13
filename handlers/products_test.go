package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo"
	"github.com/shotu/craft/config"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	c   *mongo.Client
	db  *mongo.Database
	col *mongo.Collection
	cfg config.Properties
	h   *ProductHandler
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

}
func TestProduct(t *testing.T) {

	t.Run("test create product", func(t *testing.T) {
		body := `[{
			"product_name": "ayushpaste" , 
			"is_essential":true
		}]`
		req, err := http.NewRequest("POST", "/products", strings.NewReader(body))
		if err != nil {
			fmt.Println("Err ", err)
		}

		res := httptest.NewRecorder()

		e := echo.New()

		c := e.NewContext(req, res)

		h.Col = col
		createrr := h.CreateProducts(c)
		assert.Nil(nil, createrr)
		// if assert.Error(t, h.CreateProducts(c)){

		// }
	})
}
