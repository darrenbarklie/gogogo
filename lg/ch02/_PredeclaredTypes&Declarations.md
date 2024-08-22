# Chapter 2: Predeclared Types and Declarations

One overarching princple: write your programs in a way that makes
your intentions clear.

## Predeclared Types

Go has _predeclared types_ built in the language:

- booleans
- integers
- floats
- strings

### Zero Value

A zero value is assigned to any varaible declared but not assinged a value.
This reduces code volume and bugs commonly found in C, C++.

## Literals

What I have:

```json
[
  {
    id: 1,
    name: "Organisation Name",
    organisation_type_id: 3,
    ...
  }
]
```

What I want:

```json
[
  {
    "id": 1,
    "name": "Organisation Name",
    "organisation_type": {
      "id": 3,
      "name": "Limited Company"
    }
  }
]
```
