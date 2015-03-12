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

