package targetlexicon

import "fmt"

type Architecture interface {
	Endianness() (Endianness, bool)
	PointerWidth() (PointerWidth, bool)
	IsClever() bool
	IntoStr() string
	fmt.Stringer
	isArchitecture()
}
type ArchitectureUnknown struct{}
type ArchitectureARM struct {
	A ARMArchitecture
}
type ArchitectureAMDGCN struct{}
type ArchitectureAArch64 struct {
	A AArch64Architecture
}
type ArchitectureASMJS struct{}
type ArchitectureAVR struct{}
type ArchitectureBPFEB struct{}
type ArchitectureBPFEL struct{}
type ArchitectureHexagon struct{}
type ArchitectureX8632 struct {
	A X8632Architecture
}
type ArchitectureM68K struct{}
type ArchitectureLoongArch64 struct{}
type ArchitectureMIPS32 struct {
	A MIPS32Architecture
}
type ArchitectureMIPS64 struct {
	A MIPS64Architecture
}
type ArchitectureMSP430 struct{}
type ArchitectureNVPTX64 struct{}
type ArchitecturePulley32 struct{}
type ArchitecturePulley64 struct{}
type ArchitecturePulley32BE struct{}
type ArchitecturePulley64BE struct{}
type ArchitecturePowerPC struct{}
type ArchitecturePowerPC64 struct{}
type ArchitecturePowerPC64LE struct{}
type ArchitectureRISCV32 struct {
	A RISCV32Architecture
}
type ArchitectureRISCV64 struct {
	A RISCV64Architecture
}
type ArchitectureS390X struct{}
type ArchitectureSPARC struct{}
type ArchitectureSPARC64 struct{}
type ArchitectureSPARCV9 struct{}
type ArchitectureWASM32 struct{}
type ArchitectureWASM64 struct{}
type ArchitectureX8664 struct{}
type ArchitectureX8664H struct{}
type ArchitectureXtensa struct{}
type ArchitectureClever struct {
	A CleverArchitecture
}

var _ Architecture = (*ArchitectureUnknown)(nil)
var _ Architecture = (*ArchitectureARM)(nil)
var _ Architecture = (*ArchitectureAMDGCN)(nil)
var _ Architecture = (*ArchitectureAArch64)(nil)
var _ Architecture = (*ArchitectureASMJS)(nil)
var _ Architecture = (*ArchitectureAVR)(nil)
var _ Architecture = (*ArchitectureBPFEB)(nil)
var _ Architecture = (*ArchitectureBPFEL)(nil)
var _ Architecture = (*ArchitectureHexagon)(nil)
var _ Architecture = (*ArchitectureX8632)(nil)
var _ Architecture = (*ArchitectureM68K)(nil)
var _ Architecture = (*ArchitectureLoongArch64)(nil)
var _ Architecture = (*ArchitectureMIPS32)(nil)
var _ Architecture = (*ArchitectureMIPS64)(nil)
var _ Architecture = (*ArchitectureMSP430)(nil)
var _ Architecture = (*ArchitectureNVPTX64)(nil)
var _ Architecture = (*ArchitecturePulley32)(nil)
var _ Architecture = (*ArchitecturePulley64)(nil)
var _ Architecture = (*ArchitecturePulley32BE)(nil)
var _ Architecture = (*ArchitecturePulley64BE)(nil)
var _ Architecture = (*ArchitecturePowerPC)(nil)
var _ Architecture = (*ArchitecturePowerPC64)(nil)
var _ Architecture = (*ArchitecturePowerPC64LE)(nil)
var _ Architecture = (*ArchitectureRISCV32)(nil)
var _ Architecture = (*ArchitectureRISCV64)(nil)
var _ Architecture = (*ArchitectureS390X)(nil)
var _ Architecture = (*ArchitectureSPARC)(nil)
var _ Architecture = (*ArchitectureSPARC64)(nil)
var _ Architecture = (*ArchitectureSPARCV9)(nil)
var _ Architecture = (*ArchitectureWASM32)(nil)
var _ Architecture = (*ArchitectureWASM64)(nil)
var _ Architecture = (*ArchitectureX8664)(nil)
var _ Architecture = (*ArchitectureX8664H)(nil)
var _ Architecture = (*ArchitectureXtensa)(nil)
var _ Architecture = (*ArchitectureClever)(nil)

func (ArchitectureUnknown) isArchitecture()     {}
func (ArchitectureARM) isArchitecture()         {}
func (ArchitectureAMDGCN) isArchitecture()      {}
func (ArchitectureAArch64) isArchitecture()     {}
func (ArchitectureASMJS) isArchitecture()       {}
func (ArchitectureAVR) isArchitecture()         {}
func (ArchitectureBPFEB) isArchitecture()       {}
func (ArchitectureBPFEL) isArchitecture()       {}
func (ArchitectureHexagon) isArchitecture()     {}
func (ArchitectureX8632) isArchitecture()       {}
func (ArchitectureM68K) isArchitecture()        {}
func (ArchitectureLoongArch64) isArchitecture() {}
func (ArchitectureMIPS32) isArchitecture()      {}
func (ArchitectureMIPS64) isArchitecture()      {}
func (ArchitectureMSP430) isArchitecture()      {}
func (ArchitectureNVPTX64) isArchitecture()     {}
func (ArchitecturePulley32) isArchitecture()    {}
func (ArchitecturePulley64) isArchitecture()    {}
func (ArchitecturePulley32BE) isArchitecture()  {}
func (ArchitecturePulley64BE) isArchitecture()  {}
func (ArchitecturePowerPC) isArchitecture()     {}
func (ArchitecturePowerPC64) isArchitecture()   {}
func (ArchitecturePowerPC64LE) isArchitecture() {}
func (ArchitectureRISCV32) isArchitecture()     {}
func (ArchitectureRISCV64) isArchitecture()     {}
func (ArchitectureS390X) isArchitecture()       {}
func (ArchitectureSPARC) isArchitecture()       {}
func (ArchitectureSPARC64) isArchitecture()     {}
func (ArchitectureSPARCV9) isArchitecture()     {}
func (ArchitectureWASM32) isArchitecture()      {}
func (ArchitectureWASM64) isArchitecture()      {}
func (ArchitectureX8664) isArchitecture()       {}
func (ArchitectureX8664H) isArchitecture()      {}
func (ArchitectureXtensa) isArchitecture()      {}
func (ArchitectureClever) isArchitecture()      {}

// NON EXHAUSTIVE
type ARMArchitecture int

const (
	ARMArchitectureARM ARMArchitecture = iota
	ARMArchitectureARMEB
	ARMArchitectureARMV4
	ARMArchitectureARMV4T
	ARMArchitectureARMV5T
	ARMArchitectureARMV5TE
	ARMArchitectureARMV5TEJ
	ARMArchitectureARMV6
	ARMArchitectureARMV6J
	ARMArchitectureARMV6K
	ARMArchitectureARMV6Z
	ARMArchitectureARMV6KZ
	ARMArchitectureARMV6T2
	ARMArchitectureARMV6M
	ARMArchitectureARMV7
	ARMArchitectureARMV7A
	ARMArchitectureARMV7K
	ARMArchitectureARMV7VE
	ARMArchitectureARMV7M
	ARMArchitectureARMV7R
	ARMArchitectureARMV7S
	ARMArchitectureARMV8
	ARMArchitectureARMV8A
	ARMArchitectureARMV81A
	ARMArchitectureARMV82A
	ARMArchitectureARMV83A
	ARMArchitectureARMV84A
	ARMArchitectureARMV85A
	ARMArchitectureARMV8MBase
	ARMArchitectureARMV8MMain
	ARMArchitectureARMV8R
	ARMArchitectureARMEBV7R
	ARMArchitectureThumbEB
	ARMArchitectureThumbV4T
	ARMArchitectureThumbV5TE
	ARMArchitectureThumbV6M
	ARMArchitectureThumbV7A
	ARMArchitectureThumbV7EM
	ARMArchitectureThumbV7M
	ARMArchitectureThumbV7Neon
	ARMArchitectureThumbV8MBase
	ARMArchitectureThumbV8MMain
)

// NON EXHAUSTIVE
type AArch64Architecture int

const (
	AArch64ArchitectureAArch64 AArch64Architecture = iota
	AArch64ArchitectureAArch64BE
)

func (a ARMArchitecture) IsThumb() bool {
	switch a {
	case ARMArchitectureARM:
		fallthrough
	case ARMArchitectureARMEB:
		fallthrough
	case ARMArchitectureARMV4:
		fallthrough
	case ARMArchitectureARMV4T:
		fallthrough
	case ARMArchitectureARMV5T:
		fallthrough
	case ARMArchitectureARMV5TE:
		fallthrough
	case ARMArchitectureARMV5TEJ:
		fallthrough
	case ARMArchitectureARMV6:
		fallthrough
	case ARMArchitectureARMV6J:
		fallthrough
	case ARMArchitectureARMV6K:
		fallthrough
	case ARMArchitectureARMV6Z:
		fallthrough
	case ARMArchitectureARMV6KZ:
		fallthrough
	case ARMArchitectureARMV6T2:
		fallthrough
	case ARMArchitectureARMV6M:
		fallthrough
	case ARMArchitectureARMV7:
		fallthrough
	case ARMArchitectureARMV7A:
		fallthrough
	case ARMArchitectureARMV7K:
		fallthrough
	case ARMArchitectureARMV7VE:
		fallthrough
	case ARMArchitectureARMV7M:
		fallthrough
	case ARMArchitectureARMV7R:
		fallthrough
	case ARMArchitectureARMV7S:
		fallthrough
	case ARMArchitectureARMV8:
		fallthrough
	case ARMArchitectureARMV8A:
		fallthrough
	case ARMArchitectureARMV81A:
		fallthrough
	case ARMArchitectureARMV82A:
		fallthrough
	case ARMArchitectureARMV83A:
		fallthrough
	case ARMArchitectureARMV84A:
		fallthrough
	case ARMArchitectureARMV85A:
		fallthrough
	case ARMArchitectureARMV8MBase:
		fallthrough
	case ARMArchitectureARMV8MMain:
		fallthrough
	case ARMArchitectureARMV8R:
		fallthrough
	case ARMArchitectureARMEBV7R:
		return false
	case ARMArchitectureThumbEB:
		fallthrough
	case ARMArchitectureThumbV4T:
		fallthrough
	case ARMArchitectureThumbV5TE:
		fallthrough
	case ARMArchitectureThumbV6M:
		fallthrough
	case ARMArchitectureThumbV7A:
		fallthrough
	case ARMArchitectureThumbV7EM:
		fallthrough
	case ARMArchitectureThumbV7M:
		fallthrough
	case ARMArchitectureThumbV7Neon:
		fallthrough
	case ARMArchitectureThumbV8MBase:
		fallthrough
	case ARMArchitectureThumbV8MMain:
		return true
	default:
		panic(fmt.Sprintf("unknown ARM architecture %d", a))
	}
}

func (a ARMArchitecture) PointerWidth() PointerWidth {
	switch a {
	case ARMArchitectureARM:
		fallthrough
	case ARMArchitectureARMEB:
		fallthrough
	case ARMArchitectureARMV4:
		fallthrough
	case ARMArchitectureARMV4T:
		fallthrough
	case ARMArchitectureARMV5T:
		fallthrough
	case ARMArchitectureARMV5TE:
		fallthrough
	case ARMArchitectureARMV5TEJ:
		fallthrough
	case ARMArchitectureARMV6:
		fallthrough
	case ARMArchitectureARMV6J:
		fallthrough
	case ARMArchitectureARMV6K:
		fallthrough
	case ARMArchitectureARMV6Z:
		fallthrough
	case ARMArchitectureARMV6KZ:
		fallthrough
	case ARMArchitectureARMV6T2:
		fallthrough
	case ARMArchitectureARMV6M:
		fallthrough
	case ARMArchitectureARMV7:
		fallthrough
	case ARMArchitectureARMV7A:
		fallthrough
	case ARMArchitectureARMV7K:
		fallthrough
	case ARMArchitectureARMV7VE:
		fallthrough
	case ARMArchitectureARMV7M:
		fallthrough
	case ARMArchitectureARMV7R:
		fallthrough
	case ARMArchitectureARMV7S:
		fallthrough
	case ARMArchitectureARMV8:
		fallthrough
	case ARMArchitectureARMV8A:
		fallthrough
	case ARMArchitectureARMV81A:
		fallthrough
	case ARMArchitectureARMV82A:
		fallthrough
	case ARMArchitectureARMV83A:
		fallthrough
	case ARMArchitectureARMV84A:
		fallthrough
	case ARMArchitectureARMV85A:
		fallthrough
	case ARMArchitectureARMV8MBase:
		fallthrough
	case ARMArchitectureARMV8MMain:
		fallthrough
	case ARMArchitectureARMV8R:
		fallthrough
	case ARMArchitectureARMEBV7R:
		fallthrough
	case ARMArchitectureThumbEB:
		fallthrough
	case ARMArchitectureThumbV4T:
		fallthrough
	case ARMArchitectureThumbV5TE:
		fallthrough
	case ARMArchitectureThumbV6M:
		fallthrough
	case ARMArchitectureThumbV7A:
		fallthrough
	case ARMArchitectureThumbV7EM:
		fallthrough
	case ARMArchitectureThumbV7M:
		fallthrough
	case ARMArchitectureThumbV7Neon:
		fallthrough
	case ARMArchitectureThumbV8MBase:
		fallthrough
	case ARMArchitectureThumbV8MMain:
		return PointerWidthU32
	default:
		panic(fmt.Sprintf("unknown ARM architecture %d", a))
	}
}

func (a ARMArchitecture) Endianness() Endianness {
	switch a {
	case ARMArchitectureARM:
		fallthrough
	case ARMArchitectureARMV4:
		fallthrough
	case ARMArchitectureARMV4T:
		fallthrough
	case ARMArchitectureARMV5T:
		fallthrough
	case ARMArchitectureARMV5TE:
		fallthrough
	case ARMArchitectureARMV5TEJ:
		fallthrough
	case ARMArchitectureARMV6:
		fallthrough
	case ARMArchitectureARMV6J:
		fallthrough
	case ARMArchitectureARMV6K:
		fallthrough
	case ARMArchitectureARMV6Z:
		fallthrough
	case ARMArchitectureARMV6KZ:
		fallthrough
	case ARMArchitectureARMV6T2:
		fallthrough
	case ARMArchitectureARMV6M:
		fallthrough
	case ARMArchitectureARMV7:
		fallthrough
	case ARMArchitectureARMV7A:
		fallthrough
	case ARMArchitectureARMV7K:
		fallthrough
	case ARMArchitectureARMV7VE:
		fallthrough
	case ARMArchitectureARMV7M:
		fallthrough
	case ARMArchitectureARMV7R:
		fallthrough
	case ARMArchitectureARMV7S:
		fallthrough
	case ARMArchitectureARMV8:
		fallthrough
	case ARMArchitectureARMV8A:
		fallthrough
	case ARMArchitectureARMV81A:
		fallthrough
	case ARMArchitectureARMV82A:
		fallthrough
	case ARMArchitectureARMV83A:
		fallthrough
	case ARMArchitectureARMV84A:
		fallthrough
	case ARMArchitectureARMV85A:
		fallthrough
	case ARMArchitectureARMV8MBase:
		fallthrough
	case ARMArchitectureARMV8MMain:
		fallthrough
	case ARMArchitectureARMV8R:
		fallthrough
	case ARMArchitectureThumbV4T:
		fallthrough
	case ARMArchitectureThumbV5TE:
		fallthrough
	case ARMArchitectureThumbV6M:
		fallthrough
	case ARMArchitectureThumbV7A:
		fallthrough
	case ARMArchitectureThumbV7EM:
		fallthrough
	case ARMArchitectureThumbV7M:
		fallthrough
	case ARMArchitectureThumbV7Neon:
		fallthrough
	case ARMArchitectureThumbV8MBase:
		fallthrough
	case ARMArchitectureThumbV8MMain:
		return EndiannessLittle
	case ARMArchitectureARMEB:
		fallthrough
	case ARMArchitectureARMEBV7R:
		fallthrough
	case ARMArchitectureThumbEB:
		return EndiannessBig
	default:
		panic(fmt.Sprintf("unknown ARM architecture %d", a))
	}
}

func (a ARMArchitecture) IntoStr() string {
	switch a {
	case ARMArchitectureARM:
		return "arm"
	case ARMArchitectureARMEB:
		return "armeb"
	case ARMArchitectureARMV4:
		return "armv4"
	case ARMArchitectureARMV4T:
		return "armv4t"
	case ARMArchitectureARMV5T:
		return "armv5t"
	case ARMArchitectureARMV5TE:
		return "armv5te"
	case ARMArchitectureARMV5TEJ:
		return "armv5tej"
	case ARMArchitectureARMV6:
		return "armv6"
	case ARMArchitectureARMV6J:
		return "armv6j"
	case ARMArchitectureARMV6K:
		return "armv6k"
	case ARMArchitectureARMV6Z:
		return "armv6z"
	case ARMArchitectureARMV6KZ:
		return "armv6kz"
	case ARMArchitectureARMV6T2:
		return "armv6t2"
	case ARMArchitectureARMV6M:
		return "armv6m"
	case ARMArchitectureARMV7:
		return "armv7"
	case ARMArchitectureARMV7A:
		return "armv7a"
	case ARMArchitectureARMV7K:
		return "armv7k"
	case ARMArchitectureARMV7VE:
		return "armv7ve"
	case ARMArchitectureARMV7M:
		return "armv7m"
	case ARMArchitectureARMV7R:
		return "armv7r"
	case ARMArchitectureARMV7S:
		return "armv7s"
	case ARMArchitectureARMV8:
		return "armv8"
	case ARMArchitectureARMV8A:
		return "armv8a"
	case ARMArchitectureARMV81A:
		return "armv8.1a"
	case ARMArchitectureARMV82A:
		return "armv8.2a"
	case ARMArchitectureARMV83A:
		return "armv8.3a"
	case ARMArchitectureARMV84A:
		return "armv8.4a"
	case ARMArchitectureARMV85A:
		return "armv8.5a"
	case ARMArchitectureARMV8MBase:
		return "armv8m.base"
	case ARMArchitectureARMV8MMain:
		return "armv8m.main"
	case ARMArchitectureARMV8R:
		return "armv8r"
	case ARMArchitectureThumbEB:
		return "thumbeb"
	case ARMArchitectureThumbV4T:
		return "thumbv4t"
	case ARMArchitectureThumbV5TE:
		return "thumbv5te"
	case ARMArchitectureThumbV6M:
		return "thumbv6m"
	case ARMArchitectureThumbV7A:
		return "thumbv7a"
	case ARMArchitectureThumbV7EM:
		return "thumbv7em"
	case ARMArchitectureThumbV7M:
		return "thumbv7m"
	case ARMArchitectureThumbV7Neon:
		return "thumbv7neon"
	case ARMArchitectureThumbV8MBase:
		return "thumbv8m.base"
	case ARMArchitectureThumbV8MMain:
		return "thumbv8m.main"
	case ARMArchitectureARMEBV7R:
		return "armebv7r"
	default:
		panic(fmt.Sprintf("unknown ARM architecture %d", a))
	}
}

func (a AArch64Architecture) IsThumb() bool {
	switch a {
	case AArch64ArchitectureAArch64:
		fallthrough
	case AArch64ArchitectureAArch64BE:
		return false
	default:
		panic(fmt.Sprintf("unknown AArch64 architecture %d", a))
	}
}

func (a AArch64Architecture) PointerWidth() PointerWidth {
	switch a {
	case AArch64ArchitectureAArch64:
		fallthrough
	case AArch64ArchitectureAArch64BE:
		return PointerWidthU64
	default:
		panic(fmt.Sprintf("unknown AArch64 architecture %d", a))
	}
}

func (a AArch64Architecture) Endianness() Endianness {
	switch a {
	case AArch64ArchitectureAArch64:
		return EndiannessLittle
	case AArch64ArchitectureAArch64BE:
		return EndiannessBig
	default:
		panic(fmt.Sprintf("unknown AArch64 architecture %d", a))
	}
}

func (a AArch64Architecture) IntoStr() string {
	switch a {
	case AArch64ArchitectureAArch64:
		return "aarch64"
	case AArch64ArchitectureAArch64BE:
		return "aarch64_be"
	default:
		panic(fmt.Sprintf("unknown AArch64 architecture %d", a))
	}
}

// NON EXHAUSTIVE
type CleverArchitecture int

const (
	CleverArchitectureClever CleverArchitecture = iota
	CleverArchitectureClever10
)

func (c CleverArchitecture) IntoStr() string {
	switch c {
	case CleverArchitectureClever:
		return "clever"
	case CleverArchitectureClever10:
		return "clever1.0"
	default:
		panic(fmt.Sprintf("unknown Clever architecture %d", c))
	}
}

// NON EXHAUSTIVE
type RISCV32Architecture int

const (
	RISCV32ArchitectureRISCV32 RISCV32Architecture = iota
	RISCV32ArchitectureRISCV32GC
	RISCV32ArchitectureRISCV32I
	RISCV32ArchitectureRISCV32IM
	RISCV32ArchitectureRISCV32IMA
	RISCV32ArchitectureRISCV32IMAC
	RISCV32ArchitectureRISCV32IMAFC
	RISCV32ArchitectureRISCV32IMC
)

func (r RISCV32Architecture) IntoStr() string {
	switch r {
	case RISCV32ArchitectureRISCV32:
		return "riscv32"
	case RISCV32ArchitectureRISCV32GC:
		return "riscv32gc"
	case RISCV32ArchitectureRISCV32I:
		return "riscv32i"
	case RISCV32ArchitectureRISCV32IM:
		return "riscv32im"
	case RISCV32ArchitectureRISCV32IMA:
		return "riscv32ima"
	case RISCV32ArchitectureRISCV32IMAC:
		return "riscv32imac"
	case RISCV32ArchitectureRISCV32IMAFC:
		return "riscv32imafc"
	case RISCV32ArchitectureRISCV32IMC:
		return "riscv32imc"
	default:
		panic(fmt.Sprintf("unknown RISCV32 architecture %d", r))
	}
}

// NON EXHAUSTIVE
type RISCV64Architecture int

const (
	RISCV64ArchitectureRISCV64 RISCV64Architecture = iota
	RISCV64ArchitectureRISCV64GC
	RISCV64ArchitectureRISCV64IMAC
)

func (r RISCV64Architecture) IntoStr() string {
	switch r {
	case RISCV64ArchitectureRISCV64:
		return "riscv64"
	case RISCV64ArchitectureRISCV64GC:
		return "riscv64gc"
	case RISCV64ArchitectureRISCV64IMAC:
		return "riscv64imac"
	default:
		panic(fmt.Sprintf("unknown RISCV64 architecture %d", r))
	}
}

// NON EXHAUSTIVE
type X8632Architecture int

const (
	X8632ArchitectureI386 X8632Architecture = iota
	X8632ArchitectureI586
	X8632ArchitectureI686
)

func (x X8632Architecture) IntoStr() string {
	switch x {
	case X8632ArchitectureI386:
		return "i386"
	case X8632ArchitectureI586:
		return "i586"
	case X8632ArchitectureI686:
		return "i686"
	default:
		panic(fmt.Sprintf("unknown x86_32 architecture %d", x))
	}
}

// NON EXHAUSTIVE
type MIPS32Architecture int

const (
	MIPS32ArchitectureMIPS MIPS32Architecture = iota
	MIPS32ArchitectureMIPSEL
	MIPS32ArchitectureMIPSISA32R6
	MIPS32ArchitectureMIPSISA32R6EL
)

func (m MIPS32Architecture) IntoStr() string {
	switch m {
	case MIPS32ArchitectureMIPS:
		return "mips"
	case MIPS32ArchitectureMIPSEL:
		return "mipsel"
	case MIPS32ArchitectureMIPSISA32R6:
		return "mipsisa32r6"
	case MIPS32ArchitectureMIPSISA32R6EL:
		return "mipsisa32r6el"
	default:
		panic(fmt.Sprintf("unknown MIPS32 architecture %d", m))
	}
}

// NON EXHAUSTIVE
type MIPS64Architecture int

const (
	MIPS64ArchitectureMIPS64 MIPS64Architecture = iota
	MIPS64ArchitectureMIPS64EL
	MIPS64ArchitectureMIPSISA64R6
	MIPS64ArchitectureMIPSISA64R6EL
)

func (m MIPS64Architecture) IntoStr() string {
	switch m {
	case MIPS64ArchitectureMIPS64:
		return "mips64"
	case MIPS64ArchitectureMIPS64EL:
		return "mips64el"
	case MIPS64ArchitectureMIPSISA64R6:
		return "mipsisa64r6"
	case MIPS64ArchitectureMIPSISA64R6EL:
		return "mipsisa64r6el"
	default:
		panic(fmt.Sprintf("unknown MIPS64 architecture %d", m))
	}
}

type CustomVendor interface {
	AsStr() string
	isCustomVendor()
}
type CustomVendorOwned struct {
	A string
}
type CustomVendorStatic struct {
	A string
}

var _ CustomVendor = (*CustomVendorOwned)(nil)
var _ CustomVendor = (*CustomVendorStatic)(nil)

func (CustomVendorOwned) isCustomVendor()  {}
func (CustomVendorStatic) isCustomVendor() {}

func (c CustomVendorOwned) AsStr() string {
	return c.A
}
func (c CustomVendorStatic) AsStr() string {
	return c.A
}

type Vendor interface {
	isVendor()
}
type VendorUnknown struct{}
type VendorAMD struct{}
type VendorApple struct{}
type VendorEspressif struct{}
type VendorExperimental struct{}
type VendorFortanix struct{}
type VendorIBM struct{}
type VendorKMC struct{}
type VendorNintendo struct{}
type VendorNVIDIA struct{}
type VendorPC struct{}
type VendorRumprun struct{}
type VendorSun struct{}
type VendorUWP struct{}
type VendorWRS struct{}
type VendorCustom struct {
	A CustomVendor
}

var _ Vendor = (*VendorUnknown)(nil)
var _ Vendor = (*VendorAMD)(nil)
var _ Vendor = (*VendorApple)(nil)
var _ Vendor = (*VendorEspressif)(nil)
var _ Vendor = (*VendorExperimental)(nil)
var _ Vendor = (*VendorFortanix)(nil)
var _ Vendor = (*VendorIBM)(nil)
var _ Vendor = (*VendorKMC)(nil)
var _ Vendor = (*VendorNintendo)(nil)
var _ Vendor = (*VendorNVIDIA)(nil)
var _ Vendor = (*VendorPC)(nil)
var _ Vendor = (*VendorRumprun)(nil)
var _ Vendor = (*VendorSun)(nil)
var _ Vendor = (*VendorUWP)(nil)
var _ Vendor = (*VendorWRS)(nil)
var _ Vendor = (*VendorCustom)(nil)

func (VendorUnknown) isVendor()      {}
func (VendorAMD) isVendor()          {}
func (VendorApple) isVendor()        {}
func (VendorEspressif) isVendor()    {}
func (VendorExperimental) isVendor() {}
func (VendorFortanix) isVendor()     {}
func (VendorIBM) isVendor()          {}
func (VendorKMC) isVendor()          {}
func (VendorNintendo) isVendor()     {}
func (VendorNVIDIA) isVendor()       {}
func (VendorPC) isVendor()           {}
func (VendorRumprun) isVendor()      {}
func (VendorSun) isVendor()          {}
func (VendorUWP) isVendor()          {}
func (VendorWRS) isVendor()          {}
func (VendorCustom) isVendor()       {}

func (v VendorUnknown) AsStr() string {
	return "unknown"
}
func (v VendorAMD) AsStr() string {
	return "amd"
}
func (v VendorApple) AsStr() string {
	return "apple"
}
func (v VendorEspressif) AsStr() string {
	return "espressif"
}
func (v VendorExperimental) AsStr() string {
	return "experimental"
}
func (v VendorFortanix) AsStr() string {
	return "fortanix"
}
func (v VendorIBM) AsStr() string {
	return "ibm"
}
func (v VendorKMC) AsStr() string {
	return "kmc"
}
func (v VendorNintendo) AsStr() string {
	return "nintendo"
}
func (v VendorNVIDIA) AsStr() string {
	return "nvidia"
}
func (v VendorPC) AsStr() string {
	return "pc"
}
func (v VendorRumprun) AsStr() string {
	return "rumprun"
}
func (v VendorSun) AsStr() string {
	return "sun"
}
func (v VendorUWP) AsStr() string {
	return "uwp"
}
func (v VendorWRS) AsStr() string {
	return "wrs"
}
func (v VendorCustom) AsStr() string {
	return v.A.AsStr()
}

type DeploymentTarget struct {
	Major uint16
	Minor uint8
	Patch uint8
}

// NON EXHAUSTIVE
type OperatingSystem interface {
	IsLikeDarwin() bool
	AsStr() string
	isOperatingSystem()
}
type OperatingSystemUnknown struct{}
type OperatingSystemAIX struct{}
type OperatingSystemAMDHSA struct{}
type OperatingSystemBitrig struct{}
type OperatingSystemCloudABI struct{}
type OperatingSystemCUDA struct{}
type OperatingSystemDarwin struct {
	A *DeploymentTarget
}
type OperatingSystemDragonFly struct{}
type OperatingSystemEmscripten struct{}
type OperatingSystemESPIDF struct{}
type OperatingSystemFreeBSD struct{}
type OperatingSystemFuchsia struct{}
type OperatingSystemHaiku struct{}
type OperatingSystemHermit struct{}
type OperatingSystemHorizon struct{}
type OperatingSystemHurd struct{}
type OperatingSystemIllumos struct{}
type OperatingSystemIOS struct {
	A *DeploymentTarget
}
type OperatingSystemL4RE struct{}
type OperatingSystemLinux struct{}
type OperatingSystemMacOSX struct {
	A *DeploymentTarget
}
type OperatingSystemNebulet struct{}
type OperatingSystemNetBSD struct{}
type OperatingSystemNone struct{}
type OperatingSystemOpenBSD struct{}
type OperatingSystemPSP struct{}
type OperatingSystemRedox struct{}
type OperatingSystemSolaris struct{}
type OperatingSystemSolidASP3 struct{}
type OperatingSystemTVOS struct {
	A *DeploymentTarget
}
type OperatingSystemUEFI struct{}
type OperatingSystemVisionOS struct {
	A *DeploymentTarget
}
type OperatingSystemVXWorks struct{}
type OperatingSystemWASI struct{}
type OperatingSystemWASIP1 struct{}
type OperatingSystemWASIP2 struct{}
type OperatingSystemWatchOS struct {
	A *DeploymentTarget
}
type OperatingSystemWindows struct{}
type OperatingSystemXROS struct {
	A *DeploymentTarget
}

var _ OperatingSystem = (*OperatingSystemUnknown)(nil)
var _ OperatingSystem = (*OperatingSystemAIX)(nil)
var _ OperatingSystem = (*OperatingSystemAMDHSA)(nil)
var _ OperatingSystem = (*OperatingSystemBitrig)(nil)
var _ OperatingSystem = (*OperatingSystemCloudABI)(nil)
var _ OperatingSystem = (*OperatingSystemCUDA)(nil)
var _ OperatingSystem = (*OperatingSystemDarwin)(nil)
var _ OperatingSystem = (*OperatingSystemDragonFly)(nil)
var _ OperatingSystem = (*OperatingSystemEmscripten)(nil)
var _ OperatingSystem = (*OperatingSystemESPIDF)(nil)
var _ OperatingSystem = (*OperatingSystemFreeBSD)(nil)
var _ OperatingSystem = (*OperatingSystemFuchsia)(nil)
var _ OperatingSystem = (*OperatingSystemHaiku)(nil)
var _ OperatingSystem = (*OperatingSystemHermit)(nil)
var _ OperatingSystem = (*OperatingSystemHorizon)(nil)
var _ OperatingSystem = (*OperatingSystemHurd)(nil)
var _ OperatingSystem = (*OperatingSystemIllumos)(nil)
var _ OperatingSystem = (*OperatingSystemIOS)(nil)
var _ OperatingSystem = (*OperatingSystemL4RE)(nil)
var _ OperatingSystem = (*OperatingSystemLinux)(nil)
var _ OperatingSystem = (*OperatingSystemMacOSX)(nil)
var _ OperatingSystem = (*OperatingSystemNebulet)(nil)
var _ OperatingSystem = (*OperatingSystemNetBSD)(nil)
var _ OperatingSystem = (*OperatingSystemNone)(nil)
var _ OperatingSystem = (*OperatingSystemOpenBSD)(nil)
var _ OperatingSystem = (*OperatingSystemPSP)(nil)
var _ OperatingSystem = (*OperatingSystemRedox)(nil)
var _ OperatingSystem = (*OperatingSystemSolaris)(nil)
var _ OperatingSystem = (*OperatingSystemSolidASP3)(nil)
var _ OperatingSystem = (*OperatingSystemTVOS)(nil)
var _ OperatingSystem = (*OperatingSystemUEFI)(nil)
var _ OperatingSystem = (*OperatingSystemVisionOS)(nil)
var _ OperatingSystem = (*OperatingSystemVXWorks)(nil)
var _ OperatingSystem = (*OperatingSystemWASI)(nil)
var _ OperatingSystem = (*OperatingSystemWASIP1)(nil)
var _ OperatingSystem = (*OperatingSystemWASIP2)(nil)
var _ OperatingSystem = (*OperatingSystemWatchOS)(nil)
var _ OperatingSystem = (*OperatingSystemWindows)(nil)
var _ OperatingSystem = (*OperatingSystemXROS)(nil)

func (OperatingSystemUnknown) isOperatingSystem()    {}
func (OperatingSystemAIX) isOperatingSystem()        {}
func (OperatingSystemAMDHSA) isOperatingSystem()     {}
func (OperatingSystemBitrig) isOperatingSystem()     {}
func (OperatingSystemCloudABI) isOperatingSystem()   {}
func (OperatingSystemCUDA) isOperatingSystem()       {}
func (OperatingSystemDarwin) isOperatingSystem()     {}
func (OperatingSystemDragonFly) isOperatingSystem()  {}
func (OperatingSystemEmscripten) isOperatingSystem() {}
func (OperatingSystemESPIDF) isOperatingSystem()     {}
func (OperatingSystemFreeBSD) isOperatingSystem()    {}
func (OperatingSystemFuchsia) isOperatingSystem()    {}
func (OperatingSystemHaiku) isOperatingSystem()      {}
func (OperatingSystemHermit) isOperatingSystem()     {}
func (OperatingSystemHorizon) isOperatingSystem()    {}
func (OperatingSystemHurd) isOperatingSystem()       {}
func (OperatingSystemIllumos) isOperatingSystem()    {}
func (OperatingSystemIOS) isOperatingSystem()        {}
func (OperatingSystemL4RE) isOperatingSystem()       {}
func (OperatingSystemLinux) isOperatingSystem()      {}
func (OperatingSystemMacOSX) isOperatingSystem()     {}
func (OperatingSystemNebulet) isOperatingSystem()    {}
func (OperatingSystemNetBSD) isOperatingSystem()     {}
func (OperatingSystemNone) isOperatingSystem()       {}
func (OperatingSystemOpenBSD) isOperatingSystem()    {}
func (OperatingSystemPSP) isOperatingSystem()        {}
func (OperatingSystemRedox) isOperatingSystem()      {}
func (OperatingSystemSolaris) isOperatingSystem()    {}
func (OperatingSystemSolidASP3) isOperatingSystem()  {}
func (OperatingSystemTVOS) isOperatingSystem()       {}
func (OperatingSystemUEFI) isOperatingSystem()       {}
func (OperatingSystemVisionOS) isOperatingSystem()   {}
func (OperatingSystemVXWorks) isOperatingSystem()    {}
func (OperatingSystemWASI) isOperatingSystem()       {}
func (OperatingSystemWASIP1) isOperatingSystem()     {}
func (OperatingSystemWASIP2) isOperatingSystem()     {}
func (OperatingSystemWatchOS) isOperatingSystem()    {}
func (OperatingSystemWindows) isOperatingSystem()    {}
func (OperatingSystemXROS) isOperatingSystem()       {}

func darwinVersion(name string, deploymentTarget *DeploymentTarget) string {
	if deploymentTarget != nil {
		return fmt.Sprintf("%s%d.%d.%d", name, deploymentTarget.Major, deploymentTarget.Minor, deploymentTarget.Patch)
	} else {
		return name
	}
}

func (o OperatingSystemUnknown) AsStr() string {
	return "unknown"
}
func (o OperatingSystemAIX) AsStr() string {
	return "aix"
}
func (o OperatingSystemAMDHSA) AsStr() string {
	return "amdhsa"
}
func (o OperatingSystemBitrig) AsStr() string {
	return "bitrig"
}
func (o OperatingSystemCloudABI) AsStr() string {
	return "cloudabi"
}
func (o OperatingSystemCUDA) AsStr() string {
	return "cuda"
}
func (o OperatingSystemDarwin) AsStr() string {
	return darwinVersion("darwin", o.A)
}
func (o OperatingSystemDragonFly) AsStr() string {
	return "dragonfly"
}
func (o OperatingSystemEmscripten) AsStr() string {
	return "emscripten"
}
func (o OperatingSystemESPIDF) AsStr() string {
	return "espidf"
}
func (o OperatingSystemFreeBSD) AsStr() string {
	return "freebsd"
}
func (o OperatingSystemFuchsia) AsStr() string {
	return "fuchsia"
}
func (o OperatingSystemHaiku) AsStr() string {
	return "haiku"
}
func (o OperatingSystemHermit) AsStr() string {
	return "hermit"
}
func (o OperatingSystemHorizon) AsStr() string {
	return "horizon"
}
func (o OperatingSystemHurd) AsStr() string {
	return "hurd"
}
func (o OperatingSystemIllumos) AsStr() string {
	return "illumos"
}
func (o OperatingSystemIOS) AsStr() string {
	return darwinVersion("ios", o.A)
}
func (o OperatingSystemL4RE) AsStr() string {
	return "l4re"
}
func (o OperatingSystemLinux) AsStr() string {
	return "linux"
}
func (o OperatingSystemMacOSX) AsStr() string {
	return darwinVersion("macosx", o.A)
}
func (o OperatingSystemNebulet) AsStr() string {
	return "nebulet"
}
func (o OperatingSystemNetBSD) AsStr() string {
	return "netbsd"
}
func (o OperatingSystemNone) AsStr() string {
	return "none"
}
func (o OperatingSystemOpenBSD) AsStr() string {
	return "openbsd"
}
func (o OperatingSystemPSP) AsStr() string {
	return "psp"
}
func (o OperatingSystemRedox) AsStr() string {
	return "redox"
}
func (o OperatingSystemSolaris) AsStr() string {
	return "solaris"
}
func (o OperatingSystemSolidASP3) AsStr() string {
	return "solid_asp3"
}
func (o OperatingSystemTVOS) AsStr() string {
	return darwinVersion("tvos", o.A)
}
func (o OperatingSystemUEFI) AsStr() string {
	return "uefi"
}
func (o OperatingSystemVisionOS) AsStr() string {
	return darwinVersion("visionos", o.A)
}
func (o OperatingSystemVXWorks) AsStr() string {
	return "vxworks"
}
func (o OperatingSystemWASI) AsStr() string {
	return "wasi"
}
func (o OperatingSystemWASIP1) AsStr() string {
	return "wasip1"
}
func (o OperatingSystemWASIP2) AsStr() string {
	return "wasip2"
}
func (o OperatingSystemWatchOS) AsStr() string {
	return darwinVersion("watchos", o.A)
}
func (o OperatingSystemWindows) AsStr() string {
	return "windows"
}
func (o OperatingSystemXROS) AsStr() string {
	return darwinVersion("xros", o.A)
}

func (o OperatingSystemUnknown) IsLikeDarwin() bool {
	return false
}
func (o OperatingSystemAIX) IsLikeDarwin() bool {
	return false
}
func (o OperatingSystemAMDHSA) IsLikeDarwin() bool {
	return false
}
func (o OperatingSystemBitrig) IsLikeDarwin() bool {
	return false
}
func (o OperatingSystemCloudABI) IsLikeDarwin() bool {
	return false
}
func (o OperatingSystemCUDA) IsLikeDarwin() bool {
	return false
}
func (o OperatingSystemDarwin) IsLikeDarwin() bool {
	return true
}
func (o OperatingSystemDragonFly) IsLikeDarwin() bool {
	return false
}
func (o OperatingSystemEmscripten) IsLikeDarwin() bool {
	return false
}
func (o OperatingSystemESPIDF) IsLikeDarwin() bool {
	return false
}
func (o OperatingSystemFreeBSD) IsLikeDarwin() bool {
	return false
}
func (o OperatingSystemFuchsia) IsLikeDarwin() bool {
	return false
}
func (o OperatingSystemHaiku) IsLikeDarwin() bool {
	return false
}
func (o OperatingSystemHermit) IsLikeDarwin() bool {
	return false
}
func (o OperatingSystemHorizon) IsLikeDarwin() bool {
	return false
}
func (o OperatingSystemHurd) IsLikeDarwin() bool {
	return false
}
func (o OperatingSystemIllumos) IsLikeDarwin() bool {
	return false
}
func (o OperatingSystemIOS) IsLikeDarwin() bool {
	return true
}
func (o OperatingSystemL4RE) IsLikeDarwin() bool {
	return false
}
func (o OperatingSystemLinux) IsLikeDarwin() bool {
	return false
}
func (o OperatingSystemMacOSX) IsLikeDarwin() bool {
	return true
}
func (o OperatingSystemNebulet) IsLikeDarwin() bool {
	return false
}
func (o OperatingSystemNetBSD) IsLikeDarwin() bool {
	return false
}
func (o OperatingSystemNone) IsLikeDarwin() bool {
	return false
}
func (o OperatingSystemOpenBSD) IsLikeDarwin() bool {
	return false
}
func (o OperatingSystemPSP) IsLikeDarwin() bool {
	return false
}
func (o OperatingSystemRedox) IsLikeDarwin() bool {
	return false
}
func (o OperatingSystemSolaris) IsLikeDarwin() bool {
	return false
}
func (o OperatingSystemSolidASP3) IsLikeDarwin() bool {
	return false
}
func (o OperatingSystemTVOS) IsLikeDarwin() bool {
	return true
}
func (o OperatingSystemUEFI) IsLikeDarwin() bool {
	return false
}
func (o OperatingSystemVisionOS) IsLikeDarwin() bool {
	return true
}
func (o OperatingSystemVXWorks) IsLikeDarwin() bool {
	return false
}
func (o OperatingSystemWASI) IsLikeDarwin() bool {
	return false
}
func (o OperatingSystemWASIP1) IsLikeDarwin() bool {
	return false
}
func (o OperatingSystemWASIP2) IsLikeDarwin() bool {
	return false
}
func (o OperatingSystemWatchOS) IsLikeDarwin() bool {
	return true
}
func (o OperatingSystemWindows) IsLikeDarwin() bool {
	return false
}
func (o OperatingSystemXROS) IsLikeDarwin() bool {
	return true
}

type Environment int

const (
	EnvironmentUnknown Environment = iota
	EnvironmentAMDGIZ
	EnvironmentAndroid
	EnvironmentAndroidEABI
	EnvironmentEABI
	EnvironmentEABIHF
	EnvironmentGNU
	EnvironmentGNUABI64
	EnvironmentGNUEABI
	EnvironmentGNUEABIHF
	EnvironmentGNUSPE
	EnvironmentGNUX32
	EnvironmentGNUILP32
	EnvironmentGNULLVM
	EnvironmentHermitKernel
	EnvironmentHurdKernel
	EnvironmentLinuxKernel
	EnvironmentMacABI
	EnvironmentMusl
	EnvironmentMuslEABI
	EnvironmentMuslEABIHF
	EnvironmentMuslABI64
	EnvironmentMSVC
	EnvironmentNewlib
	EnvironmentNone
	EnvironmentKernel
	EnvironmentUClibc
	EnvironmentUClibcEABI
	EnvironmentUClibcEABIHF
	EnvironmentSGX
	EnvironmentSIM
	EnvironmentSoftfloat
	EnvironmentSPE
	EnvironmentThreads
	EnvironmentOHOS
)

func (e Environment) IntoStr() string {
	switch e {
	case EnvironmentUnknown:
		return "unknown"
	case EnvironmentAMDGIZ:
		return "amdgiz"
	case EnvironmentAndroid:
		return "android"
	case EnvironmentAndroidEABI:
		return "androideabi"
	case EnvironmentEABI:
		return "eabi"
	case EnvironmentEABIHF:
		return "eabihf"
	case EnvironmentGNU:
		return "gnu"
	case EnvironmentGNUABI64:
		return "gnuabi64"
	case EnvironmentGNUEABI:
		return "gnueabi"
	case EnvironmentGNUEABIHF:
		return "gnueabihf"
	case EnvironmentGNUSPE:
		return "gnuspe"
	case EnvironmentGNUX32:
		return "gnux32"
	case EnvironmentGNUILP32:
		return "gnuilp32"
	case EnvironmentGNULLVM:
		return "gnullvm"
	case EnvironmentHermitKernel:
		return "hermit"
	case EnvironmentHurdKernel:
		return "hurd"
	case EnvironmentLinuxKernel:
		return "linux"
	case EnvironmentMacABI:
		return "macabi"
	case EnvironmentMusl:
		return "musl"
	case EnvironmentMuslEABI:
		return "musleabi"
	case EnvironmentMuslEABIHF:
		return "musleabihf"
	case EnvironmentMuslABI64:
		return "muslabi64"
	case EnvironmentMSVC:
		return "msvc"
	case EnvironmentNewlib:
		return "newlib"
	case EnvironmentNone:
		return "none"
	case EnvironmentKernel:
		return "kernel"
	case EnvironmentUClibc:
		return "uclibc"
	case EnvironmentUClibcEABI:
		return "uclibeabi"
	case EnvironmentUClibcEABIHF:
		return "uclibeabihf"
	case EnvironmentSGX:
		return "sgx"
	case EnvironmentSIM:
		return "sim"
	case EnvironmentSoftfloat:
		return "softfloat"
	case EnvironmentSPE:
		return "spe"
	case EnvironmentThreads:
		return "threads"
	case EnvironmentOHOS:
		return "ohos"
	default:
		panic(fmt.Sprintf("unknown environment %d", e))
	}
}

type BinaryFormat int

const (
	BinaryFormatUnknown BinaryFormat = iota
	BinaryFormatELF
	BinaryFormatCOFF
	BinaryFormatMachO
	BinaryFormatWASM
	BinaryFormatXCOFF
)

func (b BinaryFormat) IntoStr() string {
	switch b {
	case BinaryFormatUnknown:
		return "unknown"
	case BinaryFormatELF:
		return "elf"
	case BinaryFormatCOFF:
		return "coff"
	case BinaryFormatMachO:
		return "macho"
	case BinaryFormatWASM:
		return "wasm"
	case BinaryFormatXCOFF:
		return "xcoff"
	default:
		panic(fmt.Sprintf("unknown binary format %d", b))
	}
}

func (a ArchitectureUnknown) Endianness() (Endianness, bool) {
	var e Endianness
	return e, false
}
func (a ArchitectureARM) Endianness() (Endianness, bool) {
	return a.A.Endianness(), true
}
func (a ArchitectureAArch64) Endianness() (Endianness, bool) {
	return a.A.Endianness(), true
}
func (a ArchitectureAMDGCN) Endianness() (Endianness, bool) {
	return EndiannessLittle, true
}
func (a ArchitectureASMJS) Endianness() (Endianness, bool) {
	return EndiannessLittle, true
}
func (a ArchitectureAVR) Endianness() (Endianness, bool) {
	return EndiannessLittle, true
}
func (a ArchitectureBPFEL) Endianness() (Endianness, bool) {
	return EndiannessLittle, true
}
func (a ArchitectureHexagon) Endianness() (Endianness, bool) {
	return EndiannessLittle, true
}
func (a ArchitectureX8632) Endianness() (Endianness, bool) {
	return EndiannessLittle, true
}
func (a ArchitectureLoongArch64) Endianness() (Endianness, bool) {
	return EndiannessLittle, true
}
func (a ArchitectureMIPS32) Endianness() (Endianness, bool) {
	switch a.A {
	case MIPS32ArchitectureMIPSEL:
		fallthrough
	case MIPS32ArchitectureMIPSISA32R6EL:
		return EndiannessLittle, true
	case MIPS32ArchitectureMIPS:
		fallthrough
	case MIPS32ArchitectureMIPSISA32R6:
		return EndiannessBig, true
	default:
		panic(fmt.Sprintf("unknown MIPS32 architecture %d", a.A))
	}
}
func (a ArchitectureMIPS64) Endianness() (Endianness, bool) {
	switch a.A {
	case MIPS64ArchitectureMIPS64EL:
		fallthrough
	case MIPS64ArchitectureMIPSISA64R6EL:
		return EndiannessLittle, true
	case MIPS64ArchitectureMIPS64:
		fallthrough
	case MIPS64ArchitectureMIPSISA64R6:
		return EndiannessBig, true
	default:
		panic(fmt.Sprintf("unknown MIPS64 architecture %d", a.A))
	}
}
func (a ArchitectureMSP430) Endianness() (Endianness, bool) {
	return EndiannessLittle, true
}
func (a ArchitectureNVPTX64) Endianness() (Endianness, bool) {
	return EndiannessLittle, true
}
func (a ArchitecturePulley32) Endianness() (Endianness, bool) {
	return EndiannessLittle, true
}
func (a ArchitecturePulley64) Endianness() (Endianness, bool) {
	return EndiannessLittle, true
}
func (a ArchitecturePowerPC64LE) Endianness() (Endianness, bool) {
	return EndiannessLittle, true
}
func (a ArchitectureRISCV32) Endianness() (Endianness, bool) {
	return EndiannessLittle, true
}
func (a ArchitectureRISCV64) Endianness() (Endianness, bool) {
	return EndiannessLittle, true
}
func (a ArchitectureWASM32) Endianness() (Endianness, bool) {
	return EndiannessLittle, true
}
func (a ArchitectureWASM64) Endianness() (Endianness, bool) {
	return EndiannessLittle, true
}
func (a ArchitectureX8664) Endianness() (Endianness, bool) {
	return EndiannessLittle, true
}
func (a ArchitectureX8664H) Endianness() (Endianness, bool) {
	return EndiannessLittle, true
}
func (a ArchitectureXtensa) Endianness() (Endianness, bool) {
	return EndiannessLittle, true
}
func (a ArchitectureClever) Endianness() (Endianness, bool) {
	return EndiannessLittle, true
}
func (a ArchitectureBPFEB) Endianness() (Endianness, bool) {
	return EndiannessBig, true
}
func (a ArchitectureM68K) Endianness() (Endianness, bool) {
	return EndiannessBig, true
}
func (a ArchitecturePowerPC) Endianness() (Endianness, bool) {
	return EndiannessBig, true
}
func (a ArchitecturePowerPC64) Endianness() (Endianness, bool) {
	return EndiannessBig, true
}
func (a ArchitecturePulley32BE) Endianness() (Endianness, bool) {
	return EndiannessBig, true
}
func (a ArchitecturePulley64BE) Endianness() (Endianness, bool) {
	return EndiannessBig, true
}
func (a ArchitectureS390X) Endianness() (Endianness, bool) {
	return EndiannessBig, true
}
func (a ArchitectureSPARC) Endianness() (Endianness, bool) {
	return EndiannessBig, true
}
func (a ArchitectureSPARC64) Endianness() (Endianness, bool) {
	return EndiannessBig, true
}
func (a ArchitectureSPARCV9) Endianness() (Endianness, bool) {
	return EndiannessBig, true
}

func (ArchitectureUnknown) PointerWidth() (PointerWidth, bool) {
	var p PointerWidth
	return p, false
}
func (a ArchitectureAVR) PointerWidth() (PointerWidth, bool) {
	return PointerWidthU16, true
}
func (a ArchitectureMSP430) PointerWidth() (PointerWidth, bool) {
	return PointerWidthU16, true
}
func (a ArchitectureARM) PointerWidth() (PointerWidth, bool) {
	return a.A.PointerWidth(), true
}
func (a ArchitectureAArch64) PointerWidth() (PointerWidth, bool) {
	return a.A.PointerWidth(), true
}
func (a ArchitectureASMJS) PointerWidth() (PointerWidth, bool) {
	return PointerWidthU32, true
}
func (a ArchitectureHexagon) PointerWidth() (PointerWidth, bool) {
	return PointerWidthU32, true
}
func (a ArchitectureX8632) PointerWidth() (PointerWidth, bool) {
	return PointerWidthU32, true
}
func (a ArchitectureRISCV32) PointerWidth() (PointerWidth, bool) {
	return PointerWidthU32, true
}
func (a ArchitectureSPARC) PointerWidth() (PointerWidth, bool) {
	return PointerWidthU32, true
}
func (a ArchitectureWASM32) PointerWidth() (PointerWidth, bool) {
	return PointerWidthU32, true
}
func (a ArchitectureM68K) PointerWidth() (PointerWidth, bool) {
	return PointerWidthU32, true
}
func (a ArchitectureMIPS32) PointerWidth() (PointerWidth, bool) {
	return PointerWidthU32, true
}
func (a ArchitecturePulley32) PointerWidth() (PointerWidth, bool) {
	return PointerWidthU32, true
}
func (a ArchitecturePulley32BE) PointerWidth() (PointerWidth, bool) {
	return PointerWidthU32, true
}
func (a ArchitecturePowerPC) PointerWidth() (PointerWidth, bool) {
	return PointerWidthU32, true
}
func (a ArchitectureXtensa) PointerWidth() (PointerWidth, bool) {
	return PointerWidthU32, true
}
func (a ArchitectureAMDGCN) PointerWidth() (PointerWidth, bool) {
	return PointerWidthU64, true
}
func (a ArchitectureBPFEB) PointerWidth() (PointerWidth, bool) {
	return PointerWidthU64, true
}
func (a ArchitectureBPFEL) PointerWidth() (PointerWidth, bool) {
	return PointerWidthU64, true
}
func (a ArchitecturePowerPC64LE) PointerWidth() (PointerWidth, bool) {
	return PointerWidthU64, true
}
func (a ArchitectureRISCV64) PointerWidth() (PointerWidth, bool) {
	return PointerWidthU64, true
}
func (a ArchitectureX8664) PointerWidth() (PointerWidth, bool) {
	return PointerWidthU64, true
}
func (a ArchitectureX8664H) PointerWidth() (PointerWidth, bool) {
	return PointerWidthU64, true
}
func (a ArchitectureMIPS64) PointerWidth() (PointerWidth, bool) {
	return PointerWidthU64, true
}
func (a ArchitectureNVPTX64) PointerWidth() (PointerWidth, bool) {
	return PointerWidthU64, true
}
func (a ArchitecturePulley64) PointerWidth() (PointerWidth, bool) {
	return PointerWidthU64, true
}
func (a ArchitecturePulley64BE) PointerWidth() (PointerWidth, bool) {
	return PointerWidthU64, true
}
func (a ArchitecturePowerPC64) PointerWidth() (PointerWidth, bool) {
	return PointerWidthU64, true
}
func (a ArchitectureS390X) PointerWidth() (PointerWidth, bool) {
	return PointerWidthU64, true
}
func (a ArchitectureSPARC64) PointerWidth() (PointerWidth, bool) {
	return PointerWidthU64, true
}
func (a ArchitectureSPARCV9) PointerWidth() (PointerWidth, bool) {
	return PointerWidthU64, true
}
func (a ArchitectureLoongArch64) PointerWidth() (PointerWidth, bool) {
	return PointerWidthU64, true
}
func (a ArchitectureWASM64) PointerWidth() (PointerWidth, bool) {
	return PointerWidthU64, true
}
func (a ArchitectureClever) PointerWidth() (PointerWidth, bool) {
	return PointerWidthU64, true
}

func (a ArchitectureUnknown) IsClever() bool {
	return false
}
func (a ArchitectureARM) IsClever() bool {
	return false
}
func (a ArchitectureAArch64) IsClever() bool {
	return false
}
func (a ArchitectureAMDGCN) IsClever() bool {
	return false
}
func (a ArchitectureASMJS) IsClever() bool {
	return false
}
func (a ArchitectureAVR) IsClever() bool {
	return false
}
func (a ArchitectureBPFEL) IsClever() bool {
	return false
}
func (a ArchitectureHexagon) IsClever() bool {
	return false
}
func (a ArchitectureX8632) IsClever() bool {
	return false
}
func (a ArchitectureLoongArch64) IsClever() bool {
	return false
}
func (a ArchitectureMIPS32) IsClever() bool {
	return false
}
func (a ArchitectureMIPS64) IsClever() bool {
	return false
}
func (a ArchitectureMSP430) IsClever() bool {
	return false
}
func (a ArchitectureNVPTX64) IsClever() bool {
	return false
}
func (a ArchitecturePulley32) IsClever() bool {
	return false
}
func (a ArchitecturePulley64) IsClever() bool {
	return false
}
func (a ArchitecturePowerPC64LE) IsClever() bool {
	return false
}
func (a ArchitectureRISCV32) IsClever() bool {
	return false
}
func (a ArchitectureRISCV64) IsClever() bool {
	return false
}
func (a ArchitectureWASM32) IsClever() bool {
	return false
}
func (a ArchitectureWASM64) IsClever() bool {
	return false
}
func (a ArchitectureX8664) IsClever() bool {
	return false
}
func (a ArchitectureX8664H) IsClever() bool {
	return false
}
func (a ArchitectureXtensa) IsClever() bool {
	return false
}
func (a ArchitectureClever) IsClever() bool {
	return true
}
func (a ArchitectureBPFEB) IsClever() bool {
	return false
}
func (a ArchitectureM68K) IsClever() bool {
	return false
}
func (a ArchitecturePowerPC) IsClever() bool {
	return false
}
func (a ArchitecturePowerPC64) IsClever() bool {
	return false
}
func (a ArchitecturePulley32BE) IsClever() bool {
	return false
}
func (a ArchitecturePulley64BE) IsClever() bool {
	return false
}
func (a ArchitectureS390X) IsClever() bool {
	return false
}
func (a ArchitectureSPARC) IsClever() bool {
	return false
}
func (a ArchitectureSPARC64) IsClever() bool {
	return false
}
func (a ArchitectureSPARCV9) IsClever() bool {
	return false
}

func (a ArchitectureARM) IntoStr() string {
	return a.A.IntoStr()
}
func (a ArchitectureAArch64) IntoStr() string {
	return a.A.IntoStr()
}
func (a ArchitectureUnknown) IntoStr() string {
	return "unknown"
}
func (a ArchitectureAMDGCN) IntoStr() string {
	return "amdgcn"
}
func (a ArchitectureASMJS) IntoStr() string {
	return "asmjs"
}
func (a ArchitectureAVR) IntoStr() string {
	return "avr"
}
func (a ArchitectureBPFEB) IntoStr() string {
	return "bpfeb"
}
func (a ArchitectureBPFEL) IntoStr() string {
	return "bpfel"
}
func (a ArchitectureHexagon) IntoStr() string {
	return "hexagon"
}
func (a ArchitectureX8632) IntoStr() string {
	return a.A.IntoStr()
}
func (a ArchitectureLoongArch64) IntoStr() string {
	return "loongarch64"
}
func (a ArchitectureM68K) IntoStr() string {
	return "m68k"
}
func (a ArchitectureMIPS32) IntoStr() string {
	return a.A.IntoStr()
}
func (a ArchitectureMIPS64) IntoStr() string {
	return a.A.IntoStr()
}
func (a ArchitectureMSP430) IntoStr() string {
	return "msp430"
}
func (a ArchitectureNVPTX64) IntoStr() string {
	return "nvptx64"
}
func (a ArchitecturePulley32) IntoStr() string {
	return "pulley32"
}
func (a ArchitecturePulley64) IntoStr() string {
	return "pulley64"
}
func (a ArchitecturePulley32BE) IntoStr() string {
	return "pulley32be"
}
func (a ArchitecturePulley64BE) IntoStr() string {
	return "pulley64be"
}
func (a ArchitecturePowerPC) IntoStr() string {
	return "powerpc"
}
func (a ArchitecturePowerPC64) IntoStr() string {
	return "powerpc64"
}
func (a ArchitecturePowerPC64LE) IntoStr() string {
	return "powerpc64le"
}
func (a ArchitectureRISCV32) IntoStr() string {
	return a.A.IntoStr()
}
func (a ArchitectureRISCV64) IntoStr() string {
	return a.A.IntoStr()
}
func (a ArchitectureS390X) IntoStr() string {
	return "s390x"
}
func (a ArchitectureSPARC) IntoStr() string {
	return "sparc"
}
func (a ArchitectureSPARC64) IntoStr() string {
	return "sparc64"
}
func (a ArchitectureSPARCV9) IntoStr() string {
	return "sparcv9"
}
func (a ArchitectureWASM32) IntoStr() string {
	return "wasm32"
}
func (a ArchitectureWASM64) IntoStr() string {
	return "wasm64"
}
func (a ArchitectureX8664) IntoStr() string {
	return "x86_64"
}
func (a ArchitectureX8664H) IntoStr() string {
	return "x86_64h"
}
func (a ArchitectureXtensa) IntoStr() string {
	return "xtensa"
}
func (a ArchitectureClever) IntoStr() string {
	return a.A.IntoStr()
}

func DefaultBinaryFormat(triple Triple) BinaryFormat {
	matchArchitecture := func() BinaryFormat {
		if _, ok := triple.Architecture.(ArchitectureWASM32); ok {
			return BinaryFormatWASM
		} else if _, ok := triple.Architecture.(ArchitectureWASM64); ok {
			return BinaryFormatWASM
		} else if _, ok := triple.Architecture.(ArchitectureUnknown); ok {
			return BinaryFormatUnknown
		} else {
			return BinaryFormatELF
		}
	}
	if _, ok := triple.OperatingSystem.(OperatingSystemNone); ok {
		switch triple.Environment {
		case EnvironmentEABI:
			fallthrough
		case EnvironmentEABIHF:
			return BinaryFormatELF
		default:
			return BinaryFormatUnknown
		}
	} else if _, ok := triple.OperatingSystem.(OperatingSystemAIX); ok {
		return BinaryFormatXCOFF
	} else if triple.OperatingSystem.IsLikeDarwin() {
		return BinaryFormatMachO
	} else if _, ok := triple.OperatingSystem.(OperatingSystemWindows); ok {
		return BinaryFormatCOFF
	} else if _, ok := triple.OperatingSystem.(OperatingSystemNebulet); ok {
		return matchArchitecture()
	} else if _, ok := triple.OperatingSystem.(OperatingSystemEmscripten); ok {
		return matchArchitecture()
	} else if _, ok := triple.OperatingSystem.(OperatingSystemVXWorks); ok {
		return matchArchitecture()
	} else if _, ok := triple.OperatingSystem.(OperatingSystemWASI); ok {
		return matchArchitecture()
	} else {
		return BinaryFormatELF
	}
}

func (a ARMArchitecture) String() string {
	return a.IntoStr()
}

func (a AArch64Architecture) String() string {
	return a.IntoStr()
}

func (a CleverArchitecture) String() string {
	return a.IntoStr()
}

func (a RISCV32Architecture) String() string {
	return a.IntoStr()
}

func (a RISCV64Architecture) String() string {
	return a.IntoStr()
}

func (a X8632Architecture) String() string {
	return a.IntoStr()
}

func (a MIPS32Architecture) String() string {
	return a.IntoStr()
}

func (a MIPS64Architecture) String() string {
	return a.IntoStr()
}

func (a ArchitectureAArch64) String() string {
	return a.IntoStr()
}
func (a ArchitectureAMDGCN) String() string {
	return a.IntoStr()
}
func (a ArchitectureASMJS) String() string {
	return a.IntoStr()
}
func (a ArchitectureAVR) String() string {
	return a.IntoStr()
}
func (a ArchitectureBPFEB) String() string {
	return a.IntoStr()
}
func (a ArchitectureBPFEL) String() string {
	return a.IntoStr()
}
func (a ArchitectureHexagon) String() string {
	return a.IntoStr()
}
func (a ArchitectureLoongArch64) String() string {
	return a.IntoStr()
}
func (a ArchitectureM68K) String() string {
	return a.IntoStr()
}
func (a ArchitectureMIPS32) String() string {
	return a.IntoStr()
}
func (a ArchitectureMIPS64) String() string {
	return a.IntoStr()
}
func (a ArchitectureMSP430) String() string {
	return a.IntoStr()
}
func (a ArchitectureNVPTX64) String() string {
	return a.IntoStr()
}
func (a ArchitecturePulley32) String() string {
	return a.IntoStr()
}
func (a ArchitecturePulley64) String() string {
	return a.IntoStr()
}
func (a ArchitecturePulley32BE) String() string {
	return a.IntoStr()
}
func (a ArchitecturePulley64BE) String() string {
	return a.IntoStr()
}
func (a ArchitecturePowerPC) String() string {
	return a.IntoStr()
}
func (a ArchitecturePowerPC64) String() string {
	return a.IntoStr()
}
func (a ArchitecturePowerPC64LE) String() string {
	return a.IntoStr()
}
func (a ArchitectureRISCV32) String() string {
	return a.IntoStr()
}
func (a ArchitectureRISCV64) String() string {
	return a.IntoStr()
}
func (a ArchitectureS390X) String() string {
	return a.IntoStr()
}
func (a ArchitectureSPARC) String() string {
	return a.IntoStr()
}
func (a ArchitectureSPARC64) String() string {
	return a.IntoStr()
}
func (a ArchitectureSPARCV9) String() string {
	return a.IntoStr()
}
func (a ArchitectureWASM32) String() string {
	return a.IntoStr()
}
func (a ArchitectureWASM64) String() string {
	return a.IntoStr()
}
func (a ArchitectureX8664) String() string {
	return a.IntoStr()
}
func (a ArchitectureX8664H) String() string {
	return a.IntoStr()
}
func (a ArchitectureXtensa) String() string {
	return a.IntoStr()
}
func (a ArchitectureClever) String() string {
	return a.IntoStr()
}
func (a ArchitectureUnknown) String() string {
	return a.IntoStr()
}
func (a ArchitectureARM) String() string {
	return a.IntoStr()
}
func (a ArchitectureX8632) String() string {
	return a.IntoStr()
}

func ARMArchitectureFromStr(s string) (ARMArchitecture, bool) {
	switch s {
	case "arm":
		return ARMArchitectureARM, true
	case "armeb":
		return ARMArchitectureARMEB, true
	case "armv4":
		return ARMArchitectureARMV4, true
	case "armv4t":
		return ARMArchitectureARMV4T, true
	case "armv5t":
		return ARMArchitectureARMV5T, true
	case "armv5te":
		return ARMArchitectureARMV5TE, true
	case "armv5tej":
		return ARMArchitectureARMV5TEJ, true
	case "armv6":
		return ARMArchitectureARMV6, true
	case "armv6k":
		return ARMArchitectureARMV6K, true
	case "armv6t2":
		return ARMArchitectureARMV6T2, true
	case "armv6kz":
		return ARMArchitectureARMV6KZ, true
	case "armv6m":
		return ARMArchitectureARMV6M, true
	case "armv7":
		return ARMArchitectureARMV7, true
	case "armv7a":
		return ARMArchitectureARMV7A, true
	case "armv7ve":
		return ARMArchitectureARMV7VE, true
	case "armv7r":
		return ARMArchitectureARMV7R, true
	case "armv7m":
		return ARMArchitectureARMV7M, true
	case "armv7s":
		return ARMArchitectureARMV7S, true
	case "armv8":
		return ARMArchitectureARMV8, true
	case "armv8a":
		return ARMArchitectureARMV8A, true
	case "armv8r":
		return ARMArchitectureARMV8R, true
	case "armv8m.base":
		return ARMArchitectureARMV8MBase, true
	case "armv8m.main":
		return ARMArchitectureARMV8MMain, true
	case "armv8.1a":
		return ARMArchitectureARMV81A, true
	case "armv8.2a":
		return ARMArchitectureARMV82A, true
	case "armv8.3a":
		return ARMArchitectureARMV83A, true
	case "armv8.4a":
		return ARMArchitectureARMV84A, true
	case "armv8.5a":
		return ARMArchitectureARMV85A, true
	case "thumbeb":
		return ARMArchitectureThumbEB, true
	case "thumbv4t":
		return ARMArchitectureThumbV4T, true
	case "thumbv5te":
		return ARMArchitectureThumbV5TE, true
	case "thumbv6m":
		return ARMArchitectureThumbV6M, true
	case "thumbv7a":
		return ARMArchitectureThumbV7A, true
	case "thumbv7em":
		return ARMArchitectureThumbV7EM, true
	case "thumbv7m":
		return ARMArchitectureThumbV7M, true
	case "thumbv7neon":
		return ARMArchitectureThumbV7Neon, true
	case "thumbv8m.base":
		return ARMArchitectureThumbV8MBase, true
	case "thumbv8m.main":
		return ARMArchitectureThumbV8MMain, true
	case "armebv7r":
		return ARMArchitectureARMEBV7R, true
	default:
		var a ARMArchitecture
		return a, false
	}
}
