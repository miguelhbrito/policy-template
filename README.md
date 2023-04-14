# policy-template-open-suse

Technical Assessment

## Introduction

This repository contains a working policy written in Go.

The policy looks at the `labels` of a Kubernetes Pod and rejects the request
if the key is a palindrome.

The configuration of the policy is expressed via this structure:

```json
{
   "labels"{ 
    "env": "production", 
    "level": "debug"
   } 
}
```

## Code organization

The code that takes care of parsing the settings can be found inside of the
`settings.go` file.

The actual validation code is defined inside of the `validate.go` file.

The `main.go` contains only the code which registers the entry points of the
policy.

## Implementation details

> **DISCLAIMER:** WebAssembly is a constantly evolving topic. This document
> describes the status of the Go ecosystem at April 2021.

Currently the official Go compiler cannot produce WebAssembly binaries
that can be run **outside** of the browser. Because of that, Kubewarden Go
policies can be built only with the [TinyGo](https://tinygo.org/) compiler.

TinyGo doesn't yet support all the Go features (see [here](https://tinygo.org/lang-support/)
to see the current project status). Currently its biggest limitation
is the lack of a fully supported `reflect` package. Among other things, that
leads to the inability to use the `encoding/json` package against structures
and user defined types.

Kubewarden policies need to process JSON data like the policy settings and
the actual request received by Kubernetes.
However it's still possible to write a Kubewarden policy by using some 3rd party
libraries.

This is a list of libraries that can be useful when writing a Kubewarden
policy:

* [Kubernetes Go types](https://github.com/kubewarden/k8s-objects) for TinyGo:
  the official Kubernetes Go Types cannot be used with TinyGo. This module provides all the
  Kubernetes Types in a TinyGo-friendly way.
* [easyjson](https://github.com/mailru/easyjson/): this provides a way to
  marshal and unmarshal Go types without using reflection.
* Parsing JSON: queries against JSON documents can be written using the
  [gjson](https://github.com/tidwall/gjson) library. The library features a
  powerful query language that allows quick navigation of JSON documents and
  data retrieval.
* Generic `set` implementation: using [Set](https://en.wikipedia.org/wiki/Set_(abstract_data_type))
  data types can significantly reduce the amount of code inside of a policy,
  see the `union`, `intersection`, `difference`,... operations provided
  by a Set implementation.
  The [mapset](https://github.com/deckarep/golang-set) can be used when writing
  policies.

Last but not least, this policy takes advantage of helper functions provided
by [Kubewarden's Go SDK](https://github.com/kubewarden/policy-sdk-go).

## Testing

This policy comes with a set of unit tests implemented using the Go testing
framework.

As usual, the tests are defined inside of the `_test.go` files. Given these
tests are not part of the final WebAssembly binary, the official Go compiler
can be used to run them.

The unit tests can be run via a simple command:

```shell
make test
```

It's also important the test the final result of the TinyGo compilation:
the actual WebAssembly module.

This is done by a second set of end-to-end tests. These tests use the
`kwctl` cli provided by the Kubewarden project to load and execute
the policy.

## Automation

This project contains the following [GitHub Actions](https://docs.github.com/en/actions):

  * `unit-tests`: this action runs the Go unit tests
