# Built-in Types
Go has many of the same built-in types as other languages: booleans, integers, floats, and strings. We'll look at these types and see how they work best in Go. Additionally there are some concepts that apply to all types.

## The Zero Value
Go, assigns a default zero value to any variable that is declared by not assigned a value. Having an explicit zero value makes code clearer and removes a source of bugs found in C and C++ programs.

## Literals
A *literal* in Go refers to writing out a number, character, or string. There are 4 common kinds of literals that you'll find in Go programs. 

*Integer literals* are a sequence of numbers normally base 10, but different prefixes are used to indicate other bases(`0b` : binary, `0o` : octal, `0x` : hexadecimal). 

*Floating point literals* have decimal points to indicate the fractional portion of the value. They can also have an exponent specified with the letter e, such as `6.03e23`.

*Rune literals* represent characters and are surrounded by single quotes. Single and double quotes are not interchangeable in Go. Rune literals can be written as single Unicode characters (`'a'`), 8-bit octal numbers (`'\141'`), 8-bit hexadecimal numbers (`'\x61'`), 16-bit hexadecimal numbers (`'\u0061'`), or 32-bit Unicode numbers (`'\U00000061'`).

There are two different ways to indicate *string literals*. Most of the time, you should use double quotes to create an *interpreted string literal*. These contain zero or more rune literals, in any of the forms allowed. An example `"Greetings and\n\"Salutations\""`. If you need to include backslashes, double quotes, or newlines in your string, use a *raw string literal*. These are delimited with backquotes and can contain any literal character except a backquote:

``` Go
`Greetings and
"Salutations"`
```

Literals in Go are untyped meaning they can interact with any variable that's compatible with the literal. We can even use literals with user-defined types based on primitive types. That being said you can't assign a string literal to a variable with a numeric type or a float literal to an int. In cases where a type isn't explicitly declared Go uses the *default type* for a literal, more on this later.

## Booleans
The `bool` type represents Boolean variables. Variables of `bool` type can have on of two values: `true` or `false`. The zero value for a `bool` is `false`:

```Go
var flag bool // no value assigned, set to false
var isAwesome = true
```

## Numeric Types
Go has a large number of numeric types:

| Type name | Value range                                 |
| --------- | ------------------------------------------- |
| int8      | -128 to 127                                 |
| int16     | -32768 to 32767                             |
| int32     | –2147483648 to 2147483647                   |
| int64     | –9223372036854775808 to 9223372036854775807 |
| uint8     | 0 to 255                                    |
| uint16    | 0 to 65536                                  |
| uint32    | 0 to 4294967295                             |
| uint64    | 0 to 18446744073709551615                                            |

### The special integer types
A `byte` is an alias for `uint8`; it is legal to assign, compare, or perform mathematical operation between a `byte` and a `uint8`. However, `uint8` is rarely used, just use `byte` instead.

The second special name is `int`. On a 32-bit CPU, `int` is a 32-bit signed integer like an `int32`. On most 64-bit CPUs, `int` is a 64-bit singed integer, just like an `int64`. Because `int` insn't consistent from platform to platform, it is a compile-time error to assign, compare or perform mathematical operations between an `int` and an `int32` or `int64` without a type conversion. Integer literals default to being of `int` type.

The third special name is `uint`. It follows the same rules as `int` but it is unsigned. The two other special names for integer types are `rune` and `uintptr`.

### Choosing which integer to use
Given all of the available choices you might wonder when you should use each of them. There are 3 rules to follow here:
- If you are working with a binary format or network protocol that has an integer of a specific size or sign, use the corresponding integer type.
- If you are writing a library function that should work with any integer type, write a pari of functions, one with `int64` for the parameters and variables and the other with `uint64`. These are the idiomatic choices without generics (Go will have generics as of version 1.18) and because Go doesn't have function overloading. In this case those that call your code can use type conversion to pass values in and convert data that's returned.
- In all other cases just use `int`. 
	- Unless you *need* to be explicit about the size or sign of an integer for performance or integration purposes, use the `int` type. Consider any other type to be a premature optimization until proven otherwise.

### Floating point types
| Type name | Largest absolute value                          | Smallest (nonzero) absolute value             |
| --------- | ----------------------------------------------- | --------------------------------------------- |
| float32   | 3.40282346638528859811704183484516925440e+38    | 1.401298464324817070923729583289916131280e-45 |
| float64   | 1.797693134862315708145274237317043567981e +308 | 4.940656458412465441765687928682213723651e-324                                              |

Floating point in Go is similar to floating point math in other languages, they have a large range and limited precision. Picking with floating point type to use is simple: unless you have to be compatible with an existing format, use `float64`. Floating point literals have a default type of `float64` which helps mitigate floating point accuracy issues. Don't worry about the difference in memory size unless you have used the profiler to determine that it is a significant problem

In most cases you shouldn't use floating point numbers at all. Because floats aren't exact they can only be used in situation where inexact values are acceptable or the rules of floating point are well understood. This limits them to things like graphics and scientific operations.

## A Taste of Strings and Runes
The zero value for a string is the empty string. Like integers and floats they can be compared for equality using `==`, difference with `!=`, or ordering with `>, >=, <, <=`. They can be concatenated using the `+` operator.

Strings in Go are immutable, you can reassign the value of a string variable, but you cannot change the value of the string that is assigned to it.

Go also has a type that represents a single code point. The *rune* type is an alias for the `int32` type just like `byte` is an alias for `uint8`.

## Explicit Type Conversion
Go doesn't allow automatic type promotion between variables. Even different-sized integers and floats must be converted to the same type to interact. This makes it clear exactly what type you want without having to memorize any type conversion rules.

```Go
var x int = 10
var y float64 = 30.2
var z float64 = float64(x) + y
var d int = x + int(y)
```

Since all type conversion in Go are explicit, you cannot treat another Go type as a boolean. In many languages a nonzero number or nonempty string can be interpreted as a boolean `true`. Go doesn't allow truthiness, in fat, *no other type can be converted to a bool, implicitly or explicitly*. 

# Using const
Many languages have a way to declare a value is immutable. In Go this is done using the `const` keyword. 

```Go
const x int64 = 10

const (
	idKey = "id"
	nameKey = "name"
)

const z = 20 * 10

func main() {
	const y = "hello"
	
	fmt.Println(x)
	fmt.Println(y)

	x = x + 1
	y = "bye"

	fmt.Println(x)
	fmt.Println(y)
}
```

If you try running this code, compilation fails with the following error messages:

```shell
./const.go:20:4: cannot assign to x 
./const.go:21:4: cannot assign to y
```

Constants in Go are a way to give names to literals. They can only hold values that the compiler can figure out at compile time. This means they can be assigned:
- Numeric literals
- `true` and `false`
- String
- Runes
- The built-in functions `complex`, `real`, `imag`, `len`, and `cap`
- Expressions that consist of operators and the preceding values

Go doesn't provide a way to specify that a value calculated at runtime is immutable. There are no immutable arrays, slices, maps, or structs, and there's no way to declare that a field in a struct is immutable. 

## Typed and Untyped Constants
Constants can be typed or untyped. An untyped constant works exactly like a literal; it has no type of its own, but does have a default type that is used when no other type can be inferred. A typed constant can only be directly assigned to a variable with the same type.

# Next Section
[[Composite Types]]






















