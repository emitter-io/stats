# Emitter Stats - JavaScript Library

This subpackage contains the necessary glue to build a javascript package for emitter stats. This package is built user [GopherJS](https://github.com/gopherjs/gopherjs) and will require it to be installed.


## Compiling

Build process is pretty straightforward, the command below compiles a minified version of the library which can then be used.
```
gopherjs build -m -o stats.js
```

## Usage

The library exposes the `restore` function through a `stats` global.

```
let snapshots = stats.restore(... bytes ...)
```