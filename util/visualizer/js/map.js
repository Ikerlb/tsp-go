var cities = [];
cities.push({ name: "Honolulu", country: "United States", coordinates: [21.306900,-157.858002] });
cities.push({ name: "Vancouver", country: "Canada", coordinates: [49.250000,-123.133003] });
cities.push({ name: "Tijuana", country: "Mexico", coordinates: [32.533298,-117.016998] });
cities.push({ name: "Hermosillo", country: "Mexico", coordinates: [29.066700,-110.967003] });
cities.push({ name: "Colorado Springs", country: "United States", coordinates: [38.833900,-104.820999] });
cities.push({ name: "San Antonio", country: "United States", coordinates: [29.423901,-98.493301] });

var targetSVG = "M9,0C4.029,0,0,4.029,0,9s4.029,9,9,9s9-4.029,9-9S13.971,0,9,0z M9,15.93 c-3.83,0-6.93-3.1-6.93-6.93S5.17,2.07,9,2.07s6.93,3.1,6.93,6.93S12.83,15.93,9,15.93 M12.5,9c0,1.933-1.567,3.5-3.5,3.5S5.5,10.933,5.5,9S7.067,5.5,9,5.5 S12.5,7.067,12.5,9z";
var planeSVG = "M19.671,8.11l-2.777,2.777l-3.837-0.861c0.362-0.505,0.916-1.683,0.464-2.135c-0.518-0.517-1.979,0.278-2.305,0.604l-0.913,0.913L7.614,8.804l-2.021,2.021l2.232,1.061l-0.082,0.082l1.701,1.701l0.688-0.687l3.164,1.504L9.571,18.21H6.413l-1.137,1.138l3.6,0.948l1.83,1.83l0.947,3.598l1.137-1.137V21.43l3.725-3.725l1.504,3.164l-0.687,0.687l1.702,1.701l0.081-0.081l1.062,2.231l2.02-2.02l-0.604-2.689l0.912-0.912c0.326-0.326,1.121-1.789,0.604-2.306c-0.452-0.452-1.63,0.101-2.135,0.464l-0.861-3.838l2.777-2.777c0.947-0.947,3.599-4.862,2.62-5.839C24.533,4.512,20.618,7.163,19.671,8.11z";
var map = AmCharts.makeChart( "chartdiv", {"type": "map","theme": "light","dataProvider": {"map": "worldLow","zoomLevel": 1,"zoomLongitude": -20.1341,"zoomLatitude": 49.1712,"lines": [{"latitudes": [ cities[0]["coordinates"][0], cities[1]["coordinates"][0] ], "longitudes": [ cities[0]["coordinates"][1], cities[1]["coordinates"][1] ]},
{"latitudes": [ cities[1]["coordinates"][0], cities[2]["coordinates"][0] ], "longitudes": [ cities[1]["coordinates"][1], cities[2]["coordinates"][1] ]},
{"latitudes": [ cities[2]["coordinates"][0], cities[3]["coordinates"][0] ], "longitudes": [ cities[2]["coordinates"][1], cities[3]["coordinates"][1] ]},
{"latitudes": [ cities[3]["coordinates"][0], cities[4]["coordinates"][0] ], "longitudes": [ cities[3]["coordinates"][1], cities[4]["coordinates"][1] ]},
{"latitudes": [ cities[4]["coordinates"][0], cities[5]["coordinates"][0] ], "longitudes": [ cities[4]["coordinates"][1], cities[5]["coordinates"][1] ]},
],"images": [{ "svgPath": targetSVG, "title": cities[0]["name"]+", "+cities[0]["country"], "latitude": cities[0]["coordinates"][0],"longitude": cities[0]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[1]["name"]+", "+cities[1]["country"], "latitude": cities[1]["coordinates"][0],"longitude": cities[1]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[2]["name"]+", "+cities[2]["country"], "latitude": cities[2]["coordinates"][0],"longitude": cities[2]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[3]["name"]+", "+cities[3]["country"], "latitude": cities[3]["coordinates"][0],"longitude": cities[3]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[4]["name"]+", "+cities[4]["country"], "latitude": cities[4]["coordinates"][0],"longitude": cities[4]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[5]["name"]+", "+cities[5]["country"], "latitude": cities[5]["coordinates"][0],"longitude": cities[5]["coordinates"][1],"scale": 1},
]},"areasSettings": {"unlistedAreasColor": "#4F80E1","unlistedAreasAlpha": 0.6},"imagesSettings": {"color": "#7F2924","rollOverColor": "#7F2924","selectedColor": "#7F2924","alpha": 1,},"linesSettings": {"arc": 0,"arrow": "end","color": "#FF5349","alpha": 1,"thickness": 2,"arrowAlpha": 1,"arrowSize": 5},"zoomControl": {"gridHeight": 100,"draggerAlpha": 1,"gridAlpha": 0.4},"backgroundZoomsToTop": true,"linesAboveImages": true,"export": {"enabled": true}} );
