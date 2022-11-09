package main

import (
	"fmt"
)

type hotdog int 
var x hotdog

func main() {
	atividadeUm();
	atividadeDois();
	atividadeTres();
	atividadeQuatro();
	atividadeCinco();
}

/* Atividade 1
* Utilizando um operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y" e "z".
* Agora demostre os valores nestas variáveis, com
* 1 - Uma única declaração print
* 2 - Múltiplas declarações print
 */
func atividadeUm(){
	fmt.Println("Atividade 1:")
	x, y, z := 42, "James Bond", true
	fmt.Printf("%v, %v, %v \n",x ,y ,z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println("Finalizado atividade\n\n")
}
/* Atividade 2
* Use var para declarar três variáveis. Elas devem ter package-level scope. Não atribua valores a estas variáveis. Utilize os seguintes identificadores e tipos para estas veriáveis.
* Identificador "x" deverá ter tipo int
* Identificador "y" deverá ter tipo string
* Identificador "z" deverá ter tipo bool
 */
func atividadeDois() {
	fmt.Println("Atividade 2:")
	var x int;
	var y string;
	var z bool;
	fmt.Printf("%v,\n%v,\n%v,\n",x ,y ,z)
	fmt.Println("Finalizado atividade\n\n")
}
/* Atividade 3
* Utilizando a solução do exercicio anterior:
* Em package-level scope atribua os seguintes valores às variáveis:
* 1 - pra x atribua 42
* 2 - pra x atribua "James Bond"
* 3 - pra x atribua true

* Use fmt.Sprintf para atribuir todos esses valores a uma única variável. 
* Faça essa atribuição de tipo string a uma variável de nome "s" utilizando o operador curto de declaração.
*/
func atividadeTres() {
	fmt.Println("Atividade 3:")
	var x int  = 42
	var y string  = "James Bond"
	var z bool = true

	s := fmt.Sprintf("%v\t%v\t%v", x ,y ,z )
	fmt.Println(s)
	fmt.Println("Finalizado atividade\n\n")
}
/* Atividade 4
* Crie um tipo. O tipo subjacente deve ser int.
* Crie uma variável para esse tipo, com o identificador "x", utilizando a palavra chave var.
* 1 - Demonstre o valor da variável "x".
* 2 - Demonstre o tipo da variável "x".
* 3 - Atribua 42 à variável "x", utilizando  o operador "=".
* 4 - Demonstre o valor da variável "x".
*/
func atividadeQuatro() {
	fmt.Println("Atividade 4:")
	fmt.Printf("%v\n", x)
	fmt.Printf("%t\n", x)
	x = 42
	fmt.Printf("%v\n", x)
	fmt.Println("Finalizado atividade\n\n")
}
/* Atividade 5 
* Utilizando a solução do exercicio anterior: 
* 1. Em package-level scope, utilizando a palavra-chave var, crie uma variavel com o identiifcador "y". O tipo desta variável 
deve ser o tipo subjacente do tipo que você criou no exercicio anterior.
* 2. Na função:
		1. Isto já deve estar feito: 
		* Demonstre o valor da variável "x".
		* Demonstre o tipo da variável "x".
		* Atribua 42 à variável "x" utilizando o operador "=".
		* Demostre o valor da variável "x".
		2. Agora faça também:
		* Utilize conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente e, utilizando o operador "=", atribua o valor de "x" a "y".
		* Demonstre o valor da variável "y".
		* Demonstre o tipo de "y".
*/
func atividadeCinco(){
	fmt.Println("Atividade 5:")
	y := int(x)
	fmt.Println(y)
	fmt.Printf("%t", y)
	fmt.Println("Finalizado atividade\n\n")
}