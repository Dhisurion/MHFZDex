package main

import (
	"context"
	"fmt"
	"log"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client
var ctx context.Context
var cancel context.CancelFunc
var err error

// This is a user defined method to close resources.
// This method closes mongoDB connection and cancel context.
func close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {

	// CancelFunc to cancel to context
	defer cancel()

	// client provides a method to close
	// a mongoDB connection.
	defer func() {

		// client.Disconnect method also has deadline.
		// returns error if any,
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func connect(uri string) (*mongo.Client, context.Context,
	context.CancelFunc, error) {

	// ctx will be used to set deadline for process, here
	// deadline will of 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	return client, ctx, cancel, err
}

// This is a user defined method that accepts
// mongo.Client and context.Context
// This method used to ping the mongoDB, return error if any.
func ping(client *mongo.Client, ctx context.Context) error {

	// mongo.Client has Ping to ping mongoDB, deadline of
	// the Ping method will be determined by cxt
	// Ping method return error if any occurred, then
	// the error can be handled.
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("connected successfully")
	return nil
}

/*func exInsertOne(client *mongo.Client, ctx context.Context) error {
	coll := client.Database("Frontier").Collection("Items")
	doc := bson.D{{"Name", "Potion"}, {"Rarity", 1}, {"Qty", 10}, {"Sell", 7}, {"Buy", 66}}
	result, err := coll.InsertOne(context.TODO(), doc)

	if err != nil {
		fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	}
	return nil
}*/

func InsertOneItem(client *mongo.Client, ctx context.Context) error {
	coll := client.Database("Frontier").Collection("Items")
	doc := bson.D{{"Name", tempitem.Name}, {"Icon", tempitem.EncodedIcon}, {"Rarity", tempitem.Rarity}, {"Qty", tempitem.Qty}, {"Sell", tempitem.Sell}, {"Buy", tempitem.Buy}}
	result, err := coll.InsertOne(context.TODO(), doc)

	if err != nil {
		fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	}
	return err
}

func ReadAllItems(client *mongo.Client, ctx context.Context) ([]ItemStruct, error) {
	coll := client.Database("Frontier").Collection("Items")

	var results []ItemStruct
	cursor, err := coll.Find(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var result ItemStruct
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return results, err
}

func Count(client *mongo.Client, ctx context.Context, t string) int64 {
	coll := client.Database("Frontier").Collection(t)
	count, err := coll.CountDocuments(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	return count
}

func UpdateOneItem(client *mongo.Client, ctx context.Context, Item ItemStruct) {
	coll := client.Database("Frontier").Collection("Items")

	result, err := coll.UpdateOne(ctx,
		bson.M{"_id": Item.ID},
		bson.D{
			{"$set", bson.D{{"Name", tempitem.Name}, {"Icon", tempitem.EncodedIcon}, {"Rarity", tempitem.Rarity}, {"Qty", tempitem.Qty}, {"Sell", tempitem.Sell}, {"Buy", tempitem.Buy}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
}

func DeleteOneItem(client *mongo.Client, ctx context.Context, Item ItemStruct) {
	coll := client.Database("Frontier").Collection("Items")

	result, err := coll.DeleteOne(ctx,
		bson.M{"_id": Item.ID})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted  %v Documents!\n", result.DeletedCount)
}

//Weapons
func ReadAllWeapons(client *mongo.Client, ctx context.Context) ([]WeaponStruct, error) {
	coll := client.Database("Frontier").Collection("Weapons")

	var results []WeaponStruct
	cursor, err := coll.Find(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var result WeaponStruct
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return results, err
}

//Monsters

func InsertOneMonster(client *mongo.Client, ctx context.Context, tempmonster TempMonsterStruct) error {
	coll := client.Database("Frontier").Collection("Monsters")
	doc := bson.D{{"Name", tempmonster.Name}, {"Icon", tempmonster.EncodedIcon}, {"Pic", tempmonster.EncodedPic},
		{"FireWeakness", tempmonster.Fire}, {"ThunderWeakness", tempmonster.Thunder}, {"WaterWeakness", tempmonster.Water},
		{"IceWeakness", tempmonster.Ice}, {"DragonWeakness", tempmonster.Dragon},
		{"MaterialsLowRank", tempmonster.LRMat}, {"MaterialsHighRank", tempmonster.HRMat}, {"MaterialsGouRank", tempmonster.GouRMat},
		{"MaterialsGRank", tempmonster.GRMat}, {"MaterialsZenithRank", tempmonster.ZRMat}}
	result, err := coll.InsertOne(context.TODO(), doc)

	if err != nil {
		fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	}
	return err
}

func ReadAllMonsters(client *mongo.Client, ctx context.Context) ([]MonsterStruct, error) {
	coll := client.Database("Frontier").Collection("Monsters")

	var results []MonsterStruct
	cursor, err := coll.Find(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var result MonsterStruct
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return results, err
}

func UpdateOneMonster(client *mongo.Client, ctx context.Context, Monster MonsterStruct, tempmonster TempMonsterStruct) {
	coll := client.Database("Frontier").Collection("Monsters")

	result, err := coll.UpdateOne(ctx,
		bson.M{"_id": Monster.ID},
		bson.D{
			{"$set", bson.D{{"Name", tempmonster.Name}, {"Icon", tempmonster.EncodedIcon}, {"Pic", tempmonster.EncodedPic},
				{"FireWeakness", tempmonster.Fire}, {"ThunderWeakness", tempmonster.Thunder}, {"WaterWeakness", tempmonster.Water},
				{"IceWeakness", tempmonster.Ice}, {"DragonWeakness", tempmonster.Dragon},
				{"MaterialsLowRank", tempmonster.LRMat}, {"MaterialsHighRank", tempmonster.HRMat}, {"MaterialsGouRank", tempmonster.GouRMat},
				{"MaterialsGRank", tempmonster.GRMat}, {"MaterialsZenithRank", tempmonster.ZRMat}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
}

func DeleteOneMonster(client *mongo.Client, ctx context.Context, Monster MonsterStruct) {
	coll := client.Database("Frontier").Collection("Monsters")

	result, err := coll.DeleteOne(ctx,
		bson.M{"_id": Monster.ID})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted  %v Documents!\n", result.DeletedCount)
}
