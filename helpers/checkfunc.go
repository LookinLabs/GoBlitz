package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CheckFunc func() (bool, error)

type Check struct {
	Func         CheckFunc
	ErrorMessage string
	ShouldExist  bool
}

func NewCheck(f CheckFunc, errorMessage string, shouldExist bool) Check {
	return Check{
		Func:         f,
		ErrorMessage: errorMessage,
		ShouldExist:  shouldExist,
	}
}

func CheckAndRespond(c *gin.Context, checks []Check) bool {
	for _, check := range checks {
		exists, err := check.Func()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return true
		}

		if exists != check.ShouldExist {
			c.JSON(http.StatusBadRequest, gin.H{"error": check.ErrorMessage})
			return true
		}
	}

	return false
}
