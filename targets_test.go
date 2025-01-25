package targetlexicon

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRoundtripKnownTriples(t *testing.T) {
	// This list is constructed from:
	//  - targets emitted by "rustup target list"
	//  - targets emitted by "rustc +nightly --print target-list"
	//  - targets contributors have added
	targets := []string{
		"aarch64-apple-darwin",
		"aarch64-apple-ios",
		"aarch64-apple-ios-macabi",
		"aarch64-apple-ios-sim",
		"aarch64-apple-tvos",
		"aarch64-apple-tvos-sim",
		"aarch64-apple-visionos",
		"aarch64-apple-visionos-sim",
		"aarch64-apple-watchos",
		"aarch64-apple-watchos-sim",
		"aarch64_be-unknown-linux-gnu",
		"aarch64_be-unknown-linux-gnu_ilp32",
		"aarch64_be-unknown-netbsd",
		"aarch64-kmc-solid_asp3",
		"aarch64-linux-android",
		//"aarch64-nintendo-switch-freestanding", // TODO
		"aarch64-pc-windows-gnullvm",
		"aarch64-pc-windows-msvc",
		"aarch64-unknown-cloudabi",
		"aarch64-unknown-freebsd",
		"aarch64-unknown-fuchsia",
		"aarch64-unknown-hermit",
		"aarch64-unknown-illumos",
		"aarch64-unknown-linux-gnu",
		"aarch64-unknown-linux-gnu_ilp32",
		"aarch64-unknown-linux-musl",
		"aarch64-unknown-linux-ohos",
		"aarch64-unknown-netbsd",
		"aarch64-unknown-none",
		"aarch64-unknown-none-softfloat",
		//"aarch64-unknown-nto-qnx710", // TODO
		"aarch64-unknown-openbsd",
		"aarch64-unknown-redox",
		//"aarch64-unknown-teeos", // TODO
		"aarch64-unknown-uefi",
		"aarch64-uwp-windows-msvc",
		"aarch64-wrs-vxworks",
		//"arm64_32-apple-watchos", // TODO
		//"arm64e-apple-darwin", // TODO
		"amdgcn-amd-amdhsa",
		"amdgcn-amd-amdhsa-amdgiz",
		//"arm64e-apple-ios", // TODO
		//"arm64ec-pc-windows-msvc", // TODO
		"armeb-unknown-linux-gnueabi",
		"armebv7r-none-eabi",
		"armebv7r-none-eabihf",
		"arm-linux-androideabi",
		"arm-unknown-linux-gnueabi",
		"arm-unknown-linux-gnueabihf",
		"arm-unknown-linux-musleabi",
		"arm-unknown-linux-musleabihf",
		"armv4t-none-eabi",
		"armv4t-unknown-linux-gnueabi",
		"armv5te-none-eabi",
		"armv5te-unknown-linux-gnueabi",
		"armv5te-unknown-linux-musleabi",
		"armv5te-unknown-linux-uclibceabi",
		"armv6k-nintendo-3ds",
		"armv6-unknown-freebsd",
		"armv6-unknown-netbsd-eabihf",
		"armv7a-kmc-solid_asp3-eabi",
		"armv7a-kmc-solid_asp3-eabihf",
		"armv7a-none-eabi",
		"armv7a-none-eabihf",
		"armv7-apple-ios",
		"armv7k-apple-watchos",
		"armv7-linux-androideabi",
		"armv7r-none-eabi",
		"armv7r-none-eabihf",
		"armv7s-apple-ios",
		"armv7-unknown-cloudabi-eabihf",
		//"armv7-sony-vita-newlibeabihf", // TODO
		"armv7-unknown-freebsd",
		"armv7-unknown-linux-gnueabi",
		"armv7-unknown-linux-gnueabihf",
		"armv7-unknown-linux-musleabi",
		"armv7-unknown-linux-musleabihf",
		"armv7-unknown-linux-ohos",
		"armv7-unknown-linux-uclibceabi",
		"armv7-unknown-linux-uclibceabihf",
		"armv7-unknown-netbsd-eabihf",
		"armv7-wrs-vxworks-eabihf",
		"asmjs-unknown-emscripten",
		"armv8r-none-eabihf",
		//"avr-unknown-gnu-atmega328", // TODO
		"avr-unknown-unknown",
		"bpfeb-unknown-none",
		"bpfel-unknown-none",
		//"csky-unknown-linux-gnuabiv2", // TODO
		//"csky-unknown-linux-gnuabiv2hf", // TODO
		"hexagon-unknown-linux-musl",
		"hexagon-unknown-none-elf",
		"i386-apple-ios",
		//"i586-pc-nto-qnx700", // TODO
		"i586-pc-windows-msvc",
		"i586-unknown-linux-gnu",
		"i586-unknown-linux-musl",
		"i586-unknown-netbsd",
		"i686-apple-darwin",
		"i686-linux-android",
		"i686-apple-macosx10.7.0",
		"i686-pc-windows-gnu",
		"i686-pc-windows-gnullvm",
		"i686-pc-windows-msvc",
		"i686-unknown-cloudabi",
		"i686-unknown-dragonfly",
		"i686-unknown-freebsd",
		"i686-unknown-haiku",
		"i686-unknown-hurd-gnu",
		"i686-unknown-linux-gnu",
		"i686-unknown-linux-musl",
		"i686-unknown-netbsd",
		"i686-unknown-openbsd",
		"i686-unknown-redox",
		"i686-unknown-uefi",
		"i686-uwp-windows-gnu",
		"i686-uwp-windows-msvc",
		"i686-win7-windows-msvc",
		"i686-wrs-vxworks",
		"loongarch64-unknown-linux-gnu",
		"loongarch64-unknown-linux-musl",
		"loongarch64-unknown-none",
		"loongarch64-unknown-none-softfloat",
		"m68k-unknown-linux-gnu",
		"mips64el-unknown-linux-gnuabi64",
		"mips64el-unknown-linux-muslabi64",
		"mips64-openwrt-linux-musl",
		"mips64-unknown-linux-gnuabi64",
		"mips64-unknown-linux-muslabi64",
		"mipsel-sony-psp",
		//"mipsel-sony-psx", // TODO
		"mipsel-unknown-linux-gnu",
		"mipsel-unknown-linux-musl",
		"mipsel-unknown-linux-uclibc",
		"mipsel-unknown-netbsd",
		"mipsel-unknown-none",
		"mipsisa32r6el-unknown-linux-gnu",
		"mipsisa32r6-unknown-linux-gnu",
		"mipsisa64r6el-unknown-linux-gnuabi64",
		"mipsisa64r6-unknown-linux-gnuabi64",
		"mips-unknown-linux-gnu",
		"mips-unknown-linux-musl",
		"mips-unknown-linux-uclibc",
		"msp430-none-elf",
		"nvptx64-nvidia-cuda",
		"powerpc64-ibm-aix",
		"powerpc64le-unknown-freebsd",
		"powerpc64le-unknown-linux-gnu",
		"powerpc64le-unknown-linux-musl",
		"powerpc64-unknown-freebsd",
		"powerpc64-unknown-linux-gnu",
		"powerpc64-unknown-linux-musl",
		"powerpc64-unknown-openbsd",
		"powerpc64-wrs-vxworks",
		"powerpc-ibm-aix",
		"powerpc-unknown-freebsd",
		"powerpc-unknown-linux-gnu",
		"powerpc-unknown-linux-gnuspe",
		"powerpc-unknown-linux-musl",
		"powerpc-unknown-netbsd",
		"powerpc-unknown-openbsd",
		"powerpc-wrs-vxworks",
		"powerpc-wrs-vxworks-spe",
		"riscv32gc-unknown-linux-gnu",
		"riscv32gc-unknown-linux-musl",
		"riscv32imac-esp-espidf",
		"riscv32imac-unknown-none-elf",
		//"riscv32imac-unknown-xous-elf", // TODO
		"riscv32imafc-esp-espidf",
		"riscv32imafc-unknown-none-elf",
		"riscv32ima-unknown-none-elf",
		"riscv32imc-esp-espidf",
		"riscv32imc-unknown-none-elf",
		//"riscv32im-risc0-zkvm-elf", // TODO
		"riscv32im-unknown-none-elf",
		"riscv32i-unknown-none-elf",
		"riscv64gc-unknown-freebsd",
		"riscv64gc-unknown-fuchsia",
		"riscv64gc-unknown-hermit",
		"riscv64gc-unknown-linux-gnu",
		"riscv64gc-unknown-linux-musl",
		"riscv64gc-unknown-netbsd",
		"riscv64gc-unknown-none-elf",
		"riscv64gc-unknown-openbsd",
		"riscv64imac-unknown-none-elf",
		"riscv64-linux-android",
		"s390x-unknown-linux-gnu",
		"s390x-unknown-linux-musl",
		"sparc64-unknown-linux-gnu",
		"sparc64-unknown-netbsd",
		"sparc64-unknown-openbsd",
		"sparc-unknown-linux-gnu",
		"sparc-unknown-none-elf",
		"sparcv9-sun-solaris",
		"thumbv4t-none-eabi",
		"thumbv5te-none-eabi",
		"thumbv6m-none-eabi",
		"thumbv7a-pc-windows-msvc",
		"thumbv7a-uwp-windows-msvc",
		"thumbv7em-none-eabi",
		"thumbv7em-none-eabihf",
		"thumbv7m-none-eabi",
		"thumbv7neon-linux-androideabi",
		"thumbv7neon-unknown-linux-gnueabihf",
		"thumbv7neon-unknown-linux-musleabihf",
		"thumbv8m.base-none-eabi",
		"thumbv8m.main-none-eabi",
		"thumbv8m.main-none-eabihf",
		"wasm32-experimental-emscripten",
		"wasm32-unknown-emscripten",
		"wasm32-unknown-unknown",
		"wasm32-wasi",
		"wasm32-wasip1",
		"wasm32-wasip1-threads",
		"wasm32-wasip2",
		"wasm64-unknown-unknown",
		"wasm64-wasi",
		"x86_64-apple-darwin",
		"x86_64-apple-darwin23.6.0",
		"x86_64-apple-ios",
		"x86_64-apple-ios-macabi",
		"x86_64-apple-tvos",
		"x86_64-apple-watchos-sim",
		"x86_64-fortanix-unknown-sgx",
		"x86_64h-apple-darwin",
		"x86_64-linux-android",
		//"x86_64-pc-nto-qnx710", // TODO
		"x86_64-linux-kernel", // Changed to x86_64-unknown-none-linuxkernel in 1.53.0
		"x86_64-apple-macosx",
		"x86_64-apple-macosx10.7.0",
		"x86_64-pc-solaris",
		"x86_64-pc-windows-gnu",
		"x86_64-pc-windows-gnullvm",
		"x86_64-pc-windows-msvc",
		"x86_64-rumprun-netbsd", // Removed in 1.53.0
		"x86_64-sun-solaris",
		"x86_64-unknown-bitrig",
		"x86_64-unknown-cloudabi",
		"x86_64-unikraft-linux-musl",
		"x86_64-unknown-dragonfly",
		"x86_64-unknown-freebsd",
		"x86_64-unknown-fuchsia",
		"x86_64-unknown-haiku",
		"x86_64-unknown-hermit-kernel", // Changed to x86_64-unknown-none-hermitkernel in 1.53.0
		"x86_64-unknown-hermit",
		"x86_64-unknown-illumos",
		"x86_64-unknown-l4re-uclibc",
		"x86_64-unknown-linux-gnu",
		"x86_64-unknown-linux-gnux32",
		"x86_64-unknown-linux-musl",
		"x86_64-unknown-linux-none",
		"x86_64-unknown-linux-ohos",
		"x86_64-unknown-netbsd",
		"x86_64-unknown-none",
		"x86_64-unknown-none-hermitkernel",
		"x86_64-unknown-none-linuxkernel",
		"x86_64-unknown-openbsd",
		"x86_64-unknown-redox",
		"x86_64-unknown-uefi",
		"x86_64-uwp-windows-gnu",
		"x86_64-uwp-windows-msvc",
		"x86_64-win7-windows-msvc",
		"x86_64-wrs-vxworks",
		"xtensa-esp32-espidf",
		"clever-unknown-elf",
		"xtensa-esp32-none-elf",
		"xtensa-esp32s2-espidf",
		"xtensa-esp32s2-none-elf",
		"xtensa-esp32s3-espidf",
		"xtensa-esp32s3-none-elf",
		"zkasm-unknown-unknown",
	}

	for _, target := range targets {
		t2, ok := TripleFromStr(target)
		require.True(t, ok, "can't parse target")
		assert.NotEqual(t, t2.Architecture, ArchitectureUnknown{})
		assert.Equal(t, t2.String(), target, "%#+v", t2)
	}
}

func TestDefaultFormatToELF(t *testing.T) {
	t2, ok := TripleFromStr("riscv64")
	require.True(t, ok, "can't parse target")
	assert.Equal(t, t2.Architecture, ArchitectureRISCV64{RISCV64ArchitectureRISCV64})
	assert.Equal(t, t2.Vendor, VendorUnknown{})
	assert.Equal(t, t2.OperatingSystem, OperatingSystemUnknown{})
	assert.Equal(t, t2.Environment, EnvironmentUnknown)
	assert.Equal(t, t2.BinaryFormat, BinaryFormatELF)
}

func TestThumbV7EMNoneEABIHF(t *testing.T) {
	t2, ok := TripleFromStr("thumbv7em-none-eabihf")
	require.True(t, ok, "can't parse target")
	assert.Equal(t, t2.Architecture, ArchitectureARM{ARMArchitectureThumbV7EM})
	assert.Equal(t, t2.Vendor, VendorUnknown{})
	assert.Equal(t, t2.OperatingSystem, OperatingSystemNone{})
	assert.Equal(t, t2.Environment, EnvironmentEABIHF)
	assert.Equal(t, t2.BinaryFormat, BinaryFormatELF)
}

func TestFuchsiaRename(t *testing.T) {
	t1, ok1 := TripleFromStr("aarch64-fuchsia")
	t2, ok2 := TripleFromStr("aarch64-unknown-fuchsia")
	assert.True(t, cmp.Equal(t1, t2))
	assert.Equal(t, ok1, ok2)
	t1, ok1 = TripleFromStr("x86_64-fuchsia")
	t2, ok2 = TripleFromStr("x86_64-unknown-fuchsia")
	assert.True(t, cmp.Equal(t1, t2))
	assert.Equal(t, ok1, ok2)
}

func TestCustomVendors(t *testing.T) {
	// Test various invalid cases.
	_, ok := TripleFromStr("x86_64--linux")
	assert.False(t, ok)
	_, ok = TripleFromStr("x86_64-42-linux")
	assert.False(t, ok)
	_, ok = TripleFromStr("x86_64-__customvendor__-linux")
	assert.False(t, ok)
	_, ok = TripleFromStr("x86_64-^-linux")
	assert.False(t, ok)
	_, ok = TripleFromStr("x86_64- -linux")
	assert.False(t, ok)
	_, ok = TripleFromStr("x86_64-CustomVendor-linux")
	assert.False(t, ok)
	_, ok = TripleFromStr("x86_64-linux-linux")
	assert.False(t, ok)
	_, ok = TripleFromStr("x86_64-x86_64-linux")
	assert.False(t, ok)
	_, ok = TripleFromStr("x86_64-elf-linux")
	assert.False(t, ok)
	_, ok = TripleFromStr("x86_64-gnu-linux")
	assert.False(t, ok)
	_, ok = TripleFromStr("x86_64-linux-customvendor")
	assert.False(t, ok)
	_, ok = TripleFromStr("customvendor")
	assert.False(t, ok)
	_, ok = TripleFromStr("customvendor-x86_64")
	assert.False(t, ok)
	_, ok = TripleFromStr("x86_64-")
	assert.False(t, ok)
	_, ok = TripleFromStr("x86_64--")
	
	// Test various Unicode things.
	_, ok = TripleFromStr("x86_64-𝓬𝓾𝓼𝓽𝓸𝓶𝓿𝓮𝓷𝓭𝓸𝓻-linux")
	assert.False(t, ok, "unicode font hazard")
	_, ok = TripleFromStr("x86_64-ćúśtőḿvéńdőŕ-linux")
	assert.False(t, ok, "diacritical mark stripping hazard")
	_, ok = TripleFromStr("x86_64-customvendοr-linux")
	assert.False(t, ok, "homoglpyh hazard")
	_, ok = TripleFromStr("x86_64-customvendor-linux")
	assert.True(t, ok)
	_, ok = TripleFromStr("x86_64-ﬃ-linux")
	assert.False(t, ok, "normalization hazard")
	_, ok = TripleFromStr("x86_64-ffi-linux")
	assert.True(t, ok)
	_, ok = TripleFromStr("x86_64-custom‍vendor-linux")
	assert.False(t, ok, "zero-width character hazard")
	_, ok = TripleFromStr("x86_64-\uFEFFcustomvendor-linux")
	assert.False(t, ok, "BOM hazard")

	// Test some valid cases.
	t2, ok := TripleFromStr("x86_64-customvendor-linux")
	assert.True(t, ok, "can't parse target with custom vendor")
	assert.Equal(t, t2.Architecture, ArchitectureX8664{})
	assert.Equal(t, t2.Vendor, VendorCustom{CustomVendorStatic{"customvendor"}})
	assert.Equal(t, t2.OperatingSystem, OperatingSystemLinux{})
	assert.Equal(t, t2.Environment, EnvironmentUnknown)
	assert.Equal(t, t2.BinaryFormat, BinaryFormatELF)
	assert.Equal(t, t2.String(), "x86_64-customvendor-linux")

	t2, ok = TripleFromStr("x86_64-customvendor")
	assert.True(t, ok, "can't parse target with custom vendor")
	assert.Equal(t, t2.Architecture, ArchitectureX8664{})
	assert.Equal(t, t2.Vendor, VendorCustom{CustomVendorStatic{"customvendor"}})
	assert.Equal(t, t2.OperatingSystem, OperatingSystemUnknown{})
	assert.Equal(t, t2.Environment, EnvironmentUnknown)
	assert.Equal(t, t2.BinaryFormat, BinaryFormatELF)

	t2, ok = TripleFromStr("unknown-foo")
	assert.True(t, ok)
	assert.True(t, cmp.Equal(t2, Triple{
		Architecture: ArchitectureUnknown{},
		Vendor: 	 VendorCustom{CustomVendorStatic{"foo"}},
		OperatingSystem: OperatingSystemUnknown{},
		Environment: EnvironmentUnknown,
		BinaryFormat: BinaryFormatUnknown,
	}))
}

func TestDeploymentVersionParsing(t *testing.T) {
	t2, ok := TripleFromStr("aarch64-apple-macosx")
	assert.True(t, ok)
	assert.True(t, cmp.Equal(t2, Triple{
		Architecture: ArchitectureAArch64{AArch64ArchitectureAArch64},
		Vendor:       VendorApple{},
		OperatingSystem: OperatingSystemMacOSX{nil},
		Environment: EnvironmentUnknown,
		BinaryFormat: BinaryFormatMachO,
	}))

	t2, ok = TripleFromStr("aarch64-apple-macosx10.14.6")
	assert.True(t, ok)
	assert.True(t, cmp.Equal(t2, Triple{
		Architecture: ArchitectureAArch64{AArch64ArchitectureAArch64},
		Vendor:       VendorApple{},
		OperatingSystem: OperatingSystemMacOSX{&DeploymentTarget{10, 14, 6}},
		Environment: EnvironmentUnknown,
		BinaryFormat: BinaryFormatMachO,
	}))

	expected := &Triple{
		Architecture: ArchitectureX8664{},
		Vendor:       VendorApple{},
		OperatingSystem: OperatingSystemDarwin{&DeploymentTarget{23,0,0}},
		Environment: EnvironmentUnknown,
		BinaryFormat: BinaryFormatMachO,
	}
	t2, ok = TripleFromStr("x86_64-apple-darwin23")
	assert.True(t, ok)
	assert.True(t, cmp.Equal(expected.Clone(), t))
	t2, ok = TripleFromStr("x86_64-apple-darwin23.0")
	assert.True(t, ok)
	assert.True(t, cmp.Equal(expected.Clone(), t))
	t2, ok = TripleFromStr("x86_64-apple-darwin23.0.0")
	assert.True(t, ok)
	assert.True(t, cmp.Equal(expected, t))

	_, ok = TripleFromStr("x86_64-apple-darwin.")
	assert.False(t, ok)
	_, ok = TripleFromStr("x86_64-apple-darwin23.0.0.0")
	assert.False(t, ok)
}