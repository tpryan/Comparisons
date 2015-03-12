<?php
	
	include_once("password/php/rules.php");
	$rules = new Rules();

	if (defined('STDIN') && array_key_exists(1, $argv)) {
	  $loopcount = $argv[1];
	} else { 
	  $loopcount = 1;
	}

	if (defined('STDIN') && array_key_exists(2, $argv)) {
	  $method = $argv[2];
	} else { 
	  $method = "bruteforce";
	}

	$i=1;

	$handle = fopen("password/data/test_passwords.txt", "r");
	if ($handle) {
    	while (($line = fgets($handle)) !== false) {
    		if ($i>$loopcount){
    			break;
    		}
    		$result = $rules->validate($line, $method);
    		//echo $line . " " . implode(" ",$result) . "\n";

    		$i++;
	    }

	    fclose($handle);
	} else {
	    echo "Error opening test file";
	} 

?>