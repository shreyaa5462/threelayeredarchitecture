package taskhandler

import (
	"ThreeLayeredArchitecture/models"
	"ThreeLayeredArchitecture/services/task"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type TaskHandler struct {
	Service *taskservice.TaskService
}

func NewTaskHandler(service *taskservice.TaskService) *TaskHandler {
	return &TaskHandler{Service: service}
}

func (h *TaskHandler) HandleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		tasks, err := h.Service.GetPendingTasks()
		if err != nil {
			http.Error(w, "Error fetching tasks", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	case http.MethodPost:
		var input models.Task
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		task, err := h.Service.AddTask(input.Description)
		if err != nil {
			http.Error(w, "Failed to add", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(task)
	case http.MethodDelete:
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		err := h.Service.DeleteTask(id)
		if err != nil {
			http.Error(w, "Delete failed", http.StatusNotFound)
			return
		}
		w.Write([]byte("Deleted successfully"))
	case http.MethodPatch:
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		msg, err := h.Service.CompleteTask(id)
		if err != nil {
			http.Error(w, "Update failed", http.StatusInternalServerError)
			return
		}
		w.Write([]byte(msg))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TaskHandler) HandleTaskByID(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		parts := strings.Split(r.URL.Path, "/")
		id, _ := strconv.Atoi(parts[len(parts)-1])
		task, err := h.Service.GetTaskByID(id)
		if err != nil {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
