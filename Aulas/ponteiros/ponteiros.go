package main

import(
	 "fmt"
	 "strings"
)


func main(){
	x:=2
	y:=x

	fmt.Println(x,y)

	x =3

	fmt.Println(x,y)
	z := &x//referencia
	fmt.Println(x,z)
	fmt.Println(x, *z)//dereferencia
	x=4
	fmt.Println(x,z)
	fmt.Println(x, *z)//dereferencia

	msg1 := "ola,mundo"
	alteraMensagem(&msg1)
	fmt.Println(msg1)

	n1:= Numero{Valor:10}
	fmt.Println(n1)//10
	n1.Incremento()
	fmt.Println(n1)
}
type Numero struct{
	Valor int
}
func(n *Numero)Incremento(){
	n.Valor++
}
/*usamos ponteiros como parametros de funcoes, quando:
1 -> é necessario reduzir o espaco consumido em memoria
2 -> queremos alterar o valor dos parametros

*/
func alteraMensagem(msg *string){
	novaMensagem := strings.ReplaceAll(*msg, "mundo","turma")
	*msg = novaMensagem
}