//go:build wireinject
// +build wireinject

package app

import "github.com/google/wire"

// func InitA() (*A, error) {
// 	wire.Build(wire.Struct(new(A), "*"), wire.Struct(new(B), "*"))
// 	// wire.Build(NewA, NewB) // 函数注入
// 	return &A{}, nil
// }
// var SuperSet = wire.NewSet(NewA, NewB)

// func InitA() (*A, error) {
// 	wire.Build(SuperSet)
// 	return &A{}, nil
// }

var SuperSet = wire.NewSet(
	wire.Struct(new(B), "*"),         // 填充 B 中所有字段
	wire.Struct(new(A), "*"),         // 填充 A 中所有字段
	wire.Bind(new(IServer), new(*B)), // 使用 *B 填充 IServer 字段
)

func InitA() (A, error) {
	wire.Build(SuperSet) // 结构注入
	return A{}, nil
}
