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
    return ""
  end  



  
  def validate(candidate)

    if candidate.length == 0
      return Result.new(false, FAIL_EMPTY, "FAIL_EMPTY", "")
    end

    if candidate.length < MIN_LENGTH
      return Result.new(false, FAIL_MIN, "FAIL_MIN", "")
    end

    if candidate.length > MAX_LENGTH
      return Result.new(false, FAIL_MAX, "FAIL_MAX", "")
    end

    if !(candidate =~ /[A-Z]/)
     return Result.new(false, FAIL_UPPER, "FAIL_UPPER", "")
    end

    if !(candidate =~ /[a-z]/)
     return Result.new(false, FAIL_LOWER, "FAIL_LOWER", "")
    end

    if !(candidate =~ /[0-9]/)
      return Result.new(false, FAIL_NUMBER, "FAIL_NUMBER", "")
    end

    regex = /[#{SPECIAL.gsub(/./){|char| "\\#{char}"}}]/

    if !(candidate =~ regex)
       return Result.new(false, FAIL_SPECIAL, "FAIL_SPECIAL", "")
    end  

    word = match(candidate)
    if word.length > 0 
      return Result.new(false, FAIL_DICTIONARY, "FAIL_DICTIONARY", "")
    end  


    result = Result.new(true, SUCCESS, "SUCCESS", "")
  end 

end 