package task

import (
	"ThreeLayeredArchitecture/models"
	"gofr.dev/pkg/gofr"
)

type TaskStoreInterface interface {
	CreateTask(ctx *gofr.Context, task models.MYTask) (models.MYTask, error)
	GetPendingTasks(ctx *gofr.Context) ([]models.MYTask, error)
	GetTaskByID(ctx *gofr.Context, id int) (models.MYTask, error)
	MarkTaskCompleted(ctx *gofr.Context, id int) error
	DeleteTask(ctx *gofr.Context, id int) error
}
