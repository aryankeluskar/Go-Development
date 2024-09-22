package main

import "fmt"
import "unicode/utf8"
import "reflect"

func main() {
	fmt.Println("Hello World!")

	// <var> name type	
	var myVar int = 10
	fmt.Println(myVar)
	
	// int has int16, int32, int64
	// float has float32, float64
	var alsoVar float32 = 1378.2378692
	var res float32 = alsoVar * float32(myVar)
	fmt.Println(res)

	// strings
	var name string = "specter"
	fmt.Println(utf8.RuneCountInString(name)) // will print unicode tho, not actually character

	// char / rune
	var char rune = 'a'
	fmt.Println(char)
	fmt.Println(reflect.TypeOf(char))

	// bool and short-hand
	gh := false
	fmt.Println(gh)

	if gh{
		var out, err = print10times(name)
		if err==nil{
			fmt.Println(out)
		}
	}


}

func print10times(sentence string) (int, error){
	var err error 

	for i := 0; i < 10; i++ {
		fmt.Println(sentence)
	}
	return 10, err
}