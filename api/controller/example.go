package controller

// 下面是使用例子
import (
	"github.com/gin-gonic/gin"
	"github.com/penguinn/go-sdk/log"
	"github.com/penguinn/goframe/api/apimodel"
	"github.com/penguinn/goframe/dao"
	"github.com/penguinn/goframe/service"
)

// ExampleController 示例控制器
type ExampleController struct {
}

func (p ExampleController) Create(c *gin.Context) {
	req := apimodel.CreateExampleRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		renderInvalidRequest(c)
		log.Error(err)
		return
	}
	s := service.NewExampleService(&dao.Example{})
	if err := s.Create(req.Parse()); err != nil {
		renderInternalServerError(c)
		log.Error(err)
	} else {
		renderSuccess(c)
	}
}
