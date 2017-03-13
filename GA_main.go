package main 

import (
	"github.com/rupalbarman/genetic-algorithm-time-table/model"
	"github.com/rupalbarman/genetic-algorithm-time-table/view"

	"math/rand"
	"time"
)

var days int= 3
var periods int = 5	// per day (so in week= days x periods)

//================================================================
// Memory allocation

func init() {
	rand.Seed(time.Now().UnixNano())

	model.Teachers= make([]model.Teacher, 4)
	model.Subjects= make([]model.Subject, 5)
	model.Table= make([][]int, days)

	model.CreditCount= model.CreditCounter{}
	model.CreditCount.Sub= make([]model.Subject, 5)

	for i:=0;i<len(model.Table);i++ {
		model.Table[i]= make([]int, periods)
	}

	for i,_:= range model.Subjects {
		model.Subjects[i].Sid= i+1
		model.Subjects[i].Credits= rand.Intn(4)+1		//so that credits is never 0
	}

	copy(model.CreditCount.Sub, model.Subjects)

	for i,_:= range model.Teachers {
		model.Teachers[i].Tid= 10+i
		model.Teachers[i].TakenSubs= make([]int, 2)	//takes 2 subjects per teacher
	}
}

func main() {

	model.GeneticEngineParameters(3, 5, days, periods)

	view.DisplaySubjects()
	model.AssignSubjects()
	view.DisplayTeachers()

	model.AssignInitTable()
	view.DisplayTable()
	view.DisplayRemainingCredits()
	view.DisplayAugmentTable()	//this one is the orginal AugmentTable with [numOfsections][day1,day2] form
	model.CreateAugmentTable()
	model.CreatePopulation()
	//calcFitness()
	//displayAugmentTable()	//this one is the [orgCulture][numOfSections with each sections day data]
	model.GenerationHandler()
	
}