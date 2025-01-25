# Target lexicon for Go

🎯 Target "triple" support for Go

## Installation

```sh
go get github.com/jcbhmr/go-targetlexicon
```

## Usage

This is a library for managing targets for compilers and related tools.

Currently, the main feature is support for decoding [LLVM "triples"](https://clang.llvm.org/docs/CrossCompilation.html#target-triple), which
are strings that identify a particular target configuration. They're named
"triples" because historically they contained three fields, though over time
they've added additional fields. This library provides a `Triple` struct
containing enums for each of fields of a triple. `Triple` has an associated
`TripleFromStr()` and implements `fmt.Stringer` so it can be converted to and from the
conventional string representation of a triple.

`Triple` also has functions for querying a triple's endianness,
pointer bit width, and binary format.

And, `Triple` and the enum types have `<T>Host()` constructors, for targeting
the host.

It does not support reading JSON target files itself. To use it with a JSON
target file, construct a `Triple` using the value of the "llvm-target" field.

## Development

This Go module attempts to mirror the API of the Rust [`target-lexicon`](https://crates.io/crates/target-lexicon) crate. An attempt is also made to keep pace with the versioning of that crate. That means `target-lexicon` version 0.13.1 should correspond to `go-targetlexicon` version 0.13.1.

Tests are also ported from the Rust crate for completeness.

Some conventions:

- Rust `FromStr` is represented by `<T>FromStr()`. Since `FromStr` also extends the `str` type with a `.parse::<T>()` method, we use the equivalent Go convention of `Parse<T>(s string)` instead.
- Rust `Result<T, ()>` is a result type in Rust, but doesn't work well with the `(T, error)` "result type" in Go. Instead, we use the Go `(T, bool)` convention for these cases.
- Rust `Option<T>` in fields is a `*T` in Go.
- Rust enums with associated data are represented as Go interfaces with Go structs for the variant types. The variant types should all satisfy the enum's interface.
- Rust enums without associated data are represented as Go numeric (probably `int`) newtypes with `const ()` blocks.
- `std::fmt::Display` is represented by `fmt.Stringer` in Go.
- `std::fmt::Debug` is represented by `fmt.GoStringer` in Go. This is the `%#v` format specifier.
- `Eq`/`PartialEq` is translated to be a `(T) Equal(o T) bool` method in Go **if it needs it**. This works well with [github.com/google/go-cmp](https://github.com/google/go-cmp) and also fits the standard library convention.
- Non-`Eq` structs should have a `[0]func()` to make them uncomparable.
- Non-`Copy` structs should have a `noCopy noCopy` field where the `noCopy` struct has stub `.Lock()` and `.Unlock()` methods.
- `Clone` is represented by a `Clone() T` method in Go **if it needs it**. Simpler types can get away with the default copying behaviour.
- Things that are copiable should be passed by value unless needed. Things that are not should be passed by pointer.
- Don't know how to map Rust's `Hash` trait to Go.
- Static Rust methods like `Triple::default()` are mapped to `<T>Default()` in Go.
- No macros.
- Rust `ToString` or `.to_string()` is equivalent to Go's `fmt.Stringer` or `.String()` method.
- Tests are all in the same package, not the `<package>_test` package.

The generated `host.rs` is normally generated at build time by `build.rs` with the `TARGET` string. Instead we have to map Go's `${GOOS}/${GOARCH}` pairs to a valid LLVM target triple manually. There's also some gnu/musl and gnu/msvc detection that we do via C macros.

- ~~triple	A convenient syntax for triple literals.~~
- [x] DefaultToHost	A simple wrapper around Triple that provides an implementation of Default which defaults to Triple::host().
- [x] DefaultToUnknown	A simple wrapper around Triple that provides an implementation of Default which defaults to Triple::unknown().
- [x] DeploymentTarget	The minimum OS version that we’re compiling for.
- [ ] Triple	An LLVM target “triple”. Historically such things had three fields, though they’ve added additional fields over time.
- [ ] Aarch64Architecture
- [ ] Architecture	The “architecture” field, which in some cases also specifies a specific subarchitecture.
- [ ] ArmArchitecture
- [ ] BinaryFormat	The “binary format” field, which is usually omitted, and the binary format is implied by the other fields.
- [x] CDataModel	The C data model used on a target.
- [ ] CallingConvention	The calling convention, which specifies things like which registers are used for passing arguments, which registers are callee-saved, and so on.
- [ ] CustomVendor	A string for a Vendor::Custom that can either be used in const contexts or hold dynamic strings.
- [ ] Endianness	The target memory endianness.
- [ ] Environment	The “environment” field, which specifies an ABI environment on top of the operating system. In many configurations, this field is omitted, and the environment is implied by the operating system.
- [ ] Mips32Architecture	An enum for all 32-bit MIPS architectures (not just “MIPS32”).
- [ ] Mips64Architecture	An enum for all 64-bit MIPS architectures (not just “MIPS64”).
- [ ] OperatingSystem	The “operating system” field, which sometimes implies an environment, and sometimes isn’t an actual operating system.
- [ ] ParseError	An error returned from parsing a triple.
- [ ] PointerWidth	The width of a pointer (in the default address space).
- [ ] Riscv32Architecture	An enum for all 32-bit RISC-V architectures.
- [ ] Riscv64Architecture	An enum for all 64-bit RISC-V architectures.
- [x] Size	The size of a type.
- [ ] Vendor	The “vendor” field, which in practice is little more than an arbitrary modifier.
- [ ] X86_32Architecture	An enum for all 32-bit x86 architectures.
- [ ] HOST	The Triple of the current host.
