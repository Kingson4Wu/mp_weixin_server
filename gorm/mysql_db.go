package gorm

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//定义全局的db对象，我们执行数据库操作主要通过他实现。
var _db *gorm.DB

//create database weixin_app default character set utf8mb4 collate utf8mb4_unicode_ci;

//包初始化函数，golang特性，每个包初始化的时候会自动执行init函数，这里用来初始化gorm。
func InitDB() {
	//配置MySQL连接参数
	username := "root"     //账号
	password := "611264"   //密码
	host := "192.168.10.8" //数据库地址，可以是Ip或者域名
	port := 3306           //数据库端口
	Dbname := "weixin_app" //数据库名
	timeout := "10s"       //连接超时，10秒

	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)

	// 声明err变量，下面不能使用:=赋值运算符，否则_db变量会当成局部变量，导致外部无法访问_db变量
	var err error
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	sqlDB, _ := _db.DB()

	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(10) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(3)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
}

//获取gorm db对象，其他包需要执行数据库查询的时候，只要通过tools.getDB()获取db对象即可。
//不用担心协程并发使用同样的db对象会共用同一个连接，db对象在调用他的方法的时候会从数据库连接池中获取新的连接
func GetDB() *gorm.DB {
	return _db
}

/**
gorm调试模式
result := db.Debug().Where("username = ?", "tizi365").First(&u)
*/

func ExistExtranetIp(ip string) bool {

	db := GetDB()

	// Migrate the schema
	db.AutoMigrate(&ExtranetIp{})

	var ipList []ExtranetIp

	// 将查询出来的数据放到切片中
	db.Find(&ipList)

	for _, extranetIp := range ipList {
		if extranetIp.IP == ip {
			return true
		}
	}

	return false
}

func AddExtranetIp(ip string) {

	db := GetDB()

	// Migrate the schema
	error := db.AutoMigrate(&ExtranetIp{})

	if error != nil {
		panic(error)
	}

	//db.Create(&ExtranetIp{IP: ip})
	db.Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&ExtranetIp{IP: ip})

	//db.Clauses(clause.Insert{Modifier: "OR IGNORE"}).Create(&ExtranetIp{IP: ip})
}

func AddPhoto(image string, account string) {

	db := GetDB()

	// Migrate the schema
	error := db.AutoMigrate(&Photo{})

	if error != nil {
		panic(error)
	}

	db.Create(&Photo{Image: image, Account: account})

}
