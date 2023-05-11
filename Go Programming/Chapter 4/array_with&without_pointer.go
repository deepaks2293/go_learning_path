package array_pointer

import "fmt"

func zero(ptr [32]byte) {
    for i := range ptr {
        ptr[i] = 0
    }
}

func zero_pinter(ptr *[32]byte) {
    for i := range ptr {
        ptr[i] = 0
    }
}

func zero_pointers(ptr *[32]byte) {
    *ptr = [32]byte{}
}


var arr [32]byte
arr[0] = 1
fmt.Printf("%x\n", arr) // prints "0100000000000000000000000000000000000000000000000000000000000000"

zero(&arr)

fmt.Printf("%x\n", arr) // prints "0000000000000000000000000000000000000000000000000000000000000000"

zero(&arr)

fmt.Printf("%x\n", arr) // prints "0000000000000000000000000000000000000000000000000000000000000000"

/*

In Go, when a function is called, a copy of each argument value is assigned to the corresponding parameter variable. If the argument is an array, the copy will be a new array with the same elements as the original. Therefore, any modifications made to the array inside the function will only affect the copy, not the original array.

To avoid this inefficiency and to allow modifications to be visible to the caller, we can explicitly pass a pointer to the array. In Go, pointers are used to pass values by reference. This means that if a function modifies the value pointed to by the pointer, the changes will be visible to the caller.

Note that in Go, arrays are passed by value by default, but slices are passed by reference. This is because slices are built on top of arrays and include a pointer to the underlying array, as well as a length and capacity. When a slice is passed to a function, the pointer and other information are copied, but the underlying array is not.

We then call the zero function, passing a pointer to the arr array using the & operator. The zero function creates a new [32]byte array with all elements set to zero using an array literal, and then copies the new array into the array pointed to by the function's argument. This effectively zeroes out the contents of the array.

After calling zero, we again print the contents of the arr array in hexadecimal format, which shows that all bytes are now set to 0.

*/
