package services

import (
	"fmt"

	"github.com/blaze-d83/blog-app/pkg/mysql"
	"github.com/blaze-d83/blog-app/pkg/types"
)

type PublicService interface {
	GetAllPosts() ([]types.UserPostListView, error)
	GetPostsByID(id uint) (types.Post, error)
}

type UserRepository struct {
	db *mysql.Database
}

func (r *UserRepository) NewUserService(db *mysql.Database) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) GetAllPosts() ([]types.UserPostListView, error) {
	var userPosts []types.UserPostListView
	result := repo.db.DB.Model(types.Post{}).
		Select("title", "date", "citation", "summary").
		Scan(userPosts)
	if result.Error != nil {
		return nil, result.Error
	}
	return userPosts, nil
}

func (repo *UserRepository) GetPostByID(id uint) (types.Post, error) {
	var post types.Post
	result := repo.db.DB.First(&post, id)
	if result.Error != nil {
		if result.RowsAffected == 0 {
			return types.Post{}, fmt.Errorf("post with ID %d not found", id)
		}
		return types.Post{}, result.Error
	}
	return post, nil
}
