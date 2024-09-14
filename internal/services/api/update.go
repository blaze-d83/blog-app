package api

import (
	"fmt"

	"github.com/blaze-d83/blog-app/types"
)

func (db *DBInstance) UpdatePost(postID uint, updatedPost types.Post) error {
	updatedPost.ID = postID
	result := db.DB.Model(&types.Post{}).Where("id = ?", postID).Updates(updatedPost)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no post found with %d", postID)
	}
	return nil
}

func (db *DBInstance) UpdateCategory(categoryID uint, updatedCategory types.Category) error {
	updatedCategory.ID = categoryID
	result := db.DB.Model(&types.Category{}).Where("id = ?", categoryID).Updates(updatedCategory)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no category found with %d", categoryID)
	}
	return nil
}
