# Study Go lang

Study basic grammer about GO langauge with [this website](https://tour.golang.org).

## Packages and Imports

Every Go program is made up of packages.
Programs start running in package `main`.

Every program can use the packages with `import` paths "fmt" and "math/rand".

2 type of import

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

Go function is looks like different with C or java languages.
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







