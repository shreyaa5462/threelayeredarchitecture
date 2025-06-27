package main

import (
	"fmt"
	"log"
	"net/http"

	userds "ThreeLayeredArchitecture/datasource/user"
	userhandler "ThreeLayeredArchitecture/handlers/user"
	userservice "ThreeLayeredArchitecture/services/user"
	userstore "ThreeLayeredArchitecture/store/user"
)

func main() {
	db, err := userds.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	store := userstore.NewUserStore(db)
	service := userservice.NewUserService(store)
	handler := userhandler.NewUserHandler(service)

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handler.GetAllUsers(w, r)
		case "POST":
			handler.CreateUser(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/user/getbyid", handler.GetUserByID)

	fmt.Println("Server running at :8000")
	http.ListenAndServe(":8000", nil)
}
