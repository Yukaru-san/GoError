package errorhandle

// ErrorHandler á la Java
type ErrorHandler struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

// Exception á la Java
type Exception interface{}

// Throw á la Java
func Throw(up Exception) {
	panic(up)
}

// Do executes the Error Handler
func (tcf ErrorHandler) Do() {
	if tcf.Finally != nil {
		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}
