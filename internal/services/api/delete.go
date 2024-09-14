package api

import (
	"fmt"

	"github.com/blaze-d83/blog-app/types"
)

func (db *DBInstance) DeletePost(postID uint) error {
	result := db.DB.Delete(&types.Post{}, postID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no post found with %d", postID)
	}
	return nil
}


func (db *DBInstance) DeleteCategory(categoryID uint) error {
	result := db.DB.Delete(&types.Category{}, categoryID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no category found with %d", categoryID)
	}
	return nil
}
