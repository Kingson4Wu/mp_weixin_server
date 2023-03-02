package gorm

import (
	"gorm.io/gorm"
	"log"
)

type Memo struct {
	gorm.Model
	Account string `gorm:"column:account;type:varchar(255);index:index_account"`
	Content string `gorm:"column:content;type:varchar(2048)"`
}

func AddMemo(content string, account string) {

	db := GetDB()

	err := db.AutoMigrate(&Memo{})

	if err != nil {
		log.Println("failed to AddMemo ... " + err.Error())
		panic(err)
	}

	db.Create(&Memo{Content: content, Account: account})

	log.Println("AddMemo success ... ")

}

func SelectLatestMemo(account string) *Memo {

	defer func() {
		if err := recover(); err != nil {
			log.Printf("SelectLatestMemo error: %s \n", err)
		}
	}()

	log.Println("SelectLatestMemo :" + account)

	db := GetDB()

	var memo Memo

	db.Debug().Order("created_at DESC").Where("`account` = ? ", account).Limit(1).First(&memo)

	return &memo
}
