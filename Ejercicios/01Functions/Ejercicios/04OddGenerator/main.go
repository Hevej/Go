package main

import "fmt"

func makeOddGenerator() func() uint{
	i := uint(1)
	return func() (ret uint){
		ret = i
		i += 2
		return
	}
}

func main() {
	nextOdd := makeOddGenerator()
	for i:=0; i<=10; i++ {
		fmt.Println(nextOdd())
	}
}
