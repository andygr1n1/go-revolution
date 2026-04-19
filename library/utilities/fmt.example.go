package utilities

import (
	"fmt"
	"time"

	"github.com/andygr1n1/go-revolution/internal/globalctx"
)

func printObject(data any) {
	fmt.Printf("%+v\n", data)
}

func FmtExample(ctx globalctx.GlobalContext) {
	fmt.Println("Welcome to Go Revolution")

	currentTime := time.Now().Format("2 Jan 2006 15:04:05")

	fmt.Println("The time is:", currentTime)

	/* Object */
	type Person struct {
		Name string
		Age  int
	}

	person := Person{Name: "Andrew", Age: 35}

	fmt.Printf("Hello, %v!\n", person)
	fmt.Printf("Hello, %#v!\n", person)

	/* most useful fo objects: */
	printObject(person)
	ctx.Logger.Info("person", "value", person)
}
