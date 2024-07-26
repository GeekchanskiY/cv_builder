package usecases

import (
	"log"
	"time"
)

func BuildCV(employeeID, vacancyID int, cvChan chan int) {
	log.Println("Building CV")
	cvChan <- 1
	time.Sleep(5 * time.Second)
	log.Println("CV Build Process finished")
}
