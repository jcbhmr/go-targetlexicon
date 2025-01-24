package targetlexicon

import "fmt"

type ParseError interface {
	fmt.Stringer
	error
	isParseError()
}
type ParseErrorUnrecognizedArchitecture struct {
	A string
}
type ParseErrorUnrecognizedVendor struct {
	A string
}
type ParseErrorUnrecognizedOperatingSystem struct {
	A string
}
type ParseErrorUnrecognizedEnvironment struct {
	A string
}
type ParseErrorUnrecognizedBinaryFormat struct {
	A string
}
type ParseErrorUnrecognizedField struct {
	A string
}

var _ ParseError = (*ParseErrorUnrecognizedArchitecture)(nil)
var _ ParseError = (*ParseErrorUnrecognizedVendor)(nil)
var _ ParseError = (*ParseErrorUnrecognizedOperatingSystem)(nil)
var _ ParseError = (*ParseErrorUnrecognizedEnvironment)(nil)
var _ ParseError = (*ParseErrorUnrecognizedBinaryFormat)(nil)
var _ ParseError = (*ParseErrorUnrecognizedField)(nil)

func (ParseErrorUnrecognizedArchitecture) isParseError()    {}
func (ParseErrorUnrecognizedVendor) isParseError()          {}
func (ParseErrorUnrecognizedOperatingSystem) isParseError() {}
func (ParseErrorUnrecognizedEnvironment) isParseError()     {}
func (ParseErrorUnrecognizedBinaryFormat) isParseError()    {}
func (ParseErrorUnrecognizedField) isParseError()           {}

func (p ParseErrorUnrecognizedArchitecture) String() string {
	return "unrecognized architecture: " + p.A
}
func (p ParseErrorUnrecognizedVendor) String() string {
	return "unrecognized vendor: " + p.A
}
func (p ParseErrorUnrecognizedOperatingSystem) String() string {
	return "unrecognized operating system: " + p.A
}
func (p ParseErrorUnrecognizedEnvironment) String() string {
	return "unrecognized environment: " + p.A
}
func (p ParseErrorUnrecognizedBinaryFormat) String() string {
	return "unrecognized binary format: " + p.A
}
func (p ParseErrorUnrecognizedField) String() string {
	return "unrecognized field: " + p.A
}

func (p ParseErrorUnrecognizedArchitecture) Error() string {
	return p.String()
}
func (p ParseErrorUnrecognizedVendor) Error() string {
	return p.String()
}
func (p ParseErrorUnrecognizedOperatingSystem) Error() string {
	return p.String()
}
func (p ParseErrorUnrecognizedEnvironment) Error() string {
	return p.String()
}
func (p ParseErrorUnrecognizedBinaryFormat) Error() string {
	return p.String()
}
func (p ParseErrorUnrecognizedField) Error() string {
	return p.String()
}
