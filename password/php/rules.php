<?php 

class Rules

{

	const MINIMUM_MATCH =  4;
	const MIN_LENGTH =  8;
	const MAX_LENGTH =  24;
	const SPECIAL =  "~!@#$%^&*";
	const SUCCESS =  "Password passes policy";
	const FAIL_EMPTY =  "No password given";
	const FAIL_UPPER =  "At least one UPPERCASE character is required.";
	const FAIL_LOWER =  "At least one LOWERCASE character is required.";
	const FAIL_NUMBER =  "At least one NUMERIC character is required.";
	const FAIL_SPECIAL =  "At least one SPECIAL (~!@#$%^&*) character is required.";
	const FAIL_DICTIONARY =  "No dictionary words allowed.  Found: ";
	const FAIL_MIN =  "Password must be at least 8 characters long."  ;
	const FAIL_MAX =  "Password must be no more than 24 characters long.";

	private $dict; 

	public function __construct() {	
	  $this->dict= json_decode(file_get_contents(getcwd() . "/password/data/dict.json"));
	}

	public function validate($candidate){

		if (strlen($candidate) == 0) {
			return self::create_result(false, self::FAIL_EMPTY, "FAIL_EMPTY");	
		}

		if (strlen($candidate) <  self::MIN_LENGTH) {
			return self::create_result(false, self::FAIL_MIN, "FAIL_MIN");	
		}

		if (strlen($candidate) >  self::MAX_LENGTH) {
			return self::create_result(false, self::FAIL_MAX, "FAIL_MAX");	
		}


		if (!self::hasUpper($candidate)) {
			return self::create_result(false, self::FAIL_UPPER, "FAIL_UPPER");	
		}

		if (!self::hasLower($candidate)) {
			return self::create_result(false, self::FAIL_LOWER, "FAIL_LOWER");	
		}

		if (!self::hasNumeric($candidate)) {
			return self::create_result(false, self::FAIL_NUMBER, "FAIL_NUMBER");	
		}

		if (!self::hasSpecial($candidate)) {
			return self::create_result(false, self::FAIL_SPECIAL, "FAIL_SPECIAL");	
		}

		$dic_match = self::dictionaryMatch($candidate);
		if (strlen($dic_match) > 0) {
			return self::create_result(false, self::FAIL_DICTIONARY, "FAIL_DICTIONARY", $dic_match);	
		}
		return self::create_result(true, self::SUCCESS, "SUCCESS");

	}

	private function create_result($pass, $message, $status, $word=null){
		return ['pass'=>$pass,'message'=>$message,"status"=>$status,"word"=>$word];
	}

	private function hasUpper($candidate){
		return preg_match_all('/[A-Z]/', $candidate);
	}

	private function hasLower($candidate){
		return preg_match_all('/[a-z]/', $candidate);
	}

	private function hasNumeric($candidate){
		return preg_match_all('/[0-9]/', $candidate);
	}

	private function hasSpecial($candidate){
		return preg_match_all('/[' . self::SPECIAL . ']/', $candidate);
	}

	private function dictionaryMatch($candidate){
		
		$c_len = strlen($candidate);
		foreach ($this->dict as $key => $word){
			if (strlen($word) < self::MINIMUM_MATCH){
				continue;
			}
			if (strlen($word) > strlen($candidate)){	
				continue;
			}
			if (stripos($candidate, $word) !== false ){
				return $word;
			}

		}
		return "";

	}
}
?>