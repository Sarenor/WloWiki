package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Item represents an item document in MongoDB
type Material struct {
	ItemID      int32  `bson:"item_id"`
	Rank        string `bson:"rank"`
	Name        string `bson:"name"`
	Size        string `bson:"size"`
	Type        string `bson:"type"`
	Base        string `bson:"base"`
	Attributes  string `bson:"attributes"`
	Description string `bson:"description"`
}

// MaterialsCollection wraps the MongoDB collection
type MaterialsCollection struct {
	collection *mongo.Collection
}

// NewMaterialsCollection creates a new MaterialsCollection instance
func NewMaterialsCollection(client *Client) *MaterialsCollection {
	collection := client.Database("wlodb").Collection("material")
	return &MaterialsCollection{collection}
}

// InsertMany inserts multiple Materials into the collection
func (ic *MaterialsCollection) InsertMany(Materials []Material) error {
	var documents []interface{}
	for _, item := range Materials {
		documents = append(documents, item)
	}
	_, err := ic.collection.InsertMany(context.TODO(), documents)
	return err
}

// InsertOne inserts a single item into the collection
func (ic *MaterialsCollection) InsertOne(item Material) (*mongo.InsertOneResult, error) {
	result, err := ic.collection.InsertOne(context.TODO(), item)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// FindAll retrieves all Materials from the collection with the given filter, sort options, skip, and limit
func (ic *MaterialsCollection) FindAll(filter bson.M, sortOptions bson.D, skip int64, limit int64) ([]bson.M, error) {
	findOptions := options.Find()
	findOptions.SetSort(sortOptions)
	findOptions.SetSkip(skip)
	findOptions.SetLimit(limit)

	cursor, err := ic.collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var Materials []bson.M
	if err := cursor.All(context.TODO(), &Materials); err != nil {
		return nil, err
	}
	return Materials, nil
}

// FindOne retrieves a single item from the collection based on the given filter
func (ic *MaterialsCollection) FindOne(filter bson.M) (Material, error) {
	var item Material
	err := ic.collection.FindOne(context.TODO(), filter).Decode(&item)
	if err != nil {
		return Material{}, err
	}
	return item, nil
}

// CountDocuments counts the number of documents in the collection that match the given filter
func (ic *MaterialsCollection) CountDocuments(filter bson.M) (int64, error) {
	count, err := ic.collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// UpdateOne updates a single item matching the filter
func (ic *MaterialsCollection) UpdateOne(filter bson.M, update bson.M) error {
	_, err := ic.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

// DeleteOne deletes a single item matching the filter
func (ic *MaterialsCollection) DeleteOne(filter bson.M) error {
	_, err := ic.collection.DeleteOne(context.TODO(), filter)
	return err
}
