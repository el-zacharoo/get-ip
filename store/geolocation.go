package store

import (
	"context"
	"fmt"
	"log"

	"github.com/el-zacharoo/get-ip/model"
)

func (s *Store) AddLocation(g model.Geolocation) {
	insertResult, err := s.locaColl.InsertOne(context.Background(), g)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInserted a Single Document: %v\n", insertResult)
}
