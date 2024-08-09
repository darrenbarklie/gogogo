# Pointers

A **variable** is a piece of storage containing a value.

Variables created by declarations are identified by name:

```go
	x := 123
```

However, many variables are identified only by expressions:

```go
	y[i]
	z.f
```

All of these expressions read the value of a variable, except when
they appear on the left-hand side of an assignment, in which case a
new value is assigned to the variable (i.e. `y[i] = 4`).

A **pointer** is the address of a variable.

A pointer is thus the location at which a value is stored.

Not every value has an address, but every value does.

With a pointer, we can read or update the value of a variable _indirectly_,
without even knowing the name of the varibale, if indeed it has a name.

If a variable is declared `var x int`, the expression `&x`
("address of `x`") yields a pointer to an interger variable, that is,
a value of type `*int` ("pointer to `int`").

If this value is called `p`, we say "`p` points to `x`", or equivalently
"`p` contains the address of `x`". The variable to which `p` points is written
`*p`.

```go
	x := 1																// variable x declared, assigned 1
	p := &x																// p points to x / p contains address of x

	fmt.Printf("Variable value: %v", x)		// => "1"
	fmt.Printf("Pointer: %v\n", p)  			// => Pointer: 0xc000012028
	fmt.Printf("Pointer type: %T\n", p)		// => Pointer type: *int

	*p = 2
	fmt.Printf("Pointer value: %v", x)		// => "2"
```

The expression `*p` yields the value of that variable, an `int`, but
since `*p` denotes a variable, it may also apppear on the left-hand side of
an assignment, in which case the assignment updates the variable.

Each component of a variable of aggregate type – a field of a struct or an
element of an array – is also a variable and thuse has an address too.

Variables are sometimes described as _addressable_ values. Expressions that
denote variables are the only expressions to which the _address-of_ operator
`&` may be applied.

The zero value for a pointer of any type is `nil`.

The test `p != nil` is true is `p` points to a variable.

Pointers are comparable; two pointers are equal if and only if they point
to the same variable or both are nil.

```go
	var x, y int
	fmt.Println(&x == &x, &x == &y, &x == nil)
						 // => "true false false"
```

It is perfectly safe for a function to return the address of a local variable.
For instance, in the code below, the local variable `v` created by this
particular call to `f` will remain in existence even after the call has
returned, and the pointer `p` will still refer to it:

```go
	var p = f()

	func f() {
		v := 1
		return &v
	}
```

Each call of `f()` returns a distinct value:
`fmt.Println(f() == f()) // "false"`

Because a pointer contains the address of a variable, passing a pointer
argument to a function makes it possible for the function to update the
variable that was indirectly passed. For example, this function increments
the variable that its argument points to and returns the new value of the
variable so it may be used in an expression:

```go
	func incr(p *int) int {
		*p++ 									// increments what p points to; does not change p
		return *p
	}

	v := 1
	incr(&v)								// side effect: v in now 2
	fmt.Println(incr(&v))  // "3" (and v is 3)
```

Each time we take the address of a variable or copy a pointer, we create new
_aliases_ or ways to identify the same variable.

For example `p` is an alias for `v`.

Pointer aliasing is useful because it allows us to access a vaiable without
using its name, but this is a double-edged sword: to find all the statements
that access a variable, we have to know all its aliases.

It's not just pointers that create aliases; aliasing also occurs when we copy
values of other reference types like slices, maps, and channels, and even structs,
arrays and interfaces that contain these types.
