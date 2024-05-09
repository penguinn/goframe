package service

// 下面是使用例子
import (
	"github.com/penguinn/goframe/dao"
	"github.com/penguinn/goframe/model"
)

type ExampleService struct {
	dao *dao.Example
}

// NewExampleService creates a new PostService with the given post DAO.
func NewExampleService(dao *dao.Example) *ExampleService {
	return &ExampleService{dao: dao}
}

func (s *ExampleService) Create(r *model.Example) error {
	example := &dao.Example{
		Title:    r.Title,
		Text:     r.Text,
		Author:   r.Author,
		PostedOn: r.PostedOn,
	}
	return s.dao.Create(example)
}
