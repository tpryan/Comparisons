require_relative "rules"

candidate = ARGV[0]
method = ARGV[1]

rules = Rules.new()
result = rules.validate(candidate,method);

puts result