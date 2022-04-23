// Package mongo holds a mongoDB implementation of the ustomer Repository
package mongo

import (
	"context"
	"github.com/akolybelnikov/goddd/aggregate"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type CustomerMongoRepository struct {
	db       *mongo.Database
	customer *mongo.Collection
}

// mongoCustomer is an internal type that is used to store a CustomerAggregate
// it avoids coupling the mongo implementation to the customer aggregate
type mongoCustomer struct {
	ID   uuid.UUID `bson:"id"`
	Name string    `bson:"name"`
}

// NewFromCustomer takes in an aggregate and converts into internal structure
func NewFromCustomer(c *aggregate.Customer) mongoCustomer {
	return mongoCustomer{
		ID:   c.GetID(),
		Name: c.GetName(),
	}
}

// ToAggregate converts into an aggregate.Customer validating the values
func (c *mongoCustomer) ToAggregate() *aggregate.Customer {
	aggregateCustomer := aggregate.Customer{}
	aggregateCustomer.SetID(c.ID)
	aggregateCustomer.SetName(c.Name)

	return &aggregateCustomer
}

// New will create a new mongodb repository
func New(ctx context.Context, connectionString string) (*CustomerMongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	db := client.Database("goddd")
	customers := db.Collection("customers")

	return &CustomerMongoRepository{
		db:       db,
		customer: customers,
	}, nil
}

func (r *CustomerMongoRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result := r.customer.FindOne(ctx, bson.M{"id": id})

	var c mongoCustomer
	err := result.Decode(&c)
	if err != nil {
		return aggregate.Customer{}, err
	}

	return *c.ToAggregate(), nil
}

func (r *CustomerMongoRepository) Add(c aggregate.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	internal := NewFromCustomer(&c)
	_, err := r.customer.InsertOne(ctx, internal)
	if err != nil {
		return err
	}
	return nil
}

func (r *CustomerMongoRepository) Update(c aggregate.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"id": bson.M{
			"$eq": c.GetID(),
		},
	}
	update := bson.M{"$set": bson.M{"name": c.GetName()}}
	result := r.customer.FindOneAndUpdate(ctx, filter, update)
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}
