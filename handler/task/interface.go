package task

import (
	models "ThreeLayeredArchitecture/models/task"
	"gofr.dev/pkg/gofr"
)

type Taskservice interface {
	GetPending(ctx *gofr.Context) ([]models.Task, error)
	Add(ctx *gofr.Context, input models.Task) (models.Task, error)
	Delete(ctx *gofr.Context, id int) (string, error)
	MarkComplete(ctx *gofr.Context, id int) (string, error)
	GetByID(ctx *gofr.Context, id int) (models.Task, error)
}
