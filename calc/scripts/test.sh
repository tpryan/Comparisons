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
go build -o calc/calc  calc/go/calc.go

echo "Cleaning ruby generatated files. "
rm -rf calc/output/ruby
echo "Cleaning php generatated files. "
rm -rf calc/output/php
echo "Cleaning go generatated files. "
rm -rf calc/output/go
echo ""
echo "Executing ruby test"
time ruby calc/ruby/calc.rb $1
echo "Executing php test"
time php calc/php/calc.php $1 > /dev/null
echo "Executing go test"
time calc/calc -max=$1 