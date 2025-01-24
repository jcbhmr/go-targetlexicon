package targetlexicon

import "fmt"

type Endianness int

const (
	EndiannessLittle Endianness = iota
	EndiannessBig
)

type PointerWidth int

const (
	PointerWidthU16 PointerWidth = iota
	PointerWidthU32
	PointerWidthU64
)

func (p PointerWidth) Bits() int {
	switch p {
	case PointerWidthU16:
		return 16
	case PointerWidthU32:
		return 32
	case PointerWidthU64:
		return 64
	default:
		panic(fmt.Sprintf("unknown pointer width %d", p))
	}
}

func (p PointerWidth) Bytes() int {
	switch p {
	case PointerWidthU16:
		return 2
	case PointerWidthU32:
		return 4
	case PointerWidthU64:
		return 8
	default:
		panic(fmt.Sprintf("unknown pointer width %d", p))
	}
}

// NON EXHAUSTIVE
type CallingConvention int

const (
	CallingConventionSystemV CallingConvention = iota
	CallingConventionWASMBasicCABI
	CallingConventionWindowsFastCall
	CallingConventionAppleAArch64
)

type Triple struct {
	Architecture    Architecture
	Vendor          Vendor
	OperatingSystem OperatingSystem
	Environment     Environment
	BinaryFormat    BinaryFormat
}

func castOk[T any](v any) bool {
	_, ok := v.(T)
	return ok
}

func showBinaryFormatWithNoOS(triple Triple) bool {
	if triple.BinaryFormat == BinaryFormatUnknown {
		return false
	}

	return triple.Environment != EnvironmentEABI && triple.Environment != EnvironmentEABIHF && triple.Environment != EnvironmentSGX && !castOk[ArchitectureAVR](triple.Architecture) && !castOk[ArchitectureWASM32](triple.Architecture) && !castOk[ArchitectureWASM64](triple.Architecture) && (castOk[OperatingSystemNone](triple.OperatingSystem) || castOk[OperatingSystemUnknown](triple.OperatingSystem))
}
