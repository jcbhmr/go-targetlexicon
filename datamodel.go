package targetlexicon

// https://github.com/bytecodealliance/target-lexicon/blob/main/src/data_model.rs

import (
	"fmt"

	"github.com/gohugoio/hashstructure"
)

// The size of a type.
//
// Copyable via assignment. Clonable via copy. Debuggable via [dataModelSize.GoString]. Equalable via "==". Hashable via default [github.com/gohugoio/hashstructure.Hash].
type dataModelSize int

var _ fmt.GoStringer = (*dataModelSize)(nil)

const (
	SizeU8 dataModelSize = iota
	SizeU16
	SizeU32
	SizeU64
)

func (s dataModelSize) GoString() string {
	switch s {
	case SizeU8:
		return "U8"
	case SizeU16:
		return "U16"
	case SizeU32:
		return "U32"
	case SizeU64:
		return "U64"
	default:
		panic("invalid Size")
	}
}

// Return the number of bits this Size represents.
func (s dataModelSize) Bits() uint8 {
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
		panic("invalid Size")
	}
}

// Return the number of bytes in a size.
//
// A byte is assumed to be 8 bits.
func (s dataModelSize) Bytes() uint8 {
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
		panic("invalid Size")
	}
}

// The C data model used on a target.
//
// See also https://en.cppreference.com/w/c/language/arithmetic_types
//
// Not exhaustive. Copyable via assignment. Clonable via copy. Debuggable via [dataModelCDataModel.GoString]. Equalable via "==". Hashable via default [github.com/gohugoio/hashstructure.Hash].
type dataModelCDataModel int

var _ fmt.GoStringer = (*dataModelCDataModel)(nil)

const (
	// The data model used most commonly on Win16. long and pointer are 32 bits.
	CDataModelLP32 dataModelCDataModel = iota
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

func (c dataModelCDataModel) GoString() string {
	switch c {
	case CDataModelLP32:
		return "LP32"
	case CDataModelILP32:
		return "ILP32"
	case CDataModelLLP64:
		return "LLP64"
	case CDataModelLP64:
		return "LP64"
	case CDataModelILP64:
		return "ILP64"
	default:
		panic("invalid CDataModel")
	}
}

// The width of a pointer (in the default address space).
func (c dataModelCDataModel) PointerWidth() dataModelSize {
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
		panic("invalid CDataModel")
	}
}

// The size of a C short. This is required to be at least 16 bits.
func (c dataModelCDataModel) ShortSize() dataModelSize {
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
		panic("invalid CDataModel")
	}
}

// The size of a C int. This is required to be at least 16 bits.
func (c dataModelCDataModel) IntSize() dataModelSize {
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
		panic("invalid CDataModel")
	}
}

// The size of a C long. This is required to be at least 32 bits.
func (c dataModelCDataModel) LongSize() dataModelSize {
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
		panic("invalid CDataModel")
	}
}

// The size of a C long long. This is required (in C99+) to be at least 64 bits.
func (c dataModelCDataModel) LongLongSize() dataModelSize {
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
		panic("invalid CDataModel")
	}
}

// The size of a C float.
func (dataModelCDataModel) FloatSize() dataModelSize {
	// TODO: This is probably wrong on at least one architecture.
	return SizeU32
}

// The size of a C double.
func (dataModelCDataModel) DoubleSize() dataModelSize {
	// TODO: This is probably wrong on at least one architecture.
	return SizeU64
}
