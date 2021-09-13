package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/shotu/craft/dbiface"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	validator "gopkg.in/go-playground/validator.v9"
)

// var (
// 	v2 = validator.New()
// )

type Board struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	// Dice    Dice     `json:"dice"`
	Snakes  map[int]int      `json:"snakes" bson:"snakes" validate:"required"`
	Ladders map[int]int      `json:"ladders" bson:"ladders" validate:"required"`
	Players []map[string]int `json:"players" bson:"players" validate:"required"`
	Start   int              `json:"start" bson:"start" validate:"required,min=1"`
	End     int              `json:"end" bson:"end" validate:"required,max=100"`
}

type BoardHandler struct {
	Col dbiface.CollectionAPI
}

// ProductValidator a product validator
type BoardValidator struct {
	validator *validator.Validate
}

func insertBoard(ctx context.Context, board Board, collection dbiface.CollectionAPI) (Board, error) {

	board.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(ctx, board)
	if err != nil {
		return board, err
	}
	return board, nil
}
func findBoard(ctx context.Context, id string, collection dbiface.CollectionAPI) (Board, error) {
	var board Board

	// find if product exists, if err return 404
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("cannot convert to objectid: %v", err)
		return board, err
	}
	filter := bson.M{"_id": docID}

	res := collection.FindOne(ctx, filter)
	if err := res.Decode(&board); err != nil {

		log.Errorf("unable to  decode product: %v", err)
		return board, err
	}

	return board, nil
}

// validates product
func (b *BoardValidator) Validate(i interface{}) error {
	return b.validator.Struct(i)
}

// create tthe products in mongodb
func (b *BoardHandler) CreateBoard(c echo.Context) error {

	var board Board
	// c.Echo().Validator = &BoardValidator{validator: v}

	// fmt.Println("board is ", board)

	// err := c.Validate(board)
	// if err != nil {
	// 	log.Printf("unabale to validate product %+v %v", board, err)
	// 	return err
	// }

	if err := c.Bind(&board); err != nil {
		log.Fatalf("unable to bind: %v", err)
		return err
	}

	board, err := insertBoard(context.Background(), board, b.Col)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, board)
}

func (b *BoardHandler) GetBoard(c echo.Context) error {

	var board Board

	board, err := findBoard(context.Background(), c.Param("id"), b.Col)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, board)

}

// ROLL DICE Handler
func (b *BoardHandler) RollDice(c echo.Context) error {
	// get the board from db, 
	// rool the dice
	// update the board 
	return the board 

	return c.JSON(http.StatusOK, board)
}
