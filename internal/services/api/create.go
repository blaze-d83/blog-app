package api

import "github.com/blaze-d83/blog-app/types"

func (db *DBInstance) CreatePost(post types.Post) error {
	result := db.DB.Create(&post)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *DBInstance) CreateCategory(category types.Category) error {
	result := db.DB.Create(&category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
