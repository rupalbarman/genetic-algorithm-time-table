package view

import(
	"fmt"
	"github.com/rupalbarman/genetic-algorithm-time-table/model"
)

//=====================================================================
// Displays

func DisplayTeachers() {
	for i,_:= range model.Teachers {
		fmt.Println(model.Teachers[i].Tid, model.Teachers[i].Name, model.Teachers[i].TakenSubs)
	}
}

func DisplaySubjects() {
	for i,_:= range model.Subjects {
		fmt.Println(model.Subjects[i].Sid, model.Subjects[i].Name, model.Subjects[i].Credits)
	}
}

func DisplayTable() {
	for i:=0;i<len(model.Table);i++ {
		for j:=0;j<len(model.Table[i]);j++ {
			fmt.Printf(" %d \t", model.Table[i][j])
		}
		fmt.Printf("\n")
	}
}

func DisplayAugmentTable() {
	for i:=0;i<len(model.AugmentTable);i++ {
		for j:=0;j<len(model.AugmentTable[i]);j++ {
			fmt.Printf(" %d \t", model.AugmentTable[i][j])
		}
		fmt.Printf("\n")
	}
}

func DisplayRemainingCredits() {
	for i,_:= range model.CreditCount.Sub {
		fmt.Println(model.CreditCount.Sub[i].Sid, model.CreditCount.Sub[i].Name, model.CreditCount.Sub[i].Credits)
	}
}
//===================================================================================================
