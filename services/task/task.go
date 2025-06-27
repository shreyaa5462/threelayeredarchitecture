package taskservice

import (
	"ThreeLayeredArchitecture/models"
	"ThreeLayeredArchitecture/store/task"
)

type TaskService struct {
	Store *taskstore.TaskStore
}

func NewTaskService(store *taskstore.TaskStore) *TaskService {
	return &TaskService{Store: store}
}

func (s *TaskService) AddTask(desc string) (models.Task, error) {
	return s.Store.Add(desc)
}

func (s *TaskService) GetPendingTasks() ([]models.Task, error) {
	return s.Store.GetPending()
}

func (s *TaskService) GetTaskByID(id int) (models.Task, error) {
	return s.Store.GetByID(id)
}

func (s *TaskService) CompleteTask(id int) (string, error) {
	task, err := s.Store.GetByID(id)
	if err != nil {
		return "", err
	}
	if task.Completed {
		return "Task already completed", nil
	}
	err = s.Store.MarkComplete(id)
	if err != nil {
		return "", err
	}
	return "Task marked as complete", nil
}

func (s *TaskService) DeleteTask(id int) error {
	return s.Store.Delete(id)
}
