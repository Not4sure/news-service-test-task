package filter

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func All() bson.D {
	return bson.D{}
}

func ByID(id uuid.UUID) bson.D {
	return bson.D{
		bson.E{
			Key:   "_id",
			Value: id,
		},
	}
}
