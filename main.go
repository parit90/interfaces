//An interface type is defined as a set of method signatures.
//A value of interface type can hold any value that implements those methods.

package main

import (
	intf "github.com/parit90/interfaces/first"
	sendf "github.com/parit90/interfaces/second"
)

/*
interfaces has two main uses
 1. The first is to use it as a type that can hold anything
 2. The second is a collection of functions .
*/

func main() {
	intf.Firsteg()
	intf.Secondeg()
	sendf.SecondIntf()
}
