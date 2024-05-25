package main

import (
	"lms-crud-api/internal/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) createLessonHandler(c *gin.Context) {
	var input struct {
		Title    string `json:"title"`
		Link     string `json:"link"`
		Conspect string `json:"conspect"`
		ModuleID uint   `json:"module_id"`
	}

	err := c.BindJSON(&input)
	if err != nil {
		app.badRequestResponse(c, err)
		return
	}

	lesson := &data.Lesson{
		Title:    input.Title,
		Link:     input.Link,
		Conspect: input.Conspect,
		ModuleID: input.ModuleID,
	}

	err = app.models.Lessons.Insert(lesson)
	if err != nil {
		app.serverErrorResponse(c, err)
		return
	}

	app.writeJSON(c, http.StatusCreated, gin.H{"lesson": lesson})
}

func (app *application) showLessonHandler(c *gin.Context) {
	id, err := app.readIDParam(c)
	if err != nil {
		app.notFoundResponse(c)
		return
	}

	lesson, err := app.models.Lessons.Get(id)
	if err != nil {
		app.notFoundResponse(c)
		return
	}

	app.writeJSON(c, http.StatusOK, gin.H{"lesson": lesson})
}

func (app *application) updateLessonHandler(c *gin.Context) {
	id, err := app.readIDParam(c)
	if err != nil {
		app.notFoundResponse(c)
		return
	}

	var input struct {
		Title    string `json:"title"`
		Link     string `json:"link"`
		Conspect string `json:"conspect"`
	}

	err = c.BindJSON(&input)
	if err != nil {
		app.badRequestResponse(c, err)
		return
	}

	lesson, err := app.models.Lessons.Get(id)
	if err != nil {
		app.notFoundResponse(c)
		return
	}

	lesson.Title = input.Title
	lesson.Link = input.Link
	lesson.Conspect = input.Conspect

	err = app.models.Lessons.Update(lesson)
	if err != nil {
		app.serverErrorResponse(c, err)
		return
	}

	app.writeJSON(c, http.StatusOK, gin.H{"lesson": lesson})
}

func (app *application) deleteLessonHandler(c *gin.Context) {
	id, err := app.readIDParam(c)
	if err != nil {
		app.notFoundResponse(c)
		return
	}

	err = app.models.Lessons.Delete(id)
	if err != nil {
		app.serverErrorResponse(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
