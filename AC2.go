//Questao 1 


import (
	"fmt"
	"math/rand"
	"time"

)

func main() {
	// Seed para gerar números aleatórios diferentes a cada execução
	rand.Seed(time.Now().UnixNano())

	// Criando um vetor de inteiros com 10 elementos
	vetor := make([]int, 10)

	// Inicializando o vetor com números aleatórios
	for i := 0; i < 10; i++ {
		vetor[i] = rand.Intn(100) // Números aleatórios entre 0 e 99
	}

	// Imprimindo os elementos do vetor um abaixo do outro
	for _, num := range vetor {
		fmt.Println(num)
	}
}

//Questao 2




// Função para inverter uma string
func inverterString(str string) string {
	runes := []rune(str) // Converte a string para um slice de runes

	// Inverte o slice de runes
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes) // Converte o slice de runes de volta para uma string
}

func main() {
	// Solicita uma entrada do usuário
	var entrada string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&entrada)

	// Chama a função para inverter a string
	stringInvertida := inverterString(entrada)

	// Imprime a string invertida
	fmt.Println("String invertida:", stringInvertida)
}

//Questao 3



// Definindo o tipo Ponto
type Ponto struct {
	X, Y float64
}

// Método para calcular a distância do ponto até a origem
func (p *Ponto) DistanciaOrigem() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {
	// Criando um ponto
	p := Ponto{3, 4}

	// Calculando a distância até a origem
	distancia := p.DistanciaOrigem()

	// Imprimindo a distância
	fmt.Printf("A distância do ponto (%.2f, %.2f) até a origem é %.2f\n", p.X, p.Y, distancia)
}
//Questao 4


func main() {
	// Solicita as dimensões do retângulo ao usuário
	var base, altura float64
	fmt.Println("Digite a base e a altura do retângulo:")
	fmt.Scanln(&base, &altura)

	// Calcula a área e o perímetro usando as funções do pacote geometria
	area := geometria.AreaRetangulo(base, altura)
	perimetro := geometria.PerimetroRetangulo(base, altura)

	// Imprime os resultados
	fmt.Printf("Área do retângulo: %.2f\n", area)
	fmt.Printf("Perímetro do retângulo: %.2f\n", perimetro)
}
