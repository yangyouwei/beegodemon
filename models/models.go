package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

//字段必须是大写字母开头。
type Category struct {
	Id              int64     //默认这个字段是主键
	Title           string    //默认是255字节的string
	Created         time.Time `orm:"index"` //加上tag  设置索引。
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Countent        string `orm:"size(5000)"`
	Attachement     string
	Created         time.Time `orm:"index"`
	Updataed        time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplylastUserId int64
}


func RegisterDB()  {
	orm.RegisterModel(new(Category),new(Topic))
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql","root:Yang@1008@tcp(192.168.1.101:3306)/orm?charset=utf8",10)
}