package services

import (
	"context"
	"fmt"
	"gotest/connections"
	"gotest/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(Products models.Products) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	fmt.Println(cancel)
	defer cancel()

	result, err := connections.DB.Collection("Products").InsertOne(ctx, Products)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetAllData() ([]models.Products, error) {
	var products []models.Products
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := connections.DB.Collection("Products").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	fmt.Println(cursor)
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var product models.Products
		fmt.Println(cursor.Decode(&product))
		cursor.Decode(&product)
		products = append(products, product)
	}
	return products, nil

}

func UpdateProduct(id string, Products models.Products) (*mongo.UpdateResult, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// fmt.Println(ctx)

	update := bson.M{}
	if Products.Name != "" {
		update["name"] = Products.Name
	}
	if Products.Description != "" {
		update["description"] = Products.Description
	}
	if Products.Amount != "" {
		update["amount"] = Products.Amount
	}

	result, err := connections.DB.Collection("Products").UpdateOne(ctx, bson.M{"_id": objID}, bson.M{"$set": update})

	if err != nil {
		return nil, err
	}

	return result, nil

}

func DeleteProductById(id string) (*mongo.DeleteResult, error) {

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := connections.DB.Collection("Products").DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return nil, err
	}
	return result, nil

}
