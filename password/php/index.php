<?php 

	include_once("password/php/rules.php");

	$rules = new Rules();


	if (defined('STDIN') && array_key_exists(1, $argv)) {
	  	$candidate = $argv[1];
	} else { 
		report(['pass'=>false,'message'=>$rules::FAIL_EMPTY]);
	}

	report($rules->validate($candidate));

	function report($result){
		echo json_encode($result);
		exit;
	}
?>