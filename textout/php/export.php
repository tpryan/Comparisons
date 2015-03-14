<?php

if (defined('STDIN')) {
  $loopcount = $argv[1];
} else { 
  $loopcount = 1;
}

$sql = file_get_contents(getcwd() . "/textout/sql/entries.sql");

$db['user'] = getenv("DB_USER");
$db['pass'] = getenv("DB_PASS");
$db['host'] = getenv("DB_HOST");
$db['name'] = getenv("DB_NAME");

$output_path = getcwd() . "/textout/output/php";
$mysqli = mysqli_connect($db['host'], $db['user'], $db['pass'], $db['name'])  or die("Error " . mysqli_error($mysqli)); 


cleanDir($output_path);

$entries = getEntries($mysqli, $sql);

for ($i=1; $i<= $loopcount; $i++){
	$path_for_store = $output_path . "/" . $i;
	writeEntries($entries, $path_for_store);
}

function cleanDir($path_to_clean){
	if (!file_exists($path_to_clean)) {
	    mkdir($path_to_clean);
	} else {
		delTree($path_to_clean);
		mkdir($path_to_clean);
	}
}

function getEntries($mysqli, $sql){
	$entriesSQL = $mysqli->query($sql)  or die("Error " . mysqli_error($mysqli)); 

	$entries = [];
	while ($row = mysqli_fetch_array($entriesSQL)) {
		$entry = [];
		$entry["name"] = $row['post_name'];
		$entry["title"] = $row['post_title'];
		$entry["post_date"] = $row['post_date'];
		$entry["excerpt"] = $row['post_excerpt'];
		$entry["url"] = cleanURL($row['guid']);
		$entry["content"] = $row['post_content'];
		$entry["date"] = $row['formatted_post_date'];
		array_push($entries, $entry);
	}	
	return $entries;
}

function writeEntries($entries, $store_path){
	mkdir($store_path);

	foreach ($entries as $key => $entry){
		$item = '<article>'. "\n" .
		'	<h1><a href="' . $entry["url"] . '">' . $entry["title"] .'</a></h1>'. "\n" .
		'	<time datetime="' . $entry["post_date"] . '">' . $entry["date"] . '</time>'. "\n" .
		'	<div>' . "\n" .
		$entry["content"] . "\n" .
		'	</div>'. "\n" .
		'</article>'. "\n";

		$filename = $store_path . "/" . $entry["name"] . ".html";
		file_put_contents ($filename , $item);
	}
}

function cleanURL($url){
	$url = str_replace("blog//blog/index.php/", "", $url);
	$url = str_replace("http://http://", "http://", $url);
	return $url;
}

function delTree($dir) { 
   $files = array_diff(scandir($dir), array('.','..')); 
    foreach ($files as $file) { 
      (is_dir("$dir/$file")) ? delTree("$dir/$file") : unlink("$dir/$file"); 
    } 
    return rmdir($dir); 
  }