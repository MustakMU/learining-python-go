package main

type bot interface {
	test()
}

type bo1 struct {
}
type bo2 struct {
}

func (bo1) test() {
	println("bot1")
}
func (bo2) test() {
	println("bot2")

}
func printb(b bot) {
	b.test()
}
func main() {
	a := bo1{}
	b := bo2{}

	printb(a)
	printb(b)
}
