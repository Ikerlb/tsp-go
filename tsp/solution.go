package tsp

import (
	"fmt"
	//"math/rand"
	"github.com/willf/bitset"
)

const discFactor float64 = 100.0 

type Solution struct{
	avgDist	 float64
	maxDist	 float64
	path     []int
	Cost     float64
	Feasible *bitset.BitSet
	dists	 [][]float64
}


func NewSolution(path []int, dists [][]float64) *Solution{
	sol:=Solution{}
	sol.path=path
	sol.dists=dists
	sol.Feasible=bitset.New(uint(len(path)))
	//fmt.Println(sol.path)
	sol.calculateMaxAvg()
	sol.calculateInitialCost()
	return &sol
}


func (sol *Solution) calculateMaxAvg(){
	var max, sum float64
	var c int
	for i := 0; i < len(sol.path); i++ {
		for j := i+1; j < len(sol.path); j++ {
			d:=sol.dists[sol.path[i]][sol.path[j]]
			if d > 0.0{
				if d>max {
					max=d
				}
				sum+=d
				c++
			}
		}
	}
	sol.maxDist=max*discFactor;
	sol.avgDist=sum/(float64(c))
}


func (sol *Solution) calculateInitialCost(){
	var sum float64
	size:=len(sol.path)-1
	for i := 0; i < size; i++ {
		//if path[i] isn't connected to path[i+1] we set feasible[i]
		if sol.dists[sol.path[i]][sol.path[i+1]] == 0 {
			sol.Feasible.Set(uint(i))
		}
		sum+=sol.augmentedWeight(i,i+1)
	}
	sol.Cost=sum/(sol.avgDist*float64(size))
}


func (sol *Solution) augmentedWeight(i, j int) (dist float64){
	id1,id2:=sol.path[i],sol.path[j]
	dist=sol.dists[id1][id2]
	if dist==0 {
		dist=sol.maxDist
	}
	return
}

//peak what would the neighbor solution would be. this is to only create a full new neighbor if necessary (avoids creating new slice, bitset, etc).
func (sol *Solution) PeekNeighborCost(i,j int) float64 {
	size:=len(sol.path)-1
	previousSum:=sol.Cost*sol.avgDist*(float64(size))
	if Abs(i-j) == 1 {
		low:=Min(i,j)
		high:=Max(i,j)

		if low!=0 {
			previousSum-=sol.augmentedWeight(low-1,low)
			previousSum+=sol.augmentedWeight(low-1,high)
		}
		if high!=size {
			previousSum-=sol.augmentedWeight(high,high+1)
			previousSum+=sol.augmentedWeight(low,high+1)
		}
	} else {
	//fmt.Printf("previousSum is %f with i=%d and j=%d, path[i]=%d and path[j]=%d, maxDistance is %f and avgDist is %f\n",previousSum,i,j,sol.path[i],sol.path[j],sol.maxDist,sol.avgDist)
		if i!=0{
			previousSum-=sol.augmentedWeight(i-1,i)
			//fmt.Printf("Substracting distance: %f from indexes %d and %d, giving as a result: %f\n",sol.augmentedWeight(i-1,i),sol.path[i-1],sol.path[i], previousSum)
			previousSum+=sol.augmentedWeight(i-1,j)
			//fmt.Printf("Adding distance: %f from indexes %d and %d, giving as a result: %f\n",sol.augmentedWeight(i-1,j),sol.path[i-1],sol.path[j],previousSum)
		}
		if i!=size{
			previousSum-=sol.augmentedWeight(i,i+1)
			//fmt.Printf("Substracting distance: %f from indexes %d and %d, giving as a result: %f\n",sol.augmentedWeight(i,i+1),sol.path[i],sol.path[i+1], previousSum)
			previousSum+=sol.augmentedWeight(j,i+1)
			//fmt.Printf("Adding distance: %f from indexes %d and %d, giving as a result: %f\n",sol.augmentedWeight(j,i+1),sol.path[j],sol.path[i+1],previousSum)		
		}
		if j!=0 {
			previousSum-=sol.augmentedWeight(j-1,j)
			//fmt.Printf("Substracting distance: %f from indexes %d and %d, giving as a result: %f\n",sol.augmentedWeight(j-1,j),sol.path[j-1],sol.path[j],previousSum)
			previousSum+=sol.augmentedWeight(j-1,i)
			//fmt.Printf("Adding distance: %f from indexes %d and %d, giving as a result: %f\n",sol.augmentedWeight(j-1,i),sol.path[j-1],sol.path[i],previousSum)
		}
		if j!=size{
			previousSum-=sol.augmentedWeight(j,j+1)
			//fmt.Printf("Substracting distance: %f from indexes %d and %d, giving as a result: %f\n",sol.augmentedWeight(j,j+1),sol.path[j],sol.path[j+1],previousSum)
			previousSum+=sol.augmentedWeight(i,j+1)
			//fmt.Printf("Adding distance: %f from indexes %d and %d, giving as a result: %f\n",sol.augmentedWeight(i,j+1),sol.path[i],sol.path[j+1],previousSum)
		}
	}
	//fmt.Printf("Constant time cost is: %f20\n",previousSum/(sol.avgDist*(float64(size))))
	return previousSum/(sol.avgDist*(float64(size)))

}

func (sol *Solution) calculateNeighborFeas(i,j int){
	//sol.Feasible:=sol.Feasible.Clone()
	size:=len(sol.path)-1
	if i!=0 {
		//if i-1 to j is disconnected, set i-1
		if sol.dists[sol.path[i-1]][sol.path[j]]==0 {
			sol.Feasible.Set(uint(i-1))
		} else {
			sol.Feasible.Clear(uint(i-1))
		}
	}
	if i!=size {
		if sol.dists[sol.path[j]][sol.path[i+1]]==0 {
			sol.Feasible.Set(uint(j))
		} else {
			sol.Feasible.Clear(uint(j))
		}
	}
	if j!=0 {
		if sol.dists[sol.path[j-1]][sol.path[i]]==0 {
			sol.Feasible.Set(uint(j-1))
		} else {
			sol.Feasible.Clear(uint(j-1))
		}
	}
	if j!=size {
		if sol.dists[sol.path[i]][sol.path[j+1]]==0 {
			sol.Feasible.Set(uint(i))
		} else {
			sol.Feasible.Clear(uint(i))
		}
	}
}


func (sol *Solution) swap(i,j int){
	tmp:=sol.path[i]
	sol.path[i]=sol.path[j]
	sol.path[j]=tmp
}

//FIXX!!!!!!!
//pass two indexes as well as already peeked cost that will consist on the new solution. indexes are guaranteed to be different.
func (sol *Solution) Neighbor(i,j int,cost float64) (*Solution){
	sol.Cost=cost
	sol.calculateNeighborFeas(i,j)
	sol.swap(i,j)
	return sol
}

//the way we defined the bitset for feasibility is: if there is no path from i to i+1, we set i. thus no bit set implies feasibility
func (sol *Solution) IsFeasible() bool{
	return sol.Feasible.None()
}


func (sol Solution) String() string {
	return fmt.Sprintf("Path: %v, Cost: %f, avgDist: %f, maxDist: %f, feasible: %s",sol.path,sol.Cost,sol.avgDist,sol.maxDist,sol.Feasible)
}
