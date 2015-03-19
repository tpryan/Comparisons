export TIMEFORMAT=%R
echo "Building go executable"
go build -o password/perftest  password/go/perftest.go

echo ""
echo "Executing ruby test"
time ruby password/ruby/perftest.rb $1 $2
echo "Executing php test"
time php password/php/perftest.php $1 $2
echo "Executing go test"
time password/perftest -count=$1 -method=$2