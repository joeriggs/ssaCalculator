# Social Security Benefit Calculator
This is a tool for estimating future Social Security benefits.  The Social Security Administration (SSA) has very good tools that will show you your future benefits, but their tools assumes that you will continue to work until you file.  My tool allows you to see what your future benefits will look like if you stop working 1 or more years prior to filing for Social Security benefits.
If you create an account on the SSA website, you can download an XML file (Your_Social_Security_Statement_Data.xml) that contains your SSA history.  This tool will load the data from that file, and then it will predict future Social Security benefits based on your work history.
# How to Build
mkdir -p $GOPATH/src/github.com/joeriggs

git clone https://github.com/joeriggs/ssaCalculator $GOPATH/src/github.com/joeriggs/ssaCalculator

cd $GOPATH/src/github.com/joeriggs/ssaCalculator

go get

go build or go install

./ssaCalculator
