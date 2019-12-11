# coverage-ordering

> Testing how package alphabetical ordering impacts Golang coverage reporting

```
go get -u github.com/blend/go-sdk/cmd/coverage

cd correct/
"${GOPATH}/bin/coverage" -enforce
# coverage starting
# using covermode: set
# using coverprofile: coverage.cov
# pkg1: 100.00%
# pkg2: 50.00%
# final coverage: 71.43%
# compiling coverage report: coverage.html
# coverage complete

cd ../incorrect/
"${GOPATH}/bin/coverage" -enforce
# coverage starting
# using covermode: set
# using coverprofile: coverage.cov
# pkg1: 50.00%
# pkg2: 100.00%
# final coverage: 71.43%
# compiling coverage report: coverage.html
# coverage complete
```
