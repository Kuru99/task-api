package main

import (
	"database/sql"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	// 一覧取得
	r.GET("/tasks", func(c *gin.Context) {
		rows, err := db.Query("SELECT id, title, done FROM tasks")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		var tasks []Task
		for rows.Next() {
			var t Task
			rows.Scan(&t.ID, &t.Title, &t.Done)
			tasks = append(tasks, t)
		}
		c.JSON(200, tasks)
	})

	// 作成
	r.POST("/tasks", func(c *gin.Context) {
		var t Task
		c.BindJSON(&t)
		db.QueryRow("INSERT INTO tasks (title, done) VALUES ($1, $2) RETURNING id", t.Title, t.Done).Scan(&t.ID)
		c.JSON(201, t)
	})

	// 更新
	r.PUT("/tasks/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var t Task
		c.BindJSON(&t)

		result, err := db.Exec("UPDATE tasks SET title=$1, done=$2 WHERE id=$3", t.Title, t.Done, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "見つかりません"})
			return
		}

		t.ID = id
		c.JSON(200, t)
	})

	// 削除
	r.DELETE("/tasks/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		result, err := db.Exec("DELETE FROM tasks WHERE id=$1", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		rowsAffected, _ := result.RowsAffected()
		if rowsAffected == 0 {
			c.JSON(http.StatusNotFound, gin.H{"message": "見つかりません"})
			return
		}

		c.JSON(200, gin.H{"message": "削除しました"})
	})

	r.Run(":8080")
}
