package demo

// sayHi private function
func sayHi() string {
	return "Hello"
}

// SayHi public function
// Error, NotException
func SayHi() (string, error) {
	return sayHi(), nil
}
