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
export TIMEFORMAT=%R
echo "Building go executable"
go build -o password/perftest  password/go/perftest.go

echo ""
echo "Executing ruby test"
time ruby password/ruby/perftest.rb $1 $2
echo "Executing php test"
time php password/php/perftest.php $1 $2 > /dev/null
echo "Executing go test"
time password/perftest -count=$1 -method=$2