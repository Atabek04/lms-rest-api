package main

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (app *application) readIDParam(c *gin.Context) (uint, error) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		return 0, errors.New("invalid ID parameter")
	}
	return uint(id), nil
}
