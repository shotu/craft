package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shotu/craft/dbiface"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"product_name" bson:"product_name" validate:"required,max=10"`
	IsEssential bool               `json:"is_essential" bson:"is_essential" `
}

type ProductHandler struct {
	Col dbiface.CollectionAPI
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

// create tthe products in mongodb
func (h *ProductHandler) CreateProducts(c echo.Context) error {

	var products []Product
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
