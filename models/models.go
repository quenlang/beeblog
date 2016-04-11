package models

import (
	//"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path"
	"strconv"
	"time"
)

const (
	_DB_NAME        = "data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int32
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDB() {
	// 判断数据目录是否存在，若不存在则创建
	_, err := os.Stat(_DB_NAME)
	if !os.IsExist(err) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	// 注册模型
	orm.RegisterModel(new(Category), new(Topic))
	// 注册驱动
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	// 注册默认数据库
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

func AddCategory(name string) error {
	o := orm.NewOrm()
	category := &Category{
		Title:     name,
		Created:   time.Now(),
		TopicTime: time.Now(),
	}
	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(category)
	//fmt.Printf("One():%v\n", err)
	if err == nil {
		return err
	}
	_, err = o.Insert(category)
	if err != nil {
		return err
	}
	return nil
}

func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	o := orm.NewOrm()
	qs := o.QueryTable("category")
	_, err = qs.Filter("id", cid).Delete()
	if err != nil {
		return err
	}

	return nil
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()
	categories := make([]*Category, 0)
	_, err := o.QueryTable("category").All(&categories)
	if err != nil {
		return nil, err
	}
	return categories, err
}

func AddTopic(title, content, attachment string) error {
	o := orm.NewOrm()
	topic := &Topic{
		Title:      title,
		Content:    content,
		Attachment: attachment,
		Created:    time.Now(),
		Updated:    time.Now(),
		ReplyTime:  time.Now(),
	}
	err := o.QueryTable("topic").Filter("title", title).One(topic)
	if err == nil {
		return nil
	}
	_, err = o.Insert(topic)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTopic(tid, title, content, attachment string) error {
	o := orm.NewOrm()
	id, err := strconv.ParseInt(tid, 10, 64)
	topic := &Topic{
		Id: id,
	}

	err = o.Read(topic)
	if err == nil {
		topic.Title = title
		topic.Content = content
		topic.Attachment = attachment
		topic.Updated = time.Now()
		o.Update(topic)
	}
	return nil
}

func DelTopic(tid string) error {
	id, err := strconv.ParseInt(tid, 10, 64)
	o := orm.NewOrm()
	topic := &Topic{
		Id: id,
	}
	_, err = o.Delete(topic)
	if err != nil {
		return err
	}
	return nil
}

func GetTopic(tid string) (*Topic, error) {
	o := orm.NewOrm()
	id, err := strconv.ParseInt(tid, 10, 64)
	topic := &Topic{}
	err = o.QueryTable("topic").Filter("id", id).One(topic)
	//fmt.Printf("modles: %d\n", id)
	//fmt.Println(err)
	return topic, err

}

func GetAllTopics() ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	_, err := o.QueryTable("topic").All(&topics)
	if err != nil {
		return nil, err
	}
	return topics, err
}
