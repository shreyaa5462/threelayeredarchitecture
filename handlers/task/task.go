package task

import (
	"ThreeLayeredArchitecture/models"
	"gofr.dev/pkg/gofr/http/response"

	"strconv"

	"gofr.dev/pkg/gofr"
)

type TaskHandler struct {
	Service TaskServiceInterface
}

func NewTaskHandler(service TaskServiceInterface) *TaskHandler {
	return &TaskHandler{Service: service}
}

func (h *TaskHandler) CreateTask(ctx *gofr.Context) (any, error) {
	var input models.MYTask
	err := ctx.Bind(&input)
	if err != nil {
		return nil, err
	}

	task, err := h.Service.CreateTask(ctx, input)
	if err != nil {
		return nil, err
	}

	return response.Raw{Data: task}, nil
}

func (h *TaskHandler) GetPendingTasks(ctx *gofr.Context) (any, error) {
	tasks, err := h.Service.GetPendingTasks(ctx)
	if err != nil {
		return nil, err
	}
	return response.Raw{Data: tasks}, nil
}

func (h *TaskHandler) GetTaskByID(ctx *gofr.Context) (any, error) {
	id, err := strconv.Atoi(ctx.PathParam("id"))
	if err != nil {
		return nil, err
	}

	task, err := h.Service.GetTask(ctx, id)
	if err != nil {
		return nil, err
	}
	return response.Raw{Data: task}, nil
}

func (h *TaskHandler) CompleteTask(ctx *gofr.Context) (any, error) {
	id, err := strconv.Atoi(ctx.PathParam("id"))
	if err != nil {
		return nil, err
	}

	err = h.Service.CompleteTask(ctx, id)
	if err != nil {
		return nil, err
	}
	return "Task marked as completed", nil
}

func (h *TaskHandler) DeleteTask(ctx *gofr.Context, id int) (any, error) {
	id, err := strconv.Atoi(ctx.PathParam("id"))
	if err != nil {
		return nil, err
	}

	err = h.Service.DeleteTask(ctx, id)
	if err != nil {
		return nil, err
	}

	return response.Raw{Data: "deleted"}, nil
}
