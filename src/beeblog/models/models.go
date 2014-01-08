/**
 * Name:
 * Comment:
 * Author: Rhinux 
 * Web: http://www.rhinux.info./
 * Created: 2014-01-08 17:13:20
 * Last-Modified: 2014-01-08 18:22:39
 */
package models

import(
    "time"
    "os"
    "path"
    "github.com/Unknwon/com"
    "github.com/astaxie/beego/orm"
    _ "github.com/mattn/go-sqlite3"
)
const(
    _DB_NAME = "data/beeblog.db"
    _SQLITE3_DRIVER = "sqlite3"
)
type Category struct{
    Id          int64
    Title       string
    Created     time.Time  `orm:"index"`
    Views       int64        `orm:"index"`
    TopicTime   time.Time `orm:"index"`
    TopicCount  int64
    TopiclastUserId int64
}

type Topic struct{
    Id          int64
    Uid         int64
    Title       string
    Content     string   `orm:"size(5000)"`
    Attachment  string
    Created     time.Time  `orm:"index"`
    Updated     time.Time  `orm:"index"`
    Views       int64       `orm:"index"`
    Author      string
    ReplyTime         time.Time  `orm:"index"`
    ReplyCount        int64
    ReplyLastUserId   int64
}

func RegisterDB(){
    if !com.IsExist(_DB_NAME){
        os.MkdirAll(path.Dir(_DB_NAME),os.ModePerm)
        os.Create(_DB_NAME)
    }
    orm.RegisterModel(new(Category),new(Topic))
    orm.RegisterDriver(_SQLITE3_DRIVER, orm.DR_Sqlite)
    orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
}

