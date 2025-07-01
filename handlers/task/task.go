// package task
//
// import (
//
//	"encoding/json"
//	"fmt"
//
//	"net/http"
//	"strconv"
//
//	"ThreeLayeredArchitecture/models"
//
// )
//
//	type TaskServiceInterface interface {
//		CreateTask(description string) (models.MYTask, error)
//		GetPendingTasks() ([]models.MYTask, error)
//		GetTask(id int) (models.MYTask, error)
//		CompleteTask(id int) error
//		DeleteTask(id int) error
//	}
//
//	type TaskHandler struct {
//		Service TaskServiceInterface
//	}
//
//	func NewTaskHandler(service TaskServiceInterface) *TaskHandler {
//		return &TaskHandler{Service: service}
//	}
//
//	func (h *TaskHandler) HandleTasks(w http.ResponseWriter, r *http.Request) {
//		switch r.Method {
//		case http.MethodPost:
//			var input models.MYTask
//			json.NewDecoder(r.Body).Decode(&input)
//			task, err := h.Service.CreateTask(input.Description)
//			if err != nil {
//				fmt.Println("Error while creating task:", err)
//				http.Error(w, "Failed to create task", http.StatusInternalServerError)
//				return
//			}
//			json.NewEncoder(w).Encode(task)
//		case http.MethodGet:
//			tasks, err := h.Service.GetPendingTasks()
//			if err != nil {
//				http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
//				return
//			}
//			w.Header().Set("Content-Type", "application/json")
//			prettyJSON, _ := json.Marshal(tasks)
//			w.Write(prettyJSON)
//			json.NewEncoder(w).Encode(tasks)
//		case http.MethodPut:
//			id, _ := strconv.Atoi(r.URL.Query().Get("id"))
//			err := h.Service.CompleteTask(id)
//			if err != nil {
//				http.Error(w, "Failed to complete task", http.StatusInternalServerError)
//				return
//			}
//			w.Write([]byte("Task marked as completed"))
//		case http.MethodDelete:
//			id, _ := strconv.Atoi(r.URL.Query().Get("id"))
//			err := h.Service.DeleteTask(id)
//			if err != nil {
//				http.Error(w, "Failed to delete task", http.StatusInternalServerError)
//				return
//			}
//			w.Write([]byte("Task deleted"))
//		default:
//			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
//		}
//	}
//
//	func (h *TaskHandler) HandleTaskByID(w http.ResponseWriter, r *http.Request) {
//		if r.Method != http.MethodGet {
//			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
//			return
//		}
//		idStr := r.PathValue("id")
//		id, err := strconv.Atoi(idStr)
//		if err != nil {
//			fmt.Println("Error while parsing task ID:", err)
//			http.Error(w, "Invalid ID", http.StatusBadRequest)
//			return
//		}
//		task, err := h.Service.GetTask(id)
//		if err != nil {
//			http.Error(w, "Task not found", http.StatusNotFound)
//			return
//		}
//		json.NewEncoder(w).Encode(task)
//	}
package task

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"ThreeLayeredArchitecture/models"
)

type TaskHandler struct {
	Service TaskServiceInterface
}

func NewTaskHandler(service TaskServiceInterface) *TaskHandler {
	return &TaskHandler{Service: service}
}

func (h *TaskHandler) HandleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var input models.MYTask
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		task, err := h.Service.CreateTask(input.Description)

		if err != nil {
			fmt.Println("Error while creating task:", err)
			http.Error(w, "Failed to create task", http.StatusInternalServerError)

			return
		}

		if err := json.NewEncoder(w).Encode(task); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
		}

	case http.MethodGet:
		tasks, err := h.Service.GetPendingTasks()
		if err != nil {
			http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(tasks); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
		}

	case http.MethodPut:
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		if err := h.Service.CompleteTask(id); err != nil {
			http.Error(w, "Failed to complete task", http.StatusInternalServerError)
			return
		}
		if _, err := w.Write([]byte("Task marked as completed")); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
		}

	case http.MethodDelete:
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		if err := h.Service.DeleteTask(id); err != nil {
			http.Error(w, "Failed to delete task", http.StatusInternalServerError)
			return
		}

		if _, err := w.Write([]byte("Task deleted")); err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
		}

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TaskHandler) HandleTaskByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Fallback for Go 1.21 (use URL query param instead of PathValue)
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		fmt.Println("Error while parsing task ID:", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)

		return
	}

	task, err := h.Service.GetTask(id)
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(task); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
