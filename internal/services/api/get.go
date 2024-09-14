package api

import (
	"github.com/blaze-d83/blog-app/types"
	"gorm.io/gorm"
)

type DBInstance struct {
	DB *gorm.DB
}

func (db *DBInstance) GetAllPosts() ([]types.Post, error) {
	var posts []types.Post
	result := db.DB.Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

func (db *DBInstance) GetPostByID(id uint) (types.Post, error) {
	var post types.Post
	result := db.DB.Where("id = ?", id).Find(&post)
	if result.Error != nil {
		return types.Post{}, result.Error
	}
	return post, nil
}

func (db *DBInstance) GetPostByCategory(categoryID uint) ([]types.Post, error) {
	var posts []types.Post
	result := db.DB.Joins("JOIN post_categories ON post_categories.post_id = post.id").
		Where("post_categories.category_id = ?", categoryID).Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

func (db *DBInstance) GetAllCategories() ([]types.Category, error) {
	var categories []types.Category
	result := db.DB.Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}
