package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTask(t *Total) gin.HandlerFunc {

	return func(c *gin.Context) {
		all := t.Get()
		c.JSON(http.StatusOK, all)
	}

}

type Taskadd struct {
	Name string `json:"name"`
	Task string `json:"task"`
}

func PostTask(t *Total) gin.HandlerFunc {

	return func(c *gin.Context) {
		new := Taskadd{}
		c.Bind(&new)
		addTask := Tasks{
			Name: new.Name,
			Task: new.Task,
		}
		t.Add(addTask)
		c.Status(http.StatusNoContent)

	}
}

func Delete(t *Total) gin.HandlerFunc {

	return func(c *gin.Context) {
		Dname := c.Param("name")

		for i,v:= range t.Task{
			if v.Name == Dname {
				t.Task=append(t.Task[:i],t.Task[i+1:]...)
			}

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
	}
}
}

func Update(t *Total) gin.HandlerFunc {

	return func(c *gin.Context) {
		Dname := c.Param("name")
		Uname := Taskadd{}
		c.Bind(&Uname)
		for i, v := range t.Task {
			if v.Name == Dname {
				t.Task[i].Name = Uname.Name
			}

		}

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
	}
}
