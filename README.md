# go-intro

A collection of small, self-contained Go examples for learning core language
features. Each subfolder is an independent `package main` you can run on its
own with `go run ./path/to/folder`.

## Structure

```
01-basics/
  hello-world/               a minimal "Hello, World!" program
  variables/                 declaring and assigning variables
  level-variables/            package-level (top of file) variable declarations
  unpacking-variables/       multiple assignment / unpacking
  datatypes/                  built-in data types

02-control-flow/
  conditions/                 if / else / switch
  loops/                      for loops

03-data-structures/
  arrays/                      fixed-size arrays
  slices/                      slices (append, len, cap)
  maps/                        maps
  map-pointers/                maps of/with pointers
  structs/                     struct types

04-functions/
  functions/                            basic function declarations
  closures-and-anonymous-functions/     closures & anonymous funcs
  first-class-functions/                functions as values
  defer/                                 the defer keyword

05-interfaces-and-pointers/
  pointers/                    pointer basics
  interfaces/                  interface types
  nil-interfaces/              nil interface pitfalls

06-json-and-http/
  json-marshal/                encoding/json marshal example (+ output.json)
  http-and-json/                a separate Go module (own go.mod) covering
                                 JSON marshalling and basic HTTP handlers/clients
```

## Running an example

Most examples live under the root module, so from the repo root:

```bash
go run ./01-basics/hello-world
go run ./03-data-structures/slices
```

`06-json-and-http/http-and-json` is its own Go module (it has its own
`go.mod`), so run those from inside that folder:

```bash
cd 06-json-and-http/http-and-json
go run ./cmd/json-marshalling
go run ./cmd/http-module/http-client
go run ./cmd/http-module/prevent-method
go run ./cmd/types
```

## What changed from the original layout

- Grouped 24 flat, inconsistently-named top-level folders into 6 numbered
  topic folders (basics, control flow, data structures, functions,
  interfaces/pointers, JSON & HTTP).
- Normalized all folder names to consistent kebab-case
  (`silde-array` → `slices`, `mapping_ptr` → `map-pointers`,
  `closures_and_anonymous_func` → `closures-and-anonymous-functions`,
  `1_Json_Marshalling` → `json-marshalling`, `3.http_module` → `http-module`,
  etc.), fixing a typo (`silde` → `slice`) along the way.
- Added a root `go.mod` so every standalone example can be run with
  `go run ./path/to/folder` from the repo root.
- Left `http-and-json` as its own module (it already had a `go.mod`), just
  renamed its internal folders for consistency.
