package filter

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func ByID(id uuid.UUID) bson.D {
	return bson.D{
		bson.E{
			Key:   "_id",
			Value: id,
		},
	}
}
