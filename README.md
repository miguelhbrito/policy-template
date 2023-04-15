# policy-template-open-suse

Technical Assessment - SUSE - Software Engineer - Container Technology

## Introduction

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

The word 'level' is a palindrome, as result the request will be rejected.

## Steps

First delete the vendor file and run `make`:

```console
rm -rf vendor
make
```

We can see the result by typing:

```console
file policy.wasm
policy.wasm: WebAssembly (wasm) binary module version 0x1 (MVP)
```

## Code organization

The code that takes care of parsing the settings can be found inside of the
`settings.go` file.

The actual validation code is defined inside of the `validate.go` file.

The `main.go` contains only the code which registers the entry points of the
policy.

## Testing

The unit tests can be run via a simple command:

```shell
make test
```

To run another test with requestOne.json, just type:
```shell
kwctl run policy.wasm --request-path requestOne.json 2> /dev/null | jq
```

The second test with requestTwo.json:
```shell
kwctl run policy.wasm --request-path requestTwo.json 2> /dev/null | jq
```