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

$start = microtime(true);
	$routes = getRoutes($mysqli, $sql);
trace($start, "getRoutes");	

$start = microtime(true);
	$routes = processRoutes($routes);
trace($start, "processRoutes");



$start = microtime(true);
	writeRoutes($routes, $path_for_store);
trace($start, "writeRoutes");


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

function trace($start, $message){
	$end = microtime(true);

	$tab = "\t";
	if (strlen($message) < 12){
		$tab = "\t\t";	
	}

	echo date('Y-m-d h:i:s') . " " . $message . $tab .  " ElapsedTime in seconds: ".($end-$start) . "\n";
}

function delTree($dir) { 
	$files = array_diff(scandir($dir), array('.','..')); 
	foreach ($files as $file) { 
	  (is_dir("$dir/$file")) ? delTree("$dir/$file") : unlink("$dir/$file"); 
	} 
	return rmdir($dir); 
}
