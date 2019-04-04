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







