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
require_relative "dict"
class Rules

  MINIMUM_MATCH   = 4
  MIN_LENGTH      = 8
  MAX_LENGTH      = 24
  SPECIAL         = "~!@#$%^&*"
  SUCCESS         = "Password passes policy"
  FAIL_EMPTY      = "No password given"
  FAIL_UPPER      = "At least one UPPERCASE character is required."
  FAIL_LOWER      = "At least one LOWERCASE character is required."
  FAIL_NUMBER     = "At least one NUMERIC character is required."
  FAIL_SPECIAL    = "At least one SPECIAL (~!@#$%^&*) character is required."
  FAIL_DICTIONARY = "No dictionary words allowed."
  FAIL_MIN        = "Password must be at least 8 characters long."
  FAIL_MAX        = "Password must be no more than 24 characters long."

  Result = Struct.new(:pass, :message, :status, :word)

  def initialize()
    @dicthash = {}

    Dict::DICT.each do |word| 
      @dicthash[word] = true
    end
  end

  def break_string(str,min)
    res = {}
    len = str.length
    upstr = str.upcase
    last_start = len-min

    (min..len).each do |i|
      (0..last_start).each do |j|
        part = upstr[j,i].upcase
          res[part]=0
      end  
    end  
    res.keys
  end 

  def hashMatch(candidate)
    arr = break_string(candidate,MINIMUM_MATCH)
    arr.each do |part|
      return part if @dicthash[part] 
    end
    ""
  end  

  def match(candidate)
    uc = candidate.upcase
    
    Dict::DICT.each do |word| 
      if word.length < MINIMUM_MATCH
        next
      end  

      if word.length > candidate.length
        next
      end

      if uc.include? word
        return word
      end 

    end
    ""
  end  
  
  def validate(candidate, method="bruteforce")

    if candidate.length == 0
      return Result.new(false, FAIL_EMPTY, "FAIL_EMPTY", "")
    end

    if candidate.length < MIN_LENGTH
      return Result.new(false, FAIL_MIN, "FAIL_MIN", "")
    end

    if candidate.length > MAX_LENGTH
      return Result.new(false, FAIL_MAX, "FAIL_MAX", "")
    end

    unless candidate =~ /[A-Z]/
     return Result.new(false, FAIL_UPPER, "FAIL_UPPER", "")
    end

    unless candidate =~ /[a-z]/
     return Result.new(false, FAIL_LOWER, "FAIL_LOWER", "")
    end

    unless candidate =~ /[0-9]/
      return Result.new(false, FAIL_NUMBER, "FAIL_NUMBER", "")
    end

    unless candidate =~ /[#{SPECIAL}]/
       return Result.new(false, FAIL_SPECIAL, "FAIL_SPECIAL", "")
    end  

    if method == "bruteforce"
      word = match(candidate)
    else 
      word = hashMatch(candidate)
    end  

    if word.length > 0 
      return Result.new(false, FAIL_DICTIONARY, "FAIL_DICTIONARY", "")
    end  

    result = Result.new(true, SUCCESS, "SUCCESS", "")
  end 

end 