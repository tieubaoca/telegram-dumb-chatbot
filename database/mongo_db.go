package database

import (
	"context"
	"time"

	"github.com/tieubaoca/telegram-dumb-chatbot/types"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var _ Database = &mongoDatabase{}

const MESSAGE_COLLECTION = "messages"

type mongoDatabase struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewMongoDatabase(connectionString string) (Database, error) {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return nil, err
	}
	db := client.Database("telegram-dumb-chatbot")
	// Ensure index on SentAt field
	collection := db.Collection(MESSAGE_COLLECTION)
	indexModel := mongo.IndexModel{
		Keys:    bson.M{"sent_at": 1}, // index in ascending order
		Options: options.Index().SetUnique(false),
	}
	_, err = collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		return nil, err
	}

	return &mongoDatabase{
		client: client,
		db:     db,
	}, nil
}

func (mdb *mongoDatabase) SaveMessage(chatId string, from string, message string) error {
	messageInstance := types.Message{
		ChatId:  chatId,
		From:    from,
		Message: message,
		SentAt:  time.Now().Unix(),
	}
	_, err := mdb.db.Collection(MESSAGE_COLLECTION).InsertOne(context.Background(), messageInstance)
	return err
}

func (mdb *mongoDatabase) PaginateMessages(chatId string, page int, limit int) ([]types.Message, error) {
	var messages []types.Message
	findOptions := options.Find()
	findOptions.SetSkip(int64((page - 1) * limit))
	findOptions.SetLimit(int64(limit))
	findOptions.SetSort(bson.M{"sent_at": -1}) // sort by SentAt in descending order

	cursor, err := mdb.db.Collection(MESSAGE_COLLECTION).Find(context.Background(), bson.M{"chat_id": chatId}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var message types.Message
		err := cursor.Decode(&message)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}
