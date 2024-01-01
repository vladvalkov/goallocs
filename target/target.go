package target

func main() {
	emptyFunc()

	_ = newStruct()
	newSlice1()
	newSlice2()
	newMap1()
	newMap2()
	newMap3()
	newMap4()
	alwaysOne()
}

//go:noinline
func newStruct() *struct{ _ int } {
	s := struct{ _ int }{}
	p := &s
	return p
}

//go:noinline
func newInt() *int {
	v := 1
	return &v
}

//go:noinline
func newSlice1() {
	_ = make([]uint8, 65536)
}

//go:noinline
func newSlice2() {
	_ = make([]uint8, 65537)
}

//go:noinline
func newMap1() {
	_ = make(map[int]int)
}

//go:noinline
func newMap2() map[int]int {
	return make(map[int]int)
}

//go:noinline
func newMap3() {
	_ = make(map[int]int, 100000)
}

//go:noinline
func newMap4() map[int]int {
	return make(map[int]int, 100000)
}

//go:noinline
func emptyFunc() {

}

//go:noinline
func alwaysOne() int {
	return 1
}
