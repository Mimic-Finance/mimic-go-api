package controllers

import (
	"context"
	"mimic-api/configs"
	"mimic-api/models"
	"mimic-api/responses"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var poolCollection *mongo.Collection = configs.GetCollection(configs.DB, "pools")
var validate = validator.New()

func CreatePool(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var pool models.Pool
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&pool); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.PoolResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&pool); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.PoolResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newPool := models.Pool{
		Id:          primitive.NewObjectID(),
		Address:     pool.Address,
		Symbol:      pool.Symbol,
		Description: pool.Description,
		Type:        pool.Type,
		Token:       pool.Token,
		Apr:         pool.Apr,
		Label:       pool.Label,
		Color:       pool.Color,
		Info:        pool.Info,
	}

	result, err := poolCollection.InsertOne(ctx, newPool)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.PoolResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.PoolResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func GetAPool(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	poolAddress := c.Params("address")
	var pool models.Pool
	defer cancel()

	err := poolCollection.FindOne(ctx, bson.M{"address": poolAddress}).Decode(&pool)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.PoolResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.PoolResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": pool}})
}

func GetAllPools(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var pools []models.Pool
	defer cancel()

	results, err := poolCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.PoolResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singlePool models.Pool
		if err = results.Decode(&singlePool); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.PoolResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		pools = append(pools, singlePool)
	}

	return c.Status(http.StatusOK).JSON(
		responses.PoolResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": pools}},
	)
}
