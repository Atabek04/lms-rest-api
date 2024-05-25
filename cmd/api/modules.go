package main

import (
	"lms-crud-api/internal/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) createModuleHandler(c *gin.Context) {
	var input struct {
		Title    string `json:"title"`
		CourseID uint   `json:"course_id"`
	}

	err := c.BindJSON(&input)
	if err != nil {
		app.badRequestResponse(c, err)
		return
	}

	module := &data.Module{
		Title:    input.Title,
		CourseID: input.CourseID,
	}

	err = app.models.Modules.Insert(module)
	if err != nil {
		app.serverErrorResponse(c, err)
		return
	}

	app.writeJSON(c, http.StatusCreated, gin.H{"module": module})
}

func (app *application) showModuleHandler(c *gin.Context) {
	id, err := app.readIDParam(c)
	if err != nil {
		app.notFoundResponse(c)
		return
	}

	module, err := app.models.Modules.Get(id)
	if err != nil {
		app.notFoundResponse(c)
		return
	}

	app.writeJSON(c, http.StatusOK, gin.H{"module": module})
}

func (app *application) updateModuleHandler(c *gin.Context) {
	id, err := app.readIDParam(c)
	if err != nil {
		app.notFoundResponse(c)
		return
	}

	var input struct {
		Title string `json:"title"`
	}

	err = c.BindJSON(&input)
	if err != nil {
		app.badRequestResponse(c, err)
		return
	}

	module, err := app.models.Modules.Get(id)
	if err != nil {
		app.notFoundResponse(c)
		return
	}

	module.Title = input.Title

	err = app.models.Modules.Update(module)
	if err != nil {
		app.serverErrorResponse(c, err)
		return
	}

	app.writeJSON(c, http.StatusOK, gin.H{"module": module})
}

func (app *application) deleteModuleHandler(c *gin.Context) {
	id, err := app.readIDParam(c)
	if err != nil {
		app.notFoundResponse(c)
		return
	}

	err = app.models.Modules.Delete(id)
	if err != nil {
		app.serverErrorResponse(c, err)
		return
	}

	c.Status(http.StatusNoContent)
}
