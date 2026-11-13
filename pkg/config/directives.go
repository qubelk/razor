package config

// Directive [package] in Razor.toml
// Here's described package name and version.
// Package name is name of executable in ./build/ folder.
// Version now not used, but Razor create .VERSION file in ./build/ folder.
// Example:
// ```toml
//
// [package]
// name = "hello"
// version = "0.1.0"
//
// ```
// By default `razor init` creates package with current folder name and version 0.1.0.
// This directive is required.
type TomlPackage struct {
	Name    string `toml:"name"`
	Version string `toml:"version"`
}

// Directive [config] in Razor.toml.
// Here's described build config, like C standard, compiler or something else.
// Example:
// ```toml
//
// [config]
// compiler = "clang"
// standart = 11
//
// ```
// By default Razor create Razor.toml with clang and c11 standard.
// Now suppored only gcc or clang as compilers, and next standarts:
// [89, 99, 11, 17, 23].
// Gnu standarts not supported.
// This directive is required.
type TomlConfig struct {
	Compiler string `toml:"compiler"`
	Standard int    `toml:"standart"`
}
