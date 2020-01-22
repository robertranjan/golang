package mix

import "fmt"

func printArray(strArray [20]string) {
	fmt.Printf("\n\t(printArray) Mem loc of strArray: %x", strArray)
	for _, value := range strArray {
		fmt.Print(value)
	}
}

func updateArray(strArray [20]string) {
	for i := 0; i < 20; i++ {
		strArray[i] = string(65 + i)
	}
	fmt.Printf("\n\t(updateArray) Mem loc of strArray: %x", strArray)
}

func updateInt(anInt *int) {
	*anInt = *anInt * 2
	fmt.Printf("\n\tvalue of anInt: %v | address of anInt: %x", *anInt, anInt)
}

func printInt(anInt *int) {
	fmt.Printf("\n\tvalue of anInt: %v | address of anInt: %x", *anInt, anInt)
}

// ArrayExample ...
func ArrayExample() int {
	var x [20]string
	fmt.Printf("\n\tMem loc of x: %x", x)
	printArray(x)
	updateArray(x)
	printArray(x)
	var intVar int
	intVar = 20
	printInt(&intVar)
	updateInt(&intVar)
	printInt(&intVar)
	fmt.Printf("\n\n")
	return 0
}
