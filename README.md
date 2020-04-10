# GoError
In go you don't handle errors the way you do in java - and you shouldn't. But if you are new to golang and want to use the
good 'ol try-catch from java:<br><b>here you go</b>

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
