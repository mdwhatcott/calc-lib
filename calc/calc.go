package calc

type Addition struct{}
type Subtraction struct{}
type Division struct{}
type Multiplication struct{}

func (Addition) Calculate(a, b int) int       { return a + b }
func (Subtraction) Calculate(a, b int) int    { return a - b }
func (Multiplication) Calculate(a, b int) int { return a * b }
func (Division) Calculate(a, b int) int       { return a / b }
