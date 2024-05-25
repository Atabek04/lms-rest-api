package main

import (
	"lms-crud-api/internal/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) createCourseHandler(c *gin.Context) {
	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course := &data.Course{
		Title:       input.Title,
		Description: input.Description,
	}

	err = app.models.Courses.Insert(course)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"course": course})
}

func (app *application) showCourseHandler(c *gin.Context) {
	id, err := app.readIDParam(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid ID"})
		return
	}

	course, err := app.models.Courses.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"course": course})
}

func (app *application) updateCourseHandler(c *gin.Context) {
	id, err := app.readIDParam(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid ID"})
		return
	}

	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	err = c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course, err := app.models.Courses.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	course.Title = input.Title
	course.Description = input.Description

	err = app.models.Courses.Update(course)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"course": course})
}

func (app *application) deleteCourseHandler(c *gin.Context) {
	id, err := app.readIDParam(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid ID"})
		return
	}

	err = app.models.Courses.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
