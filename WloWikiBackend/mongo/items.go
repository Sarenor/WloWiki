package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// FindAll retrieves all items from the collection with the given filter, sort options, skip, and limit
func (ic *ItemsCollection) FindAll(filter bson.M, sortOptions bson.D, skip int64, limit int64) ([]bson.M, error) {
	findOptions := options.Find()
	findOptions.SetSort(sortOptions)
	findOptions.SetSkip(skip)
	findOptions.SetLimit(limit)

	cursor, err := ic.collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var items []bson.M
	if err := cursor.All(context.TODO(), &items); err != nil {
		return nil, err
	}
	return items, nil
}

// FindOne retrieves a single item from the collection based on the given filter
func (ic *ItemsCollection) FindOne(filter bson.M) (bson.M, error) {
	var item bson.M
	err := ic.collection.FindOne(context.TODO(), filter).Decode(&item)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// CountDocuments counts the number of documents in the collection that match the given filter
func (ic *ItemsCollection) CountDocuments(filter bson.M) (int64, error) {
	count, err := ic.collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0, err
	}
	return count, nil
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
