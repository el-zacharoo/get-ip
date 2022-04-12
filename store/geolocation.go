package store

import (
	"context"
	"fmt"
	"log"

	"github.com/ip-address/model"
)

func (s *Store) Add(g model.Geolocation) {
	insertResult, err := s.locaColl.InsertOne(context.Background(), g)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInserted a Single Document: %v\n", insertResult)
}
