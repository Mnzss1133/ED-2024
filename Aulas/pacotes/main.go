package main
import(
	"fmt"
	"projeto/modelos"
	"projeto/modelos/academico"
)
// ciontrl shift p -> buil taks -> executar pacote user
//go mod init projeto
func main(){
	fmt.Println("ola,mundo")

	aluno := modelos.Aluno{}
	aluno.Nome = "abcd"
	aluno.Matricula = "1234"

	curso :=academico.Curso{Nome:"Engenharia"}

	fmt.Println(aluno, curso)

}