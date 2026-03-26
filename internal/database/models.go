package database

import (
	"time"
)


type Items struct {
	Id	                        uint             `gorm:"primaryKey;autoIncrement"`
	Title                       string           `gorm:"type:varchar(255);not null"`
	Content                     string           `gorm:"type:text;not null"`
	CreatedAt                   time.Time        `gorm:"autoCreateTime"`
	UpdatedAt                   time.Time        `gorm:"autoUpdateTime"`
}

func (Items) TableName() string {
	return "items"
}

func GetAllItems(search string) ([]Items, error) {
	var items []Items
	result := DB.Where("title LIKE ? OR content LIKE ?", "%"+search+"%", "%"+search+"%").Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

func CreateItem(title, content string) (*Items, error) {
	item := &Items{
		Title:   title,
		Content: content,
	}
	result := DB.Create(item)
	if result.Error != nil {
		return nil, result.Error
	}
	return item, nil
}

func UpdateItem(id uint, title, content string) (*Items, error) {
	var item Items
	result := DB.First(&item, id)
	if result.Error != nil {
		return nil, result.Error
	}
	if title != "" {
		item.Title = title
	}
	if content != "" {
	item.Content = content
	}
	result = DB.Save(&item)
	if result.Error != nil {
		return nil, result.Error
	}
	return &item, nil
}

func DeleteItem(id uint) error {
	result := DB.Delete(&Items{}, id)
	return result.Error
}