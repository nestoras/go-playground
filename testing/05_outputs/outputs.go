package docs

import "fmt"

//godoc -http=:3000

type Doc struct{}

func (d Doc) Hello(){}


func Hello(name string) (string ,error){
	return fmt.Sprintf("Hello, %s", name), nil
}


func Print(checkIns map[string]bool) {
	for name, checkIn := range checkIns {
		if !checkIn {
			fmt.Printf("Paging %s; please see the front desk to check in.\n", name)
		}
	}
}