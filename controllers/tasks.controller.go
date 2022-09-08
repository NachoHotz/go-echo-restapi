package controllers

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/nachohotz/go-echo-restapi/models"
)

var tasks = []models.Task{}

func GetTasks(c echo.Context) error {
  if len(tasks) == 0 {
    return c.JSON(404, "No tasks found")
  }

  return c.JSON(200, tasks)
}

func GetTask(c echo.Context) error {
  taskId, err := strconv.Atoi(c.Param("id"))

  if err != nil {
    return c.JSON(400, map[string]string{
      "status":  "error",
      "message": "Invalid task ID",
    })
  }

  taskFound := models.Task{}

  for _, task := range tasks {
    if task.ID == taskId {
      taskFound = task
    }
  }

  if taskFound.ID == 0 {
    return c.JSON(404, map[string]string{
      "status":  "error",
      "message": "Task not found",
    })
  }

  return c.JSON(200, taskFound)
}

func CreateTask(c echo.Context) error {
  taskBody := new(models.Task)

  var taskExists bool = false

  if err := c.Bind(taskBody); err != nil {
    return err
  }

  for _, task := range tasks {
    if task.ID == taskBody.ID {
      taskExists = true
    }
  }

  if taskExists {
    return c.JSON(400, map[string]string{
      "message": "Task already exists",
    })
  }

  tasks = append(tasks, *taskBody)

  return c.JSON(201, taskBody)
}

func UpdateTask(c echo.Context) error {
  taskId, err := strconv.Atoi(c.Param("id"))

  if err != nil {
    return c.JSON(400, map[string]string{
      "status":  "error",
      "message": "Invalid task ID",
    })
  }

  taskBody := new(models.Task)

  if err := c.Bind(taskBody); err != nil {
    return err
  }

  taskFound := models.Task{}

  for index, task := range tasks {
    if task.ID == taskId {
      taskFound = task
      tasks[index] = *taskBody
    }

    if taskFound.ID == 0 {
      return c.JSON(404, map[string]string{
        "status":  "error",
        "message": "Task not found",
      })
    }
  }

  return c.JSON(200, taskBody)
}

func DeleteTask(c echo.Context) error {
  taskId, err := strconv.Atoi(c.Param("id"))

  if err != nil {
    return c.JSON(400, map[string]string{
      "status":  "error",
      "message": "Invalid task ID",
    })
  }

  taskFound := models.Task{}

  for index, task := range tasks {
    if task.ID == taskId {
      taskFound = task
      tasks = append(tasks[:index], tasks[index+1:]...)
    }

  }

  if taskFound.ID == 0 {
    return c.JSON(404, map[string]string{
      "status":  "error",
      "message": "Task not found. Nothing to delete",
    })
  }

  return c.JSON(200, map[string]string{
    "status":  "success",
    "message": "Task deleted",
  })
}
