# Scopes

1) The scope of predefined variables is the universe block.

Predefined variables are as follows:
    Types:
        bool byte complex64 complex128 error float32 float64
        int int8 int16 int32 int64 rune string
        uint uint8 uint16 uint32 uint64 uintptr

    Constants:
        true false iota

    Zero value:
        nil

    Functions:
        append cap close complex copy delete imag len
        make new panic print println real recover

2) The scope of an identifier denoting a constant, type, variable, or function (but not method) declared at a top level is the package block. Example in 01.

3) The scope of the package name of an imported package is the file block of the file containing the import declaration.

4) The scope of an identifier denoting a method receiver, function parameter, or result variable is the function body.

5) The scope of a constant or variable identifier declared inside a function begins at the end of the ConstSpec or VarSpec (ShortVarDecl for short variable declarations) and ends at the end of the innermost containing block.

6) The scope of a type identifier declared inside a function begins at the identifier in the TypeSpec and ends at the end of the innermost containing block.