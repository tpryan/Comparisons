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

loopcount = ARGV[0].to_i
method = ARGV[1]

i=1
rules = Rules.new()
File.open("password/data/test_passwords.txt").each do |line|
  if i > loopcount
    break
  end

  result = rules.validate(line,method);
  # puts result.to_s

  i+=1
end