package targetlexicon

import (
	"fmt"
	"runtime"
)

// CONST
var Host Triple

func init() {
	triple, _ := map[string]string{
		"aix/ppc64":       "",
		"android/arm64":   "",
		"darwin/amd64":    "x86_64-apple-darwin",
		"darwin/arm64":    "aarch64-apple-darwin",
		"dragonfly/amd64": "",
		"freebsd/386":     "",
		"freebsd/amd64":   "",
		"freebsd/arm":     "",
		"freebsd/arm64":   "",
		"illumos/amd64":   "",
		"ios/amd64":       "x86_64-apple-ios",
		"js/wasm":         "wasm32-unknown-unknown",
		"linux/386":       "",
		"linux/amd64":     "x86_64-unknown-linux-gnu",
		"linux/arm":       "",
		"linux/arm64":     "aarch64-unknown-linux-gnu",
		"linux/mips":      "",
		"linux/mips64":    "",
		"linux/mips64le":  "",
		"linux/mipsle":    "",
		"linux/ppc64":     "",
		"linux/ppc64le":   "",
		"linux/riscv64":   "",
		"linux/s390x":     "",
		"netbsd/386":      "",
		"netbsd/amd64":    "",
		"netbsd/arm":      "",
		"netbsd/arm64":    "",
		"openbsd/386":     "",
		"openbsd/amd64":   "",
		"openbsd/arm":     "",
		"openbsd/arm64":   "",
		"openbsd/mips64":  "",
		"plan9/386":       "",
		"plan9/amd64":     "",
		"plan9/arm":       "",
		"solaris/amd64":   "",
		"windows/386":     "",
		"windows/amd64":   "x86_64-pc-windows-msvc",
		"windows/arm":     "",
		"windows/arm64":   "",
	}[runtime.GOOS+"/"+runtime.GOARCH]
	if triple == "" {
		panic(fmt.Sprintf("could not determine host triple for %s/%s", runtime.GOOS, runtime.GOARCH))
	}
	triple2, err := TripleFromStr(triple)
	if err != nil {
		panic(err)
	}
	Host = triple2
}

func ArchitectureHost() Architecture {
	return Host.Architecture
}

func VendorHost() Vendor {
	return Host.Vendor
}

func OperatingSystemHost() OperatingSystem {
	return Host.OperatingSystem
}

func EnvironmentHost() Environment {
	return Host.Environment
}

func BinaryFormatHost() BinaryFormat {
	return Host.BinaryFormat
}

func TripleHost() Triple {
	return Host
}