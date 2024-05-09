## 前言
项目为AI和大数据交付团队的Go项目初始化规范。

## 目录结构
```    
├── Dockerfile
├── Makefile
├── api #用于处理HTTP请求，并调用 Service 进行业务处理
│  ├── apimodel #http request&response数据结构，以及与model的相互转换
│  │  └── example.go
│  ├── controller #http handler
│  │  ├── example.go
│  │  └── render.go
│  └── middleware #中间件
│      └── cors.go
├── ci.yml
├── cmd #非HTTP server的启动入口，比如k8s init、cronjob等启动入口可以放在这里
│  └── init #init启动入口文件夹
│      └── main.go
├── component #依赖的第三方微服务调用，统计放在comment管理
│  ├── cron
│  │  └── cron.go
│  └── k8s
│      └── init.go
├── config #主程序的配置类
│  └── config.go
├── config.yaml #配置文件
├── constant #常量存放目录
│  └── dao.go
├── dao #封装负责与底层的数据资源，使用gorm格式保存表
│  ├── example.go
│  └── init.go
├── go.mod
├── go.sum
├── main.go
├── model #各个对象的字段，定义了实体数据struct和数据表结构的映射
│  └── example.go
└── service #业务组装逻辑层，封装对多个数据资源（比如数据表、第三方微服务等）的操作。如果是定时任务微服务，初始化scanner
    ├── example.go
    └── scanner
        └── init.go
```
## 分层
本demo实现了一个基本的MVC风格的业务逻辑框架，适合大部分的RESTful APIs中小型项目，逻辑分层分为：
 - API层：负责处理与http handler逻辑,请求参数以及response格式相关的处理工作
 - service层：处理业务逻辑
 - 基础依赖层：处理数据访问逻辑，一般包括数据库（也就是DAO层）、缓存层（一般也就是Redis）、第三方微服务等
除此之外，model层负责实体定义相关的逻辑，贯穿api，Service，Dao/Component这三层

## 测试
本demo包含了完整的单元测试体系，覆盖了基础依赖层、service、API层（又可以称为集成测试）。
### 基础依赖层测试

1. DAO数据库依赖层测试
DAO层测试一般有三种方式：
a. 建立真实的数据库环境，虽然传统建立数据库麻烦，但好在有docker
b. 使用内存数据库sqlite3,在go里面采用Sqlite的驱动实例化一个sql.DB出来，可以兼容大部分的mysql语法
c. database/sql是一个接口层，有一个包go-sqlmock，需要自己模拟输入输出
本demo采用*第二种方案*，基于内存数据库sqlite3和Gorm AutoMigrate函数模拟业务实际采用的数据库环境
```
func init() {
	var err error
	// 使用 file::memory:?cache=shared 替代文件路径。 这会告诉 SQLite 在系统内存中使用一个临时数据库
	config.Config.DB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	config.Config.DB.AutoMigrate(&model.User{}, &model.Post{})
}
```

2. 缓存层、第三方微服务(component)等基础依赖
在应用中缓存层是少不了的。现在的缓存大部分都是Redis，类似一个内存数据库，因此测试方法和上述一样，推荐在Docker中真实的建立一个Redis服务，如果不能建立，可以考虑使用 miniredis项目，和上述DAO中的第二种方法类似，模拟了一个小的Redis，兼容大部分常见命令。
在微服务中，一个服务调用多个服务是很常见的情况。而这些服务我们又不可能在本地建立一个真实的环境。因此只能想办法去Mock掉。根据服务提供的API抽象出一个接口文件，然后使用适配器模式或代理模式进行Wrap一层。

### service层测试
本demo中基础依赖层和Service已经解藕，由于基础依赖层涉及的变动比较少，所以接口化实现，依赖注入方式初始化service。本demo采用对基础依赖层进行Mock的方式实现对service层的测试，常用的Mock工具是 go-mock。这里采用手动Mock的方式：
```go
func newMockUserDAO() dao.UserDAO {
	return &mockUserDAO{
		records: []model.User{
			{Model: model.Model{ID: 1}, FirstName: "John", LastName: "Smith", Email: "john.smith@gmail.com", Address: "Dummy Value"},
			{Model: model.Model{ID: 2}, FirstName: "Ben", LastName: "Doe", Email: "ben.doe@gmail.com", Address: "Dummy Value"},
		},
	}
}

// Mock Get function that replaces real User DAO
func (m *mockUserDAO) Get(id uint) (*model.User, error) {
	for _, record := range m.records {
		if record.ID == id {
			return &record, nil
		}
	}
	return nil, errors.New("not found")
}

type mockUserDAO struct {
	records []model.User
}
```
### API层测试

API层测试，也称为集成测试，对接口层面进行测试。基于go+gin提供的http测试工具,在`api_test.go`中建立了API测试的脚手架函数，使用方法如下：
```go
func TestUser(t *testing.T) {
	path := test_data.GetTestCaseFolder()
	runAPITests(t, []apiTestCase{
		{"t1 - get a User", "GET", "/users/:id", "/users/1", "", GetUser, http.StatusOK, path + "/user_t1.json"},
		{"t2 - get a User not Present", "GET", "/users/:id", "/users/9999", "", GetUser, http.StatusOK, path + "/user_t2.json"},
	})
}
```
## 参考

* [Building RESTful APIs in Golang](https://towardsdatascience.com/building-restful-apis-in-golang-e3fe6e3f8f95)
* [https://github.com/MartinHeinz/go-project-blueprint/tree/rest-api](https://github.com/MartinHeinz/go-project-blueprint/tree/rest-api)
* https://stoneqi.github.io/2021/02/22/Go-Test-Used/[https://stoneqi.github.io/2021/02/22/Go-Test-Used/](https://stoneqi.github.io/2021/02/22/Go-Test-Used/)
