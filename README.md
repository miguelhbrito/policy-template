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

The work 'level' is a palindrome, the request will be rejected.

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
