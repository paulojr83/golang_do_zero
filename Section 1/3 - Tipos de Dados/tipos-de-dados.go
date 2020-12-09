package main

import (
	"errors"
	"fmt"
)

func main() {

	var numeroInt8 int8 = 127 //(-128 to 127)
	var numeroInt16 int16 = -32768 //(-32768 to 32767)
	var numeroInt32 int32 = 2147483647 //(-2147483648 to 2147483647)
	var numeroInt64 int64 = -9223372036854775808 //(-9223372036854775808 to 9223372036854775807)
	fmt.Println(numeroInt8,"\n", numeroInt16,"\n", numeroInt32,"\n", numeroInt64)

	var numeroUint8    uint8 = 255// (0 to 255)
	var numeroUint16  uint16 = 65535// (0 to 65535)
	var numeroUint32  uint32 = 4294967295// (0 to 4294967295)
	var numeroUint64  uintptr = 18446744073709551615 // (0 to 18446744073709551615)
	fmt.Println(numeroUint8, "\n",numeroUint16, "\n",numeroUint32, "\n",numeroUint64)
	/*
		byte        alias for uint8
		rune        alias for int32
	*/
	var numeroFloat32 float32 = 323.33    //the set of all IEEE-754 32-bit floating-point numbers
	var numeroFloat64 float64 =33213123.33  //  the set of all IEEE-754 64-bit floating-point numbers
	fmt.Println(numeroFloat32, "\n", numeroFloat64)

	var srt string = "teste"
	fmt.Println(srt)

	var c = 'D'
	fmt.Println(c)

	var booleano bool
	fmt.Println(booleano)
	boo := true
	fmt.Println(boo)

	var err error = errors.New("algo deu errado!")
	fmt.Println(err)
}
