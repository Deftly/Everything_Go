Previously we looked at the builtin types: numbers, booleans, and strings. Now we'll look at the composite types in Go, the built-in functions that support them, and some related best practices.

# Arrays - Too Rigid to Use Directly
Go has arrays, however, they are rarely used directly for a few reasons. First we'll cover array declaration syntax and use.

All of the elements in the array must be of the type that's specified and there are a few different declaration styles:

```Go
var x [3]int // specify the size and type, all indexes set to the zero value
var x = [3]int{10, 20, 30} // initial values specified with array literal
var x = [12]int{1, 5: 4, 6, 10: 100, 15} // sparse array
```

You can use `==` and `!=` to compare arrrays:

```Go
var x = [...]int{1, 2, 3}
var y = [3]int{1, 2, 3}
fmt.Println(x == y) // prints true
```

Go only has one-dimensional arrays, but you can simulate multidimensional arrays:

```Go
var x [2][3]int
```

This declares `x` to be an array of length 2 whose type is an array of `ints` of length 3. This sounds pedantic but there are languages with true matrix support; Go isn't one of them. 

You cannot read or write past the end of an array or use negative index. Doing so with a constant or literal index is a compile time error and an out of bounds read with a variable index will compile but fails at runtime with a *panic*.

The builtin function `len` takes in an array and returns its length:

```Go
fmt.Println(len(x))
```

Arrays are rarely used explicitly because the come with an unusual limitation. Go considers the size of the array to be part of the *type* of the array. This makes an array that's declared to be `[3]int` a different type from an array that's declared `[4]int`. This also means that you cannot use a variable to specify the size of an array because types must be resolved at compile time, not at runtime.

You also cannot use type conversion to convert arrays of different sizes to identical types. because of this you can't write a function that works with array of any size and you can't assign arrays of different sizes to the same variable.

Because of these restrictions don't use arrays unless you know the exact length you need ahead of time. 

This raises the question: why is such a limited feature in the language? Arrays exist in Go to provide a backing store for *slices*, which are on of the most useful features in Go.

# Slices
Most of the time, when you want a data structure that holds a sequence of values, a slice is what you will use. With slices the length is *not* part of the type for a slice. This removes the limitations of arrays, allowing us write a single function that processes slices of any size and we can grow slices as needed. 

Working with slices looks quite a bit like working with arrays with some subtle differences:

```Go
var x = []int{10, 20, 30} // Using [...] makes an arary. Use [] to make a slice.
var x = []int{1, 5: 4, 6, 10: 100, 15} // [1, 0, 0, 0, 0, 4, 6, 0, 0, 0, 100, 15]
```

So far slices seem identical to arrays. We start to see the differences between arrays and slices when we look at declaring slices without using a literal:

```Go
var x []int
```

This creates a slice of `int`s. Since no value is assigned, `x` is assigned the zero value for a slice which is `nil`. In Go `nil` is an identifier that represents the lack of a value for some types. Like the untyped constants `nil` has no type so it can be assigned or compared against values of different types. We'll discuss `nil` further in [[Pointers|Section 6]]

A slice is the firs type we've seen that isn't *comparable*. It is a compile-time error to use `==` to see if two slices are identical or `!=` to see if they are different. The only thing you can compare slice with is `nil`.

> The reflect package contains a function called `DeepEqual` that can compare almost anything including slices. This is primarily intended for testing, but you could use it to compare slices if you needed to. We'll look at it more in [[Here There Be Dragons: Reflect, Unsafe, and Cgo|Section 14]]

## len
Go provides several built-in functions to work with its builtin types. We already saw the builtin `len` function when looking at arrays and it works for slices too. 

>Functions like `len` are built in to Go because they can do things that can't be done by the functions that you can write. `len`'s parameter can be any type of array or any type of slice. It can also work for strings, maps, and channels too. Trying to pass a variable of any other type to `len` is a compile time error. As we'll see in [[Functions|Section 5]], Go doesn't let developers write functions that behave this way. 

## append
The builtin `append` function is used to grow slices:

```Go
var x []int
x = append(x, 10)
x = append(x, 5, 6, 7)
y := []int{20, 30, 40}
x = append(x, y...)
```

The `append` function takes at least two parameters, a slice of any type and value of that type. It returns a slice of the same type. The returned slice is assigned back to the slice that's passed in. It is a compile-time error if you forget to assign the value returned from `append`. Go is a *call by value* language. Every time you pass a parameter to a function, Go makes a copy of the value that's passed in. Passing a slice to the `append` function actually passes a copy of the slice to the function. The function adds the values to the copy of the slice and returns the copy. You then assign the returned slice back to the variable in the calling function.

## Capacity
Each element in a slice is assigned to consecutive memory locations, which it makes it quick to read or write these values. Every slice has a *capacity*, which is the number of consecutive memory locations reserved and can be larger than the length. Each time you append to a slice one or more values are added to the end of the slice, and the length increases by the number of values. When the length reaches the capacity, there's no more room to put values. If you try to add additional values when the capacity is reached, the `append` function uses the Go runtime to allocate a new slice with a larger capacity. The values in the original slice are copied to the new slice, the new values are added to the end, and the new slice is returned.

>Every high-level language relies on a set of libraries to enable programs written in that language to run, and Go is no exception. The Go runtime provides services like memory allocation and garbage collection, concurrency support, networking, and implementations of built-in types and functions.
>
>The Go runtime is compiled into every Go binary. This is different from languages that use a virtual machine, which must be installed separately to allow programs written in those languages to function. Including the runtime in the binary makes it easier to distribute Go programs and avoid worries about compatibility issues between the runtime and the program.

Just as the builtin `len` function returns the current length of a slice, the builtin `cap` function returns the current capacity of a slice.

## make
`make` lets us specify the type, length, and optionally, the capacity of a slice:

```Go
x := make([]int, 5)
```

This creates an `int` slice with a length of 5 and a capacity of 5. Since it has a length of 5, `x[0]` through `x[4]` are valid elements, and they are all initialized to 0. We can also specify the initial capacity with make:

```Go
x := make([]int, 5, 10)
```

>Never specify a capacity that's less than the length! It is a compile time error to do so with a constant or numeric literal and runtime error with a variable.

## Declaring Your Slice
The primary goal is to minimize the number of times the slice needs to grow. If it's possible that the slice won't need to grow at all (the function might return nothing), use a `var` declaration with no assigned value to create a `nil` slice.

```Go
var data []int
```

If you have some starting values, or if a slice's values aren't going to change, then a slice literal is a good choice.

```Go
data := []int{2, 4, 6, 8}
```

If you have a good idea of how large your slice needs to be but don't know the values we use `make`. Now we have to decide if we should specify a nonzero length or a zero length and nonzero capacity.
- If you are using a slice as a buffer then specify a nonzero length.
- If you are *sure* you know the exact size you want, you can specify the length and index into the slice to set values. This is often done when transforming values in one slice and storing them in a second. The downside is that if you have the wrong size you'll end up with either zero values at the end or a panic from trying to access elements that don't exist.
- In other situations use `make` with a zero length and a specified capacity. This allows you to use append to add items to the slice. If the number of items turns out to be smaller you won't have extraneous zero values at the end. If the number of items is larger, you code will not panic.

## Slicing Slices
A *slice expression* creates a slice from a slice. It's written inside brackets and consists of a starting offset and an ending offset, separated by a colon(:). Leaving off the starting offset, 0 is assumed, likewise leaving off the ending the end of the slice is substituted.

```Go
x := []int{1, 2, 3, 4}
y := x[:2]
z := x[1:]
d := x[1:3]
e := x[:]
fmt.Println("x:", x) // x: [1 2 3 4]
fmt.Println("y:", y) // y: [1 2]
fmt.Println("z:", z) // z: [2 3 4]
fmt.Println("d:", d) // d: [2 3]
fmt.Println("e:", e) // e: [1 2 3 4]
```

### Slices share storage sometimes
When you take a slice from a slice, you are *not* making a copy of the data. Instead you now have two variables that share memory, so changes to an element affects all slices that share that element.

```Go
x := []int{1, 2, 3, 4}
y := x[:2]
z := x[1:]
x[1] = 20
y[0] = 10
z[1] = 30
fmt.Println("x:", x) // x: [10 20 30 4]
fmt.Println("y:", y) // y: [10 20]
fmt.Println("z:", z) // z: [20 30 4]
```

This can get extra confusing when combined with `append`. 

```Go
x := []int{1, 2, 3, 4}
y := x[:2]
fmt.Println(cap(x), cap(y)) // 4 4
y = append(y, 30)
fmt.Println("x:", x) // x: [1 2 30 4]
fmt.Println("y:", y) // y: [1 2 30]
```

Whenever you take a slice from another slice, the sub-slice's capacity is set to the capacity of the original slice, minus the offset of the sub-slice within the original slice. This means that any unused capacity in the original slice is also shared with any sub-slices.

To avoid complicated slice situations, you should either never use `append` with a sub-slice or make sure that `append` doesn't cause an overwrite by using a *full slice expression*. This makes it clear how much memory is shared between the parent slice and sub-slice. The full slice expression includes a third part, which indicated the last position in the parent slice's capacity that's available for the sub-slice.

```Go
y := x[:2:2]
z := x[2:4:4]
```

## copy
If you need to create a slice that's independent of the original, use the built-in `copy` function. 

```Go
x := []int{1, 2, 3, 4}
y := make([]int, 4)
num := copy(y, x)
fmt.Println(y, num) // [1 2 3 4] 4
```

`copy` takes two parameters, the first is the destination slice and the second is the source slice. It copies as many values as it can from source to destination, limited by whichever slice is smaller, and returns the number of elements copied. 

You can also copy from the middle of the source slice and between two slice that cover overlapping sections of an underlying slice:

```Go
x := []int{1, 2, 3, 4}
num = copy(x[:3], x[1:])
fmt.Println(x, num) // [2 3 4 4] 3
```

# Strings, Runes, and Bytes
You might think that a string in Go is made out of runes, but that's not the case. 