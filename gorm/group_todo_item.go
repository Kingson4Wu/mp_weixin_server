package gorm

import "log"

func AddGroupTodoItem(content string, sort int, account string, group string) {

	db := GetDB()

	// Migrate the schema
	err := db.AutoMigrate(&GroupTodoItem{})

	if err != nil {
		log.Println("failed to AddGroupTodoItem ... " + err.Error())
		panic(err)
	}

	db.Create(&GroupTodoItem{Content: content, Sort: sort, Account: account, Group: group})

	log.Println("AddGroupTodoItem success ... ")

}

func CompleteGroupTodoItem(id int) {

	db := GetDB()

	// Migrate the schema
	err := db.AutoMigrate(&GroupTodoItem{})

	if err != nil {
		log.Println("failed to CompleteGroupTodoItem ... " + err.Error())
		panic(err)
	}

	db.Model(&GroupTodoItem{}).Where("id = ? AND completed = ?", id, false).Update("completed", true)

	log.Println("CompleteGroupTodoItem success ... ")

}

func DeleteGroupTodoItem(id int) {

	db := GetDB()

	// Migrate the schema
	err := db.AutoMigrate(&GroupTodoItem{})

	if err != nil {
		log.Println("failed to DeleteGroupTodoItem ... " + err.Error())
		panic(err)
	}

	db.Delete(&GroupTodoItem{}, id)

	log.Println("DeleteGroupTodoItem success ... ")

}

func SelectGroupTodoList(group string) []GroupTodoItem {

	log.Println("SelectGroupTodoList ... :" + group + ":")

	db := GetDB()

	var todoList []GroupTodoItem

	//db.Debug().Find(&todoList)

	db.Debug().Order("sort DESC").Where("`group` = ? AND completed = false", group).Find(&todoList)

	return todoList
}
