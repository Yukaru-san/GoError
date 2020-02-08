# GoError
A simpel Java-oriented library to handle errors in Go

# Installation
```go
go get github.com/Yukaru-san/GoError
```

# Usage
```go
errorhandle.ErrorHandler{
	Try: func() {
		 // Try your code here
	},
	Catch: func(e errorhandle.Exception) {
		 // Handle errors here
	},
	Finally: func() {
		// Conclude handling
	},
}.Do()
```
