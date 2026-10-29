package main

import (
	"errors"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func studentRoute(app *fiber.App, db *gorm.DB) {
	app.Post("/api/students", func(c fiber.Ctx) error {
		student := new(Student)

		if error := c.Bind().Body(student); error != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": " Invalid request body",
			})
		}

		if error := createStudent(db, student); error != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": " Internal server error",
			})
		}

		return c.Status(201).JSON(fiber.Map{
			"message": "Student Create successful",
			"data":    student,
		})
	})

	app.Get("/api/students", func(c fiber.Ctx) error {
		var students []Student

		StuName := c.Query("StuName")
		StuLevel := c.Query("StuLevel")

		var error error
		if StuName != "" {
			students, error = getStudentByName(db, StuName)
		} else if StuLevel != "" {

			StuLevelInt, err := strconv.Atoi(StuLevel)
			if err != nil {
				return c.Status(400).JSON(fiber.Map{
					"message": "Invalid request parameter",
				})
			}
			students, error = getStudentByLevel(db, StuLevelInt)
		} else {
			students, error = getAllStudent(db)
		}

		if errors.Is(error, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"message": "Student not found",
			})
		}

		if error != nil {
			log.Printf("%v", error)
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "Success",
			"data":    students,
		})
	})

	app.Get("/api/students/:id", func(c fiber.Ctx) error {
		id := c.Params("id")

		idInt, error := strconv.Atoi(id)
		if error != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid request parameter",
			})
		}

		student, error := getStudentByID(db, idInt)

		if errors.Is(error, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"message": "Student not found",
			})
		}
		if error != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "success",
			"data":    student,
		})

	})

	app.Put("/api/students/:id", func(c fiber.Ctx) error {
		id := c.Params("id")

		idInt, error := strconv.Atoi(id)
		if error != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid request parameter",
			})
		}
		student := new(Student)
		if error := c.Bind().Body(student); error != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid request body",
			})
		}
		if error := updateStudent(db, idInt, student); error != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal Server error",
			})
		}
		return c.Status(200).JSON(fiber.Map{
			"message": "Student Update successful",
			"data":    student,
		})
	})
	app.Delete("/api/students/:id", func(c fiber.Ctx) error {
		id := c.Params("id")

		idInt, error := strconv.Atoi(id)
		if error != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid request parameter",
			})
		}

		if error := deleteStudent(db, idInt); error != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}
		return c.Status(200).JSON(fiber.Map{
			"message": "student delete success",
		})
	})
}
