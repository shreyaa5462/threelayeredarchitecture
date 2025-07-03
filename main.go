package main

import (
	taskds "ThreeLayeredArchitecture/datasource/task"
	taskhandler "ThreeLayeredArchitecture/handlers/task"
	taskservice "ThreeLayeredArchitecture/services/task"
	taskstore "ThreeLayeredArchitecture/store/task"
	"log"

	userds "ThreeLayeredArchitecture/datasource/user"
	userhandler "ThreeLayeredArchitecture/handlers/user"
	userservice "ThreeLayeredArchitecture/services/user"
	userstore "ThreeLayeredArchitecture/store/user"

	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	// Init task DB
	db, err := taskds.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to task DB:", err)
	}
	defer db.Close()

	taskStore := taskstore.NewTaskStore(db)
	taskService := taskservice.NewTaskService(taskStore)
	taskHandler := taskhandler.NewTaskHandler(taskService)

	app.POST("/task", taskHandler.CreateTask)
	app.GET("/task", taskHandler.GetPendingTasks)
	app.GET("/task/{id}", taskHandler.GetTaskByID)
	app.PUT("/task/{id}", taskHandler.CompleteTask)
	app.DELETE("/task/{id}", taskHandler.DeleteTask)

	// Init user DB
	userDB, err := userds.InitUserDB()
	if err != nil {
		log.Fatal("Failed to connect to user DB:", err)
	}
	userStore := userstore.NewUserStore(userDB)
	userService := userservice.NewUserService(userStore)
	userHandler := userhandler.NewUserHandler(userService)

	app.POST("/user", userHandler.CreateUser)
	app.GET("/user", userHandler.GetAllUsers)
	app.GET("/user/{id}", userHandler.GetUserByID)

	app.Run()
}
