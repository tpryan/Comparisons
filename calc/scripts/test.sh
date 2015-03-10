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
echo ""
echo ""
echo "Executing php test"
time php calc/php/calc.php $1
echo ""
echo ""
echo "Executing go test"
time calc/calc $1 w