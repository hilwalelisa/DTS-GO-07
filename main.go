package main

// Buatlah sebuah program go dengan :
//menampilkan nilai i : 21 fmt.Printf("%T \n", i) // contoh : fmt.Printf("%v \n", i)
//menampilkan tipe data dari variabel i
//menampilkan tanda %
//menampilkan nilai boolean j : true
//menampilkan nilai boolean j : true
//menampilkan unicode russia : Я (ya)
//menampilkan nilai base 10 : 21
//menampilkan nilai base 8 :25
//menampilkan nilai base 16 : f
//menampilkan nilai base 16 : F 13
//menampilkan unicode karakter Я : U+042F var k float64 = 123.456;
//menampilkan float : 123.456000
//menampilkan float scientific : 1.234560E+02

import "fmt"

func main() {
	i := 21
    fmt.Printf("%d\n", i)
    fmt.Printf("%T\n", i)
    fmt.Printf("%%\n")
    j := true
    fmt.Printf("%t\n\n", j)
    fmt.Printf("%b \n", i)
    fmt.Printf("%c \n", '\u042F')
    fmt.Printf("%d \n", i)       
	fmt.Printf("%o \n", i)      
	fmt.Printf("%x \n", 15)       
	fmt.Printf("%X \n", 15)      
	fmt.Printf("%U \n\n", 'Я')   
	var k float64 = 123.456
	fmt.Printf("%f \n", k)
	fmt.Printf("%E \n", k)
}
