package main

import "fmt"

// Interface A
type a interface {
	funcA() string
	funcB() int
}

// Interface B
type b interface {
	funcB() int
	funcC() bool
}

// Class implementing a superset of interface A and B.
type impl struct{}

func (i impl) funcA() string     { return "working" }
func (i impl) funcB() int        { return 42 }
func (i impl) funcC() bool       { return true }
func (i impl) funcNotInterface() {}

// Functions using interface a / b and non-interface class.
func expectingA(i a) {
	fmt.Printf("%T\n", i)
}

func expectingB(i b) {
	fmt.Printf("%T\n", i)
}

func expectingImpl(i impl) {
	fmt.Printf("%T\n", i)
}

func main() {
	// Instantiate a new implementing class.
	i := impl{}

	// Since the class implements a it is possible to create an a interface directly.
	var ainterface a = i

	// It is not possible to directly assign the created a interface eventhough its underlaying
	// type points to a class which implements b interface.
	//var binterface b = ainterface

	// Assignment requires a type assertion.
	var binterface b
	var ok bool
	if binterface, ok = ainterface.(impl); !ok {
		panic("Don't use panic for this :)")
	}

	// It is not possible to call a function expecting a non-interface without
	// type asserting.
	expectingImpl(i)
	//expectingImpl(ainterface)
	//expectingImpl(binterface)

	// It is possible to call a function expecting an interface using both the interface
	// and non-interface values.
	expectingA(i)
	expectingA(ainterface)
	//expectingA(binterface)

	// It is possible to call a function expecting an interface using both the interface
	// and non-interface values.
	expectingB(i)
	//expectingB(ainterface)
	expectingB(binterface)

	// Implementing classes does not need to impl. only functions required by the interface.
	i.funcNotInterface()
}
