package model

import "go.mongodb.org/mongo-driver/v2/bson"

type Subscriber struct {
	ID     bson.ObjectID `bson:"_id,omitempty" json:"id"`
	IMSI   string        `bson:"imsi" json:"imsi"`
	MSISDN string        `bson:"msisdn" json:"msisdn"`
	Plan   string        `bson:"plan" json:"plan"`
	Status string        `bson:"status" json:"status"`
}
