# TDD 
This is a practice project based on the book/website [TDD with GO](https://quii.gitbook.io) 

# Run 
```
 go run ./src/hello.go  
```

# Test
```
go test ./src
```

# Initital Setup
```
go mod init 
```

#GoLangCLI-lint
[Install Instructions](https://golangci-lint.run/usage/install/)
```
# binary will be $(go env GOPATH)/bin/golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.53.3

golangci-lint --version
```

# Documentation
Install GODOC it used to be included in go < 1.14
```
go install golang.org/x/tools/cmd/godoc@latest
```

Start the doc server
```
godoc -http :8000
open http://localhost:8000/pkg
```
