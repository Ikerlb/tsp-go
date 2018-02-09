package tsp

import (
	"fmt"
	"math"
	"math/rand"
)

type Annealer struct{
	Rng				*rand.Rand
	Path 	  		[]int
	Dists  			[][]float64
	Phi  	  		float64
	BatchSize 		int
	InitTemp  		float64
	EpsilonT  		float64
	EpsilonP  		float64
	AcceptedPercent float64
	InitialTempN	int
	// CurrentSolution *Solution
}


func NewAnnealer(path []int,dists [][]float64,phi,initTemp, epsilonT, epsilonP, acceptedPercent float64, batchSize, initialTempN int) *Annealer {
	annealer:=Annealer{}
	annealer.Path=path
	annealer.Dists=dists
	annealer.Phi=phi
	annealer.BatchSize=batchSize
	annealer.InitTemp=initTemp
	annealer.EpsilonT=epsilonT
	annealer.EpsilonP=epsilonP
	annealer.AcceptedPercent=acceptedPercent
	annealer.InitialTempN=initialTempN
	return &annealer
}

func (annealer *Annealer) AnnealWithSeed(seed int64){
	annealer.Rng = rand.New(rand.NewSource(seed))
	//Shuffle array.
	//fmt.Println("Path to anneal is",annealer.Path)
	newPath := make([]int, len(annealer.Path))
	perm := annealer.Rng.Perm(len(annealer.Path))
	for i, v := range perm {
	    newPath[v] = annealer.Path[i]
	}
	//fmt.Println("Shuffled path is",newPath)
	//Given seed initialize random current solution
	s:=NewSolution(newPath,annealer.Dists)
	t:=annealer.initialTemperature(s,annealer.InitTemp)
	//fmt.Printf("Starting tresholdAccept with temp %f and initial solution: %v",t,s)
	annealer.tresholdAccept(t,s)
}

func (annealer *Annealer) tresholdAccept(t float64, s *Solution){
	//InitTemp=initialTemperature()
	var p float64
	for t>annealer.EpsilonT {
		q:=math.Inf(0)
		for p<=q {
			q=p
			p,s=annealer.calculateBatch(t,s)
		}
		//fmt.Printf("Decreasing temperature t: %f by %f: %f\n",t,annealer.Phi,annealer.Phi*t)
		t*=annealer.Phi
	}
	if s.IsFeasible()&&s.Cost<0.9{
		fmt.Println(s)
	}
}

func (annealer *Annealer) calculateBatch(t float64, sol *Solution) (float64,*Solution){
	var r,sPrimeCost float64
	c:=0
	var s *Solution
	s=sol
	//TODO implement max number of tries
	for c<annealer.BatchSize {
		i,j:=generateRandomIdx(len(annealer.Path),annealer.Rng)
		sPrimeCost=s.PeekNeighborCost(i,j)
		if sPrimeCost<=(s.Cost+t) {
			s=s.Neighbor(i,j,sPrimeCost)
			c++
			r+=sPrimeCost
		}
	}
	return (r/float64(annealer.BatchSize)),s
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