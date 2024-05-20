package main

/*Questao1
func main() {
	var horaDormir, minutoDormir, horaAcordar, minutoAcordar int

	for {
		// Lê os valores de entrada
		fmt.Scan(&horaDormir, &minutoDormir, &horaAcordar, &minutoAcordar)


		if horaDormir == 0 && minutoDormir == 0 && horaAcordar == 0 && minutoAcordar == 0 {
			return
		}

		// Converte as horas e minutos em minutos totais desde a meia-noite
		ho := horaDormir*60 + minutoDormir
		mi := horaAcordar*60 + minutoAcordar

		// Calcula a diferença de tempo em minutos
		if ho < mi {
			fmt.Println(mi - ho)
		} else {
			fmt.Println(1440 + (mi - ho))
		}
	}
}*/

/*Questao2


func main() {
	var n int

	// Leitura de entrada
	for {
		_, err := fmt.Scanf("%d", &n)
		if err != nil {
			break
		}


		switch {
		case n <= 1:
			fmt.Println("Top 1")
		case n <= 3:
			fmt.Println("Top 3")
		case n <= 5:
			fmt.Println("Top 5")
		case n <= 10:
			fmt.Println("Top 10")
		case n <= 25:
			fmt.Println("Top 25")
		case n <= 50:
			fmt.Println("Top 50")
		case n <= 100:
			fmt.Println("Top 100")
		}
	}
}
*/

/* Questao 3


func main() {
	for {
		var N, H, C, L int
		// Leitura do número de degraus
		_, err := fmt.Scan(&N)
		if err != nil {
			break
		}

		// Leitura das medidas dos degraus
		_, err = fmt.Scan(&H, &C, &L)
		if err != nil {
			break
		}

		// Calcula o comprimento da rampa usando o teorema de Pitágoras
		totalHorizontalLength := float64(N * C)
		totalVerticalHeight := float64(N * H)
		rampLength := math.Sqrt(totalHorizontalLength*totalHorizontalLength + totalVerticalHeight*totalVerticalHeight)


		area := (rampLength * float64(L)) / 10000


		fmt.Printf("%.4f\n", area)
	}
}

*/
