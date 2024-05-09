package apimodel

// 下面是使用例子
import (
	"time"

	"github.com/penguinn/goframe/model"
)

type CreateExampleRequest struct {
	Title  string `json:"title" binding:"required"`
	Text   string `json:"text" binding:"required"`
	Author string `json:"author"`
}

func (req *CreateExampleRequest) Parse() *model.Example {
	example := &model.Example{
		Title:    req.Title,
		Text:     req.Text,
		Author:   req.Author,
		PostedOn: time.Now(),
	}
	return example
}
