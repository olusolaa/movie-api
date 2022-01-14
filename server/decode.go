package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/olusolaa/movieApi/servererrors"
	"github.com/pkg/errors"
	"log"
)

// decode the body of c into v
func (s *Server) decode(c *gin.Context, v interface{}) []string {
	if err := c.ShouldBindJSON(v); err != nil {
		errs := []string{}
		verr, ok := err.(validator.ValidationErrors)
		if ok {
			for _, fieldErr := range verr {
				errs = append(errs, servererrors.NewFieldError(fieldErr))
			}
		} else {
			log.Println(errors.Wrap(err, "could not decode body for request"))
			errs = append(errs, "internal server error")
		}
		return errs
	}
	return nil
}
