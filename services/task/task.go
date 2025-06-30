package task

import (
	"ThreeLayeredArchitecture/models"
)

type TaskStoreInterface interface {
	CreateTask(description string) (models.MYTask, error)
	GetPendingTasks() ([]models.MYTask, error)
	GetTaskByID(id int) (models.MYTask, error)
	MarkTaskCompleted(id int) error
	DeleteTask(id int) error
}

type TaskService struct {
	Store TaskStoreInterface
}

func NewTaskService(store TaskStoreInterface) *TaskService {
	return &TaskService{Store: store}
}

func (s *TaskService) CreateTask(description string) (models.MYTask, error) {
	return s.Store.CreateTask(description)
}

func (s *TaskService) GetPendingTasks() ([]models.MYTask, error) {
	return s.Store.GetPendingTasks()
}

func (s *TaskService) GetTask(id int) (models.MYTask, error) {
	return s.Store.GetTaskByID(id)
}

func (s *TaskService) CompleteTask(id int) error {
	return s.Store.MarkTaskCompleted(id)
}

func (s *TaskService) DeleteTask(id int) error {
	return s.Store.DeleteTask(id)
}
