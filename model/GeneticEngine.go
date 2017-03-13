package model

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MUTATION_RATE float32= 0.01
	CROSSOVER_RATE float32= 1.0
	GENERATIONS int= 10
	ALLELE int = 5
)

var numOfSections int
var AugmentTable [][]int
var fitness []int
var population_size int= 5
var CurrPop [][]int
var days int
var periods int	// per day (so in week= days x periods)

func GeneticEngineParameters(numSections, pop_size, day, period int) {
	numOfSections= numSections
	population_size= pop_size
	days= day
	periods= period
}

func CreateAugmentTable() {
	// we need an augTable as [numOfSections][each section's day1, day2, day3...]
	// then we convert the augTable to have just one row
	// like augTable1= {sec1}day1..day2..day3..{sec2}day1..day2..day3..
	// make n augTables to represent a population.

	AugmentTable= make([][]int, numOfSections)
	fitness= make([]int, population_size)

	for i:=0; i<len(Table); i++ {
		AugmentTable[i]= make([]int, 0)
	}

	for i,_ := range AugmentTable {
		for j,_ := range Table {
			AugmentTable[i]= append(AugmentTable[i], Table[j]...)
		}
	}

	// till line 24, augTable looks like:
	//	[sec1]..day1..day2..day3
	//	[sec2]..day1..day2..day3
	//	[sec3]..day1..day2..day3
	//
	// but now we change all those sections in one row of augTable
	// to get 1 ORGANISM which would look like
	// augTable [org1]..<sec1>..day1..day2..day3 <sec2> day1..day2..day3 <sec3>..day1..day2..day3
	// augTable [org2]..<sec1>..day1..day2..day3 <sec2> day1..day2..day3 <sec3>..day1..day2..day3
	// augTable [org3]..<sec1>..day1..day2..day3 <sec2> day1..day2..day3 <sec3>..day1..day2..day3
	//
	// these 3 orgs make up 1 POPULATION

	for i:= 1; i< numOfSections; i++ {
		AugmentTable[0]= append(AugmentTable[0], AugmentTable[i]...)
	}

	// convert AugmentTable to have only one row with data of 3 sections
	// each AugmentTable now represents an Organism
	AugmentTable= AugmentTable[0:1]		
	// Now, create the population of many such AugmentTable/ Organism
}

func CreatePopulation() {
	rand.Seed(time.Now().UnixNano())
	CurrPop= make([][]int, population_size)

	for i,_ := range CurrPop {
		CurrPop[i]= make([]int, len(AugmentTable[0]))
	}

	org_culture:= AugmentTable[0]

	copy(CurrPop[0], org_culture)

	initRandomOrgs:= func () {
		// starts from 1, since CurrPop[0] is the Organism Culture sample, from which 
		// other orgs will be initialized
		for org:=1; org < len(CurrPop); org++ {
			for chromosome,_ := range CurrPop[org] {
				randChromosome:= rand.Intn(len(org_culture))
				CurrPop[org][chromosome]= org_culture[randChromosome]
			}
		}
	}

	initRandomOrgs()

	for i,_ := range CurrPop {
		for j,_ := range CurrPop[i] {
			fmt.Printf("%d ", CurrPop[i][j])
		}
		fmt.Println()
	}
}

func calcFitness() {
	// fitness based on:
	// 1)	In each org, if we have 3 sections, then a subject should
	// 	  	appear only 'credits' times in that section. Then reset CreditCounter, and check
	//	  	for same in other 2 sections of that org only
	// 2)	In each org, each sections' subject1's position must be different. (no coinciding of sub)
	CreditCount.ResetCreditCounter()
	// resetting fitness values too
	for i,_:= range fitness{
		fitness[i]=0
	}

	for i,_ := range CurrPop {
		week_periods:= 1 	// each section in 1 org must have this many chromosome
		// 1) Credits validation
		for j,_ := range CurrPop[i] {

			week_periods++

			if CurrPop[i][j]!=0 {
					sid:= GetSidFromAugment(CurrPop[i][j])
					CreditCount.Sub[sid-1].Credits-=1
					//fmt.Printf(" sid: %d has credits %d \n", sid, CreditCount.Sub[sid-1].Credits)
					if CreditCount.Sub[sid-1].Credits >=0 {
						fitness[i]+=1
					} else {
						fitness[i]-=1
					}
				}

			if week_periods> days * periods {
				week_periods=1
				CreditCount.ResetCreditCounter()
				//fmt.Printf("\n\nCounter Resetted, and week_periods= %d\n\n", week_periods)
			}
		}
	}
	fmt.Println("Credit fit:", fitness)
	// 2) Uniqueness checker
	// WORK ON THIS
	for i,_ := range CurrPop {
		for j:=0; j< len(CurrPop[i])- days* periods; j++ {
			if CurrPop[i][j]!=0 {
				if CurrPop[i][j]!= CurrPop[i][j+ days* periods] {
					fitness[i]+=1
				} else {
					fitness[i]-=1
				}
			}
		}
	}

	fmt.Println("Uniqueness fit:", fitness)
}

func bestParentSelection() (parentIndex1, parentIndex2 int) {
	fitness_copy:= make([]int, len(fitness))
	copy(fitness_copy, fitness)
	// sort
	for i:= 0; i< len(fitness_copy)-1; i++ {
		for j:=0; j<len(fitness_copy)-i-1; j++ {

			if fitness_copy[j] < fitness_copy[j+1] {
				fitness_copy[j], fitness_copy[j+1]= fitness_copy[j+1], fitness_copy[j]
			}
		}
	}

	fmt.Println(fitness_copy)

	for i,_ := range fitness_copy {
		if fitness[i]== fitness_copy[0] {
			parentIndex1= i;
		} else if fitness[i]== fitness_copy[1] {
			parentIndex2= i;
		}
	} 

	return 	//returns both indices, golang is so nice, right?
}

func newGeneration() {			
	genes:= len(CurrPop[0])
	nextPopulation:= make([][]int, len(CurrPop))

	for i,_:= range nextPopulation {
		nextPopulation[i]= make([]int, len(CurrPop[0]))
	}

	copy(nextPopulation, CurrPop)

	for org,_:= range CurrPop {

		parent1, parent2:= bestParentSelection()
		fmt.Println("Parent1:", parent1)
		fmt.Println("Parent2:", parent2)

		crossover_probability:= rand.Intn((int)(1/CROSSOVER_RATE)) 
		crossover_point:= rand.Intn(genes)
		fmt.Println("crossover_probabilty: ",crossover_probability)

		for gene:=0; gene<genes; gene++{

			mutation:= rand.Intn((int)(1/MUTATION_RATE)) //rand.Float32()* MUTATION_RATE
			//fmt.Println("\tGene value: ", CurrPop[org][gene])

			if mutation==0 {
				fmt.Printf("\nMutated gene %d with new allele ", gene)
				nextPopulation[org][gene]= rand.Intn(ALLELE)+1
				//fmt.Printf("%d\n\n", nextPopulation[org].chromosome[gene])
			}

			if crossover_probability==0 {
				if gene <= crossover_point {
					nextPopulation[org][gene]= CurrPop[parent1][gene]
				}else {
					nextPopulation[org][gene]= CurrPop[parent2][gene]
				}
			}
		}
		fmt.Println("Crossed chromosome at point", crossover_point)
		//fmt.Println("New:", nextPopulation[org])
	}

	copy(CurrPop, nextPopulation)
	calcFitness()
}

func GenerationHandler () {
	current_gen:=1

		for current_gen<= GENERATIONS {
			fmt.Println("GENERATION ", current_gen)
			newGeneration()
			current_gen++
		}

	for i,_ := range CurrPop {
		for j,_ := range CurrPop[i] {
			fmt.Printf("%d ", CurrPop[i][j])
		}
		fmt.Println()
	}
}