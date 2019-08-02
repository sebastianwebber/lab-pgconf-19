package web

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/sebastianwebber/app/model"
)

type loadInput struct {
	TotalDepartments  int  `json:"totalDepartments"`
	DeleteBeforeStart bool `json:"deleteBeforeStart"`
}

func displayError(c *gin.Context, err error) {
	if err != nil {
		c.JSON(400, gin.H{
			"ERROR": err,
		})
	}
}

func loadData(c *gin.Context) {

	input := loadInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		displayError(c, fmt.Errorf("Invalid json: %v", err))
		return
	}

	log.Printf("input: %#v\n", input)

	if err := populateDept(input.TotalDepartments, input.DeleteBeforeStart); err != nil {
		displayError(c, fmt.Errorf("Error processing dept: %v", err))
		return
	}

	c.JSON(200, gin.H{
		"OK": "I guess that is everything ok.",
	})
}

var (
	populateDeptOpsCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
	populateDeptErrCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_err_total",
		Help: "The total number of error events",
	})
)

func populateDept(maxRows int, deleteBefore bool) error {
	defer populateDeptOpsCounter.Inc()

	itens := []model.Department{}

	if deleteBefore {

		_, err := dbConn.Model(&itens).Where("id >= 0").Delete()

		if err != nil {
			populateDeptErrCounter.Inc()
			return fmt.Errorf("Could not delete rows: %v", err)
		}
	}

	for i := 0; i < maxRows; i++ {
		itens = append(itens, model.Department{
			Name: fmt.Sprintf("Department %d", i),
		})
	}

	err := dbConn.Insert(&itens)
	if err != nil {
		populateDeptErrCounter.Inc()
		return fmt.Errorf("Could not insert rows: %v", err)
	}

	return nil
}
