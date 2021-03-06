<!-- 
   Copyright 2015, Google, Inc.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
 -->
<?php

ini_set('memory_limit', '8186M');
$loopcount = $argv[1];

$db['user'] = getenv("OF_USER");
$db['pass'] = getenv("OF_PASS");
$db['host'] = getenv("OF_HOST");
$db['name'] = getenv("OF_NAME");

$sql = file_get_contents(getcwd() . "/calc/sql/prepstatement.sql");
$sql .= "\n" . "Limit 0," . $loopcount . "\n";

$output_path = getcwd() . "/calc/output/php";
$path_for_store = $output_path . "/1/";

$mysqli = mysqli_connect($db['host'], $db['user'], $db['pass'], $db['name'])  or die("Error " . mysqli_error($mysqli)); 



cleanDir($output_path);
$routes = getRoutes($mysqli, $sql);
$routes = processRoutes($routes);
writeRoutes($routes, $path_for_store);


function writeRoutes($routes, $store_path){
	
	mkdir($store_path);

	$routeText = "<table>". "\n";
	
	$routeText .= '	<tr>'. "\n" .
		"		<th>Airline</th>" . "\n" .
		"		<th>Origin Aiport Code</th>" . "\n" .
		"		<th>Origin Aiport Name</th>" . "\n" .
		"		<th>Origin Latitude</th>" . "\n" .
		"		<th>Origin Longitude</th>" . "\n" .
		"		<th>Destination Aiport Code</th>" . "\n" .
		"		<th>Destination Aiport Name</th>" . "\n" .
		"		<th>Destination Latitude</th>" . "\n" .
		"		<th>Destination Longitude</th>" . "\n" .
		"		<th>Distance</th>" . "\n" .
		'	</tr>'. "\n";


	foreach ($routes as $key => $route){
		$item = '	<tr>'. "\n" .
		"		<td>" . $route['airline'] . "</td>" . "\n" .
		"		<td>" . $route['source_code'] . "</td>" . "\n" .
		"		<td>" . $route['source_name'] . "</td>" . "\n" .
		"		<td>" . $route['source_lat'] . "</td>" . "\n" .
		"		<td>" . $route['source_lon'] . "</td>" . "\n" .
		"		<td>" . $route['dest_code'] . "</td>" . "\n" .
		"		<td>" . $route['dest_name'] . "</td>" . "\n" .
		"		<td>" . $route['dest_lat'] . "</td>" . "\n" .
		"		<td>" . $route['dest_lon'] . "</td>" . "\n" .
		"		<td>" . $route['distance'] . "</td>" . "\n" .
		'	</tr>'. "\n";

		$routeText .=  $item;
	}
	$routeText .= "</table>". "\n";
	$filename = $store_path . "/table.html";
	file_put_contents ($filename , $routeText);
}


function cleanDir($path_to_clean){
	if (!file_exists($path_to_clean)) {
	    mkdir($path_to_clean);
	} else {
		delTree($path_to_clean);
		mkdir($path_to_clean);
	}
}


function getRoutes($mysqli, $sql){
	$routesSQL = $mysqli->query($sql)  or die("Error " . mysqli_error($mysqli)); 

	$routes = [];
	while ($row = mysqli_fetch_array($routesSQL)) {
		$i = [];
		$i["airline"] = $row['airline'];
		$i["source_code"] = $row['source_code'];
		$i["source_name"] = $row['source_name'];
		$i["source_lat"] = $row['source_lat'];
		$i["source_lon"] = $row['source_lon'];
		$i["dest_code"] = $row['dest_code'];
		$i["dest_name"] = $row['dest_name'];
		$i["dest_lat"] = $row['dest_lat'];
		$i["dest_lon"] = $row['dest_lon'];
		array_push($routes, $i);
	}	
	return $routes;
}

function processRoutes($routes){
	foreach ($routes as $key => $r){
		$r['distance'] = getDistance($r['source_lat'],$r['source_lon'],$r['dest_lat'],$r['dest_lon']);
		$routes[$key] = $r;
	}
	return $routes;
}

function getDistance($latitude1, $longitude1, $latitude2, $longitude2) {
	$earth_radius = 3963;
	
	$dLat = deg2rad($latitude2 - $latitude1);
	$dLon = deg2rad($longitude2 - $longitude1);
	
	$a = sin($dLat/2) * sin($dLat/2) + cos(deg2rad($latitude1)) * cos(deg2rad($latitude2)) * sin($dLon/2) * sin($dLon/2);
	$c = 2 * asin(sqrt($a));
	$d = $earth_radius * $c;
	return $d;
}

function delTree($dir) { 
	$files = array_diff(scandir($dir), array('.','..')); 
	foreach ($files as $file) { 
	  (is_dir("$dir/$file")) ? delTree("$dir/$file") : unlink("$dir/$file"); 
	} 
	return rmdir($dir); 
}
