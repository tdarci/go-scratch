package engine

type Engine struct {
}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) Add(a, b int) int {
	return a + b
}
