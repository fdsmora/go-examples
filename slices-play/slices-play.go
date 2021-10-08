package main

import "fmt"

/* Experimenting with slices
   Also contains:
   - Type assertions
   - Simulating variable and optional function arguments
*/

func main() {

	temp := [7]int{1,2,3,4} 
	print("temp := [7]int{1,2,3,4}")
	print("temp ->", temp)
	print("cap(temp) ->", cap(temp))
	m := temp[:2]
	print("m := temp[:2]; m ->", m)
	print("cap(m) ->", cap(m))
	n := temp[2:5]
	print("n := temp[2:5]; n ->", n)
	print("cap(n) ->", cap(n))
	fmt.Printf("%q\n", "cap(slice)=n-i, where slice=array[i:m] for array:=[m]T {...} and m<=n and some type T")
	br()

	primes := [6]int{2, 3, 5, 7, 11, 13}
	print("primes := [6]int{2, 3, 5, 7, 11, 13}")

	var s []int = primes[1:4]
	print("s=primes[1:4] ->",s)
	print("cap(s) ->", cap(primes))
	br()

	primes[1] = 9
	print("primes[1]=9")
	print("s ->",s)
	s[0] = -2
	print("s[0] = -2")
	print("s ->", s)
	print("primes ->", primes)
	// s[4] = 99 //fails
	print("s[4] = 99 -> fails, even if underlying primes has this index, the slice `s` doesn't")
	br()

	print("cap(s) ->", cap(s))
	r := append(s, 45, 64, 32)
	print("r := append(s, 45, 64, 32)")
	print("cap(s) ->", cap(s))
	print("r ->", r)
	print("s ->", s)
	print("Explanation\n`s` doesn't have enough capacity to house the 3 new elements, so `s` is copied to a new array with enought capacity for the new 3 elements and returned and assigned to `r`. \nIf `r` had enough capacity, the new elems would be appended to `s` and no new copy created, so the `r` slice would hold a reference to the underlying array of `s`")
	br()
	
	print("primes ->",primes)
	
	r[1] = 0
	print("r[1] = 0")
	print("r ->", r)
	print("s ->",s)
	print("primes ->", primes)
	print("cap(s) ->", cap(s))
	print("Explanation\n`r[1]=0` doesn't affect `s`'s underlying array (`primes`) as recall that `r` is a fresh copy of `s`'s underlying array as per the `append` operation above")
	br()
	
	t := append(s, 10)
	print("t after t:=append(s, 10) ->",t)
	print("s ->", s)
	t[0] = -1
	print("t[0]=-1 ->",t)
	print("s ->",s)
	s[2] = 70
	print("t after s[2] = 70 ->",t)
	print("s ->", s)
	
	
	a := []int{1,2,3}
	b := append(a, 5)
	print("a := []int{1,2,3}\nb := append(a, 5)\nb ->",b)
	
	a[0]=0
	print("a[0]=0 ->",a)
	print("b := append(a, 5) ->", b)
	/*
	x := make([]int, 3, 5)
	x[0] = 1
	x[1] = 2
	x[2] = 3
	y := []int{4,5}
	
	z := append(x, y...)
	print(z)
	
	x[0] = 9
	print(x)
	print(z) */
}

func br() {
	fmt.Println()
}

func print(args ...interface{}){
    var caption string
	if len(args) == 0 {
		return 
	}else if len(args) == 1 {
		fmt.Println(args[0])
		return
	}else if len(args) > 1 {
	    if val, ok := args[0].(string); !ok {
			fmt.Println("error: 1st arg must be an string")
			return
		}else{
			caption = val
		}
	}
	fmtStr := "%s %v\n"
	fmt.Printf(fmtStr, caption, args[1])
	/*
	switch data := args[1].(type) {
	case []int:
		fmt.Printf(fmtStr, caption, data)
	case [6]int:
		fmt.Printf(fmtStr, caption, data)
	case string:
		fmt.Printf(fmtStr, caption, data)
	case int:
	    fmt.Printf(fmtStr, caption, data)
	default:
		fmt.Println("error: can't convert")
	}*/
}


