var cities = [];
cities.push({ name: "Toyonaka", country: "Japan", coordinates: [34.783298,135.466995] });
cities.push({ name: "Yanji", country: "China", coordinates: [42.771400,129.423004] });
cities.push({ name: "Yingkou", country: "China", coordinates: [40.637299,122.502998] });
cities.push({ name: "Dezhou", country: "China", coordinates: [37.448601,116.292999] });
cities.push({ name: "Guilin", country: "China", coordinates: [25.281900,110.286003] });
cities.push({ name: "Johor Bahru", country: "Malaysia", coordinates: [1.465500,103.758003] });
cities.push({ name: "Mawlamyine", country: "Myanmar", coordinates: [16.491400,97.625603] });
cities.push({ name: "Dhaka", country: "Bangladesh", coordinates: [23.723101,90.408600] });
cities.push({ name: "Ghaziabad", country: "India", coordinates: [28.666700,77.433296] });
cities.push({ name: "Taraz", country: "Kazakhstan", coordinates: [42.900002,71.366699] });
cities.push({ name: "Barnaul", country: "Russian Federation", coordinates: [53.360600,83.763603] });
cities.push({ name: "Kurgan", country: "Russian Federation", coordinates: [55.450001,65.333298] });
cities.push({ name: "Kerman", country: "Iran", coordinates: [30.283199,57.078800] });
cities.push({ name: "Qom", country: "Iran", coordinates: [34.640099,50.876400] });
cities.push({ name: "Kirkuk", country: "Iraq", coordinates: [35.468102,44.392200] });
cities.push({ name: "Malatya", country: "Turkey", coordinates: [38.350201,38.316700] });
cities.push({ name: "Port Said", country: "Egypt", coordinates: [31.266701,32.299999] });
cities.push({ name: "Bucharest", country: "Romania", coordinates: [44.433300,26.100000] });
cities.push({ name: "Dresden", country: "Germany", coordinates: [51.049999,13.750000] });
cities.push({ name: "Tirana", country: "Albania", coordinates: [41.327499,19.818899] });
cities.push({ name: "Kaduna", country: "Nigeria", coordinates: [10.522200,7.438280] });
cities.push({ name: "Toulouse", country: "France", coordinates: [43.599499,1.433190] });
cities.push({ name: "Valladolid", country: "Spain", coordinates: [41.655201,-4.723720] });
cities.push({ name: "Monrovia", country: "Liberia", coordinates: [6.310560,-10.804700] });
cities.push({ name: "Dakar", country: "Senegal", coordinates: [14.670800,-17.438101] });
cities.push({ name: "Aracaju", country: "Brazil", coordinates: [-10.916700,-37.066700] });
cities.push({ name: "Ribeirão das Neves", country: "Brazil", coordinates: [-19.783300,-44.099998] });
cities.push({ name: "Porto Alegre", country: "Brazil", coordinates: [-30.033300,-51.200001] });
cities.push({ name: "Asunción", country: "Paraguay", coordinates: [-25.266701,-57.666698] });
cities.push({ name: "Pôrto Velho", country: "Brazil", coordinates: [-8.766670,-63.900002] });
cities.push({ name: "Maracaibo", country: "Venezuela", coordinates: [10.631700,-71.640602] });
cities.push({ name: "Camagüey", country: "Cuba", coordinates: [21.380800,-77.916901] });
cities.push({ name: "Mixco", country: "Guatemala", coordinates: [14.633300,-90.606400] });
cities.push({ name: "Atlanta", country: "United States", coordinates: [33.748901,-84.388100] });
cities.push({ name: "San Antonio", country: "United States", coordinates: [29.423901,-98.493301] });
cities.push({ name: "Colorado Springs", country: "United States", coordinates: [38.833900,-104.820999] });
cities.push({ name: "Hermosillo", country: "Mexico", coordinates: [29.066700,-110.967003] });
cities.push({ name: "Tijuana", country: "Mexico", coordinates: [32.533298,-117.016998] });
cities.push({ name: "Vancouver", country: "Canada", coordinates: [49.250000,-123.133003] });
cities.push({ name: "Honolulu", country: "United States", coordinates: [21.306900,-157.858002] });

var targetSVG = "M9,0C4.029,0,0,4.029,0,9s4.029,9,9,9s9-4.029,9-9S13.971,0,9,0z M9,15.93 c-3.83,0-6.93-3.1-6.93-6.93S5.17,2.07,9,2.07s6.93,3.1,6.93,6.93S12.83,15.93,9,15.93 M12.5,9c0,1.933-1.567,3.5-3.5,3.5S5.5,10.933,5.5,9S7.067,5.5,9,5.5 S12.5,7.067,12.5,9z";
var planeSVG = "M19.671,8.11l-2.777,2.777l-3.837-0.861c0.362-0.505,0.916-1.683,0.464-2.135c-0.518-0.517-1.979,0.278-2.305,0.604l-0.913,0.913L7.614,8.804l-2.021,2.021l2.232,1.061l-0.082,0.082l1.701,1.701l0.688-0.687l3.164,1.504L9.571,18.21H6.413l-1.137,1.138l3.6,0.948l1.83,1.83l0.947,3.598l1.137-1.137V21.43l3.725-3.725l1.504,3.164l-0.687,0.687l1.702,1.701l0.081-0.081l1.062,2.231l2.02-2.02l-0.604-2.689l0.912-0.912c0.326-0.326,1.121-1.789,0.604-2.306c-0.452-0.452-1.63,0.101-2.135,0.464l-0.861-3.838l2.777-2.777c0.947-0.947,3.599-4.862,2.62-5.839C24.533,4.512,20.618,7.163,19.671,8.11z";
var map = AmCharts.makeChart( "chartdiv", {"type": "map","theme": "light","dataProvider": {"map": "worldLow","zoomLevel": 1,"zoomLongitude": -20.1341,"zoomLatitude": 49.1712,"lines": [{"latitudes": [ cities[0]["coordinates"][0], cities[1]["coordinates"][0] ], "longitudes": [ cities[0]["coordinates"][1], cities[1]["coordinates"][1] ]},
{"latitudes": [ cities[1]["coordinates"][0], cities[2]["coordinates"][0] ], "longitudes": [ cities[1]["coordinates"][1], cities[2]["coordinates"][1] ]},
{"latitudes": [ cities[2]["coordinates"][0], cities[3]["coordinates"][0] ], "longitudes": [ cities[2]["coordinates"][1], cities[3]["coordinates"][1] ]},
{"latitudes": [ cities[3]["coordinates"][0], cities[4]["coordinates"][0] ], "longitudes": [ cities[3]["coordinates"][1], cities[4]["coordinates"][1] ]},
{"latitudes": [ cities[4]["coordinates"][0], cities[5]["coordinates"][0] ], "longitudes": [ cities[4]["coordinates"][1], cities[5]["coordinates"][1] ]},
{"latitudes": [ cities[5]["coordinates"][0], cities[6]["coordinates"][0] ], "longitudes": [ cities[5]["coordinates"][1], cities[6]["coordinates"][1] ]},
{"latitudes": [ cities[6]["coordinates"][0], cities[7]["coordinates"][0] ], "longitudes": [ cities[6]["coordinates"][1], cities[7]["coordinates"][1] ]},
{"latitudes": [ cities[7]["coordinates"][0], cities[8]["coordinates"][0] ], "longitudes": [ cities[7]["coordinates"][1], cities[8]["coordinates"][1] ]},
{"latitudes": [ cities[8]["coordinates"][0], cities[9]["coordinates"][0] ], "longitudes": [ cities[8]["coordinates"][1], cities[9]["coordinates"][1] ]},
{"latitudes": [ cities[9]["coordinates"][0], cities[10]["coordinates"][0] ], "longitudes": [ cities[9]["coordinates"][1], cities[10]["coordinates"][1] ]},
{"latitudes": [ cities[10]["coordinates"][0], cities[11]["coordinates"][0] ], "longitudes": [ cities[10]["coordinates"][1], cities[11]["coordinates"][1] ]},
{"latitudes": [ cities[11]["coordinates"][0], cities[12]["coordinates"][0] ], "longitudes": [ cities[11]["coordinates"][1], cities[12]["coordinates"][1] ]},
{"latitudes": [ cities[12]["coordinates"][0], cities[13]["coordinates"][0] ], "longitudes": [ cities[12]["coordinates"][1], cities[13]["coordinates"][1] ]},
{"latitudes": [ cities[13]["coordinates"][0], cities[14]["coordinates"][0] ], "longitudes": [ cities[13]["coordinates"][1], cities[14]["coordinates"][1] ]},
{"latitudes": [ cities[14]["coordinates"][0], cities[15]["coordinates"][0] ], "longitudes": [ cities[14]["coordinates"][1], cities[15]["coordinates"][1] ]},
{"latitudes": [ cities[15]["coordinates"][0], cities[16]["coordinates"][0] ], "longitudes": [ cities[15]["coordinates"][1], cities[16]["coordinates"][1] ]},
{"latitudes": [ cities[16]["coordinates"][0], cities[17]["coordinates"][0] ], "longitudes": [ cities[16]["coordinates"][1], cities[17]["coordinates"][1] ]},
{"latitudes": [ cities[17]["coordinates"][0], cities[18]["coordinates"][0] ], "longitudes": [ cities[17]["coordinates"][1], cities[18]["coordinates"][1] ]},
{"latitudes": [ cities[18]["coordinates"][0], cities[19]["coordinates"][0] ], "longitudes": [ cities[18]["coordinates"][1], cities[19]["coordinates"][1] ]},
{"latitudes": [ cities[19]["coordinates"][0], cities[20]["coordinates"][0] ], "longitudes": [ cities[19]["coordinates"][1], cities[20]["coordinates"][1] ]},
{"latitudes": [ cities[20]["coordinates"][0], cities[21]["coordinates"][0] ], "longitudes": [ cities[20]["coordinates"][1], cities[21]["coordinates"][1] ]},
{"latitudes": [ cities[21]["coordinates"][0], cities[22]["coordinates"][0] ], "longitudes": [ cities[21]["coordinates"][1], cities[22]["coordinates"][1] ]},
{"latitudes": [ cities[22]["coordinates"][0], cities[23]["coordinates"][0] ], "longitudes": [ cities[22]["coordinates"][1], cities[23]["coordinates"][1] ]},
{"latitudes": [ cities[23]["coordinates"][0], cities[24]["coordinates"][0] ], "longitudes": [ cities[23]["coordinates"][1], cities[24]["coordinates"][1] ]},
{"latitudes": [ cities[24]["coordinates"][0], cities[25]["coordinates"][0] ], "longitudes": [ cities[24]["coordinates"][1], cities[25]["coordinates"][1] ]},
{"latitudes": [ cities[25]["coordinates"][0], cities[26]["coordinates"][0] ], "longitudes": [ cities[25]["coordinates"][1], cities[26]["coordinates"][1] ]},
{"latitudes": [ cities[26]["coordinates"][0], cities[27]["coordinates"][0] ], "longitudes": [ cities[26]["coordinates"][1], cities[27]["coordinates"][1] ]},
{"latitudes": [ cities[27]["coordinates"][0], cities[28]["coordinates"][0] ], "longitudes": [ cities[27]["coordinates"][1], cities[28]["coordinates"][1] ]},
{"latitudes": [ cities[28]["coordinates"][0], cities[29]["coordinates"][0] ], "longitudes": [ cities[28]["coordinates"][1], cities[29]["coordinates"][1] ]},
{"latitudes": [ cities[29]["coordinates"][0], cities[30]["coordinates"][0] ], "longitudes": [ cities[29]["coordinates"][1], cities[30]["coordinates"][1] ]},
{"latitudes": [ cities[30]["coordinates"][0], cities[31]["coordinates"][0] ], "longitudes": [ cities[30]["coordinates"][1], cities[31]["coordinates"][1] ]},
{"latitudes": [ cities[31]["coordinates"][0], cities[32]["coordinates"][0] ], "longitudes": [ cities[31]["coordinates"][1], cities[32]["coordinates"][1] ]},
{"latitudes": [ cities[32]["coordinates"][0], cities[33]["coordinates"][0] ], "longitudes": [ cities[32]["coordinates"][1], cities[33]["coordinates"][1] ]},
{"latitudes": [ cities[33]["coordinates"][0], cities[34]["coordinates"][0] ], "longitudes": [ cities[33]["coordinates"][1], cities[34]["coordinates"][1] ]},
{"latitudes": [ cities[34]["coordinates"][0], cities[35]["coordinates"][0] ], "longitudes": [ cities[34]["coordinates"][1], cities[35]["coordinates"][1] ]},
{"latitudes": [ cities[35]["coordinates"][0], cities[36]["coordinates"][0] ], "longitudes": [ cities[35]["coordinates"][1], cities[36]["coordinates"][1] ]},
{"latitudes": [ cities[36]["coordinates"][0], cities[37]["coordinates"][0] ], "longitudes": [ cities[36]["coordinates"][1], cities[37]["coordinates"][1] ]},
{"latitudes": [ cities[37]["coordinates"][0], cities[38]["coordinates"][0] ], "longitudes": [ cities[37]["coordinates"][1], cities[38]["coordinates"][1] ]},
{"latitudes": [ cities[38]["coordinates"][0], cities[39]["coordinates"][0] ], "longitudes": [ cities[38]["coordinates"][1], cities[39]["coordinates"][1] ]},
],"images": [{ "svgPath": targetSVG, "title": cities[0]["name"]+", "+cities[0]["country"], "latitude": cities[0]["coordinates"][0],"longitude": cities[0]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[1]["name"]+", "+cities[1]["country"], "latitude": cities[1]["coordinates"][0],"longitude": cities[1]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[2]["name"]+", "+cities[2]["country"], "latitude": cities[2]["coordinates"][0],"longitude": cities[2]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[3]["name"]+", "+cities[3]["country"], "latitude": cities[3]["coordinates"][0],"longitude": cities[3]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[4]["name"]+", "+cities[4]["country"], "latitude": cities[4]["coordinates"][0],"longitude": cities[4]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[5]["name"]+", "+cities[5]["country"], "latitude": cities[5]["coordinates"][0],"longitude": cities[5]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[6]["name"]+", "+cities[6]["country"], "latitude": cities[6]["coordinates"][0],"longitude": cities[6]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[7]["name"]+", "+cities[7]["country"], "latitude": cities[7]["coordinates"][0],"longitude": cities[7]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[8]["name"]+", "+cities[8]["country"], "latitude": cities[8]["coordinates"][0],"longitude": cities[8]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[9]["name"]+", "+cities[9]["country"], "latitude": cities[9]["coordinates"][0],"longitude": cities[9]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[10]["name"]+", "+cities[10]["country"], "latitude": cities[10]["coordinates"][0],"longitude": cities[10]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[11]["name"]+", "+cities[11]["country"], "latitude": cities[11]["coordinates"][0],"longitude": cities[11]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[12]["name"]+", "+cities[12]["country"], "latitude": cities[12]["coordinates"][0],"longitude": cities[12]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[13]["name"]+", "+cities[13]["country"], "latitude": cities[13]["coordinates"][0],"longitude": cities[13]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[14]["name"]+", "+cities[14]["country"], "latitude": cities[14]["coordinates"][0],"longitude": cities[14]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[15]["name"]+", "+cities[15]["country"], "latitude": cities[15]["coordinates"][0],"longitude": cities[15]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[16]["name"]+", "+cities[16]["country"], "latitude": cities[16]["coordinates"][0],"longitude": cities[16]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[17]["name"]+", "+cities[17]["country"], "latitude": cities[17]["coordinates"][0],"longitude": cities[17]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[18]["name"]+", "+cities[18]["country"], "latitude": cities[18]["coordinates"][0],"longitude": cities[18]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[19]["name"]+", "+cities[19]["country"], "latitude": cities[19]["coordinates"][0],"longitude": cities[19]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[20]["name"]+", "+cities[20]["country"], "latitude": cities[20]["coordinates"][0],"longitude": cities[20]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[21]["name"]+", "+cities[21]["country"], "latitude": cities[21]["coordinates"][0],"longitude": cities[21]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[22]["name"]+", "+cities[22]["country"], "latitude": cities[22]["coordinates"][0],"longitude": cities[22]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[23]["name"]+", "+cities[23]["country"], "latitude": cities[23]["coordinates"][0],"longitude": cities[23]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[24]["name"]+", "+cities[24]["country"], "latitude": cities[24]["coordinates"][0],"longitude": cities[24]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[25]["name"]+", "+cities[25]["country"], "latitude": cities[25]["coordinates"][0],"longitude": cities[25]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[26]["name"]+", "+cities[26]["country"], "latitude": cities[26]["coordinates"][0],"longitude": cities[26]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[27]["name"]+", "+cities[27]["country"], "latitude": cities[27]["coordinates"][0],"longitude": cities[27]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[28]["name"]+", "+cities[28]["country"], "latitude": cities[28]["coordinates"][0],"longitude": cities[28]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[29]["name"]+", "+cities[29]["country"], "latitude": cities[29]["coordinates"][0],"longitude": cities[29]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[30]["name"]+", "+cities[30]["country"], "latitude": cities[30]["coordinates"][0],"longitude": cities[30]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[31]["name"]+", "+cities[31]["country"], "latitude": cities[31]["coordinates"][0],"longitude": cities[31]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[32]["name"]+", "+cities[32]["country"], "latitude": cities[32]["coordinates"][0],"longitude": cities[32]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[33]["name"]+", "+cities[33]["country"], "latitude": cities[33]["coordinates"][0],"longitude": cities[33]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[34]["name"]+", "+cities[34]["country"], "latitude": cities[34]["coordinates"][0],"longitude": cities[34]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[35]["name"]+", "+cities[35]["country"], "latitude": cities[35]["coordinates"][0],"longitude": cities[35]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[36]["name"]+", "+cities[36]["country"], "latitude": cities[36]["coordinates"][0],"longitude": cities[36]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[37]["name"]+", "+cities[37]["country"], "latitude": cities[37]["coordinates"][0],"longitude": cities[37]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[38]["name"]+", "+cities[38]["country"], "latitude": cities[38]["coordinates"][0],"longitude": cities[38]["coordinates"][1],"scale": 1},
{ "svgPath": targetSVG, "title": cities[39]["name"]+", "+cities[39]["country"], "latitude": cities[39]["coordinates"][0],"longitude": cities[39]["coordinates"][1],"scale": 1},
]},"areasSettings": {"unlistedAreasColor": "#4F80E1","unlistedAreasAlpha": 0.6},"imagesSettings": {"color": "#7F2924","rollOverColor": "#7F2924","selectedColor": "#7F2924","alpha": 1,},"linesSettings": {"arc": 0,"arrow": "end","color": "#FF5349","alpha": 1,"thickness": 2,"arrowAlpha": 1,"arrowSize": 5},"zoomControl": {"gridHeight": 100,"draggerAlpha": 1,"gridAlpha": 0.4},"backgroundZoomsToTop": true,"linesAboveImages": true,"export": {"enabled": true}} );
