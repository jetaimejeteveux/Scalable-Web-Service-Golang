package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	cron2 "github.com/robfig/cron"
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

type DisasterIndex struct {
	WaterValue  int    `json:"water_value"`
	WaterStatus string `json:"water_status"`
	WindValue   int    `json:"wind_value"`
	WindStatus  string `json:"wind_status"`
}

func randomIntGenerator(start, end int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(end-start) + start
}
func createJSON() {
	var disasterIndex DisasterIndex
	disasterIndex.WaterValue = randomIntGenerator(1, 100)
	disasterIndex.WindValue = randomIntGenerator(1, 100)
	switch {
	case disasterIndex.WaterValue <= 5:
		disasterIndex.WaterStatus = "Aman"
	case disasterIndex.WaterValue > 5 && disasterIndex.WaterValue <= 8:
		disasterIndex.WaterStatus = "Siaga"
	case disasterIndex.WaterValue > 8:
		disasterIndex.WaterStatus = "Bahaya"
	default:
		disasterIndex.WaterStatus = "Error"
	}
	switch {
	case disasterIndex.WindValue <= 5:
		disasterIndex.WindStatus = "Aman"
	case disasterIndex.WindValue > 5 && disasterIndex.WindValue <= 8:
		disasterIndex.WindStatus = "Siaga"
	case disasterIndex.WindValue > 8:
		disasterIndex.WindStatus = "Bahaya"
	default:
		disasterIndex.WindStatus = "Error"
	}

	router := gin.Default()
	router.GET("/status", func(c *gin.Context) {
		var (
			result gin.H
			w      http.ResponseWriter = c.Writer
		)
		result = gin.H{
			"Status": disasterIndex,
		}
		c.JSON(http.StatusOK, result)
		var t, err = template.ParseFiles("disaster.html")
		if err != nil {
			fmt.Println(err)
			return
		}

		t.Execute(w, disasterIndex)
	})
	fmt.Println("running at localhost:8080")
	http.ListenAndServe(":8080", router)
}
func main() {
	cron := cron2.New()
	cron.AddFunc("@every 15s", createJSON)
	cron.Start()

}
