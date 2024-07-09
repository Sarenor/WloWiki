package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Item represents an item document in MongoDB
type Item struct {
	ItemID      int32  `bson:"item_id"`
	Rank        string `bson:"rank"`
	Name        string `bson:"name"`
	Size        string `bson:"size"`
	Type        string `bson:"type"`
	Base        string `bson:"base"`
	Attributes  string `bson:"attributes"`
	Description string `bson:"description"`
}

// ItemsCollection wraps the MongoDB collection
type ItemsCollection struct {
	collection *mongo.Collection
}

// NewItemsCollection creates a new ItemsCollection instance
func NewItemsCollection(client *Client) *ItemsCollection {
	collection := client.Database("wlodb").Collection("equipment")
	return &ItemsCollection{collection}
}

// InsertMany inserts multiple items into the collection
func (ic *ItemsCollection) InsertMany(items []Item) error {
	var documents []interface{}
	for _, item := range items {
		documents = append(documents, item)
	}
	_, err := ic.collection.InsertMany(context.TODO(), documents)
	return err
}

// InsertOne inserts a single item into the collection
func (ic *ItemsCollection) InsertOne(item Item) (*mongo.InsertOneResult, error) {
	result, err := ic.collection.InsertOne(context.TODO(), item)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindAll fetches all items from the collection
func (ic *ItemsCollection) FindAll() ([]Item, error) {
	cursor, err := ic.collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	var results []Item
	if err := cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

// UpdateOne updates a single item matching the filter
func (ic *ItemsCollection) UpdateOne(filter bson.M, update bson.M) error {
	_, err := ic.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

// DeleteOne deletes a single item matching the filter
func (ic *ItemsCollection) DeleteOne(filter bson.M) error {
	_, err := ic.collection.DeleteOne(context.TODO(), filter)
	return err
}
