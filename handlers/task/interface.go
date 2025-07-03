package task

import (
	"ThreeLayeredArchitecture/models"
	"gofr.dev/pkg/gofr"
)

type TaskServiceInterface interface {
	CreateTask(ctx *gofr.Context, ask models.MYTask) (models.MYTask, error)
	GetPendingTasks(ctx *gofr.Context) ([]models.MYTask, error)
	GetTask(ctx *gofr.Context, id int) (models.MYTask, error)
	CompleteTask(ctx *gofr.Context, id int) error
	DeleteTask(ctx *gofr.Context, id int) error
}
