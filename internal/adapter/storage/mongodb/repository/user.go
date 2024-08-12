package repository

import (
	"api/internal/core/domain"
	"api/internal/core/util"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
	"time"
)

type UserRepository struct {
	db *mongo.Client
}

func NewUserRepository(db *mongo.Client) *UserRepository {
	return &UserRepository{
		db,
	}
}

type UserDB struct {
	ID        string    `bson:"_id,omitempty"`
	Name      string    `bson:"name"`
	Email     string    `bson:"email"`
	Password  string    `bson:"password"`
	Role      string    `bson:"role"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

func (ur *UserRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	now := util.TimeNowInSaoPaulo()

	userToInsert := UserDB{
		Name:      user.Name,
		Email:     user.Email,
		Password:  hashedPassword,
		Role:      string(user.Role),
		CreatedAt: now,
		UpdatedAt: now,
	}

	collection := ur.db.Database("matchday").Collection("users")
	_, err = collection.InsertOne(ctx, userToInsert)

	if err != nil {
		slog.Error(err.Error())
		return nil, domain.ErrInternal
	}

	return user, nil
}

func (ur *UserRepository) ExistsUserByEmail(ctx context.Context, email string) bool {
	collection := ur.db.Database("matchday").Collection("users")

	var result UserDB
	err := collection.FindOne(ctx, bson.D{{"email", email}}).Decode(&result)

	return !errors.Is(err, mongo.ErrNoDocuments)
}
