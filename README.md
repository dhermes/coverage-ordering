# coverage-ordering

> Testing how package alphabetical ordering impacts Golang coverage reporting

```
go get -u github.com/blend/go-sdk/cmd/coverage
go get -u github.com/ory/go-acc

cd correct/
"${GOPATH}/bin/coverage" -enforce -keep-coverage-out -covermode=atomic -coverprofile=coverage.cov
# coverage starting
# using covermode: set
# using coverprofile: coverage.cov
# pkg1: 100.00%
# pkg2: 50.00%
# final coverage: 71.43%
# compiling coverage report: coverage.html
# coverage complete
"${GOPATH}/bin/go-acc" --output=coverage.combined.cov $(go list ./...)
go tool cover -html=coverage.combined.cov -o coverage.html

cd ../incorrect/
"${GOPATH}/bin/coverage" -enforce -keep-coverage-out -covermode=atomic -coverprofile=coverage.cov
# coverage starting
# using covermode: set
# using coverprofile: coverage.cov
# pkg1: 50.00%
# pkg2: 100.00%
# final coverage: 71.43%
# compiling coverage report: coverage.html
# coverage complete
"${GOPATH}/bin/go-acc" --output=coverage.combined.cov $(go list ./...)
go tool cover -html=coverage.combined.cov -o coverage.html
```
