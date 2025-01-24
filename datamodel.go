package targetlexicon

import "fmt"

// The size of a type.
type Size int

const (
	SizeU8 Size = iota
	SizeU16
	SizeU32
	SizeU64
)

// Return the number of bits this Size represents.
func (s Size) Bits() uint8 {
	switch s {
	case SizeU8:
		return 8
	case SizeU16:
		return 16
	case SizeU32:
		return 32
	case SizeU64:
		return 64
	default:
		panic(fmt.Sprintf("unknown size variant %d", s))
	}
}

// Return the number of bytes in a size.
//
// A byte is assumed to be 8 bits.
func (s Size) Bytes() uint8 {
	switch s {
	case SizeU8:
		return 1
	case SizeU16:
		return 2
	case SizeU32:
		return 4
	case SizeU64:
		return 8
	default:
		panic(fmt.Sprintf("unknown size variant %d", s))
	}
}

// The C data model used on a target.
//
// See also https://en.cppreference.com/w/c/language/arithmetic_types
//
// NOT EXHAUSTIVE
type CDataModel int

const (
	// The data model used most commonly on Win16. long and pointer are 32 bits.
	CDataModelLP32 CDataModel = iota
	// The data model used most commonly on Win32 and 32-bit Unix systems.
	//
	// int, long, and pointer are all 32 bits.
	CDataModelILP32
	// The data model used most commonly on Win64
	//
	// long long, and pointer are 64 bits.
	CDataModelLLP64
	// The data model used most commonly on 64-bit Unix systems.
	//
	// long, and pointer are 64 bits.
	CDataModelLP64
	// A rare data model used on early 64-bit Unix systems.
	//
	// int, long, and pointer are all 64 bits.
	CDataModelILP64
)

// The width of a pointer (in the default address space).
func (c CDataModel) PointerWidth() Size {
	switch c {
	case CDataModelLP32:
		fallthrough
	case CDataModelILP32:
		return SizeU32
	case CDataModelLLP64:
		fallthrough
	case CDataModelLP64:
		fallthrough
	case CDataModelILP64:
		return SizeU64
	default:
		panic(fmt.Sprintf("unknown data model %d", c))
	}
}

// The size of a C short. This is required to be at least 16 bits.
func (c CDataModel) ShortSize() Size {
	switch c {
	case CDataModelLP32:
		fallthrough
	case CDataModelILP32:
		fallthrough
	case CDataModelLLP64:
		fallthrough
	case CDataModelLP64:
		fallthrough
	case CDataModelILP64:
		return SizeU16
	default:
		panic(fmt.Sprintf("unknown data model %d", c))
	}
}

// The size of a C int. This is required to be at least 16 bits.
func (c CDataModel) IntSize() Size {
	switch c {
	case CDataModelLP32:
		return SizeU16
	case CDataModelILP32:
		fallthrough
	case CDataModelLLP64:
		fallthrough
	case CDataModelLP64:
		fallthrough
	case CDataModelILP64:
		return SizeU32
	default:
		panic(fmt.Sprintf("unknown data model %d", c))
	}
}

// The size of a C long. This is required to be at least 32 bits.
func (c CDataModel) LongSize() Size {
	switch c {
	case CDataModelLP32:
		fallthrough
	case CDataModelILP32:
		fallthrough
	case CDataModelLLP64:
		fallthrough
	case CDataModelILP64:
		return SizeU32
	case CDataModelLP64:
		return SizeU64
	default:
		panic(fmt.Sprintf("unknown data model %d", c))
	}
}

// The size of a C long long. This is required (in C99+) to be at least 64 bits.
func (c CDataModel) LongLongSize() Size {
	switch c {
	case CDataModelLP32:
		fallthrough
	case CDataModelILP32:
		fallthrough
	case CDataModelLLP64:
		fallthrough
	case CDataModelLP64:
		fallthrough
	case CDataModelILP64:
		return SizeU64
	default:
		panic(fmt.Sprintf("unknown data model %d", c))
	}
}

func (CDataModel) FloatSize() Size {
	return SizeU32
}

func (CDataModel) DoubleSize() Size {
	return SizeU64
}
