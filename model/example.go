package model

// 下面是使用例子
import (
	"time"
)

// Example 储存service层的结构体
type Example struct {
	Title    string    `json:"title"`
	Text     string    `json:"text"`
	Author   string    `json:"author"`
	PostedOn time.Time `json:"postedOn"`
}
