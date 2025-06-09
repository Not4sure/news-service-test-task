package sort

import "go.mongodb.org/mongo-driver/bson"

type Direction int

const (
	LowerFirst Direction = iota
	GreaterFirst
)

// ByCreatedAt returns sort option to sort Articles by created_at field
func ByCreatedAt(d Direction) bson.D {
	var v int8

	switch d {
	case LowerFirst:
		v = 1
	case GreaterFirst:
		v = -1
	}

	return bson.D{
		bson.E{
			Key:   "created_at",
			Value: v,
		},
	}
}
