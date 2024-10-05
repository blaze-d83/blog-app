package services

import (
	"fmt"

	"github.com/blaze-d83/blog-app/pkg/mysql"
	"github.com/blaze-d83/blog-app/pkg/types"
	"gorm.io/gorm"
)

type AdminRepository struct {
	db *mysql.Database
}

type AdminService interface {
	CheckAdminExists(username string) (*types.Admin, error)
	GetAllPostsForAdmin() ([]types.AdminPostListView, error)
	GetPostByID(id uint) (types.Post, error)
	CreatePost(post types.Post) error
	UpdatePost(postID uint, updatedPost types.Post) error
	DeletePost(postID uint) error
	AdminGetAllCategories() ([]types.Category, error)
	CreateCategory(category types.Category) error
	DeleteCategory(id uint) error
}

func NewAdminService(db *mysql.Database) *AdminRepository {
	return &AdminRepository{db: db}
}

func (repo *AdminRepository) CheckAdminExists(username string) (*types.Admin, error) {
	var admin types.Admin
	err := repo.db.DB.Where("username = ?", username).Find(&admin).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("admin with username %s not found", username)
		}
		return nil, err
	}
	return &admin, nil
}

func (repo *AdminRepository) GetAllPostsForAdmin() ([]types.AdminPostListView, error) {
	var adminPosts []types.AdminPostListView
	result := repo.db.DB.Model(&types.Post{}).
		Select("title", "date", "created_at").
		Scan(&adminPosts)
	if result.Error != nil {
		return nil, result.Error
	}
	return adminPosts, nil
}

func (repo *AdminRepository) AdminGetPostByID(id uint) (types.Post, error) {
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

func (repo *AdminRepository) CreatePost(post types.Post) error {
	result := repo.db.DB.Create(post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *AdminRepository) UpdatePost(postID uint, updatedPost types.Post) error {
	updatedPost.ID = postID
	result := repo.db.DB.Model(&types.Post{}).Where("id = ?", postID).Updates(updatedPost)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no post found with ID %d", postID)
	}
	return nil
}

func (repo *AdminRepository) DeletePost(postID uint) error {
	result := repo.db.DB.Delete(types.Post{}, postID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no post found with ID %d", postID)
	}
	return nil
}

func (repo *AdminRepository) AdminGetAllCategories() ([]types.Category, error) {
	var categories []types.Category
	result := repo.db.DB.Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

func (repo *AdminRepository) CreateCategory(category types.Category) error {
	result := repo.db.DB.Create(&category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *AdminRepository) DeleteCategory(id uint) error {
	result := repo.db.DB.Delete(types.Category{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no post found with ID %d", id)
	}
	return nil
}
