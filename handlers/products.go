package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/gommon/log"

	"github.com/labstack/echo/v4"
	"github.com/shotu/craft/dbiface"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	validator "gopkg.in/go-playground/validator.v9"
)

var (
	v = validator.New()
)

type Product struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"product_name" bson:"product_name" validate:"required,max=10"`
	IsEssential bool               `json:"is_essential" bson:"is_essential" `
}

type ProductHandler struct {
	Col dbiface.CollectionAPI
}

// ProductValidator a product validator
type ProductValidator struct {
	validator *validator.Validate
}

// validates product
func (p *ProductValidator) Validate(i interface{}) error {
	return p.validator.Struct(i)
}

func insertProducts(ctx context.Context, products []Product, collection dbiface.CollectionAPI) ([]interface{}, error) {

	var insertedIds []interface{}
	for _, product := range products {
		product.ID = primitive.NewObjectID()
		insertID, err := collection.InsertOne(ctx, product)
		if err != nil {
			log.Printf("unable to insert: %v", err)
			return nil, err
		}
		insertedIds = append(insertedIds, insertID)
	}

	return insertedIds, nil

}

func findProducts(ctx context.Context, collection dbiface.CollectionAPI) ([]Product, error) {
	var products []Product
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Errorf("Unable to get products: %v", err)
	}

	curerr := cursor.All(ctx, &products)
	if curerr != nil {
		log.Errorf("Unable to read the cursor : %v", err)
		return nil, curerr
	}
	return products, nil

}

// create tthe products in mongodb
func (h *ProductHandler) CreateProducts(c echo.Context) error {

	var products []Product
	c.Echo().Validator = &ProductValidator{validator: v}
	for _, product := range products {
		err := c.Validate(product)
		if err != nil {
			log.Printf("unabale to validate product %+v %v", product, err)
			return err
		}
	}

	if err := c.Bind(&products); err != nil {
		log.Fatalf("unable to bind: %v", err)
		return err
	}

	IDs, err := insertProducts(context.Background(), products, h.Col)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, IDs)
}

// Gets list of products
func (h *ProductHandler) GetProducts(c echo.Context) error {

	products, err := findProducts(context.Background(), h.Col)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, products)
}
