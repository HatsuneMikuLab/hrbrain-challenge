package services

import (
	"math/rand"
	"time"
)

type IEvaluationService interface {
	GenEvaluation() string
}

type evaluationService struct {}

var evaluationTable = []string{ "F", "E", "D", "C", "B", "A" }

func NewEvaluationService() *evaluationService {
	return &evaluationService{}
}

func (es *evaluationService) GenEvaluation() string {
	// simulate some work to make sure that we can see how caching works
	time.Sleep(500 * time.Millisecond) 
	randomNum := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(6)
	return evaluationTable[randomNum]
}