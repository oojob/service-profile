package db

import (
	"context"

	"github.com/oojob/service-profile/src/model"
)

// CreateProfile create profile entity
func (db *Database) CreateProfile(in *model.Profile) (string, error) {
	var inerstionID string
	oojob := db.Database("test")
	// client := oojob.Client()
	companyCollection := oojob.Collection("profile")

	_, err := companyCollection.InsertOne(context.Background(), in)

	// start the session for transaction
	// session, err := client.StartSession()
	// if err != nil {
	// 	return "", err
	// }
	// defer session.EndSession(context.Background())

	// _, err = session.WithTransaction(context.Background(), func(sessionContext mongo.SessionContext) (interface{}, error) {
	// 	result, err := companyCollection.InsertOne(sessionContext, in)
	// 	if err != nil {
	// 		return "", err
	// 	}

	// 	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
	// 		inerstionID = oid.Hex()
	// 	}

	// 	return "", nil
	// })

	return inerstionID, err
}
