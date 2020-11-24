package view

import "fmt"

func ShowMessage(text string, params ...interface{}) {
	fmt.Printf("INFO: "+text+"\r\n", params...)
}
