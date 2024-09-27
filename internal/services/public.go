package services

import (
	"fmt"

	"github.com/blaze-d83/blog-app/internal/db"
	"github.com/blaze-d83/blog-app/types"
)

type PublicService interface {
	UsersGetAllPosts() ([]types.UserPostListView, error)
	UsersGetPostsByID(id uint) (types.Post, error)
}

type UserRepository struct {
	*db.Database
}

func (repo *UserRepository) UsersGetAllPosts() ([]types.UserPostListView, error) {
	var userPosts []types.UserPostListView
	result := repo.DB.Model(types.Post{}).
		Select("title", "date", "citation", "summary").
		Scan(userPosts)
	if result.Error != nil {
		return nil, result.Error
	}
	return userPosts, nil
}

func (repo *UserRepository) UsersGetPostByID(id uint) (types.Post, error) {
	var post types.Post
	result := repo.DB.First(&post, id)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return types.Post{}, fmt.Errorf("post with ID %d not found", id)
		}
		return types.Post{}, result.Error
	}
	return post, nil
}
