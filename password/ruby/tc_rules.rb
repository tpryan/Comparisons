# Copyright 2015, Google, Inc.

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at

#     http://www.apache.org/licenses/LICENSE-2.0

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
require_relative "rules"
require "test/unit"
 
class TestRules < Test::Unit::TestCase
 
  def test_empty
    candidate = ""
    result = Rules.new().validate(candidate)
    assert_false(result.pass )
    assert_equal(result.message, Rules::FAIL_EMPTY )
  end

  def test_min
    candidate = "dasdsfg"
    result = Rules.new().validate(candidate)
    assert_false(result.pass )
    assert_equal(result.message, Rules::FAIL_MIN )
  end

  def test_max
    candidate = "1234567890123456789012345"
    result = Rules.new().validate(candidate)
    assert_false(result.pass )
    assert_equal(result.message, Rules::FAIL_MAX )
  end

  def test_upper
    candidate = "dasdasdasdasd"
    result = Rules.new().validate(candidate)
    assert_false(result.pass )
    assert_equal(result.message, Rules::FAIL_UPPER )
  end

  def test_lower
    candidate = "DKRKASDKEKASKD"
    result = Rules.new().validate(candidate)
    assert_false(result.pass )
    assert_equal(result.message, Rules::FAIL_LOWER )
  end

  def test_numeric
    candidate = "Drdfjflrmg"
    result = Rules.new().validate(candidate)
    assert_false(result.pass )
    assert_equal(result.message, Rules::FAIL_NUMBER )
  end

  def test_special
    candidate = "Drdfjflr9mg"
    result = Rules.new().validate(candidate)
    assert_false(result.pass )
    assert_equal(result.message, Rules::FAIL_SPECIAL )
  end

  def test_dictionary
    candidate = "Drdfjflr9mg&Apple"
    result = Rules.new().validate(candidate)
    assert_false(result.pass )
    assert_equal(result.message, Rules::FAIL_DICTIONARY )
  end

  def test_valid
    candidate = "Drdfjflr9mg&"
    result = Rules.new().validate(candidate)
    assert_true(result.pass )
    assert_equal(result.message, Rules::SUCCESS )
  end

  def test_dictionary_hash
    candidate = "Drdfjflr9mg&Apple"
    result = Rules.new().validate(candidate, "hash")
    assert_false(result.pass )
    assert_equal(result.message, Rules::FAIL_DICTIONARY )
  end

  def test_valid_hash
    candidate = "Drdfjflr9mg&"
    result = Rules.new().validate(candidate, "hash")
    assert_true(result.pass )
    assert_equal(result.message, Rules::SUCCESS )
  end
 
end

