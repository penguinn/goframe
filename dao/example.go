package dao

// 下面是使用例子
import (
	"time"

	"gorm.io/gorm"
)

type Example struct {
	gorm.Model
	Title    string    `gorm:"column:title"`
	Text     string    `gorm:"column:text"`
	Author   string    `gorm:"column:author"`
	PostedOn time.Time `gorm:"column:posted_on"`
}

func (dao *Example) TableName() string {
	return "example"
}

func (dao *Example) Create(example *Example) error {
	err := GetDefault().Create(example).Error
	return err
}
