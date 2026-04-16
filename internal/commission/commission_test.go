package commission

import "fmt"

func Example_mask() {
	fmt.Println(Mask("****"))
	fmt.Println(Mask("8888888888888888"))
	fmt.Println(Mask("abdsderfvderfswe"))
	//Output:
	//
	//8888 **** **** 8888
	//abds **** **** fswe
}
