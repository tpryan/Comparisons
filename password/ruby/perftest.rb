require_relative "rules"

loopcount = ARGV[0].to_i

i=1
rules = Rules.new()
File.open("password/data/test_passwords.txt").each do |line|
  if i > loopcount
    break
  end

  result = rules.validate(line);
  # puts result.to_s

  i+=1
end