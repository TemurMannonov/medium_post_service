package service

import (
	"context"
	"time"

	pb "github.com/TemurMannonov/medium_post_service/genproto/post_service"
	"github.com/TemurMannonov/medium_post_service/storage/repo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/TemurMannonov/medium_post_service/storage"
)

type PostService struct {
	pb.UnimplementedPostServiceServer
	storage storage.StorageI
}

func NewPostService(strg storage.StorageI) *PostService {
	return &PostService{
		storage: strg,
	}
}

func (s *PostService) Create(ctx context.Context, req *pb.Post) (*pb.Post, error) {
	user, err := s.storage.Post().Create(&repo.Post{
		Title:       req.Title,
		Description: req.Description,
		ImageUrl:    req.ImageUrl,
		UserID:      req.UserId,
		CategoryID:  req.CategoryId,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}
	return parsePostModel(user), nil
}

func parsePostModel(post *repo.Post) *pb.Post {
	return &pb.Post{
		Id:          post.ID,
		Title:       post.Title,
		Description: post.Description,
		ImageUrl:    post.ImageUrl,
		UserId:      post.UserID,
		CategoryId:  post.CategoryID,
		CreatedAt:   post.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   post.UpdatedAt.Format(time.RFC3339),
		ViewsCount:  post.ViewsCount,
	}
}
