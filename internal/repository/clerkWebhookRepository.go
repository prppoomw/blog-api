package repository

import (
	"context"

	"github.com/prppoomw/blog-api/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type clerkWebhookRepository struct {
	userCollection *mongo.Collection
	postCollection *mongo.Collection
}

func NewClerkWebhookRepository(userCollection, postCollection *mongo.Collection) domain.ClerkWebhookRepository {
	return &clerkWebhookRepository{
		userCollection: userCollection,
		postCollection: postCollection,
	}
}

func (r *clerkWebhookRepository) CreateUser(c context.Context, user *domain.User) error {
	_, err := r.userCollection.InsertOne(c, user)
	return err
}

func (r *clerkWebhookRepository) DeleteUser(c context.Context, clerkUserId string) error {
	filer := bson.M{"clerkUserId": clerkUserId}
	_, err := r.userCollection.DeleteOne(c, filer)
	return err
}

func (r *clerkWebhookRepository) DeletePostsByUser(c context.Context, userId string) error {
	filter := bson.M{"userId": userId}
	_, err := r.postCollection.DeleteMany(c, filter)
	return err
}
