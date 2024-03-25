package main

import (
    "fmt"
    "strings"
)

// Função para calcular o maior divisor comum (MDC)
func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a%b
    }
    return a
}

// Função para simplificar uma fração
func simplifyFraction(num, den int) (int, int) {
    divisor := gcd(num, den)
    return num / divisor, den / divisor
}

// Função para realizar operações em frações
func performOperation(n1, d1, n2, d2 int, op rune) (int, int) {
    var num, den int

    switch op {
    case '+':
        num = n1*d2 + n2*d1
        den = d1 * d2
    case '-':
        num = n1*d2 - n2*d1
        den = d1 * d2
    case '*':
        num = n1 * n2
        den = d1 * d2
    case '/':
        num = n1 * d2
        den = n2 * d1
    }

    return simplifyFraction(num, den)
}

func main() {
    var N int
    fmt.Scanln(&N)

    for i := 0; i < N; i++ {
        var x, y string
        var op rune

        fmt.Scan(&x, &op, &y)

        xParts := strings.Split(x, "/")
        yParts := strings.Split(y, "/")

        xNum := parseFractionPart(xParts[0])
        xDen := parseFractionPart(xParts[1])
        yNum := parseFractionPart(yParts[0])
        yDen := parseFractionPart(yParts[1])

        resultNum, resultDen := performOperation(xNum, xDen, yNum, yDen, op)

        fmt.Printf("%d/%d = ", resultNum, resultDen)
        simplifiedNum, simplifiedDen := simplifyFraction(resultNum, resultDen)
        fmt.Printf("%d/%d\n", simplifiedNum, simplifiedDen)
    }
}

// Função para converter uma parte de fração (numerador ou denominador) em um número inteiro
func parseFractionPart(part string) int {
    num, _ := fmt.Scanf("%d", &num)
    return num
}





//Questao2

func main() {
	var N int     // Número de pessoas no círculo
	var L float64 // Volume de água no termo (em litros)
	var Q float64 // Volume de água em um mate (em litros)

	// Ler entradas
	fmt.Scanln(&N)
	fmt.Scanln(&L)
	fmt.Scanln(&Q)

	// Ler os nomes dos participantes
	participantes := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&participantes[i])
	}

	// Calcular a quantidade de água que cada pessoa bebe
	quantidadePorPessoa := Q

	// Calcular a quantidade de água restante após cada pessoa beber
	aguaRestante := L
	for i := 0; i < N-1; i++ {
		aguaRestante -= quantidadePorPessoa
	}

	// Calcular a quantidade de água que a última pessoa bebe
	aguaUltimaPessoa := aguaRestante

	// Encontrar o nome da última pessoa
	indiceUltimaPessoa := (N - 1) % N
	nomeUltimaPessoa := participantes[indiceUltimaPessoa]

	// Imprimir o resultado
	fmt.Printf("%s, o rico do chimarrão, %.1f\n", nomeUltimaPessoa, aguaUltimaPessoa)
}

//Questao 3


//import "fmt"

func main() {
    // Declare variables to store input values
    var code1, units1, code2, units2 int
    var price1, price2 float64

    // Read input values
    fmt.Scanln(&code1, &units1, &price1)
    fmt.Scanln(&code2, &units2, &price2)

    // Calculate the total amount to be paid
    total := (float64(units1) * price1) + (float64(units2) * price2)

    // Print the result
    fmt.Printf("VALOR A PAGAR: R$ %.2f\n", total)
}
