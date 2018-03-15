package tsp

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	//"log"
)

type SolutionLite struct{
	Path	 string
	Cost	 float64
	Feasible string
}

type Annealer struct{
	Rng			*rand.Rand
	Path			[]int
	Dists			[][]float64
	Phi			float64
	BatchSize		int
	InitTemp		float64
	EpsilonT		float64
	EpsilonP		float64
	AcceptedPercent		float64
	InitialTempN		int
	BestSolution		*SolutionLite
	Sweeping		bool
	VeryVerbose		bool
	BatchLimit		int
	// CurrentSolution *Solution
}

func NewAnnealer(path []int,dists [][]float64,phi,initTemp, epsilonT, epsilonP, acceptedPercent float64, batchSize, initialTempN, batchLimit int, sweeping, veryverbose bool) *Annealer {
	annealer:=Annealer{}
	annealer.Path=path
	annealer.Dists=dists
	annealer.Phi=phi
	annealer.BatchSize=batchSize
	annealer.BatchLimit=batchLimit
	annealer.InitTemp=initTemp
	annealer.EpsilonT=epsilonT
	annealer.EpsilonP=epsilonP
	annealer.AcceptedPercent=acceptedPercent
	annealer.InitialTempN=initialTempN
	annealer.BestSolution=&SolutionLite{Cost:math.Inf(0)}
	annealer.Sweeping=sweeping
	annealer.VeryVerbose=veryverbose
	return &annealer
}


//TODO: Implement logging mechanism eg verbose flag -v very verbose flag -vv
func (annealer *Annealer) AnnealWithSeed(seed int64){
	annealer.Rng = rand.New(rand.NewSource(seed))

	//Shuffle array.
	newPath := make([]int, len(annealer.Path))
	perm := annealer.Rng.Perm(len(annealer.Path))
	for i, v := range perm {
	    newPath[v] = annealer.Path[i]
	}

	//Given seed initialize random current solution
	s:=NewSolution(newPath,annealer.Dists)
	t:=annealer.initialTemperature(s,annealer.InitTemp)
	//fmt.Printf("Starting tresholdAccept with temp %f and initial solution: %v",t,s)
	annealer.tresholdAccept(t,s,seed)
}

func (annealer *Annealer) tresholdAccept(t float64, s *Solution, seed int64){
	annealer.BestSolution.Cost=math.Inf(0)
	annealer.BestSolution.Path=""
	annealer.BestSolution.Feasible=""
	var p float64
	stop:=false
	//Sweep initial solution,
	if annealer.Sweeping{
		annealer.sweepSolution(s)
	}
	for t>annealer.EpsilonT&&(!stop) {
		q:=math.Inf(0)
		for p<=q&&(!stop) {
			q=p
			p,s,stop=annealer.calculateBatch(t,s)
		}
		//fmt.Printf("Decreasing temperature t: %f by %f: %f\n",t,annealer.Phi,annealer.Phi*t)
		t*=annealer.Phi
	}

	//Sweep for the last time.
	if annealer.Sweeping{
		annealer.sweepSolution(s)
	}
	fmt.Printf("\nSeed: %d, Solution: %s\n",seed,annealer.BestSolution)
}

func (annealer *Annealer) calculateBatch(t float64, sol *Solution) (float64,*Solution,bool){
	var r,sPrimeCost float64
	c:=0
	var s *Solution
	s=sol
	k:=0
	//for k=0;c<annealer.BatchSize&&k<annealer.BatchSize*annealer.BatchSize;k++ {
	for k=0;c<annealer.BatchSize&&k<annealer.BatchSize*annealer.BatchLimit;k++ {
		i,j:=generateRandomIdx(len(annealer.Path),annealer.Rng)
		sPrimeCost=s.PeekNeighborCost(i,j)
		if sPrimeCost<=(s.Cost+t) {
			s=s.Neighbor(i,j,sPrimeCost)
			c++
			r+=sPrimeCost
			if annealer.VeryVerbose {
				if sPrimeCost<annealer.BestSolution.Cost {
					//annealer.GraphCosts=append(annealer.GraphCosts,-sPrimeCost)
					fmt.Printf("E: %f\n",sPrimeCost)
				} else {
					//annealer.GraphCosts=append(annealer.GraphCosts,sPrimeCost)
					fmt.Printf("E: %f\n",sPrimeCost)
				}
			}
			if sPrimeCost<annealer.BestSolution.Cost {
				annealer.BestSolution=NewSolutionLite(s)
				if annealer.Sweeping {
					//fmt.Printf("Sweeping with k=%d...\n",k)
					annealer.sweepSolution(s)
				}
			}
		}
	}
	return (r/float64(annealer.BatchSize)),s,k>=(annealer.BatchSize*annealer.BatchLimit)-1
	//return (r/float64(annealer.BatchSize)),s,k>=(annealer.BatchSize*annealer.BatchSize)-1
}

//ask the following!!
//Are you to update the solution with the neighbor when given a better one? doesn't that conflict with what the sweeping is supposed to do?
func (annealer *Annealer) sweepSolution(sol *Solution){
	size:=len(annealer.Path)
	var sPrimeCost float64
	for i:=0;i<size;i++{
		for j:=i+1;j<size;j++{
			sPrimeCost=sol.PeekNeighborCost(i,j)
			if sPrimeCost < annealer.BestSolution.Cost {
				sol=sol.Neighbor(i,j,sPrimeCost)
				annealer.BestSolution=NewSolutionLite(sol)
				if annealer.VeryVerbose {
					fmt.Printf("E: %f\n",sPrimeCost)
				}
			}
		}
	}
}

func (annealer *Annealer) initialTemperature(sol *Solution,t float64) float64 {
	var t1,t2 float64
	//fmt.Println("Going into acceptedPercent with P=",annealer.AcceptedPercent)
	p:=annealer.acceptedPercent(sol,t)
	//fmt.Printf("Returned from acceptedPercent with p=%f\n",p)
	if math.Abs(annealer.AcceptedPercent-p) <= annealer.EpsilonP {
		return t
	}
	if p<annealer.AcceptedPercent {
		for p<annealer.AcceptedPercent {
			t*=2
			//fmt.Printf("Multiplying T by 2: Current Temperature is: %f ... Entering acceptedPercent again\n",t)
			p=annealer.acceptedPercent(sol,t)
		}
		t1=t/2
		t2=t
		//fmt.Printf("Entered case 1, t1: %f, t2: %f, t: %f\n",t1,t2,t)
	} else {
		for p>annealer.AcceptedPercent {
			t/=2
			//fmt.Printf("Diving T by 2: Current Temperature is: %f ... Entering acceptedPercent again\n",t)
			p=annealer.acceptedPercent(sol,t)
		}
		t1=t
		t2=2*t
		//fmt.Printf("Entered case 2, t1: %f, t2: %f, t: %f\n",t1,t2,t)
	}
	return annealer.binarySearch(sol,t1,t2)
}

func (annealer *Annealer) acceptedPercent(sol *Solution, t float64) float64{
	var sPrimeCost float64
	c:=0
	var s *Solution
	s=sol
	for k := 0; k < annealer.InitialTempN; k++ {
		i,j:=generateRandomIdx(len(annealer.Path),annealer.Rng)
		sPrimeCost=s.PeekNeighborCost(i,j)
		if sPrimeCost<=(s.Cost+t) {
			//fmt.Printf("On %dth iteration we got random indexes (%d,%d). Solution is %v and neighbor solution cost is: %f\n",k,i,j,s,sPrimeCost)
			s=s.Neighbor(i,j,sPrimeCost)
			c++
		}
	}
	return float64(c)/float64(annealer.InitialTempN)
}

func (annealer *Annealer) binarySearch(sol *Solution,t1, t2 float64) float64{
	tm:=(t1+t2)/2.0
	if t2-t1 < annealer.EpsilonT {
		return tm
	}
	p:=annealer.acceptedPercent(sol,tm)
	if math.Abs(annealer.AcceptedPercent-p) < annealer.EpsilonP {
		return tm
	}
	if p>annealer.AcceptedPercent{
		return annealer.binarySearch(sol,t1,tm)
	}
	return annealer.binarySearch(sol,tm,t2)
}

func generateRandomIdx(size int,rng *rand.Rand) (int,int){
	i,j:=rng.Intn(size),rng.Intn(size)
	for i==j {
		i=rng.Intn(size)
		j=rng.Intn(size)
	}
	return i,j
}

func NewSolutionLite(sol *Solution) *SolutionLite{
	solLite:=SolutionLite{}
	solLite.Path=strings.Replace(fmt.Sprintf("%v",sol.Path)," ",",",-1)
	solLite.Cost=sol.Cost
	solLite.Feasible=fmt.Sprintf("%v",sol.Feasible)
	return &solLite
}

func (sol *SolutionLite) String() string{
	return fmt.Sprintf("Path: %s, Cost: %f, Feasible: %s",sol.Path,sol.Cost,sol.Feasible)
}
