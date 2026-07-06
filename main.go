package main

import "fmt"

type koda struct{
	are a
}

type a struct{
	the b
}

type b struct{
	best string
}

type say struct{
	world string
}

func main() {
	we := koda{
		are: a{
			the: b{
				best: "Koda",
			},
		},
	}
	fmt.Println(we.are.the.best)

	hello := say{
		world: "Hello World",
	}
	fmt.Println(hello.world)
}