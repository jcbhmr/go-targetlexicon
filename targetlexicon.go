package targetlexicon

// https://github.com/bytecodealliance/target-lexicon/blob/main/src/lib.rs

import (
	"encoding/json"
	"fmt"
)

// TODO: Keep the documentation in sync with the underlying types & functions.

// The C data model used on a target.
//
// See also https://en.cppreference.com/w/c/language/arithmetic_types
//
// Not exhaustive. Copyable via assignment. Clonable via copy. Debuggable via [dataModelCDataModel.GoString]. Equalable via "==". Hashable via default [github.com/gohugoio/hashstructure.Hash].
//
// https://docs.rs/target-lexicon/latest/target_lexicon/enum.CDataModel.html
type CDataModel = dataModelCDataModel

// The size of a type.
//
// https://docs.rs/target-lexicon/latest/target_lexicon/enum.Size.html
type Size = dataModelSize

// Constant. Do not mutate. Use [Triple.Clone] to get a mutable clone.
var Host = hostHost

type ParseError = parseErrorParseError
type AArch64Architecture = targetsAArch64Architecture
type Architecture = targetsArchitecture
type ARMArchitecture = targetsARMArchitecture
type BinaryFormat = targetsBinaryFormat
type CustomVendor = targetsCustomVendor
type DeploymentTarget = targetsDeploymentTarget
type Environment = targetsEnvironment
type MIPS32Architecture = targetsMIPS32Architecture
type MIPS64Architecture = targetsMIPS64Architecture
type OperatingSystem = targetsOperatingSystem
type RISCV32Architecture = targetsRISCV32Architecture
type RISCV64Architecture = targetsRISCV64Architecture
type Vendor = targetsVendor
type X8664Architecture = targetsX8664Architecture
type CallingConvention = tripleCallingConvention
type Endianness = tripleEndianness
type PointerWidth = triplePointerWidth
type Triple = tripleTriple

// "go vet" will complain about attempts to copy anything that contains a value of this type due to the presence of the Lock() and Unlock() methods.
//
//	type NotCopyable struct {
//	    _ noCopy
//	    Name string
//	}
//	x := NotCopyable{Name: "Alan Turing"}
//	y := x
//	// "go vet" error: assignment copies lock value to y: NotCopyable contains noCopy
type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

// A simple wrapper around Triple that provides an accompanying DefaultToHostDefault() which defaults to TripleHost().
//
// Not copyable. Has a non-trivial [DefaultToHost.Clone]. Has a non-trivial [DefaultToHost.Equal].
type DefaultToHost struct {
	_ noCopy
	A Triple
}

var _ interface{ Clone() *DefaultToHost } = (*DefaultToHost)(nil)
var _ fmt.GoStringer = (*DefaultToHost)(nil)
var _ interface{ Equal(*DefaultToHost) bool } = (*DefaultToHost)(nil)

// Use this instead of
func (d *DefaultToHost) Clone() *DefaultToHost {
	return &DefaultToHost{
		A: d.A.Clone(),
	}
}

func (d *DefaultToHost) GoString() string {
	return fmt.Sprintf("DefaultToHost{%#v}", d.A)
}

// Use this instead of "=="
func (d *DefaultToHost) Equal(o *DefaultToHost) bool {
	return d.A.Equal(&o.A)
}

func DefaultToHostDefault() *DefaultToHost {
	return &DefaultToHost{
		A: TripleHost(),
	}
}

// A simple wrapper around Triple that provides an accompanying DefaultToUnknownDefault() which defaults to TripleUnknown().
//
// Has a non-trivial [DefaultToUnknown.Clone]. Has a non-trivial [DefaultToUnknown.Equal].
type DefaultToUnknown struct {
	A Triple
}

var _ interface{ Clone() *DefaultToUnknown } = (*DefaultToUnknown)(nil)
var _ fmt.GoStringer = (*DefaultToUnknown)(nil)
var _ interface{ Equal(*DefaultToUnknown) bool } = (*DefaultToUnknown)(nil)

func (d *DefaultToUnknown) Clone() *DefaultToUnknown {
	return &DefaultToUnknown{
		A: d.A.Clone(),
	}
}

func (d *DefaultToUnknown) GoString() string {
	return fmt.Sprintf("DefaultToUnknown{%#v}", d.A)
}

func (d *DefaultToUnknown) Equal(o *DefaultToUnknown) bool {
	return d.A.Equal(&o.A)
}

func DefaultToUnknownDefault() *DefaultToUnknown {
	return &DefaultToUnknown{
		A: TripleUnknown(),
	}
}

func (t *Triple) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

func (t *Triple) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	t2, ok := ParseTriple(s)
	if !ok {
		return fmt.Errorf("could not parse triple %q", s)
	}
	*t = *t2
	return nil
}
