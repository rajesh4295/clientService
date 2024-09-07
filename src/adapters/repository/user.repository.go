package repository

import (
	"context"
	"fmt"
	"log"

	"etnasys.com/clientService/protogen/go/user"
	"etnasys.com/clientService/src/common"
	"etnasys.com/clientService/src/core/interfaces"
	"etnasys.com/clientService/src/core/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(client *mongo.Client, dbName string, collection string) interfaces.IUserRepository {
	mongoCollection := client.Database(dbName).Collection(collection)
	return &UserRepository{collection: mongoCollection}
}

func (repo *UserRepository) Create(createUser *user.CreateUserRequest) (*user.User, error) {
	userModel := common.CreateUserRequestToUsermodel(createUser)
	result, err := repo.collection.InsertOne(context.TODO(), userModel)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprint("Could not create new user"), err)
	}
	userModel.Id = result.InsertedID.(primitive.ObjectID).Hex()
	userRes := common.UserModelToUser(userModel)
	return userRes, nil
}

func (repo *UserRepository) Get(params *user.IdRequest) (*user.User, error) {
	oid, err := primitive.ObjectIDFromHex(params.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	result := repo.collection.FindOne(context.TODO(), bson.M{"_id": oid})
	userModel := &model.UserModel{}
	if err := result.Decode(&userModel); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("could not find user with id: %v", params.GetId()), err)
	}
	userRes := common.UserModelToUser(userModel)
	return userRes, nil
}

func (repo *UserRepository) Search(params *user.SearchUserRequest) (*user.PaginatedUserResponse, error) {
	log.Printf("params: %+v", params)           // Log the entire request
	log.Printf("query: %+v", params.GetQuery()) // Log only the query
	log.Printf("pagination: %+v", params.GetPagination())
	filter, err := common.ConvertProtoFieldsToBson(params.GetQuery())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Could not parse search filter: %v", err)
	}

	page := params.Pagination.GetPage()
	limit := params.Pagination.GetLimit()

	total, err := repo.collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not count documents: %v", err)
	}

	options := options.Find().
		SetSkip(int64((page - 1) * limit)).
		SetLimit(int64(limit))

	cursor, err := repo.collection.Find(context.TODO(), filter, options)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not find users: %v", err)
	}
	defer cursor.Close(context.TODO())

	var userModels []model.UserModel
	if err = cursor.All(context.TODO(), &userModels); err != nil {
		return nil, status.Errorf(codes.Internal, "Could not decode users: %v", err)
	}

	var users []*user.User
	for _, userModel := range userModels {
		users = append(users, common.UserModelToUser(&userModel))
	}

	return &user.PaginatedUserResponse{
		Result:      users,
		Pagination:  &user.Pagination{Page: page, Limit: limit},
		TotalResult: int32(total),
	}, nil
}

func (repo *UserRepository) Update(user *user.User) (*user.User, error) {
	oid, err := primitive.ObjectIDFromHex(user.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	userModel := common.CreateUserToUserModel(user)
	result := repo.collection.FindOneAndUpdate(context.TODO(), bson.M{"_id": oid}, bson.M{"$set": userModel}, options.FindOneAndUpdate().SetReturnDocument(1))
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprint("Could not update user"), err)
	}
	if err := result.Decode(&userModel); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("could not find user with id: %v", user.GetId()), err)
	}
	userRes := common.UserModelToUser(userModel)
	return userRes, nil
}

func (repo *UserRepository) Delete(req *user.IdRequest) (*user.DeleteUserResponse, error) {
	oid, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Could not convert to ObjectId: %v", err))
	}
	result, err := repo.collection.DeleteOne(context.TODO(), bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprint("Could not delete user"), err)
	}
	deleteRes := &user.DeleteUserResponse{
		DeletedCount: result.DeletedCount,
		Acknowledged: true,
	}
	return deleteRes, nil
}
