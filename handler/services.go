package handler

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getPathParamUint(c *gin.Context, key string) (*uint, error) {
	par := c.Param(key)
	fmt.Println(par)
	uint64Param, err := strconv.ParseUint(par, 10, 64)
	if err != nil {
		return nil, err
	}
	uintParam := uint(uint64Param)
	return &uintParam, nil
}
