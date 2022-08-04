package godigitalent

import (
	"database/sql"
	"godigitalent/mysqldata"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App     *fiber.App
	DB      *sql.DB
	Queries *mysqldata.Queries
}

func (s *Server) Routes() {

	s.App.Get("/api/list", s.getAllTasks)
	s.App.Post("/api/store", s.storeTask)
	s.App.Get("/api/:id", s.getTaskById)
	s.App.Put("/api/update/:id", s.updateTask)
	s.App.Put("/api/update/progress/:id", s.updateTaskProgress)
	s.App.Delete("/api/:id", s.deleteTask)
}
