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