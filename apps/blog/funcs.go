package blog

import (
	"gorm.io/gorm"
)

func CreateBlog(db *gorm.DB, blog *Blog) (Blog, error) {
	result := db.Create(&blog)
	if result.Error != nil {
		return Blog{}, result.Error
	}
	return *blog, nil
}

func GetAllBlogs(db *gorm.DB) ([]Blog, error) {
	var blogs []Blog

	result := db.Find(&blogs)
	if result.Error != nil {
		return blogs, result.Error
	}
	return blogs, nil
}

func GetBlogById(db *gorm.DB, id int) (Blog, error) {
	var blog Blog

	result := db.First(&blog, id)
	if result.Error != nil {
		return Blog{}, result.Error
	}
	return blog, nil
}

func DeleteBlog(db *gorm.DB, id int) (bool, error) {
	var blog Blog

	result := db.First(&blog, id).Delete(&Blog{})

	if result.Error != nil {
		return false, result.Error
	}
	if result.RowsAffected == 0 {
		return false, nil
	}
	return true, nil

}

func UpdateBlog(db *gorm.DB, id int, updates map[string]interface{}) (Blog, error) {
	var blog Blog
	if err := db.First(&blog, id).Error; err != nil {
		return Blog{}, err
	}

	if err := db.Model(&blog).Updates(updates).Error; err != nil {
		return Blog{}, err
	}

	return blog, nil
}
