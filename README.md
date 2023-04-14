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

## Testing

The unit tests can be run via a simple command:

```shell
make test
```

## Automation

This project contains the following [GitHub Actions](https://docs.github.com/en/actions):

  * `unit-tests`: this action runs the Go unit tests
