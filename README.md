# GoError
A simple Java-oriented library to handle errors in Go

# Installation
```go
go get github.com/Yukaru-san/GoError
```

# Usage
```go
goerror.ErrorHandler{
   Try: func() {
      // Try your code here
   },
   Catch: func(e goerror.Exception) {
      // Handle errors here
   },
   Finally: func() {
      // Conclude handling
   },
}.Do()
```
