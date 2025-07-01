package task

import "ThreeLayeredArchitecture/models"

type TaskServiceInterface interface {
	CreateTask(description string) (models.MYTask, error)
	GetPendingTasks() ([]models.MYTask, error)
	GetTask(id int) (models.MYTask, error)
	CompleteTask(id int) error
	DeleteTask(id int) error
}
