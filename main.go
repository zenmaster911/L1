package main

import (
	"fmt"
)

type exclaimer interface {
	exclaim()
}

type example1 struct{}

type example2 struct{}

type balerina struct{}

func (ex example1) experiment() {
	fmt.Println("Меня сейчас адаптируют")
}

func (ex example2) experimentResult() {
	fmt.Println("меня адпатировали")
}

func (bal balerina) piruet() {
	fmt.Println("А я - балерина!")
}

type adapter1 struct {
	*example1
}

type adapter2 struct {
	*example2
}

type baleriner struct {
	*balerina
}

func (a1 adapter1) exclaim() {
	a1.experiment()
}

func NewAdapter1(ex *example1) *adapter1 {
	return &adapter1{ex}
}

func (a2 adapter2) exclaim() {
	a2.experimentResult()
}

func NewAdapter2(ex *example2) *adapter2 {
	return &adapter2{ex}
}

func (b baleriner) exclaim() {
	b.piruet()
}

func NewBaleriner(bal *balerina) *baleriner {
	return &baleriner{bal}
}

func main() {
	madHouse := []exclaimer{NewAdapter1(&example1{}), NewAdapter2(&example2{}), NewBaleriner(&balerina{})}
	for _, ward := range madHouse {
		ward.exclaim()
	}
}
