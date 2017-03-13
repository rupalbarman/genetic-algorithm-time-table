package model

import (
	"fmt"
	"math/rand"
)

//================================================================
// Model

type Subject struct {
	Sid int			//1,2 etc
	Name string
	Credits int		//per week
}

type Teacher struct {
	Tid int			//10,11 etc
	Name string
	TakenSubs []int	//max is 3
}

type CreditCounter struct {		//keeps count of how the subjects appearing per week per section
	Sub []Subject
}

func(cc *CreditCounter) ResetCreditCounter() {
	// After using CreditCounter in checking whether a Faculty is allowed to pick a particular sub or not
	// ie., he/she only picks the sub if it has remainiing credits.
	// Another use is that we can use this CreditCounter to evaluate Fitness of Organisms later
	copy(cc.Sub,Subjects)
	fmt.Println(cc.Sub)
}

//================================================================
//Globals

var Teachers []Teacher
var Subjects []Subject
var Table [][]int 	//3x5 AugmentID matrix
var CreditCount CreditCounter

//=======================================================================
// Driver functions

func AssignSubjects() {
	//Augment id in TakenSubs looks like 113 ie, Tid is 11, and subject is 3

	for i,_ := range Teachers {

		//for random subject 1
		check1:
		randSub:= rand.Intn(len(Subjects))+1
		if CreditCount.Sub[randSub-1].Credits>0 {
			Teachers[i].TakenSubs[0]= Teachers[i].Tid*10 + randSub
			CreditCount.Sub[randSub-1].Credits-=1
		} else {
			goto check1
		}

		//for random subject 2
		check2:
		randSub= rand.Intn(len(Subjects))+1

		if CreditCount.Sub[randSub-1].Credits>0 && randSub!= GetSidFromAugment(Teachers[i].TakenSubs[0]) {
			Teachers[i].TakenSubs[1]= Teachers[i].Tid*10 + randSub
			CreditCount.Sub[randSub-1].Credits-=1
		} else {
			goto check2
		}
	}
}
//========================================================================

func AssignInitTable() {
	for i,_ := range Teachers {
		for j,_:= range Teachers[i].TakenSubs {
			subject:= GetSidFromAugment(Teachers[i].TakenSubs[j])
			credit:= Subjects[subject-1].Credits
			fmt.Printf("%d takenSub %d subject %d credit %d\n", Teachers[i].Tid, Teachers[i].TakenSubs[j], subject, credit)

			check:
			x:= rand.Intn(len(Table))
			y:= rand.Intn(len(Table[0]))
			if Table[x][y]==0 {
				Table[x][y]= Teachers[i].TakenSubs[j]
			} else {
				goto check
			}
		}
	}

}

func GetSidFromAugment(augment int) int {
	return augment % 10
}
