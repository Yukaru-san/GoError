# GoError
A simpel Java-oriented library to handle errors in Go

# Installation
```go
go get github.com/Yukaru-san/GoError
```

# Usage
```go
Block{
		Try: func() { 
      // Try your code here
		},
		Catch: func(e Exception) {
			// Handle errors here
		},
		Finally: func() {
			// Conclude handling
		},
	}.Do()
```
