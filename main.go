package main

//run annealer with 10 seeds starting with 0: tsp-go -n 10 -f 0 
//visualizing an already given solution (ideally generated by this program) stored in iker40.json : tsp-go -v utils/iker40.json

import(
    "encoding/json"
    "github.com/ikerlb/tsp-go/tspvis"
    "github.com/ikerlb/tsp-go/tsp"
    //"math/rand"
    "flag"
    "database/sql"
    "fmt"
    "os"
    _ "github.com/mattn/go-sqlite3"
)

type SolutionSet struct {
    Set []int `json:"Set"`
}

type Settings struct{
    Phi             float64 `json:"Phi"`
    BatchSize       int     `json:"BatchSize"`
    InitTemp        float64 `json:"InitTemp"`
    EpsilonT        float64 `json:"EpsilonT"`
    EpsilonP        float64 `json:"EpsilonP"`
    AcceptedPercent float64 `json:"AcceptedPercent"`
    InitialTempN    int     `json:"InitialTempN"`
    PathToAnneal    []int   `json:"PathToAnneal"`
}

const citiesLength=1092

var (
        SeedNumber      int    //Number of seeds to run.
        FirstSeed       int    //Defines range of seeds: [FirstSeed,FirstSeed+SeedNumber]
        CreateGraphs    bool   //Create annealing graph in util/graphs. (1 per seed)
        Visualizer      string //json file with 'solution' to parse and show in map
        SettingsFile    string //json file with annealer parameters and cities set. defaults to ./settings.json
        Cities          []tspvis.City
        Distances       [][]float64
        SolutionIds     []int
        Phi             float64
        BatchSize       int     
        InitTemp        float64 
        EpsilonT        float64 
        EpsilonP        float64 
        AcceptedPercent float64 
        InitialTempN    int
        PathToAnneal    []int     
)

func main() {

    //TODO: Errors!!!!
    flag.Parse()

    //ok := true
    //ok = errorMessage("ERROR: input argument required")     
    fmt.Println("Extracting info from database...");
    getDistances()
    fmt.Println("Succesfully extracted adjacency matrix from ./db/tsp.db")

    //if visualizer has a path, show map with its input.
    if Visualizer != "" {
        getCities()
        parseSolution()
        fmt.Printf("Succesfully extracted Cities information from db ./db/tsp.db\n")
        tspvis.Visualizer(Cities,Distances,SolutionIds)
        os.Exit(3) //exit as you are in visualization mode.
    }
    fmt.Println("Extracting parameters from settings file:",SettingsFile)
    parseSettings()
    //NewAnnealer(path []int,dists [][]float64,phi,initTemp, epsilonT, epsilonP, acceptedPercent float64, batchSize, initialTempN int)
    ann:=tsp.NewAnnealer(PathToAnneal, Distances, Phi, InitTemp, EpsilonT, EpsilonP, AcceptedPercent, BatchSize, InitialTempN)
    for i := FirstSeed; i < FirstSeed+SeedNumber; i++ {
        ann.AnnealWithSeed(int64(i))
    }
}

/*
* Init function. Parses flags.
*/
func init() {
    //IDEA: -vis or -v command flag ignores the rest of the flags and opens visualizer GUI
    flag.IntVar(&SeedNumber, "n", 1, "number of seeds to run, defaults to 1")
    flag.IntVar(&FirstSeed, "f", 0, "first seed to run. defines range [s,n+s). defaults to 0")
    flag.BoolVar(&CreateGraphs, "g", false, "create annealing graphs in util/graphs (1 per seed)")
    flag.StringVar(&Visualizer, "v", "", "Json file with 'solution' to parse and show in map.")
    flag.StringVar(&SettingsFile, "s", "settings.json", "Json file with annealer parameters and cities set to solve.")
}

func check(e error) {
    if e != nil {
        fmt.Println(e)
        panic(e)
    }
}

//close db
func getDistances(){
    db, err := sql.Open("sqlite3", "db/tsp.db")
    check(err)
    rows, err := db.Query("SELECT id_city_1, id_city_2,distance FROM CONNECTIONS")
    check(err)
    
    //initialize matrix, from different function.
    Distances=make([][]float64,citiesLength+1)
    for i := range Distances {
        Distances[i] = make([]float64, citiesLength+1)
    }   


    for i:=0;rows.Next();i++ {
        var id_city_1,id_city_2 int
        var distance float64
        err=rows.Scan(&id_city_1,&id_city_2,&distance)
        Distances[id_city_1][id_city_2]=distance
        Distances[id_city_2][id_city_1]=distance
    }   
}

//close db, rows
func getCities(){
    db, err := sql.Open("sqlite3", "db/tsp.db")
    check(err)
    rows, err := db.Query("SELECT id,name,country,latitude,longitude FROM CITIES")
    check(err)
    Cities=make([]tspvis.City,citiesLength)
    for i:=0;rows.Next();i++ {
        city:=tspvis.City{}
        err=rows.Scan(&(city.Id),&(city.Name),&(city.Country),&(city.Latitude),&(city.Longitude))
        Cities[i]=city
    }
}
//Phi,BatchSize,InitialTemp,EpsilonP,EpsilonT,AcceptedPercent,InitialTempN,PathToAnneal
func parseSettings(){
    file, err := os.Open(SettingsFile)
    check(err)
    fmt.Println("Successfully Opened",SettingsFile)
    sett:=Settings{}
    decoder := json.NewDecoder(file) 
    err = decoder.Decode(&sett) 
    check(err)
    Phi=sett.Phi
    BatchSize=sett.BatchSize
    InitTemp=sett.InitTemp
    EpsilonT=sett.EpsilonT
    EpsilonP=sett.EpsilonP
    AcceptedPercent=sett.AcceptedPercent
    InitialTempN=sett.InitialTempN
    PathToAnneal=sett.PathToAnneal
    // defer the closing of our jsonFile so that we can parse it later on
    file.Close()
    fmt.Println("Succesfully extracted Settings:",sett," from json file:",SettingsFile)
}

func parseSolution(){
    file, err := os.Open(Visualizer)
    check(err)
    fmt.Println("Successfully Opened",Visualizer)
    sol:=SolutionSet{}
    decoder := json.NewDecoder(file) 
    err = decoder.Decode(&sol) 
    check(err)
    SolutionIds=sol.Set
    // defer the closing of our jsonFile so that we can parse it later on
    file.Close()
    fmt.Println("Succesfully extracted SolutionIds:",SolutionIds,"from json file:",Visualizer)
}
