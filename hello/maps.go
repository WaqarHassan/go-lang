// package main

// import "fmt"

// import "sort"

// func myFunc()  {
// 	i := 0
// 	Here:
// 	i++
// 	fmt.Println(i)
// 	goto Here
// }

// func Solution(N int) []int {
//     // write your code in Go 1.4

//     a := make([]int, N)
//     ind := 0
//     half := N / 2
//     ai := 0
//     if (N %2 != 0) {
//         a[0] = 0
//         ind = 1
//         half = N / 2 + 1
//         ai = 1
//     }

//     for i := ind  ; i < half && ai < N  ; i++ {
//         a[ai] = (i + 1) * -1
//         a[ai + 1] = i + 1
//         ai = ai + 2
//     }

//     return a
// }
// func Solution3(N int, K int) int {
//     // write your code in Go 1.4

//     a := 0
//     counter : = 0
//     x2arr := make([]int, N)
//     if (K == 0) {
//         return N - 1
//     } else {
//         divided := K
//         i := 0
//         for (divided > 1 && counter <= K) {
//             divided = N / 2
//             arr[i] = divided
//             i++
//             counter ++
//         }
//     }
//     return counter
// }
// func main() {
// 	// var m1 map[int]int

// 	// var m2 map[int]int

// 	// m3 := make(map[string]int)

// 	// m4 := make(map[string]int , 30)

// 	// m5 := make(map[string]int, 0)

// 	// var m6 map[string]int = map[string]int{"foo": 20, "Bar": 233}

// 	// fmt.Println(m6["foo"])

// 	// var mm = map[string]string{"a": "b", "c": "d"}

// 	// fmt.Println(mm["c"])

// 	// myFunc()

// 	// i := 1
// 	// for i < 10 {
// 	// 	fmt.Println(i)
// 	// 	i++
// 	// }

// 	// for i:=0; i < 10 ; i++ {
// 	// 	if i == 3 || i == 9 || i == 8 {
// 	// 		continue
// 	// 	}
// 	// 	fmt.Println(i)
// 	// }

// 	// list := []string{"a", "b", "c", "d", "e", "f"}
// 	// for i := 0 ; i <  len(list) ; i++ {
// 		// fmt.Println(list[i])
// 	// }

// 	// Ranges
// 	// for k, v := range "list" {
//  //    fmt.Printf("character '%c' starts at byte position %d\n", v, k)
// 	// }

// 	// c := 'A'
// 	// switch {
// 	// 	case '0' <= c && c <= '9':
// 	// 	    fmt.Println( c - '0')
// 	// 	case 'a' <= c && c <= 'f':
// 	// 	    fmt.Println( c - 'a' + 10)
// 	// 	case 'A' <= c && c <= 'F':
// 	// 	    fmt.Println( c - 'A' + 10)
// 	// }
// 	// fmt.Println()
// 	// fmt.Println()
// 	// fmt.Println()
// 	// fmt.Println()
// 		// return 0

// 		// var a := []int {4,2,55,4}
// 		// sort.Ints(a)
// 		fmt.Println(2 > 3 ? 2 : 3)

// }