package main

import (
	"context"
	"database/sql"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type League struct {
	bun.BaseModel `bun:"table:leagues,alias:l" bson:"-"`

	ID        string    `bun:",pk" bson:"_id"`
	Name      string    `bun:"name" bson:"name"`
	Type      string    `bun:"type" bson:"type"`
	Logo      string    `bun:"logo" bson:"logo"`
	CreatedAt time.Time `bun:"created_at" bson:"created_at"`
	UpdatedAt time.Time `bun:"updated_at" bson:"updated_at"`
}

func main() {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN("postgres://postgres:postgres@localhost:5432/footballapp?sslmode=disable")))
	db := bun.NewDB(sqldb, pgdialect.New())
	var leagues = make([]*League, 0)
	err := db.NewRaw("SELECT * FROM leagues").Scan(context.Background(), &leagues)
	if err != nil {
		panic(err)
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://admin:admin@localhost:27017/?authSource=admin"))
	if err != nil {
		panic(err)
	}
	collection := client.Database("test").Collection("leagues")
	for _, league := range leagues {
		_, err := collection.InsertOne(context.Background(), league)
		if err != nil {
			panic(err)
		}
	}
}
