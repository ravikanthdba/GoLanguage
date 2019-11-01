package main

import (
    "fmt"
    "os"
    "strings"
    "io"
)

func main() {
    f, err := os.Create("names.txt");
    if err != nil {
        fmt.Println(err);
        return;
    }
    defer f.Close()

    fmt.Printf("The type of f is : %T\n", f);

    r := strings.NewReader("Hello World!!");
    io.Copy(f, r);

    r = strings.NewReader("Hello World 123!!");
    io.Copy(f, r);


    fileInfo, err := os.Stat("names.txt");

    fmt.Println("File name:", fileInfo.Name());
    fmt.Println("Size in bytes:", fileInfo.Size());
    fmt.Println("Permissions:", fileInfo.Mode());
    fmt.Println("Last modified:", fileInfo.ModTime());
    fmt.Println("Is Directory: ", fileInfo.IsDir());
    fmt.Printf("System interface type: %T\n", fileInfo.Sys());
    fmt.Printf("System info: %+v\n\n", fileInfo.Sys());
    
}