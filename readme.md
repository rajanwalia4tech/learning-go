# Getting Started

## Creating go.mod file

```
go mod init hello
```

## Compile and run go file (Below command will not create executable)
```
go run main.go
```

## Check the go installed folder using gopath
```
go env GOPATH
```

## Build for different OS

### Type the below command
```
go env
```

### Using the GOOS command, we can build the executable for different OS
```
GOOS="windows" go build // build for windows

GOOS="linux" go build // build for linux

GOOS="darwin" go build // build for mac

```




## Golang DataTypes
- String
- Bool
- Integer - uint8, uint64, int8, int64, uintptr
- Floating - float32, float64
- Complex

## Advance Types 
- Arrays
- Slices
- structs
- Maps
- Pointers
- functions
- channels
- mutex

## Variables
- Case Insensitive - almost
- Variable type should be known in advance.
- Almost everything is type here