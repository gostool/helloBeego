package main

import (
	"fmt"
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"helloBee/models"
)

func init() {
	// set default database
	username, err := web.AppConfig.String("mysqluser")
	if err != nil {
		panic(err)
	}
	passwd, err := web.AppConfig.String("mysqlpass")
	if err != nil {
		panic(err)
	}
	address, err := web.AppConfig.String("mysqlurls")
	if err != nil {
		panic(err)
	}
	dbname, err := web.AppConfig.String("mysqldb")

	orm.RegisterDataBase("default", "mysql",
		username+":"+passwd+"@"+"tcp("+address+")/"+dbname+"?charset=utf8mb4&loc=Asia%2FShanghai")
	orm.RegisterModel(
		new(models.User),
		new(models.Post),
		new(models.Profile),
		new(models.Tag),
	)
}

func main() {
	orm.Debug = true
	name := "default"
	// drop table 后再建表
	force := true
	// 打印执行过程
	verbose := true
	// 遇到错误立即返回
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		panic(err)
	}
	// 默认使用 default，你可以指定为其他数据库
	o := orm.NewOrm()

	posts := new(models.Post)
	posts.Title = "post1"
	posts.User = &models.User{Id: 1}

	tags := new(models.Tag)
	tags.Name = "tag1"

	profile := new(models.Profile)
	profile.Age = 30

	user := new(models.User)
	user.Profile = profile
	user.Name = "slene"

	fmt.Println(o.Insert(posts))
	fmt.Println(o.Insert(tags))
	fmt.Println(o.Insert(profile))
	fmt.Println(o.Insert(user))

	tag := models.Tag{Id: 1}
	// 1 获取多对多操对象
	m2m := o.QueryM2M(&tag, "Posts")
	// 2 获取插入对象
	title := "tagpost"
	post := models.Post{Id: 1, Title: title}
	_, err = m2m.Add(&post)
	if err != nil {
		logs.Info("插入多对多失败")
	}

	//// validate test
	//models.Validate()
	//
	//// we start admin service
	//// Prometheus will fetch metrics data from admin service's port
	//web.BConfig.Listen.EnableAdmin = true
	//web.BConfig.AppName = "my app"
	//web.Router("/get", &controllers.MainController{}, "get:Get")
	//web.Router("/post", &controllers.MainController{}, "post:Post")
	//web.Router("/put", &controllers.MainController{}, "put:Put")
	//web.Router("/delete", &controllers.MainController{}, "delete:Delete")
	//web.Router("/hello", &controllers.MainController{}, "get:Hello")
	//web.SetStaticPath("/static", "static")
	//fb := &prometheus.FilterChainBuilder{}
	//web.InsertFilterChain("/*", fb.FilterChain)
	//web.ErrorController(&controllers.ErrorController{})
	//web.Run()
	// after you start the server
	// and GET http://localhost:8080/hello
	// access http://localhost:8088/metrics
	// you can see something looks like:
	// http_request_web_sum{appname="my app",duration="1002",env="prod",method="GET",pattern="/hello",server="webServer:1.12.1",status="200"} 1002
	// http_request_web_count{appname="my app",duration="1002",env="prod",method="GET",pattern="/hello",server="webServer:1.12.1",status="200"} 1
	// http_request_web_sum{appname="my app",duration="1004",env="prod",method="GET",pattern="/hello",server="webServer:1.12.1",status="200"} 1004
	// http_request_web_count{appname="my app",duration="1004",env="prod",method="GET",pattern="/hello",server="webServer:1.12.1",status="200"} 1
}
