package gorm

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/kingson4wu/mp_weixin_server/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

//定义全局的db对象，我们执行数据库操作主要通过他实现。
var _db *gorm.DB

//create database weixin_app default character set utf8mb4 collate utf8mb4_unicode_ci;

//包初始化函数，golang特性，每个包初始化的时候会自动执行init函数，这里用来初始化gorm。
func InitDB() {
	//配置MySQL连接参数

	_database := config.GetDatabaseConfig()

	username := _database.Username //账号
	password := _database.Password //密码
	host := _database.Host         //数据库地址，可以是Ip或者域名
	port := _database.Port         //数据库端口
	Dbname := _database.Dbname     //数据库名
	timeout := _database.Timeout   //连接超时，10秒

	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)

	// 声明err变量，下面不能使用:=赋值运算符，否则_db变量会当成局部变量，导致外部无法访问_db变量
	var err error
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
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
		log.Println("failed to AddExtranetIp ... " + error.Error())
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
		log.Println("failed to AddPhoto ... " + error.Error())
		panic(error)
	}

	//base64Image := savePhoto(image)

	db.Create(&Photo{Image: image, Account: account})

	log.Println("AddPhoto success ... ")

}

func savePhoto(url string) string {
	resp, err := http.Get(url)

	if err != nil {
		return ""
	}

	defer resp.Body.Close()

	if err != nil {
		return ""
	}

	pix, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return ""
	}

	encodeString := base64.StdEncoding.EncodeToString(pix)

	//fmt.Println(encodeString)

	return encodeString

}

func SelectTodayPhotos(account string) []string {
	return SelectPhotos(account, time.Now())
}

func SelectPhotos(account string, day time.Time) []string {

	db := GetDB()

	var photoList []Photo

	// 将查询出来的数据放到切片中
	db.Find(&photoList)

	startTime, end := GetDateTime(day)
	db.Where("created_at BETWEEN ? AND ? AND Account = ?", startTime, end, account).Find(&photoList)

	resultList := []string{}

	for _, photo := range photoList {
		if photo.Image != "" {
			resultList = append(resultList, photo.Image)
		}
	}

	return resultList
}

func GetDateTime(day time.Time) (*time.Time, *time.Time) {

	//date := time.Now().AddDate(0, 0, -1).Local().Format("2006-01-02")
	date := day.Local().Format("2006-01-02")

	//获取当前时区
	loc, _ := time.LoadLocation("Local")

	//日期当天0点时间戳(拼接字符串)
	startDate := date + "_00:00:00"
	startTime, _ := time.ParseInLocation("2006-01-02_15:04:05", startDate, loc)

	//日期当天23时59分时间戳
	endDate := date + "_23:59:59"
	end, _ := time.ParseInLocation("2006-01-02_15:04:05", endDate, loc)

	//返回当天0点和23点59分的时间戳
	return &startTime, &end
}

func AddTodoItem(content string, sort int, account string) {

	db := GetDB()

	// Migrate the schema
	error := db.AutoMigrate(&TodoItem{})

	if error != nil {
		log.Println("failed to AddTodoItem ... " + error.Error())
		panic(error)
	}

	db.Create(&TodoItem{Content: content, Sort: sort, Account: account})

	log.Println("AddTodoItem success ... ")

}

func CompleteTodoItem(id int) {

	db := GetDB()

	// Migrate the schema
	error := db.AutoMigrate(&TodoItem{})

	if error != nil {
		log.Println("failed to CompleteTodoItem ... " + error.Error())
		panic(error)
	}

	db.Model(&TodoItem{}).Where("id = ? AND completed = ?", id, false).Update("completed", true)

	log.Println("CompleteTodoItem success ... ")

}

func DeleteTodoItem(id int) {

	db := GetDB()

	// Migrate the schema
	error := db.AutoMigrate(&TodoItem{})

	if error != nil {
		log.Println("failed to DeleteTodoItem ... " + error.Error())
		panic(error)
	}

	db.Delete(&TodoItem{}, id)

	log.Println("DeleteTodoItem success ... ")

}

func SelectTodoList(account string) []TodoItem {

	db := GetDB()

	var todoList []TodoItem

	db.Find(&todoList)

	db.Order("sort DESC").Where(" account = ? AND completed = false", account).Find(&todoList)

	return todoList
}
