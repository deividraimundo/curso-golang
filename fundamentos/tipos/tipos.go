package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	/* TABELA
	uint8       unsigned  8-bit integers (0 to 255)
	uint16      unsigned 16-bit integers (0 to 65535)
	uint32      unsigned 32-bit integers (0 to 4294967295)
	uint64      unsigned 64-bit integers (0 to 18446744073709551615)
	int8        signed  8-bit integers (-128 to 127)
	int16       signed 16-bit integers (-32768 to 32767)
	int32       signed 32-bit integers (-2147483648 to 2147483647)
	int64       signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
	float32     IEEE-754 32-bit floating-point numbers
	float64     IEEE-754 64-bit floating-point numbers
	*/

	//números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// sem sinal (só positivos)... uint8 uint16 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	// com sinal... int8 int16 int32 int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo de i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// números reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é", reflect.TypeOf(49.99)) // float64

	// boolean
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// string
	s1 := "Olá, meu nome é Deivid"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da string é", len(s1))

	// string com múltiplas linhas
	s2 := `Olá
	meu
	nome
	é
	Deivid`
	fmt.Println("O tamanho da string é", len(s2))

	// char??? não existe
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char)
}
