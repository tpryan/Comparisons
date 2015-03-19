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
include "password/php/rules.php";
class rules_test extends PHPUnit_Framework_TestCase
{
    

    public function testEmpty() {
       $rules = new Rules();
       $candiate = "";
       $result = $rules->validate($candiate);
       $this->assertFalse($result['pass']);
       $this->assertEquals($result['message'], $rules::FAIL_EMPTY);
    }

    public function testMin() {
       $rules = new Rules();
       $candiate = "dasdsfg";
       $result = $rules->validate($candiate);
       $this->assertFalse($result['pass']);
       $this->assertEquals($result['message'], $rules::FAIL_MIN);
    }

    public function testMax() {
       $rules = new Rules();
       $candiate = "1234567890123456789012345";
       $result = $rules->validate($candiate);
       $this->assertFalse($result['pass']);
       $this->assertEquals($result['message'], $rules::FAIL_MAX);
    }

    public function testNoUpper() {
       $rules = new Rules();
       $candiate = "dasdasdasdasd";
       $result = $rules->validate($candiate);
       $this->assertFalse($result['pass']);
       $this->assertEquals($result['message'], $rules::FAIL_UPPER);
    }

    public function testNoLower() {
       $rules = new Rules();
       $candiate = "DKRKASDKEKASKD";
       $result = $rules->validate($candiate);
       $this->assertFalse($result['pass']);
       $this->assertEquals($result['message'], $rules::FAIL_LOWER);
    }

    public function testNumeric() {
       $rules = new Rules();
       $candiate = "Drdfjflrmg";
       $result = $rules->validate($candiate);
       $this->assertFalse($result['pass']);
       $this->assertEquals($result['message'], $rules::FAIL_NUMBER);
    }

    public function testSpecial() {
       $rules = new Rules();
       $candiate = "Drdfjflr9mg";
       $result = $rules->validate($candiate);
       $this->assertFalse($result['pass']);
       $this->assertEquals($result['message'], $rules::FAIL_SPECIAL);
    }

    public function testDictionary() {
       $rules = new Rules();
       $candiate = "Drdfjflr9mg&Apple";
       $result = $rules->validate($candiate);
       $this->assertFalse($result['pass']);
       $this->assertEquals($result['message'], $rules::FAIL_DICTIONARY);
       $this->assertEquals($result['word'], "APPLE");
    }

    public function testValid() {
       $rules = new Rules();
       $candiate = "Drdfjflr9mg&";
       $result = $rules->validate($candiate);
       $this->assertTrue($result['pass']);
       $this->assertEquals($result['message'], $rules::SUCCESS);
    }

    public function testDictionaryHash() {
       $rules = new Rules();
       $candiate = "Drdfjflr9mg&Apple";
       $result = $rules->validate($candiate, "hash");
       $this->assertFalse($result['pass']);
       $this->assertEquals($result['message'], $rules::FAIL_DICTIONARY);
       $this->assertEquals($result['word'], "APPLE");
    }

    public function testValidHash() {
       $rules = new Rules();
       $candiate = "Drdfjflr9mg&";
       $result = $rules->validate($candiate, "hash");
       $this->assertTrue($result['pass']);
       $this->assertEquals($result['message'], $rules::SUCCESS);
    }

}
;?>