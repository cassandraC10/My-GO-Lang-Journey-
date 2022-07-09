package main

//import "fmt"
//
//func main() {
//fmt.Println("Hello, World!")
//fmt.Print("my name is sandra")
//fmt.Println()
//
//dog := "max"
//fmt.Println(dog)
//fmt.Println("my dog's name is ", dog)
//
//age := 21 	short statement declaration
//fmt.Println("i am ", age, " years old")
//fmt.Println("1 + 1 =", 1+1)
//
//fmt.Println(len("Hello World"))
//fmt.Println("Hello World"[1])
//fmt.Println("Hello " + "World")
//fmt.Println(321325 * 424521)
//
//var x string = "How To Create/Store A Variable" long statement declaration
//fmt.Println(x)
//
//var y string
//y = "first "
//fmt.Println(x)
//y = y + "second"
//fmt.Println(y)

//var x string = "hello"
//var y string = "hello"
//fmt.Println(x == y)

//const x string = "Hello World"
//fmt.Println(x)
//
//var (
//	a = 5
//	b = 10
//	c = 15
//)
//fmt.Println(a, b, c)

//fmt.Print("Enter a number: ")
//var input float64
//fmt.Scanf("%f", &input)
//output := input * 2
//fmt.Println(output)

//fmt.Print("try a number again: ")
//var try float64
//fmt.Scanf("%f", &try)
//fmt.Println(try * 2)

//fmt.Print("calculate a number: ")
//var dorcas float64
//fmt.Scan(&dorcas)
//fmt.Print(dorcas)

// Exercise 4.5
//fmt.Print("Enter temperature in fahrenheit: ")
//var f float64
//fmt.Scan(&f)
//fmt.Println("Temperature in fahrenheit is ", f) // i used 75
//
//fmt.Print("Enter temperature in celsius: ") // i used 36, try urs
//var c float64
//fmt.Scan(&c)
//fmt.Println("Temperature in celsius is ", c)
//fmt.Println()
////fmt.Println(f, " degrees fahrenheit to", c, "celsius is: ", 5.0/9.0*(f-32))
//C := (f - 32) * 5 / 9
//fmt.Print(f, " degree fahrenheit to", c, "celsius is: ", C)

//Exercise 4.6
//fmt.Print("Enter a value for feet: ") // i used 5.8
//var feet float64
//fmt.Scan(&feet)
//meters := feet * 0.3048
//fmt.Println(feet, "ft is ", meters, "m.")

//i := 1
//for i <= 10 {
//	fmt.Println(i)
//	i = i + 1

//for i := 1; i <= 10; i++ {
//	fmt.Print(i)
//if i%2 == 0 {
//	fmt.Println(i, "even")
//} else {
//	fmt.Println(i, "odd")
//}
//	if i == 0 {
//		fmt.Println("Zero")
//	} else if i == 1 {
//		fmt.Println("One")
//	} else if i == 2 {
//		fmt.Println("Two")
//	} else if i == 3 {
//		fmt.Println("Three")
//	} else if i == 4 {
//		fmt.Println("Four")
//	} else if i == 5 {
//		fmt.Println("Five")
//	}
//	switch i {
//	case 0:
//		fmt.Println(" Zero")
//	case 1:
//		fmt.Println(" One")
//	case 2:
//		fmt.Println(" Two")
//	case 3:
//		fmt.Println(" Three")
//	case 4:
//		fmt.Println(" four")
//	case 5:
//		fmt.Println(" five")
//	case 6:
//		fmt.Println(" six")
//	case 7:
//		fmt.Println(" seven")
//	case 8:
//		fmt.Println(" eight")
//	case 9:
//		fmt.Println(" nine")
//	//case 10:
//	//	fmt.Println(" Ten")
//	default:
//		fmt.Println(" Unknown number")
//	}
//}

// exercise 5.1
//i := 10
//if i > 10 {
//	fmt.Println("Big")
//} else {
//	fmt.Println("Small")
//}

//Exercise 5.2
//for i := 1; i <= 100; i++ {
//	if i%3 == 0 {
//		fmt.Println(i)
//	}
//}

//	Exercise 5.3
//for i := 1; i <= 100; i++ {
//	if i%15 == 0 {
//		fmt.Println("FizzBuzz")
//	} else if (i % 5) == 0 {
//		fmt.Println("Buzz")
//	} else if (i % 3) == 0 {
//		fmt.Println("Fizz")
//	} else {
//		fmt.Println(i)
//	}
//	//fmt.Print(i)
//}

//	CHAPTER 6
//var x [5]float64
//x[0] = 98
//x[1] = 93
//x[2] = 77
//x[3] = 82
//x[4] = 83 // sum all of them,

//var total float64 = 0
//for i := 0; i < 5; i++ {
//	total += x[i] // total is 0, so total is the sum = 433
//}
//fmt.Println(total / 5) // 433/ 5 = 86.6

//var total float64 = 0
//for i := 0; i < len(x); i++ {
//	total += x[i]
//}
//fmt.Println(total / float64(len(x)))
// converting len int to float, because total is a float

// special for loop
//x := [5]float64{98, 93, 77, 82, 83}
//var total float64 = 0
//for _, value := range x {
//	total += value
//}
//fmt.Println(total / float64(len(x)))

//slice1 := []int{1, 2, 3}
//slice2 := append(slice1, 4, 5)
//fmt.Println(slice1, slice2)  // on append

//slice1 := []int{1, 2, 3}
//slice2 := make([]int, 2)
//copy(slice2, slice1)
//fmt.Println(slice1, slice2)  // on copy

//elements := make(map[string]string)  // first kind of map
//elements["H"] = "Hydrogen"
//elements["He"] = "Helium"
//elements["Li"] = "Lithium"
//elements["Be"] = "Beryllium"
//elements["B"] = "Boron"
//elements["C"] = "Carbon"
//elements["N"] = "Nitrogen"
//elements["O"] = "Oxygen"
//elements["F"] = "Fluorine"
//elements["Ne"] = "Neon"

//elements := map[string]string{
//	"H":  "Hydrogen",
//	"He": "Helium",
//	"Li": "Lithium",
//	"Be": "Beryllium",
//	"B":  "Boron",
//	"C":  "Carbon",
//	"N":  "Nitrogen",
//	"O":  "Oxygen",
//	"F":  "Fluorine",
//	"Ne": "Neon",
//}
//fmt.Println(elements["Li"])

//elements := map[string]map[string]string{
//	"H": map[string]string{
//		"name":  "Hydrogen",
//		"state": "gas",
//	},
//	"He": map[string]string{
//		"name":  "Helium",
//		"state": "gas",
//	},
//	"Li": map[string]string{
//		"name":  "Lithium",
//		"state": "solid",
//	},
//	"Be": map[string]string{
//		"name":  "Beryllium",
//		"state": "solid",
//	},
//	"B": map[string]string{
//		"name":  "Boron",
//		"state": "solid",
//	},
//	"C": map[string]string{
//		"name":  "Carbon",
//		"state": "solid",
//	},
//	"N": map[string]string{
//		"name":  "Nitrogen",
//		"state": "gas",
//	},
//	"O": map[string]string{
//		"name":  "Oxygen",
//		"state": "gas",
//	},
//	"F": map[string]string{
//		"name":  "Fluorine",
//		"state": "gas",
//	},
//	"Ne": map[string]string{
//		"name":  "Neon",
//		"state": "gas",
//	},
//}
//if el, ok := elements["Li"]; ok {
//	fmt.Println(el["name"], el["state"])
//}

// Exercise 6.2
//slice := make([]int, 3, 9)
//fmt.Println(len(slice))

// Exercise 6.3
//x := [6]string{"a", "b", "c", "d", "e", "f"}
//fmt.Print(x[2:5])	// [c d e]

// Exercise 6.4
//x := []int{
//	48, 96, 86, 68,
//	57, 82, 63, 70,
//	37, 34, 83, 27,
//	19, 97, 9, 17,
//}
//smallest := x[0]
//for _, num := range x[1:] {
//	if num < smallest {
//		smallest = num
//	}
//}
//fmt.Println(smallest)

//	xs := []float64{98, 93, 77, 82, 83}
//	fmt.Println(average(xs))
//	x := 0
//	increment := func() int {
//		x++
//		return x
//	}
//	fmt.Println(increment())
//	fmt.Println(increment())
//}
//
//func average(xs []float64) float64 {
//	total := 0.0
//	for _, v := range xs {
//		total += v
//	}
//	return total / float64(len(xs))
//}
