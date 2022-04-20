package store

import (
	"context"
	"fmt"
	"log"

	"github.com/el-zacharoo/get-ip/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Store) AddLocation(g model.Geolocation) {
	insertResult, err := s.locaColl.InsertOne(context.Background(), g)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInserted a Single Document: %v\n", insertResult)
}

func (s *Store) Getlocation(id string) (model.Geolocation, error) {

	var g model.Geolocation
	if err := s.locaColl.FindOne(
		context.Background(),
		bson.M{"id": id},
	).Decode(&g); err != nil {
		return model.Geolocation{}, err
	}

	return g, nil
}

func (s *Store) Getlocations(cn, searchText string, limit, skip *int64) (model.Page, error) {

	filter := bson.M{}

	if cn != "" {
		filter = bson.M{"$and": bson.A{filter,
			bson.M{"country": cn},
		}}
	}

	if searchText != "" {
		filter = bson.M{"$and": bson.A{filter,
			bson.M{"$text": bson.M{"$search": searchText}},
		}}
	}

	opt := options.FindOptions{
		Skip:  skip,
		Limit: limit,
		Sort:  bson.M{"date": -1},
	}

	mctx := context.Background()
	cursor, err := s.locaColl.Find(mctx, filter, &opt)
	if err != nil {
		return model.Page{}, err
	}

	// unpack results
	var pg model.Page
	if err := cursor.All(mctx, &pg.Data); err != nil {
		return model.Page{}, err
	}
	if pg.Matches, err = s.locaColl.CountDocuments(mctx, filter); err != nil {
		return model.Page{}, err
	}
	return pg, nil
}
