# Problem statement

Start with https://github.com/GoesToEleven/go-programming/tree/master/code_samples/010-ninja-level-thirteen/02/01-code-starting. Get the code ready to BET on (add benchmarks, examples, tests) however do not write an example for the func that returns a map; and only write a test for it as an extra challenge. Add documentation to the code. Run the following in this order:
tests - done - go test -v
benchmarks - done - go test -bench .
coverage - go test -cover 16.7% tests done, as we are not doing on map[strings]. 
coverage writing to c.out - go test -coverprofile c.out
coverage shown in web browser - go tool cover -html=c.out 
examples shown in documentation in a web browser - godoc -http=127.0.0.1:8080
