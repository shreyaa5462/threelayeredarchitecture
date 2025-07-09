package task

import (
	"ThreeLayeredArchitecture/models/task"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/http/response"
	"strconv"
)

type TaskHandler struct {
	Service Taskservice
}

func NewTaskHandler(service Taskservice) *TaskHandler {
	return &TaskHandler{Service: service}
}

// GetAllTasks handles GET /tasks.
func (h *TaskHandler) GetAllTasks(ctx *gofr.Context) (any, error) {
	tasks, err := h.Service.GetPending(ctx)
	if err != nil {
		return nil, err
	}

	return response.Raw{Data: tasks}, nil
}

// CreateTask handles POST /tasks.
func (h *TaskHandler) CreateTask(ctx *gofr.Context) (any, error) {
	var input models.Task

	err := ctx.Bind(&input)
	if err != nil {
		return nil, err
	}

	task, err := h.Service.Add(ctx, input)
	if err != nil {
		return nil, err
	}

	return response.Raw{Data: task}, nil
}

// DeleteTask handles DELETE /tasks?id=1
func (h *TaskHandler) DeleteTask(ctx *gofr.Context) (any, error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return nil, err
	}

	msg, err := h.Service.Delete(ctx, id)
	if err != nil {
		return msg, err
	}

	return response.Raw{Data: msg}, nil
}

// MarkTaskComplete handles PATCH /tasks?id=1
func (h *TaskHandler) MarkTaskComplete(ctx *gofr.Context) (any, error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return nil, err
	}

	msg, err := h.Service.MarkComplete(ctx, id)
	if err != nil {
		return nil, err
	}

	return response.Raw{Data: msg}, nil
}

func (h *TaskHandler) HandleTaskByID(ctx *gofr.Context) (any, error) {
	id, err := strconv.Atoi(ctx.PathParam("id"))
	if err != nil {
		return nil, err
	}

	task, err := h.Service.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return response.Raw{Data: task}, nil
}
