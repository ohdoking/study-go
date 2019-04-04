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