package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices")
	var array1 [5]int
	array1[0]=1
	array1[3]=3
	fmt.Println(array1)

	array2 :=[5]string{"Posição 1", "Posição 2", "Posição 3"}
	array2[3]= "Posição 4"
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(array3)

	slice1 :=[]int{1,2}
	fmt.Println(slice1)
	slice1 = append(slice1, 20)
	fmt.Println(slice1)

	slice2 := array2[1:3]
	array2[1]="Posição alterada"
	fmt.Println(slice2)

	slice3 := make([]float32,10,15)
	slice3 = append(slice3, 1.1)
	fmt.Println(slice3, len(slice3), cap(slice3))
}
