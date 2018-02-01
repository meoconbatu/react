package react

const testVersion = 5

type Reactive struct {
}

func (r *Celler) SetValue(input int) {
	r.i = input
	if r.callBack != nil {
		r.callBack(r.i)
	}
}

type Celler struct {
	i        int
	callBack func(int)
	c        *ComputeCeller
}

func (c Celler) Value() int {
	if c.c.f != nil {
		return c.c.f(c.c.c.Value())
	}
	if c.c.ff != nil {
		return c.c.ff(c.c.c.Value(), c.c.c2.Value())
	}
	return c.i
}
func (cc *Celler) AddCallback(f func(int)) Canceler {
	cc.callBack = f
	return *new(Reactive)
}

type ComputeCeller struct {
	c  *Celler
	c2 *Celler
	f  func(int) int
	ff func(int, int) int
}

func (c Reactive) Cancel() {

}
func (r *Reactive) CreateInput(input int) InputCell {
	ic := new(Celler)
	ic.c = new(ComputeCeller)
	ic.SetValue(input)
	return ic
}
func (r *Reactive) CreateCompute1(c Cell, f func(int) int) ComputeCell {
	cc := new(Celler)
	cc.c = new(ComputeCeller)
	cc.c.f = f
	cc.c.c = c.(*Celler)
	return cc
}
func (r *Reactive) CreateCompute2(c Cell, c1 Cell, f func(int, int) int) ComputeCell {
	cc := new(Celler)
	cc.c = new(ComputeCeller)
	cc.c.ff = f
	cc.c.c = c.(*Celler)
	cc.c.c2 = c1.(*Celler)
	return cc
}
func New() *Reactive {
	r := new(Reactive)
	return r
}
