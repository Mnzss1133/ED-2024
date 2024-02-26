package main
import(
	"fmt"
	"projeto/modelos"
	"projeto/modelos/academico"
)
//go mod init projeto
func main(){
	fmt.Println("ola,mundo")

	aluno := modelos.Aluno{}
	aluno.Nome = "abcd"
	aluno.Matricula = "1234"

	curso :=academico.Curso{Nome:"Engenharia"}

	fmt.Println(aluno, curso)

}