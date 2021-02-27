# console-go

A Go library for colored console output using ANSI/VT100 escape sequences


## Usage

```go
package main

import (
	"fmt"
	"github.com/cwkr/console-go"
	"github.com/cwkr/console-go/color"
)

func main() {
	console.Init()
	defer console.Reset()
	console.SetBackgroundColor(color.Navy)
	console.Clear()

	console.SetColor(color.Black)
	fmt.Println("Black")
	console.SetColor(color.Navy)
	fmt.Println("Navy")
	console.SetColor(color.Green)
	fmt.Println("Green")
	console.SetColor(color.Teal)
	fmt.Println("Teal")
	console.SetColor(color.Maroon)
	fmt.Println("Maroon")
	console.SetColor(color.Purple)
	fmt.Println("Purple")
	console.SetColor(color.Olive)
	fmt.Println("Olive")
	console.SetColor(color.Silver)
	fmt.Println("Silver")

	console.SetColor(color.Grey)
	fmt.Println("Grey")
	console.SetColor(color.Blue)
	fmt.Println("Blue")
	console.SetColor(color.Lime)
	fmt.Println("Lime")
	console.SetColor(color.Aqua)
	fmt.Println("Aqua")
	console.SetColor(color.Red)
	fmt.Println("Red")
	console.SetColor(color.Pink)
	fmt.Println("Pink")
	console.SetColor(color.Yellow)
	fmt.Println("Yellow")
	console.SetColor(color.White)
	fmt.Println("White")

	console.SetBackgroundColor(color.Black)
	fmt.Println("Black")
	console.SetBackgroundColor(color.Navy)
	fmt.Println("Navy")
	console.SetBackgroundColor(color.Green)
	fmt.Println("Green")
	console.SetBackgroundColor(color.Teal)
	fmt.Println("Teal")
	console.SetBackgroundColor(color.Maroon)
	fmt.Println("Maroon")
	console.SetBackgroundColor(color.Purple)
	fmt.Println("Purple")
	console.SetBackgroundColor(color.Olive)
	fmt.Println("Olive")
	console.SetBackgroundColor(color.Silver)
	fmt.Println("Silver")

	console.SetColor(color.Black)
	console.SetBackgroundColor(color.Grey)
	fmt.Println("Grey")
	console.SetBackgroundColor(color.Blue)
	fmt.Println("Blue")
	console.SetBackgroundColor(color.Lime)
	fmt.Println("Lime")
	console.SetBackgroundColor(color.Aqua)
	fmt.Println("Aqua")
	console.SetBackgroundColor(color.Red)
	fmt.Println("Red")
	console.SetBackgroundColor(color.Pink)
	fmt.Println("Pink")
	console.SetBackgroundColor(color.Yellow)
	fmt.Println("Yellow")
	console.SetBackgroundColor(color.White)
	fmt.Println("White")

	console.ReadLine()
	console.SetColor(color.Silver)
	console.SetBackgroundColor(color.Navy)
	console.Clear()
	console.SetCursorPosition(5, 5)
	fmt.Println("Hello world")
	console.ReadLine()
}
```


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
