// package main
//
// import (
//
//	userds "ThreeLayeredArchitecture/datasource/user"
//	userhandler "ThreeLayeredArchitecture/handlers/user"
//	userservice "ThreeLayeredArchitecture/services/user"
//	userstore "ThreeLayeredArchitecture/store/user"
//	"fmt"
//	"log"
//	"net/http"
//
//	taskds "ThreeLayeredArchitecture/datasource/task"
//	taskhandler "ThreeLayeredArchitecture/handlers/task"
//	taskservice "ThreeLayeredArchitecture/services/task"
//	taskstore "ThreeLayeredArchitecture/store/task"
//
// )
//
//	func main() {
//		db, err := taskds.InitDB()
//		if err != nil {
//			log.Fatal("Failed to connect to database:", err)
//		}
//		defer db.Close()
//
//		store := taskstore.NewTaskStore(db)
//		service := taskservice.NewTaskService(store)
//		handler := taskhandler.NewTaskHandler(service)
//
//		http.HandleFunc("/task", handler.HandleTasks)
//		http.HandleFunc("/task/{id}", handler.HandleTaskByID)
//
//		if err != nil {
//			log.Fatal("Server failed to start:", err)
//		}
//		userDB, err := userds.InitUserDB()
//		if err != nil {
//			log.Fatal("Failed to connect to user database:", err)
//		}
//		defer userDB.Close()
//
//		userStore := userstore.NewUserStore(userDB)
//		userService := userservice.NewUserService(userStore)
//		userHandler := userhandler.NewUserHandler(userService)
//
//		http.HandleFunc("/user", userHandler.HandleUsers)
//		http.HandleFunc("/user/", userHandler.HandleUserByID)
//
//		fmt.Println("Server starting on :8000")
//		err = http.ListenAndServe(":8000", nil)
//		if err != nil {
//			log.Fatal("Server failed to start:", err)
//		}
//	}
package main

import (
	"fmt"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"

	userds "ThreeLayeredArchitecture/datasource/user"
	userhandler "ThreeLayeredArchitecture/handlers/user"
	userservice "ThreeLayeredArchitecture/services/user"
	userstore "ThreeLayeredArchitecture/store/user"

	taskds "ThreeLayeredArchitecture/datasource/task"
	taskhandler "ThreeLayeredArchitecture/handlers/task"
	taskservice "ThreeLayeredArchitecture/services/task"
	taskstore "ThreeLayeredArchitecture/store/task"
)

func main() {
	db, err := taskds.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to task database:", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Println("Error closing task database:", err)
		}
	}()

	store := taskstore.NewTaskStore(db)
	service := taskservice.NewTaskService(store)
	handler := taskhandler.NewTaskHandler(service)

	http.HandleFunc("/task", handler.HandleTasks)
	http.HandleFunc("/task/{id}", handler.HandleTaskByID)

	userDB, err := userds.InitUserDB()

	defer func() {
		if err := userDB.Close(); err != nil {
			log.Println("Error closing user database:", err)
		}
	}()

	userStore := userstore.NewUserStore(userDB)
	userService := userservice.NewUserService(userStore)
	userHandler := userhandler.NewUserHandler(userService)

	http.HandleFunc("/user", userHandler.HandleUsers)
	http.HandleFunc("/user/", userHandler.HandleUserByID)

	http.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(http.Dir("docs"))))
	http.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/docs/task.json"),
	))
	fmt.Println("Server starting on :8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
