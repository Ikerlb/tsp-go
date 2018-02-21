package tsp

import (
	"fmt"
	"github.com/willf/bitset"
)

type Solution struct{
	avgDist	 float64
	maxDist	 float64
	Path     []int
	Cost     float64
	Feasible *bitset.BitSet
	dists	 [][]float64
}

const discFactor float64 = 50.0 


func NewSolution(path []int, dists [][]float64) *Solution{
	sol:=Solution{}
	sol.Path=path
	sol.dists=dists
	sol.Feasible=bitset.New(uint(len(path)))
	sol.calculateMaxAvg()
	sol.calculateInitialCost()
	return &sol
}


func (sol *Solution) calculateMaxAvg(){
	var max, sum float64
	var c int
	for i := 0; i < len(sol.Path); i++ {
		for j := i+1; j < len(sol.Path); j++ {
			d:=sol.dists[sol.Path[i]][sol.Path[j]]
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
	size:=len(sol.Path)-1
	for i := 0; i < size; i++ {
		//if path[i] isn't connected to path[i+1] we set feasible[i]
		if sol.dists[sol.Path[i]][sol.Path[i+1]] == 0 {
			sol.Feasible.Set(uint(i))
		}
		sum+=sol.augmentedWeight(i,i+1)
	}
	sol.Cost=sum/(sol.avgDist*float64(size))
}


func (sol *Solution) augmentedWeight(i, j int) (dist float64){
	id1,id2:=sol.Path[i],sol.Path[j]
	dist=sol.dists[id1][id2]
	if dist==0 {
		dist=sol.maxDist
	}
	return
}

//peak what would the neighbor solution would be. this is to only create a full new neighbor if necessary (avoids creating new slice, bitset, etc).
func (sol *Solution) PeekNeighborCost(i,j int) float64 {
	size:=len(sol.Path)-1
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
		if i!=0{
			previousSum-=sol.augmentedWeight(i-1,i)
			previousSum+=sol.augmentedWeight(i-1,j)
		}
		if i!=size{
			previousSum-=sol.augmentedWeight(i,i+1)
			previousSum+=sol.augmentedWeight(j,i+1)
		}
		if j!=0 {
			previousSum-=sol.augmentedWeight(j-1,j)
			previousSum+=sol.augmentedWeight(j-1,i)
		}
		if j!=size{
			previousSum-=sol.augmentedWeight(j,j+1)
			previousSum+=sol.augmentedWeight(i,j+1)
		}
	}
	return previousSum/(sol.avgDist*(float64(size)))

}

func (sol *Solution) calculateNeighborFeas(i,j int){
	//sol.Feasible:=sol.Feasible.Clone()
	size:=len(sol.Path)-1
	if Abs(i-j) == 1 {
		low:=Min(i,j)
		high:=Max(i,j)
		if low!=0 {
			if sol.dists[sol.Path[low-1]][sol.Path[high]]==0 {
				sol.Feasible.Set(uint(low-1))
			} else {
				sol.Feasible.Clear(uint(low-1))
			}
		}
		if high!=size {
			if sol.dists[sol.Path[low]][sol.Path[high+1]]==0 {
				sol.Feasible.Set(uint(high))
			} else {
				sol.Feasible.Clear(uint(high))
			}
		}
	} else {
		if i!=0 {
			if sol.dists[sol.Path[i-1]][sol.Path[j]]==0 {
				sol.Feasible.Set(uint(i-1))
			} else {
				sol.Feasible.Clear(uint(i-1))
			}
		}
		if i!=size {
			if sol.dists[sol.Path[j]][sol.Path[i+1]]==0 {
				sol.Feasible.Set(uint(i))
			} else {
				sol.Feasible.Clear(uint(i))
			}
		}
		if j!=0 {
			if sol.dists[sol.Path[j-1]][sol.Path[i]]==0 {
				sol.Feasible.Set(uint(j-1))
			} else {
				sol.Feasible.Clear(uint(j-1))
			}
		}
		if j!=size {
			if sol.dists[sol.Path[i]][sol.Path[j+1]]==0 {
				sol.Feasible.Set(uint(j))
			} else {
				sol.Feasible.Clear(uint(j))
			}
		}
	}
}


func (sol *Solution) isConnected(i,j int) bool{
	id1,id2:=sol.Path[i],sol.Path[j]
	return sol.dists[id1][id2]!=0
}


func (sol *Solution) swap(i,j int){
	tmp:=sol.Path[i]
	sol.Path[i]=sol.Path[j]
	sol.Path[j]=tmp
}

//pass two indexes as well as already peeked cost that will consist on the new solution. indexes are guaranteed to be different.
func (sol *Solution) Neighbor(i,j int,cost float64) (*Solution){
	sol.Cost=cost
	sol.calculateNeighborFeas(i,j)
	sol.swap(i,j)
	return sol
}

//the way we defined the bitset for feasibility is: if there is no path from i to i+1, we set i. thus no bit set implies feasible i to i+1 path
func (sol *Solution) IsFeasible() bool{
	return sol.Feasible.None()
}


func (sol Solution) String() string {
	return fmt.Sprintf("Path: %v, Cost: %f, avgDist: %f, maxDist: %f, feasible: %s",sol.Path,sol.Cost,sol.avgDist,sol.maxDist,sol.Feasible)
}
