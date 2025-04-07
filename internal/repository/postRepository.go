package repository

import (
	"context"
	"errors"
	"time"

	"github.com/prppoomw/blog-api/internal/domain"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type postRepository struct {
	db         *mongo.Database
	collection string
}

func NewPostRepository(db *mongo.Database, collection string) domain.PostRepository {
	return &postRepository{
		db:         db,
		collection: collection,
	}
}

func (r *postRepository) FindBySlug(c context.Context, slug string) (*domain.Post, error) {
	collection := r.db.Collection(r.collection)
	var post domain.Post
	err := collection.FindOne(c, bson.M{"slug": slug}).Decode(&post)
	return &post, err
}

func (r *postRepository) Create(c context.Context, post *domain.Post) (*domain.Post, error) {
	userCollection := r.db.Collection(domain.CollectionUsers)
	var user domain.User
	e := userCollection.FindOne(c, bson.M{"clerkUserId": post.UserId}).Decode(&user)
	if e != nil {
		return nil, e
	}
	postCollection := r.db.Collection(r.collection)
	post.ID = bson.NewObjectID()
	now := time.Now()
	post.CreatedAt, post.UpdatedAt = now, now
	post.Username = user.Username
	post.UserImg = user.Img
	_, err := postCollection.InsertOne(c, post)
	return post, err
}

func (r *postRepository) Delete(c context.Context, id bson.ObjectID, userId string) error {
	collection := r.db.Collection(r.collection)
	filter := bson.M{"_id": id}
	filter["userId"] = userId
	result, err := collection.DeleteOne(c, filter)
	if result.DeletedCount == 0 {
		return errors.New("you can delete only your post")
	}
	return err
}

func (r *postRepository) FindByQuery(c context.Context, query *domain.PostListQueryRequest) (*domain.PostListResponse, error) {
	collection := r.db.Collection(r.collection)

	filter := bson.M{}
	if query.Category != "" {
		filter["category"] = query.Category
	}
	if query.Search != "" {
		filter["title"] = bson.M{"$regex": query.Search, "$options": "i"}
	}

	opts1 := options.Find().SetSkip(int64((query.Page - 1) * query.Limit))
	opts2 := options.Find().SetLimit(int64(query.Limit))
	opts := []options.Lister[options.FindOptions]{opts1, opts2}

	cursor, err := collection.Find(c, filter, opts...)

	var postListRes domain.PostListResponse
	if err != nil {
		return &postListRes, err
	}

	var posts []domain.Post

	if err := cursor.All(context.TODO(), &posts); err != nil {
		return &postListRes, err
	}

	totalPosts, err := collection.CountDocuments(c, bson.M{})
	if err != nil {
		return &postListRes, err
	}

	postListRes.Posts = posts
	postListRes.HasMore = query.Page*query.Limit < int(totalPosts)
	return &postListRes, nil
}
