package resthandler

import "github.com/gin-gonic/gin"

type PeriodResthandler interface {
	CreatePeriod(c *gin.Context)
	FindAllPeriod(c *gin.Context)
	FindAllPeriodByStudent(c *gin.Context)
	FindPeriodById(c *gin.Context)
	UpdatePeriod(c *gin.Context)
	DeletePeriodById(c *gin.Context)
}
