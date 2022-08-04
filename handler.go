package godigitalent

import (
	"godigitalent/mysqldata"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Task struct {
	ID          int32     `json:"id"`
	Description string    `json:"description"`
	Assignee    string    `json:"assignee"`
	IsDone      bool      `json:"is_done"`
	DeadlineAt  time.Time `json:"deadline_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (s *Server) getAllTasks(c *fiber.Ctx) error {
	ToDo, err := s.Queries.GetToDoList(c.Context())
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"data":   err.Error(),
		})
	}

	Done, err := s.Queries.GetDoneList(c.Context())
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"data":   err.Error(),
		})
	}

	data := struct {
		ToDo []mysqldata.Task `json:"todo"`
		Done []mysqldata.Task `json:"done"`
	}{
		ToDo: ToDo,
		Done: Done,
	}

	return c.Status(200).JSON(fiber.Map{
		"status": 200,
		"data":   data,
	})
}

func (s *Server) storeTask(c *fiber.Ctx) error {
	payload := struct {
		Description string `json:"description"`
		Assignee    string `json:"assignee"`
		DeadlineAt  string `json:"deadline_at"`
	}{}
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	DeadlineAT, err := time.Parse("2006-01-02", payload.DeadlineAt)
	if err != nil {
		return err
	}

	data, err := s.Queries.TaskInsert(c.Context(), mysqldata.TaskInsertParams{
		Description: payload.Description,
		Assignee:    payload.Assignee,
		DeadlineAt:  DeadlineAT,
	})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"data":   err.Error(),
		})
	}

	ID, err := data.LastInsertId()
	if err != nil {
		return err
	}

	result, err := s.Queries.TaskGetById(c.Context(), int32(ID))
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{
		"status": 200,
		"data":   result,
	})
}

func (s *Server) getTaskById(c *fiber.Ctx) error {
	ID, err := strconv.ParseInt(c.Params("id"), 0, 32)
	if err != nil {
		return err
	}

	data, err := s.Queries.TaskGetById(c.Context(), int32(ID))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"data":   err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": 200,
		"data":   data,
	})
}

func (s *Server) updateTask(c *fiber.Ctx) error {
	payload := struct {
		Description string `json:"description"`
		Assignee    string `json:"assignee"`
		DeadlineAt  string `json:"deadline_at"`
		ID          int32  `json:"id"`
	}{}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"data":   "Failed to parse data",
		})
	}

	ID, err := strconv.ParseInt(c.Params("id"), 0, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"data":   "Failed to convert ID",
		})
	}

	DeadlineAT, err := time.Parse("2006-01-02", payload.DeadlineAt)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"data":   "Failed to convert Deadline",
		})
	}

	err = s.Queries.TaskUpdate(c.Context(), mysqldata.TaskUpdateParams{
		Description: payload.Description,
		Assignee:    payload.Assignee,
		DeadlineAt:  DeadlineAT,
		ID:          int32(ID),
	})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"data":   "Failed to update data",
		})
	}

	result, err := s.Queries.TaskGetById(c.Context(), int32(ID))
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"data":    result,
		"message": "OK",
	})
}

func (s *Server) updateTaskProgress(c *fiber.Ctx) error {
	payload := struct {
		IsDone bool  `json:"is_done"`
		ID     int32 `json:"id"`
	}{}
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"data":   "Failed to parse data",
		})
	}

	ID, err := strconv.ParseInt(c.Params("id"), 0, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"data":   "Failed to convert ID",
		})
	}

	err = s.Queries.TaskUpdateProgress(c.Context(), mysqldata.TaskUpdateProgressParams{
		IsDone: payload.IsDone,
		ID:     int32(ID),
	})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"data":   "Failed to update data",
		})
	}

	result, err := s.Queries.TaskGetById(c.Context(), int32(ID))
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"data":    result,
		"message": "OK",
	})
}

func (s *Server) deleteTask(c *fiber.Ctx) error {
	ID, err := strconv.ParseInt(c.Params("id"), 0, 32)
	if err != nil {
		return err
	}

	err = s.Queries.TaskDelete(c.Context(), int32(ID))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": 400,
			"data":   "Failed to delete data",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status": 200,
		"data":   "Success to delete data",
	})
}
