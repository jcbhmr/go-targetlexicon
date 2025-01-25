package targetlexicon_test

import (
	"fmt"

	"github.com/jcbhmr/go-targetlexicon"
)

func ExampleHost() {
	pw, ok := targetlexicon.Host.PointerWidth()
	if !ok {
		panic("architecture should be known")
	}
	fmt.Printf("%v\n", pw.Bytes())
}

func ExampleMisc() {
	fmt.Printf("The host triple is %v.\n", targetlexicon.Host)

	t, ok := targetlexicon.TripleFromStr("riscv32-unknown-unknown")
	if !ok {
		panic("expected to recognize the RISC-V target")
	}
	e, ok := t.Endianness()
	if !ok {
		panic("expected to know the endianness of RISC-V")
	}
	fmt.Printf("The endianness of RISC-V is %#v.\n", e)
}