//package main

//import "fmt"

/* Questao 1 */

func calculaMedia(valores []float64) float64 {
    total := 0.0
    for _, v := range valores {
        total += v
    }
    return total / float64(len(valores))
}

func main() {
    
    valores := []float64{10.5, 15.3, 20.7, 8.9}
    media := calculaMedia(valores)
    fmt.Printf("A média é: %.2f\n", media)
}



/* Questao 2 */
func verificaParidade(numero int) string {
    if numero%2 == 0 {
        return "par"
    } else {
        return "impar"
    }
}

func main() {
    numero := 22
    resultado := verificaParidade(numero)
    fmt.Printf("%d é %s.\n", numero, resultado)
}


/*Questao 3*/

func minhaPotencia(base, expoente int) int {
    resultado := 1
    for expoente > 0 {
        if expoente%2 == 1 {
            resultado *= base
        }
        base *= base
        expoente /= 2
    }
    return resultado
}

func main() {
    base := 2
    expoente := 3
    resultado := minhaPotencia(base, expoente)
    fmt.Printf("%d elevado à %d é %d.\n", base, expoente, resultado)
}


/* Questao 4  */
func converteCelsiusParaFahrenheit(celsius float64) float64 {
    fahrenheit := (celsius * 9 / 5) + 32
    return fahrenheit
}

func main() {
   
    temperaturaCelsius := 22.0
    temperaturaFahrenheit := converteCelsiusParaFahrenheit(temperaturaCelsius)
    fmt.Printf("%.1f°C corresponde a %.1f°F\n", temperaturaCelsius, temperaturaFahrenheit)
}
