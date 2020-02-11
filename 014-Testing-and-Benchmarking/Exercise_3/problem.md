# Problem Statement

Start with https://github.com/GoesToEleven/go-programming/tree/master/code_samples/010-ninja-level-thirteen/03/01-code-starting. Get the code ready to BET on (add benchmarks, examples, tests). Write a table test. Add documentation to the code. Run the following in this order:
tests - go test -v
benchmarks - go test -bench .
coverage - go test -cover
coverage shown in web browser - go test -coverprofile c.out; go tool cover -html=c.out; file:///private/var/folders/_0/gymzcft569z40g4_s8b2j740000kx5/T/cover600871259/coverage.html#file0
examples shown in documentation in a web browser -  godoc -http=127.0.0.1:8080; Open 127.0.0.1:8080 in the browser; Navigate to packages - Search for mymath. 
