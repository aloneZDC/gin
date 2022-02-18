package blog

import (
	"bytes"
	"database/sql"
	"fmt"
	"gin/config"
	_ "gin/config"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func Routers(e *gin.Engine) {
	e.GET("/", postHandler)
	e.GET("/comment", commentHandler)
	e.GET("/getinfo", getInfo)
	e.GET("/getall", getAllInfo)
	e.POST("/add", add)
	e.POST("/update", update)
	e.POST("/del", del)
}

func postHandler(e *gin.Context) {
	e.JSON(http.StatusOK, gin.H{
		"message": "Hello www.aaa.com",
	})
}

//创建表
func commentHandler(e *gin.Context) {
	//e.JSON(http.StatusOK, gin.H{
	//	"message": "Hello www.bbb.com",
	//})
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}
	stmt, err := db.Prepare("CREATE TABLE person (id int NOT NULL AUTO_INCREMENT, first_name varchar(40), last_name varchar(40), PRIMARY KEY (id))ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';;")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Printf("Person Table successfully migrated....")
	}
}

//获取单条数据
func getInfo(e *gin.Context) {
	db := config.InitDB()
	defer db.Close()
	type Ecm_persons struct {
		Id         int
		First_Name string
		Last_Name  string
	}
	var (
		person Ecm_persons
		result gin.H
	)
	err := db.First(&person, "Id=1").Error
	if err != nil {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  1,
		}
	}
	e.JSON(http.StatusOK, result)
}

//获取多条数据
func getAllInfo(e *gin.Context) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()

	type Person struct {
		Id         int
		First_Name string
		Last_Name  string
	}
	var (
		person  Person
		persons []Person
	)
	rows, err := db.Query("select id, first_name, last_name from person;")
	if err != nil {
		fmt.Print(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&person.Id, &person.First_Name, &person.Last_Name)
		persons = append(persons, person)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	defer rows.Close()
	e.JSON(http.StatusOK, gin.H{
		"result": persons,
		"count":  len(persons),
	})
}

//新增数据
func add(e *gin.Context) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()

	var buffer bytes.Buffer
	first_name := e.PostForm("first_name")
	last_name := e.PostForm("last_name")
	if first_name == "" || last_name == "" {
		e.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("error created"),
		})
		return
	}
	stmt, err := db.Prepare("insert into person (first_name, last_name) values(?,?);")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(first_name, last_name)
	if err != nil {
		fmt.Print(err.Error())
	}
	// Fastest way to append strings
	buffer.WriteString(first_name)
	buffer.WriteString(" ")
	buffer.WriteString(last_name)
	defer stmt.Close()
	name := buffer.String()
	e.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf(" %s successfully created", name),
	})
}

//更新数据
func update(e *gin.Context) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	var buffer bytes.Buffer
	id := e.PostForm("id")
	if id == "" {
		e.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("error update,id is null"),
		})
		return
	}
	first_name := e.PostForm("first_name")
	last_name := e.PostForm("last_name")
	stmt, err := db.Prepare("update person set first_name= ?, last_name= ? where id= ?;")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(first_name, last_name, id)
	if err != nil {
		fmt.Print(err.Error())
	}
	// Fastest way to append strings
	buffer.WriteString(first_name)
	buffer.WriteString(" ")
	buffer.WriteString(last_name)
	defer stmt.Close()
	name := buffer.String()
	e.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully updated to %s", name),
	})
}

//删除数据
func del(e *gin.Context) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	id := e.PostForm("id")
	if id == "" {
		e.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("error delete,id is null"),
		})
		return
	}
	stmt, err := db.Prepare("delete from person where id= ?;")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	e.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully deleted user: %s", id),
	})
}
