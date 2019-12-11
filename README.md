# coverage-ordering

> Testing who package alphabetical ordering impacts Golang coverage reporting

```
go get -u github.com/blend/go-sdk/cmd/coverage
cd correct/
"${GOPATH}/bin/coverage" -update

cd ../incorrect/
"${GOPATH}/bin/coverage" -update
```
