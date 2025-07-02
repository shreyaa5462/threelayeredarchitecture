package task

import (
	"ThreeLayeredArchitecture/models"
)

type TaskService struct {
	Store TaskStoreInterface
}

func NewTaskService(store TaskStoreInterface) *TaskService {
	return &TaskService{Store: store}
}

func (s *TaskService) CreateTask(task models.MYTask) (models.MYTask, error) {
	return s.Store.CreateTask(task)
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
