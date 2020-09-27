package db_mysql

import (
	"beego02/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

//链接数据库
//创建数据库的配置
func init() {
	fmt.Println("链接mysql数据库")
	config := beego.AppConfig
	dbDriver := config.String("db_driverName")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")

	connUrl := dbUser + ":" + dbPassword + "@tcp("+dbIp+")/"+dbName+"?charset=utf8"
	db, err := sql.Open(dbDriver,connUrl)
	if err != nil {
		fmt.Println(err.Error())
		panic("数据库连接错误，请检查配置")
	}
	//为全局变量赋值
	Db = db
}
//将用户信息保存到数据表当中
func InsertUser(user models.User) (int64,error) {
	//1.将用户密码进行hash加密,使用md5计算哈希值
	hashMd5 := md5.New()
	hashMd5.Write([]byte(user.Psaaword))
	bytes := hashMd5.Sum(nil)
	user.Psaaword = hex.EncodeToString(bytes)
	fmt.Println("将要保存的用户名:",user.Nick,"密码:",user.Psaaword)
	result , err := Db.Exec("insert into user(nick, password, email, birthday, address) values(?,?,?,?,?) ",
		user.Nick,user.Psaaword,user.Email,user.Birthday,user.Address,
	)
	if err != nil {//保存数据时遇到错误
		return -1 ,err
	}
	id, err := result.RowsAffected()
	if err != nil {
		return -1 ,err
	}
	return id ,nil
}
//查询用户
func QueryUser()  {
	Db.QueryRow("select * from ")
}