# razor

Simple Cargo style build system for C.

## Overview

Razor command list:

```bash
# Initialize project, create Razor.toml and src/ directory.
razor init

# Build project into build/ directory, with name from Razor.toml.
razor build # or `razor run` for build and run.

# Remove build/ directory.
razor clean
```

Razor use incremental build with compare modify time of source and object.
Also razor use parallel build, by create one goroutine for each source file to object build.

## Made over the weekend just for fun.
