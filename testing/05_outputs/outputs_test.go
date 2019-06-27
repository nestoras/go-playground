package docs

import "fmt"

func ExampleHello() {
	greeting, err := Hello("Nestoras")
	if err != nil {
		panic(err)
	}

	fmt.Println(greeting)

	// Output:
	// Hello, Nestoras
}

// go test -v -run ExamplePrint
func ExamplePrint() {
	checkIns := map[string]bool{
		"Nestoras": false,
		"George": false,
		"Maria": true,
		"Chris": true,
		"Mike": false,
	}
	Print(checkIns)

	// Unordered output:
	// Paging George; please see the front desk to check in.
	// Paging Mike; please see the front desk to check in.
	// Paging Nestoras; please see the front desk to check in.
}
