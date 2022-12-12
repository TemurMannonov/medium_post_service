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

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	storage storage.StorageI
}

func NewCategoryService(strg storage.StorageI) *CategoryService {
	return &CategoryService{
		storage: strg,
	}
}

func (s *CategoryService) Create(ctx context.Context, req *pb.Category) (*pb.Category, error) {
	user, err := s.storage.Category().Create(&repo.Category{
		Title: req.Title,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server error: %v", err)
	}
	return parseCategoryModel(user), nil
}

func parseCategoryModel(c *repo.Category) *pb.Category {
	return &pb.Category{
		Id:        c.ID,
		Title:     c.Title,
		CreatedAt: c.CreatedAt.Format(time.RFC3339),
	}
}
