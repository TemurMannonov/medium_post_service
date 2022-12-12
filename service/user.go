package service

import (
	"context"

	pb "github.com/TemurMannonov/medium_post_service/genproto/user_service"

	"github.com/TemurMannonov/medium_post_service/storage"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	storage storage.StorageI
}

func NewUserService(strg storage.StorageI) *UserService {
	return &UserService{
		storage: strg,
	}
}

func (s *UserService) Create(ctx context.Context, req *pb.User) (*pb.User, error) {
	// user, err := s.storage.User().Create(&repo.User{
	// 	FirstName:       req.FirstName,
	// 	LastName:        req.LastName,
	// 	PhoneNumber:     req.PhoneNumber,
	// 	Email:           req.Email,
	// 	Gender:          req.Gender,
	// 	Password:        req.Password,
	// 	Username:        req.Username,
	// 	ProfileImageUrl: req.ProfileImageUrl,
	// 	Type:            req.Type,
	// })
	// if err != nil {
	// 	return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	// }
	return nil, nil
	// return parseUserModel(user), nil
}

// func parseUserModel(user *repo.User) *pb.User {
// 	return &pb.User{
// 		Id:              user.ID,
// 		FirstName:       user.FirstName,
// 		LastName:        user.LastName,
// 		PhoneNumber:     user.PhoneNumber,
// 		Email:           user.Email,
// 		Gender:          user.Gender,
// 		Password:        user.Password,
// 		Username:        user.Username,
// 		ProfileImageUrl: user.ProfileImageUrl,
// 		Type:            user.Type,
// 		CreatedAt:       user.CreatedAt.Format(time.RFC3339),
// 	}
// }
