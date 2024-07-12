package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// AlchemyItem represents an alchemy item document in MongoDB
type AlchemyItem struct {
	Level          int    `json:"level"`
	RankDifference string `json:"rank_difference"`
	Result         int    `json:"result"`
	Ingredients    []int  `json:"ingredients"`
}

// AlchemyCollection wraps the MongoDB collection
type AlchemyCollection struct {
	collection *mongo.Collection
}

// NewAlchemyCollection creates a new AlchemyCollection instance
func NewAlchemyCollection(client *Client) *AlchemyCollection {
	collection := client.Database("wlodb").Collection("alchemy")
	return &AlchemyCollection{collection}
}

// InsertMany inserts multiple items into the collection
func (ac *AlchemyCollection) InsertMany(items []AlchemyItem) error {
	var documents []interface{}
	for _, item := range items {
		documents = append(documents, item)
	}
	_, err := ac.collection.InsertMany(context.TODO(), documents)
	return err
}

func (ac *AlchemyCollection) Find(filter bson.M, opts ...*options.FindOptions) ([]AlchemyItem, error) {
	// Define the default sort option (sorting by Level in ascending order)
	defaultSort := options.Find().SetSort(bson.D{{Key: "level", Value: 1}})

	// If no options are provided, use the default sort option
	if len(opts) == 0 {
		opts = append(opts, defaultSort)
	}

	// Create a cursor for the query
	cursor, err := ac.collection.Find(context.TODO(), filter, opts...)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	// Retrieve all documents from the cursor
	var items []AlchemyItem
	if err := cursor.All(context.TODO(), &items); err != nil {
		return nil, err
	}

	return items, nil
}

// InsertOne inserts a single item into the collection
func (ac *AlchemyCollection) InsertOne(item AlchemyItem) (*mongo.InsertOneResult, error) {
	result, err := ac.collection.InsertOne(context.TODO(), item)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindAll retrieves all items from the collection without filters or pagination
func (ac *AlchemyCollection) FindAll() ([]AlchemyItem, error) {
	// Create a filter to match all documents
	filter := bson.M{}

	// Define sort options (sorting by Level in ascending order)
	sortOptions := options.Find().SetSort(bson.D{{Key: "level", Value: 1}})

	cursor, err := ac.collection.Find(context.TODO(), filter, sortOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var items []AlchemyItem
	if err := cursor.All(context.TODO(), &items); err != nil {
		return nil, err
	}
	return items, nil
}

// FindOne retrieves a single item from the collection based on the given filter
func (ac *AlchemyCollection) FindOne(filter bson.M) (AlchemyItem, error) {
	var item AlchemyItem
	err := ac.collection.FindOne(context.TODO(), filter).Decode(&item)
	if err != nil {
		return item, err
	}
	return item, nil
}

// CountDocuments counts the number of documents in the collection that match the given filter
func (ac *AlchemyCollection) CountDocuments(filter bson.M) (int64, error) {
	count, err := ac.collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// UpdateOne updates a single item matching the filter
func (ac *AlchemyCollection) UpdateOne(filter bson.M, update bson.M) error {
	_, err := ac.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

// DeleteOne deletes a single item matching the filter
func (ac *AlchemyCollection) DeleteOne(filter bson.M) error {
	_, err := ac.collection.DeleteOne(context.TODO(), filter)
	return err
}
