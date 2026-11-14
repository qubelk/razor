# razor

Simple Cargo style build system for C.

## Overview

Razor command list:

```bash
# Initialize project, create Razor.toml and src/ directory.
razor init

# Create <workdir> folder and init project in it.
razor new <workdir>

# Build project into build/ directory, with name from Razor.toml.
razor build # or `razor run` for build and run.

# Remove build/ directory.
razor clean
```

Razor use incremental build with compare modify time of source and object.
Also razor use parallel build, by create one goroutine for each source file to object build.

Razor.toml configuration:

```toml
[package]
name = "example" # Name of executable
version = "0.1.0" # Razor create .VERSION file in build/ folder.

[config]
compiler = "clang" # Supports only gcc or clang.
standard = 11 # C standart.
```

## Made over the weekend just for fun.
