package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

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

func findProducts(ctx context.Context, q url.Values, collection dbiface.CollectionAPI) ([]Product, error) {
	var products []Product

	filter := make(map[string]interface{})

	for k, v := range q {
		filter[k] = v[0]
	}

	if filter["_id"] != "" {
		docID, err := primitive.ObjectIDFromHex(filter["_id"].(string))
		if err != nil {
			return products, err
		}
		filter["_id"] = docID

	}
	// cursor, err := collection.Find(ctx, bson.M{})
	cursor, err := collection.Find(ctx, bson.M(filter))

	if err != nil {
		log.Errorf("Unable to get products: %v", err)
		return products, err
	}

	curerr := cursor.All(ctx, &products)
	if curerr != nil {
		log.Errorf("Unable to read the cursor : %v", err)
		return nil, curerr
	}
	return products, nil

}

func modifyProduct(ctx context.Context, id string, reqBody io.ReadCloser, collection dbiface.CollectionAPI) (Product, error) {
	var product Product
	// find if product exists, if err return 404
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("cannot convert to objectid: %v", err)
		return product, err
	}
	filter := bson.M{"_id": docID}

	res := collection.FindOne(ctx, filter)
	if err := res.Decode(&product); err != nil {

		log.Errorf("unable to  decode product: %v", err)
		return product, err
	}

	//decode the req payload, if return 500
	decodeError := json.NewDecoder(reqBody).Decode(&product)
	if err != nil {
		return product, decodeError
	}

	//validate the req, if err return 400

	if err := v.Struct(product); err != nil {
		return product, echo.NewHTTPError(500, "unable to decode")
	}

	// update the product
	_, err = collection.UpdateOne(ctx, filter, bson.M{"$set": product})
	if err != nil {
		return product, err
	}

	return product, nil

}

// create tthe products in mongodb
func (h *ProductHandler) CreateProducts(c echo.Context) error {

	var products []Product

	fmt.Println("product is ", products)

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

	fmt.Println("product are ", products)

	IDs, err := insertProducts(context.Background(), products, h.Col)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, IDs)
}

// Gets list of products
func (h *ProductHandler) GetProducts(c echo.Context) error {

	products, err := findProducts(context.Background(), c.QueryParams(), h.Col)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) UpdateProduct(c echo.Context) error {

	product, err := modifyProduct(context.Background(), c.Param("id"), c.Request().Body, h.Col)

	if err != nil {
		log.Errorf("unable to bind the product: %v", err)
		return err
	}
	return c.JSON(http.StatusOK, product)
}
