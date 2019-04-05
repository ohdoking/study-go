# Study Go lang

Study basic grammer about GO langauge with [this website](https://tour.golang.org).

## Packages and Imports

Every Go program is made up of packages.
Programs start running in package `main`.

Every program can use the packages with `import` paths "fmt" and "math/rand".

#### 2 type of import

```
1. 
import (
	"fmt"
	"math"
)

2.
import "fmt"
import "math"

```

## Functions

Go function looks like different with C or java languages.

```
func [method_name]([parameters]) [types]
```
### Characteristic

#### Omit the type in consecutive named function parameters case

When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
```
//example
1. 
func add(x int, y int) int{
	return x + y
}

2.
func add(x, y int) int{
	return x + y
}
```

#### Multiple return 

A function can return any number of results.
```
//example
func test(x, y int) (int, int){
	return x, y
}
```

#### Named return values

A return statement without arguments returns the named return values. This is known as a "naked" return.

```
//example
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
```

## Variables

Go Variables also looks like different with C or java languages.
```
var [variable_name] [type]
```

### Characteristic

#### Can omit type when an initializer is present, 
If an initializer is present, the type can be omitted; the variable will take the type of the initializer.

```
//example
1.
var i, j int = 1, 2
2.
var c, python, java = true, false, "no!"
```

#### Short variable declarations in inside a function
Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

```
//example
1.
var i, j int = 1, 2
2.
k := 3
```

#### Basic types

Go's basic types are
```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

#### Zero values

Variables declared without an explicit initial value are given their zero value.

The zero value is:

0 for numeric types,
false for the boolean type, and
"" (the empty string) for strings.

#### Type conversions

The expression T(v) converts the value v to the type T.

Some numeric conversions:
```
//example
1.
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
2.
i := 42
f := float64(i)
u := uint(f)
```

#### Constants

Constants are declared like variables, but with the const keyword.
Constants can be character, string, boolean, or numeric values.

```
//example
const Truth = true
fmt.Println("Go rules?", Truth)
```

#### Numeric Constants

Numeric constants are high-precision values.
An untyped constant takes the type needed by its context.
Try printing `needInt(Big)` too.
(An int can store at maximum a 64-bit integer, and sometimes less.)

```
//example
package main

import "fmt"

const (
    // Create a huge number by shifting a 1 bit left 100 places.
    // In other words, the binary number that is 1 followed by 100 zeroes.
    Big = 1 << 100
    // Shift it right again 99 places, so we end up with 1<<1, or 2.
    Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
    return x * 0.1
}

func main() {
    fmt.Println(needInt(Small))
    fmt.Println(needFloat(Small))
    fmt.Println(needFloat(Big))
}
```

## Loop and statements 

### For loop

Go has only one looping construct, the for loop.

The basic for loop has three components separated by semicolons:
    * the init statement: executed before the first iteration
    * the condition expression: evaluated before every iteration
    * the post statement: executed at the end of every iteration

```
for i := 0; i < 10; i++ {
    sum += i
}
```

The init and post statements are optional.

```
for ; sum < 1000;  {
    sum += sum
}
```

At that point you can drop the semicolons: For is Go's "while"

```
for sum < 1000 {
    sum += sum
}
```

If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.

```
for {

}
```

### IF statements

Go's `if` statements are like its `for` loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.

```
if x < 0 {
    return sqrt(-x) + "i"
}
```

#### If with a short statement

Like for, the if statement can start with a short statement to execute before the condition.

Variables declared by the statement are only in scope until the end of the if.

```
if v := math.Pow(x, n); v < lim {
    return v
}
```


#### If and else

Variables declared inside an if short statement are also available inside any of the else blocks.

```
if v := math.Pow(x, n); v < lim {
    return v
} else {
    fmt.Printf("%g >= %g\n", v, lim)
}
```


### Switch statements

A switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression.

```
switch os := runtime.GOOS; os {
case "darwin":
    fmt.Println("OS X.")
case "linux":
    fmt.Println("Linux.")
default:
    // freebsd, openbsd,
    // plan9, windows...
    fmt.Printf("%s.\n", os)
}
```

Switch without a condition is the same as switch true.
This construct can be a clean way to write long if-then-else chains.

```
switch {
case t.Hour() < 12:
    fmt.Println("Good morning!")
case t.Hour() < 17:
    fmt.Println("Good afternoon.")
default:
    fmt.Println("Good evening.")
}
```

### defer statement

A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
```
i := 1
defer fmt.Println(i,"'s world")

i = i+1
fmt.Println(i,"'s step")

//result
2 's step
1 's world
```

Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

```
for i := 0; i < 10; i++ {
    defer fmt.Println(i)
}
```

## More types

### Pointer

Go has pointers. A pointer holds the memory address of a value.
The type *T is a pointer to a T value. Its zero value is nil.
```
var p *int
```

The & operator generates a pointer to its operand.
```
i := 42
p = &i
```

The * operator denotes the pointer's underlying value.
```
fmt.Println(*p) // read i through the pointer p
*p = 21         // set i through the pointer p
```

This is known as "dereferencing" or "indirecting".

### Structs

A struct is a collection of fields.

```
type [struct_name] struct {
    [struct_fields]
}
```

#### Pointers to structs

Struct fields can be accessed through a struct pointer.

To access the field X of a struct when we have the struct pointer p we could write (*p).X. However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.

```
type Vertex struct {
    X int
    Y int
}

func main() {
    v := Vertex{1, 2}
    p := &v
    p.X = 1e9
    fmt.Println(v)
    
    z := v
    z.X = 1
    fmt.Println(v)
    fmt.Println(z)
}
//result
{1000000000 2}
{1000000000 2}
{1 2}
```

#### Struct Literals

A struct literal denotes a newly allocated struct value by listing the values of its fields.

You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)

The special prefix & returns a pointer to the struct value.

```
type Vertex struct {
    X, Y int
}

var (
    v1 = Vertex{1, 2}  // has type Vertex
    v2 = Vertex{X: 1}  // Y:0 is implicit
    v3 = Vertex{}      // X:0 and Y:0
    p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
    fmt.Println(v1, p, v2, v3)
    //result : {1 2} &{1 2} {1 0} {0 0}
}
```

### Arrays

The type [n]T is an array of n values of type T.

The expression:
```
var a [10]int
```
declares a variable a as an array of ten integers.

#### Slices
An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

The type []T is a slice with elements of type T.

A slice is formed by specifying two indices, a low and high bound, separated by a colon:
```
a[low : high]
```

This selects a half-open range which includes the first element, but excludes the last one.

The following expression creates a slice which includes elements 1 through 3 of a:
```
a[1:4]
```

#### Slices are like references to arrays

A slice does not store any data, it just describes a section of an underlying array.
Changing the elements of a slice modifies the corresponding elements of its underlying array.
Other slices that share the same underlying array will see those changes.

```
func main() {
    names := [4]string{
        "John",
        "Paul",
        "George",
        "Ringo",
    }
    fmt.Println(names)

    a := names[0:2]
    b := names[1:3]
    fmt.Println(a, b)

    b[0] = "XXX"
    fmt.Println(a, b)
    fmt.Println(names)
}
//result
[John Paul George Ringo]
[John Paul] [Paul George]
[John XXX] [XXX George]
[John XXX George Ringo]
```

#### Slice literals

A slice literal is like an array literal without the length.

#### Slice defaults

When slicing, you may omit the high or low bounds to use their defaults instead.
The default is zero for the low bound and the length of the slice for the high bound.

For the array
```
var a [10]int
```

these slice expressions are equivalent:
```
a[0:10]
a[:10]
a[0:]
a[:]
```

#### Slice length and capacity
A slice has both a length and a capacity.

The length of a slice is the number of elements it contains.

The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

The length and capacity of a slice s can be obtained using the expressions `len(s)` and `cap(s)`.

You can extend a slice's length by re-slicing it, provided it has sufficient capacity. Try changing one of the slice operations in the example program to extend it beyond its capacity and see what happens.

```
func main() {
    s := []int{2, 3, 5, 7, 11, 13}
    printSlice(s)

    // Slice the slice to give it zero length.
    s = s[:0]
    printSlice(s)

    // Extend its length.
    s = s[:4]
    printSlice(s)

    // Drop its first two values.
    s = s[2:]
    printSlice(s)
}

func printSlice(s []int) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
//result
len=6 cap=6 [2 3 5 7 11 13]
len=0 cap=6 []
len=4 cap=6 [2 3 5 7]
len=2 cap=4 [5 7]
```

#### Nil slices

The zero value of a slice is nil.
A nil slice has a length and capacity of 0 and has no underlying array.



