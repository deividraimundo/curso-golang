# Go Lang

## CAPÍTULO 1 - INTRODUÇÃO
- [x] História/Características
- [ ] Retornar em terminal do linux

## CAPÍTULO 2 - CONFIGURAÇÕES DO AMBIENTE
- [ ] Configurações do ambiente no Linux

## CAPÍTULO 3 - FUNDAMENTOS
- [x] Estrutura básica
- [x] Ferramentas
- [x] Const/Var
1. Sempre que declarar uma variável é necessário usá-la
2. Declarações de variáveis são feitas de duas formas:
	1.  const/var nomeDaVariavel tipo = valor
	2. nomeDaVariavel := valor (forma reduzida de se criar uma variável em Go)
- [x] Tipos básicos
- [x] Funções
- [x] Operadores/Ponteiros...

## CAPÍTULO 4 - ESTRUTURAS DE CONTROLE
- [x] if/else
- [x] for
- [x] switch

## CAPÍTULO 5 - ARRAY/SLICE/MAP
- [x] Array
	Tamanho pré-definido
- [x] Slice
	Têm tamanhos dinâmicos. São partes de um array.
- [x] Map
	Lista com chave e valor

## CAPÍTULO 6 - FUNÇÕES
- [x] Múltiplos retornos
- [x] Retorno nomeado
- [x] Funções variáticas
	Número de parâmetros variáveis
- [x] Closure
	Declaração de variáveis com nomes iguais em diversos escopos
- [x] Recursividade
	Chamando uma função dentro da própria função
- [x] Defer/init/ponteiro...
	Defer: atrasa a execução
	
	Ponteiro: na declaração o ponteiro é criado com um *, após isso, o * na variável será usado para desreferenciar o ponteiro, ou seja, para pegar seu valor.
	Exemplo:    
```
    package main 
     
    import "fmt" 
     
    func main() { 
    	n := 1 
     
    	inc1(n) // passagem por valor 
    	fmt.Println(n) 
     
    	// revisão: & usado para obter o endereço da variavel 
    	inc2(&n) 
    	fmt.Println(n) 
    } 
     
    func inc1(n int) { 
    	n++ 
    } 
     
    // revisão: um ponteiro é representado por um * 
     
    func inc2(n *int) { 
    	// revisar: * é usado para desreferenciar, ou seja, ter acesso ao valor no qual o ponteiro aponta 
    	*n++ 
    }
```

Init: é a prioridade de execução (executada automaticamente primeiro)
Exemplo:
```
    package main

    import "fmt"

    func main() {
        fmt.Println("Main...")
    }

    func init() {
        fmt.Println("Inicializando...")
    }
```

## CAPÍTULO 7 - SISTEMA DE TIPOS
- [x] Struct
	É uma estrutura que guarda valores
	Struct aninhada, exemplo:
```
    package main

    import "fmt"

    type item struct {
        produtoID  int
        quantidade int
        preco      float64
    }

    type pedido struct {
        userID int
        itens  []item
    }

    func main() {
        pedido := pedido{
            userID: 1,
            itens: []item{
                {
                    produtoID:  1,
                    quantidade: 2,
                    preco:      12.10,
                },
                {
                    produtoID:  2,
                    quantidade: 1,
                    preco:      23.49,
                },
                {
                    produtoID:  11,
                    quantidade: 100,
                    preco:      3.13,
                },
            },
        }

        fmt.Printf("Valor total do pedido é R$ %.2f\n", pedido.valorTotal())
    }

    func (p pedido) valorTotal() float64 {
        total := 0.0
        for _, item := range p.itens {
            total += item.preco * float64(item.quantidade)
        }

        return total
    }
```
	
	Utilizando ponteiros no Struct
```
    package main

    import (
        "fmt"
        "strings"
    )

    type pessoa struct {
        nome      string
        sobrenome string
    }

    func main() {
        p1 := pessoa{
            nome:      "Pedro",
            sobrenome: "Silva",
        }
        fmt.Println(p1.getNomeCompleto()) // "Pedro Silva"

        p1.setNomeCompleto("Maria Pereira")
        fmt.Println(p1.getNomeCompleto()) // "Pedro Silva" "Maria Pereira"
    }
    func (p pessoa) getNomeCompleto() string {
        return p.nome + " " + p.sobrenome
    }

    func (p *pessoa) setNomeCompleto(nomeCompleto string) {
        partes := strings.Split(nomeCompleto, " ")
        p.nome = partes[0]
        p.sobrenome = partes[1]
    }
```
	
	Pseudo-Herança em Struct: (Struct não tem herança)
```
    package main

    import "fmt"

    type carro struct {
        nome            string
        velocidadeAtual int
    }

    type ferrari struct {
        carro       // campos anônimos (tudo que é de carro fica disponivel)
        turboLigado bool
    }

    func main() {
        f := ferrari{}
        f.nome = "F40"
        f.velocidadeAtual = 0
        f.turboLigado = true

        fmt.Printf("A ferrari %s está com turbo ligado? %v", f.nome, f.turboLigado)
    }
```
	
- [x] Interface
- [x] Métodos
- [x] Polimorfismo
- [x] Composição

## CAPÍTULO 8 - PACOTES
- [x] Pacotes
- [x] Pacotes no Github

## CAPÍTULO 9 - CONCORRÊNCIA
- [x] Paral. vs Concor.
- [x] Goroutine/Channel
	Goroutine: é uma função que executa simultaneamente com outras goroutine. Causa um assincronismo.
	Channel: é a forma de comunicação entre goroutines. Faz sincronismo baseado nos dados que esta esperando receber.
- [x] Padrões de concorrência

## CAPÍTULO 10 - TESTES
- [x] Testes
- [x] Cobertura

## CAPÍTULO 11 - BANCO DE DADOS
- [ ] Criar Schema/Tabela
- [ ] Insert/Transação
- [ ] Update
- [ ] Select

## CAPÍTULO 12 - HTTP
- [ ] Servidor estático
- [ ] Servidor dinâmico
- [ ] Acessando BD

## CAPÍTULO 13 - CONCLUSÃO
- [ ] Conclusão
