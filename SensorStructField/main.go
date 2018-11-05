package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type Person struct {
	Name       string
	Age        int
	Sex        string
	CreditCard string
}

func (p Person) String() string {
	type anon Person
	p2 := anon(p)
	p2.CreditCard = sensor(p.CreditCard)
	return fmt.Sprintf("%+v", p2)
}

func (p Person) Format(st fmt.State, verb rune) {
	res := p.String()
	if verb == 'q' {
		res = fmt.Sprintf("%q", res)
	}
	io.WriteString(st, res)
}

func main() {
	customer := &Person{
		Name:       "John",
		Age:        27,
		Sex:        "male",
		CreditCard: "2753 7453 1221 6444",
	}
	fmt.Printf("%#v\n", customer)
	fmt.Printf("%+v\n", customer)
	fmt.Printf("%s\n", customer)
	fmt.Printf("%q\n", customer)
	spew.Dump(customer)
}

func sensor(str string) string {
	if str == "" {
		return str
	}
	fields := strings.Fields(str)
	n := len(fields) - 1
	return strings.Repeat("**** ", n) + fields[n]
}
