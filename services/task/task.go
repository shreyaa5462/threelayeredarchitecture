package task

import (
	"ThreeLayeredArchitecture/models"
	"gofr.dev/pkg/gofr"
)

type TaskService struct {
	Store TaskStoreInterface
}

func NewTaskService(store TaskStoreInterface) *TaskService {
	return &TaskService{Store: store}
}

func (s *TaskService) CreateTask(ctx *gofr.Context, task models.MYTask) (models.MYTask, error) {
	return s.Store.CreateTask(ctx, task)
}

func (s *TaskService) GetPendingTasks(ctx *gofr.Context) ([]models.MYTask, error) {
	return s.Store.GetPendingTasks(ctx)
}

func (s *TaskService) GetTask(ctx *gofr.Context, id int) (models.MYTask, error) {
	return s.Store.GetTaskByID(ctx, id)
}

func (s *TaskService) CompleteTask(ctx *gofr.Context, id int) error {
	return s.Store.MarkTaskCompleted(ctx, id)
}

func (s *TaskService) DeleteTask(ctx *gofr.Context, id int) error {
	return s.Store.DeleteTask(ctx, id)
}
