package services

import (
	"github.com/blaze-d83/blog-app/internal/db"
	"github.com/blaze-d83/blog-app/types"
)

type PublicService interface {
	GetAllPostsForUsers() ([]types.UserPostListView, error)
}

type UserRepository struct {
	db.Database
}

func (ur *UserRepository) GetAllPostsForUsers() ([]types.UserPostListView, error) {
	var userPosts []types.UserPostListView
	result := ur.DB.Model(types.Post{}).
		Select("title", "date", "citation", "summary").
		Scan(userPosts)
	if result.Error != nil {
		return nil, result.Error
	}
	return userPosts, nil
}
