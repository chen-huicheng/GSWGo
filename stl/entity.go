package stl

type User struct {
	Name   string `init:"john"`
	Age    int    `ini1t:"18"`
	Gender uint8  `init:"1"`
}

type Student struct {
	User  `init:"magicKey"`
	Grade string `init:"初一"`
	StuNo string `init:"20220801"`
}
