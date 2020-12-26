package main

import ("fmt"
			 "encoding/json")

func main() {
	
	Data := make(map[string]string)
	var Name,Address string
	fmt.Println("Enter a name")
	fmt.Scan(&Name)
	fmt.Println("Enter a address")
	fmt.Scan(&Address)
	Data["name"] = Name
	Data["address"] = Address
	//fmt.Println(Data)

	b, _ := json.Marshal(Data)

	fmt.Println(string(b))

}