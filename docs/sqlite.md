$ sqlite3 test.db
SQLite version 3.32.2 2020-06-04 12:58:43
Enter ".help" for usage hints.

sqlite> .databases
main: /gospace/go-demos/go-gorm-sqlite/test.db

sqlite> .tables
products

sqlite> .schema products
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

sqlite> .indices products
idx_products_deleted_at

sqlite> .header on
sqlite> .mode list
sqlite> SELECT * FROM products;
id|created_at|updated_at|deleted_at|title|code|price
1|2021-07-11 13:12:42.867814+08:00|2021-07-11 13:12:42.871312+08:00
|2021-07-11 13:12:42.872189+08:00|新款手机|F42|2000

sqlite> .mode line
sqlite> SELECT * FROM products;
        id = 1
created_at = 2021-07-11 13:12:42.867814+08:00
updated_at = 2021-07-11 13:12:42.871312+08:00
deleted_at = 2021-07-11 13:12:42.872189+08:00
     title = 新款手机
      code = F42
     price = 2000

---


Mac 上交叉编译Golang项目到Linux（sqlite3）
<https://yryz.net/post/mac-cross-compile-golang-for-linux-sqlite3/>

go没法在Mac OSX上交叉编译到Linux，原因是go-sqlite3使用了cgo。

brew install FiloSottile/musl-cross/musl-cross