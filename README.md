# GoLanguage

Setting up Go Lang on the system

Linux: Assume: current-location - /home/rgarimel/Programming
1) Download and Install Go - 
    a) yum install epel-release
    b) yum install golang
    c) Download the tar file from https://golang.org/dl/, marked as Linux
        - Unzip the file
        - Untar the file
2) Create 3 directories - src pkg bin
    - All the source code files are stored in "src"
    - All the binaries are stored in "bin"
    - All the packages are stored in "pkg". go get -u <package-name> gets the code in "pkg" directory
3) Environmental Variables - (in .bashrc or .zshrc)
    a) export GOPATH="/home/rgarimel/Programming"
    b) export PATH=$PATH:$GOPATH/bin


Mac: - Assume: current-location - /Users/rgarimel/Documents/Programming/GoLang/goworkspace/
1) Download and Install Go - 
    a) brew install go
    b) Download the .pkg file from "https://golang.org/dl/" and double click and keep going next
2) Create 3 directories - src pkg bin
    - All the source code files are stored in "src"
    - All the binaries are stored in "bin"
    - All the packages are stored in "pkg". go get -u <package-name> gets the code in "pkg" directory
3) Environmental Variables -
    a) export GOPATH="/Users/rgarimel/Documents/Programming/GoLang/goworkspace/"
    b) export PATH=$PATH:$GOPATH/bin


Some of the basic Commands:
which go - will show the binary location of where go is installed
go version - gives the version of go
go help - gives us all the options and its details
