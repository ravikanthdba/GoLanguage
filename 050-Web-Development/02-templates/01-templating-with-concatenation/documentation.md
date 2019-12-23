In this case, we are writing data to a file through different formats:

- Writing data to a variable, and concatenating with string and standard output.
- Writing data to a variable, and concatenating with body and writing it to a file.
- Reading data through CLI arguments, concatenating with the body and writing it to a file.

Bottlenecks:
- If there is a very large program, this is not scalable solution.
- All the code is in func main, we need to make it re-usable.
