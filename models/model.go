package models

import (
	"github.com/beego/beego/v2/adapter/logs"
	//"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
	"strings"
)

type User struct {
	Id      int
	Name    string
	Profile *Profile `orm:"rel(one)"`      // OneToOne relation
	Post    []*Post  `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct {
	Id   int
	Age  int16
	User *User `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"` //设置一对多关系
	Tags  []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"` //设置多对多反向关系
}

func init() {
	// 需要在init中注册定义的model
}

type user struct {
	Id     int
	Name   string `valid:"Required;Match(/^Bee.*/)"` // Name 不能为空并且以 Bee 开头
	Age    int    `valid:"Range(1, 140)"`            // 1 <= Age <= 140，超出此范围即为不合法
	Email  string `valid:"Email; MaxSize(100)"`      // Email 字段需要符合邮箱格式，并且最大长度不能大于 100 个字符
	Mobile string `valid:"Mobile" orm:"size(11)"`    // Mobile 必须为正确的手机号
	IP     string `valid:"IP"`                       // IP 必须为一个正确的 IPv4 地址
}

// 如果你的 struct 实现了接口 validation.ValidFormer
// 当 StructTag 中的测试都成功时，将会执行 Valid 函数进行自定义验证
func (u *user) Valid(v *validation.Validation) {
	if strings.Index(u.Name, "admin") != -1 {
		// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
		v.SetError("Name", "名称里不能含有 admin")
	}
}
func Validate() {
	valid := validation.Validation{}
	u := user{Name: "Beegoadmin", Age: 2, Email: "dev@web.me", Mobile: "18892701521", IP: "192.168.10.56"}
	b, err := valid.Valid(&u)
	if err != nil {
		// handle error
	}
	if !b {
		// validation does not pass
		// blabla...
		for _, err := range valid.Errors {
			logs.Error(err.Key, err.Message)
		}
	}
}
