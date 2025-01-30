# Target lexicon for Go

🎯 Target "triple" support for Go

<table align=center><td>

```go
pw, ok := targetlexicon.Host.PointerWidth()
if !ok {
    log.Fatal("architecture should be known")
}
fmt.Println(pw.Bytes())
// Possible output: 8
```

```go
triple, err := targetlexicon.ParseTriple("riscv32-unknown-unknown")
if err != nil {
    panic("riscv32-unknown-unknown should be a valid triple")
}
e, ok = triple.Endianness()
if !ok {
    panic("endianness of riscv32-unknown-unknown should be known")
}
fmt.Printf("The endianness of RISC-V is %#v.\n", e)
// Output: The endianness of RISC-V is <TODO>.
```

</table>

## Installation

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=Go&logoColor=FFFFFF)

```sh
go get github.com/jcbhmr/go-targetlexicon
```

## Usage

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=Go&logoColor=FFFFFF)

This is a library for managing targets for compilers and related tools.

Currently, the main feature is support for decoding [LLVM "triples"](https://clang.llvm.org/docs/CrossCompilation.html#target-triple), which
are strings that identify a particular target configuration. They're named
"triples" because historically they contained three fields, though over time
they've added additional fields. This library provides a `Triple` struct
containing enums for each of fields of a triple. `Triple` has an associated
`ParseTriple()` and implements `fmt.Formatter` so it can be converted to and from the
conventional string representation of a triple.

`Triple` also has functions for querying a triple's endianness,
pointer bit width, and binary format.

And, `Triple` and the enum types have `<T>Host()` constructors, for targeting
the host.

It does not support reading JSON target files itself. To use it with a JSON
target file, construct a `Triple` using the value of the "llvm-target" field.

## Development

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=Go&logoColor=FFFFFF)

This Go module attempts to mirror the API of the Rust [`target-lexicon`](https://crates.io/crates/target-lexicon) crate. An attempt is also made to keep pace with the versioning of that crate. That means `target-lexicon` version 0.13.1 should correspond to `go-targetlexicon` version 0.13.1.

Tests are also ported from the Rust crate for completeness.

Some conventions:

- Try to map rust `std::*` types to existing Go standard library types.
- Try to maintain Rust's copy, move, and cloning semantincs on ported Go types.
- The Rust `FromStr` trait is mapped to a Go `Parse<T>(s string) (T, error)`.
- The Rust `Error` trait is mapped to a Go `(T) Error() string` method.
- The Rust `Display` trait means implementing the Go `fmt.Formatter` interface. Since `Display` automatically implements `ToString` in Rust make sure to also provide a `.String()` method.
- The Rust `Debug` trait is mapped to a Go `fmt.GoStringer` interface. If the type also implements `fmt.Formatter` then it should properly handle the `%#v` format specifier.
- The Rust `Eq` and `PartialEq` traits are mapped to a Go `(T) Equal(o T) bool` method. Make sure to document the guarentees that `.Equal()` provides for `Eq` vs `PartialEq`. This works well with [github.com/google/go-cmp](https://github.com/google/go-cmp).
- Rust types not marked with `Copy` should include a `[0]sync.Mutex` field to warn against copying.
- Rust types not marked with `Eq` or `PartialEq` should include a `[0]func()` field to prevent comparison.
- Rust types marked with `Clone` may clarify that they permit "copy is clone". If cloning is non-trivial then a `.Clone() T` method should be provided.
- Rust types with `#[derive(Hash)]` are fine as-is. If the type has a custom hash implementation then a `.Hash() (uint64, error)` method should be provided to satisfy [github.com/gohugoio/hashstructure.Hashable](https://pkg.go.dev/github.com/gohugoio/hashstructure#Hashable).
- Static Rust methods like `Triple::default()` are mapped to `<T>Default()` in Go.
- Since Rust unit tests have total access to the entire package's internals, ported unit tests should use `package targetlexicon` instead of `package targetlexicon_test`.
- Rust crates are equivalent to Go packages. Avoid splitting things into `internal/*` Go packages since this causes circular import problems. Just keep everything in the root package.
- Rust `my_crate::my_module::my_submodule::MyType` is equivalent to Go `mycrate.MyModuleMySubmoduleMyType`.
- Each Rust file maps to a Go file with a similar name.
- Rust uses camel case like "HtmlId" while Go uses "HTMLID". The Go convention should be followed.

The generated `host.rs` is normally generated at build time by `build.rs` with the `TARGET` string. Instead we have to map Go's `$GOOS/$GOARCH` pairs to a valid LLVM target triple manually. There's also some `gnu`/`musl` and `gnu`/`msvc` detection that we do via C macros.

- [x] ~~**triple:** A convenient syntax for triple literals.~~
- [x] **DefaultToHost:** A simple wrapper around Triple that provides an implementation of Default which defaults to Triple::host().
- [x] **DefaultToUnknown:** A simple wrapper around Triple that provides an implementation of Default which defaults to Triple::unknown().
- [x] **DeploymentTarget:** The minimum OS version that we’re compiling for.
- [ ] **Triple:** An LLVM target “triple”. Historically such things had three fields, though they’ve added additional fields over time.
- [ ] Aarch64Architecture
- [ ] **Architecture:** The “architecture” field, which in some cases also specifies a specific subarchitecture.
- [ ] ArmArchitecture
- [ ] **BinaryFormat:** The “binary format” field, which is usually omitted, and the binary format is implied by the other fields.
- [x] **CDataModel:** The C data model used on a target.
- [ ] **CallingConvention:** The calling convention, which specifies things like which registers are used for passing arguments, which registers are callee-saved, and so on.
- [ ] **CustomVendor:** A string for a Vendor::Custom that can either be used in const contexts or hold dynamic strings.
- [ ] **Endianness:** The target memory endianness.
- [ ] **Environment:** The “environment” field, which specifies an ABI environment on top of the operating system. In many configurations, this field is omitted, and the environment is implied by the operating system.
- [ ] **Mips32Architecture:** An enum for all 32-bit MIPS architectures (not just “MIPS32”).
- [ ] **Mips64Architecture:** An enum for all 64-bit MIPS architectures (not just “MIPS64”).
- [ ] **OperatingSystem:** The “operating system” field, which sometimes implies an environment, and sometimes isn’t an actual operating system.
- [ ] **ParseError:** An error returned from parsing a triple.
- [ ] **PointerWidth:** The width of a pointer (in the default address space).
- [ ] **Riscv32Architecture:** An enum for all 32-bit RISC-V architectures.
- [ ] **Riscv64Architecture:** An enum for all 64-bit RISC-V architectures.
- [x] **Size:** The size of a type.
- [ ] **Vendor:** The “vendor” field, which in practice is little more than an arbitrary modifier.
- [ ] **X86_32Architecture:** An enum for all 32-bit x86 architectures.
- [ ] **HOST:** The Triple of the current host.
