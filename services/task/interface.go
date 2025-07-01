package task

import "ThreeLayeredArchitecture/models"

type TaskStoreInterface interface {
	CreateTask(description string) (models.MYTask, error)
	GetPendingTasks() ([]models.MYTask, error)
	GetTaskByID(id int) (models.MYTask, error)
	MarkTaskCompleted(id int) error
	DeleteTask(id int) error
}
