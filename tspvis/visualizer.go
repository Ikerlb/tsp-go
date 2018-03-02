package tspvis

import (
	"fmt"
    "os"
    //"math"
    // "github.com/wcharczuk/go-chart"
    // "github.com/wcharczuk/go-chart/seq"
)

func check(e error) {
    if e != nil {
    	fmt.Println(e)
        panic(e)
    }
}

//TODO: Get info on solution
func Visualizer(cities []City, distances [][]float64, citiesIds []int) {
	pushCities:=""
	lines:=""
	//images:=""
	images:=fmt.Sprintf("{ \"svgPath\": targetSVG, \"title\": cities[%d][\"name\"]+\", \"+cities[%d][\"country\"], \"latitude\": cities[%d][\"coordinates\"][0],\"longitude\": cities[%d][\"coordinates\"][1],\"scale\": 1},\n",0,0,0,0)
	images=images+fmt.Sprintf("{ \"svgPath\": targetSVG, \"title\": cities[%d][\"name\"]+\", \"+cities[%d][\"country\"], \"latitude\": cities[%d][\"coordinates\"][0],\"longitude\": cities[%d][\"coordinates\"][1],\"scale\": 1},\n",len(citiesIds)-1,len(citiesIds)-1,len(citiesIds)-1,len(citiesIds)-1)
	for i := 0; i < len(citiesIds); i++ {
		pushCities=pushCities+fmt.Sprintf("cities.push({ name: \"%s\", country: \"%s\", coordinates: [%f,%f] });\n",cities[citiesIds[i]-1].Name,cities[citiesIds[i]-1].Country,cities[citiesIds[i]-1].Latitude,cities[citiesIds[i]-1].Longitude)
		//!!!images=images+fmt.Sprintf("{ \"svgPath\": targetSVG, \"title\": cities[%d][\"name\"]+\", \"+cities[%d][\"country\"], \"latitude\": cities[%d][\"coordinates\"][0],\"longitude\": cities[%d][\"coordinates\"][1],\"scale\": 1},\n",i,i,i,i)
		if i<len(citiesIds)-1 {
			//fmt.Printf("City id1 is: %d, city id2 is: %d and its distance is %f\n",citiesIds[i], citiesIds[i+1], distances[citiesIds[i]][citiesIds[i+1]])
			if distances[citiesIds[i]][citiesIds[i+1]]!=0 {
				lines=lines+fmt.Sprintf("{\"latitudes\": [ cities[%d][\"coordinates\"][0], cities[%d][\"coordinates\"][0] ], \"longitudes\": [ cities[%d][\"coordinates\"][1], cities[%d][\"coordinates\"][1] ]},\n",i,i+1,i,i+1)
			}
		}
	}

	s := fmt.Sprintf("var cities = [];\n%s\nvar targetSVG = \"M9,0C4.029,0,0,4.029,0,9s4.029,9,9,9s9-4.029,9-9S13.971,0,9,0z M9,15.93 c-3.83,0-6.93-3.1-6.93-6.93S5.17,2.07,9,2.07s6.93,3.1,6.93,6.93S12.83,15.93,9,15.93 M12.5,9c0,1.933-1.567,3.5-3.5,3.5S5.5,10.933,5.5,9S7.067,5.5,9,5.5 S12.5,7.067,12.5,9z\";\nvar planeSVG = \"M19.671,8.11l-2.777,2.777l-3.837-0.861c0.362-0.505,0.916-1.683,0.464-2.135c-0.518-0.517-1.979,0.278-2.305,0.604l-0.913,0.913L7.614,8.804l-2.021,2.021l2.232,1.061l-0.082,0.082l1.701,1.701l0.688-0.687l3.164,1.504L9.571,18.21H6.413l-1.137,1.138l3.6,0.948l1.83,1.83l0.947,3.598l1.137-1.137V21.43l3.725-3.725l1.504,3.164l-0.687,0.687l1.702,1.701l0.081-0.081l1.062,2.231l2.02-2.02l-0.604-2.689l0.912-0.912c0.326-0.326,1.121-1.789,0.604-2.306c-0.452-0.452-1.63,0.101-2.135,0.464l-0.861-3.838l2.777-2.777c0.947-0.947,3.599-4.862,2.62-5.839C24.533,4.512,20.618,7.163,19.671,8.11z\";\nvar map = AmCharts.makeChart( \"chartdiv\", {\"type\": \"map\",\"theme\": \"light\",\"dataProvider\": {\"map\": \"worldLow\",\"zoomLevel\": 1,\"zoomLongitude\": -20.1341,\"zoomLatitude\": 49.1712,\"lines\": [%s],\"images\": [%s]},\"areasSettings\": {\"unlistedAreasColor\": \"#4F80E1\",\"unlistedAreasAlpha\": 0.6},\"imagesSettings\": {\"color\": \"#545454\",\"rollOverColor\": \"#7F2924\",\"selectedColor\": \"#7F2924\",\"alpha\": 1,},\"linesSettings\": {\"arc\": 0,\"arrow\": \"end\",\"color\": \"#FF5349\",\"alpha\": 1,\"thickness\": 2,\"arrowAlpha\": 1,\"arrowSize\": 8},\"zoomControl\": {\"gridHeight\": 100,\"draggerAlpha\": 1,\"gridAlpha\": 0.4},\"backgroundZoomsToTop\": true,\"linesAboveImages\": true,\"export\": {\"enabled\": true}} );\n", pushCities,lines,images)
	f, err := os.Create("./util/visualizer/js/map.js")
	check(err)
	defer f.Close()
	n, err := f.WriteString(s)
	check(err)
	fmt.Printf("Created file ./util/visualizer/js/map.js. Wrote %d bytes\n",n)
}


/*
				XValues: seq.Range(0, 127),
				YValues: seq.New(seq.NewRandom().WithLen(128).WithMax(1024)).Array(),
				res.Header().Set("Content-Type", chart.ContentTypeSVG)
				err := pie.Render(chart.SVG, res)
*/
// func MakeGraph(costs []float64, seed int) {
// 	batchSize:=50000
// 	hinit:="<!DOCTYPE HTML><html><head><script type=\"text/javascript\">"
// 	fp:=fmt.Sprintf("./util/graphs/%d.html",seed)
// 	f, err := os.Create(fp)
// 	check(err)
// 	onl:=fmt.Sprintf("window.onload = function() {\nvar chart = new CanvasJS.Chart(\"chartContainer\",\n{\nzoomEnabled: true,\ntitle:{\ntext: \"Seed %d\"\n},\naxisY:{\nincludeZero: false\n},\ndata: data,\n});\nchart.render();\n}\n",seed)
// 	linit:="var data = [];\nvar dataSeries = {type:\"line\",};\nvar dataPoints = [];\n"
// 	fmt.Println("Entering main loop with costs array of size:",len(costs))
// 	s:=hinit+onl+linit
// 	fmt.Println("Writing header...")
// 	_, err = f.WriteString(s)
// 	check(err)
// 	i:=0
// 	for ;i<len(costs);i++{
// 		lmain:=""
// 		for j:=0;j<batchSize;j++{
// 			if costs[i] < 0.0 {
// 				lmain+=fmt.Sprintf("dataPoints.push({x:%d,y:%f,color:\"Red\",markerSize: 5,markerType:\"circle\",});\n",i,math.Abs(costs[i]))
// 			} else{
// 				lmain+=fmt.Sprintf("dataPoints.push({x:%d,y:%f});\n",i,costs[i])
// 			}
// 			i++
// 		}
// 		fmt.Printf("Writing main loop %d lines out of %d lines...\n",i,len(costs))
// 		_, err = f.WriteString(lmain)
// 		check(err)
// 	}
// 	hend:="dataSeries.dataPoints = dataPoints;\ndata.push(dataSeries);\n</script>\n<script type=\"text/javascript\" src=\"https://canvasjs.com/assets/script/canvasjs.min.js\"></script>\n</head>\n<body>\n<div id=\"chartContainer\" style=\"height: 100%; width: 100%;\"></div>\n</body>\n</html>\n"
// 	// graph := chart.Chart{
// 	// 	Series: []chart.Series{
// 	// 		chart.ContinuousSeries{
// 	// 			XValues: seq.RangeWithStep(0, float64(len(costs))-1, 1),
// 	// 			//XValues: seq.Range(0,float64(len(costs))),
// 	// 			YValues: costs,
// 	// 		},
// 	// 	},
// 	// }
// 	// graph.Render(chart.SVG, f)
// 	fmt.Println("Writing rest...")
// 	n, err := f.WriteString(hend)
// 	check(err)
// 	fmt.Printf("Created file %s. Wrote %d bytes\n",fp,n)
//}
