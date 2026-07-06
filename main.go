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

// Kedua
type say struct{
	world string
}

// Ketiga
type a1 struct{
	str [][][]b2
}
type tech struct{
	tech c2
}

type b2 struct{
	man []tech
}
type c2 struct{
	academy string
}

// Empat
type d1 struct {
	is string
}

type b3 struct {
	fruit d1
}

type a2 struct {
	favourite []b3
}

// lima
type bil struct{
	first []int
	second []int
}

func main() {
	// Pertama we.are.the.best
	we := koda{
		are: a{
			the: b{
				best: "Koda",
			},
		},
	}
	fmt.Println(we.are.the.best)

	// Kedua hello.world
	hello := say{
		world: "Hello World",
	}
	fmt.Println(hello.world)

	// Ketiga obj.str[3][1][2].man[0].tech.academy

	obj := a1{
		str : [][][]b2{
			3: {
				1:{
					2: b2{
							man: []tech{
								{
									tech: c2{academy: "Tech Academy"},
								},
							},
					},
				},
			},
		},
	}
	fmt.Println(obj.str[3][1][2].man[0].tech.academy)

	my := []a2{
		{
			favourite: []b3{
				3:{
					fruit: d1{
						is: "Apple",
					},
				},
			},
		},
	}
	fmt.Println(my[0].favourite[3].fruit.is)

	// Lima
	num := bil{
		first: []int{
			1: 16,
		},
		second: []int{
			2: 16,
		},
	}

	fmt.Println(num.first[1] + num.second[2])
}