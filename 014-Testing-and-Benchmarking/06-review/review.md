# BET

Benchmarking
Example
Testing

# Syntax of the calls

BenchmarkXxx
ExampleXxx
TestXxx

Xxx - is the function

# Commands

godoc -http=:127.0.0.1:8080

go test  
go test -bench .  
go test -cover  
go test -coverprofile c.out  
go tool cover -html=c.out