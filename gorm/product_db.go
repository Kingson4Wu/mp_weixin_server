package gorm

import (
	"os"

	"github.com/kingson4wu/weixin-app/common"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title string
	Code  string
	Price uint
}

type ExtranetIp struct {
	gorm.Model
	IP string
}

func openDatabase() *gorm.DB {

	dbDirPath := common.AppDataDir() + "/db"

	if !common.Exists(dbDirPath) {
		os.Mkdir(dbDirPath, os.ModePerm)
	}

	dbPath := dbDirPath + "/wexin_app.db"

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func ExistExtranetIp(ip string) bool {

	db := openDatabase()

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

	db := openDatabase()

	// Migrate the schema
	db.AutoMigrate(&ExtranetIp{})

	db.Create(&ExtranetIp{IP: ip})

}

func Operate() {

	db := openDatabase()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// 插入内容
	db.Create(&Product{Title: "新款手机", Code: "D42", Price: 1000})
	db.Create(&Product{Title: "新款电脑", Code: "D43", Price: 3500})

	// 读取内容
	var product Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// 更新操作：更新单个字段
	db.Model(&product).Update("Price", 2000)

	// 更新操作：更新多个字段
	db.Model(&product).Updates(Product{Price: 2000, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 2000, "Code": "F42"})

	// 删除操作：
	db.Delete(&product, 1)
}

/**
CREATE TABLE `products` (
	`id` integer,`created_at` datetime,
	`updated_at` datetime,
	`deleted_at` datetime,
	`title` text,
	`code` text,
	`price` integer,
	PRIMARY KEY (`id`)
);
CREATE INDEX `idx_products_deleted_at` ON `products`(`deleted_at`);
*/

//./sqlite3 ~/Personal/go-src/weixin-app/test.db

//https://blog.csdn.net/cnwyt/article/details/118904882
//运行结束后，查看当前目录，发现项目里会多一个 test.db 文件，就是生产的 sqlite 数据库文件。
//DBeaver 连接 SQLite 数据库
//https://www.jianshu.com/p/0df6f38b221d
//使用go-git备份数据库文件

///Users/kingsonwu/soft/sqlite-tools-osx-x86-3380100/sqlite3 ~/.weixin_app/db/test.db
