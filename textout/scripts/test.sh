export TIMEFORMAT=%R
echo "Building go executable"
go build -o textout/export  textout/go/export.go

echo "Cleaning ruby generatated files. "
rm -rf textout/output/ruby
echo "Cleaning php generatated files. "
rm -rf textout/output/php
echo "Cleaning go generatated files. "
rm -rf textout/output/go
echo ""
echo "Executing ruby test"
time ruby textout/ruby/export.rb $1
echo "Executing php test"
time php textout/php/export.php $1
echo "Executing go test"
time textout/export -count=$1 