package react

const testVersion = 5

type ReactorImpl struct {
	Cell
}
type InputCellImpl struct {
	CellImpl
}
type ComputeCellImpl struct {
	Cell
	f func(int) int
}
type CellImpl struct {
	v int
}
type CancelerImpl struct {
}

func (c CancelerImpl) Cancel() {
}
func (c CellImpl) Value() int {
	return c.v
}
func (c *ComputeCellImpl) Value() int {
	//	if c.f != nil {
	f := c.f
	//	}
	return f(c.Cell.Value())
}
func (i *InputCellImpl) SetValue(in int) {
	i.v = in
}
func New() Reactor {
	r := ReactorImpl{}
	return r
}

func (r ReactorImpl) CreateInput(input int) InputCell {
	i := new(InputCellImpl)
	i.SetValue(input)
	return i
}

func (r ReactorImpl) CreateCompute1(c Cell, f func(int) int) ComputeCell {
	cc := new(ComputeCellImpl)
	//	cc.CellImpl.v = f(c.Value())
	//cc.AddCallback(f)
	cc.Cell = c
	cc.f = f
	return cc
}

func (r ReactorImpl) CreateCompute2(c1 Cell, c2 Cell, f func(int, int) int) ComputeCell {
	cc := new(ComputeCellImpl)
	//	cc.CellImpl.v = f(c1.Value(), c2.Value())
	return cc
}

func (c *ComputeCellImpl) AddCallback(f func(int)) Canceler {
	//c.funcs = append(c.funcs, f)
	cc := new(CancelerImpl)
	return cc
}
