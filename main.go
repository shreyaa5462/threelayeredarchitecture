package main

import (
	taskhandler "ThreeLayeredArchitecture/handlers/task"
	taskservice "ThreeLayeredArchitecture/services/task"
	taskstore "ThreeLayeredArchitecture/store/task"

	userhandler "ThreeLayeredArchitecture/handlers/user"
	userservice "ThreeLayeredArchitecture/services/user"
	userstore "ThreeLayeredArchitecture/store/user"

	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()
	taskStore := taskstore.NewTaskStore()
	taskService := taskservice.NewTaskService(taskStore)
	taskHandler := taskhandler.NewTaskHandler(taskService)

	app.POST("/task", taskHandler.CreateTask)
	app.GET("/task", taskHandler.GetPendingTasks)
	app.GET("/task/{id}", taskHandler.GetTaskByID)
	app.PUT("/task/{id}", taskHandler.CompleteTask)
	//app.DELETE("/task/{id}", taskHandler.DeleteTask)

	userStore := userstore.NewUserStore()
	userService := userservice.NewUserService(userStore)
	userHandler := userhandler.NewUserHandler(userService)

	app.POST("/user", userHandler.CreateUser)
	app.GET("/user", userHandler.GetAllUsers)
	app.GET("/user/{id}", userHandler.GetUserByID)

	app.Run()
}
