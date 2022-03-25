package appError

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type appHandler func (c *gin.Context) error

func Middleware(handler appHandler) func (c *gin.Context){
	return func (c *gin.Context){
		err:=handler(c)
		var appErr *AppError
		if err!=nil{
			if errors.As(err,&appErr){
				if errors.Is(err, ErrOrderNotFound) {
					c.HTML(http.StatusNotFound,"ErrOrderNotFound.html",nil)
					return
				}
				appErr=err.(*AppError)
				c.AbortWithStatusJSON(http.StatusBadRequest,appErr)
				return
			}
			c.AbortWithStatusJSON(http.StatusTeapot, systemError(err))
		}
	}
}