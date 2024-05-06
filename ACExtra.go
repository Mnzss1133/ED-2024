package main

import "fmt"

type No struct {
	valor int
	prox  *No
}

type Fila struct {
	tamanho int
	inicio  *No
	fim     *No
}

// Percorrer a fila
func (f *Fila) percorrer() {
	no := f.inicio
	if no == nil {
		fmt.Println("Fila vazia!")
	} else {
		for no.prox != nil {
			fmt.Print(no.valor, " -> ")
			no = no.prox
		}
		fmt.Println(no.valor)
	}
}

// Enfileirar um elemento na fila
func (f *Fila) enfileirar(novoValor int) {
	no := &No{valor: novoValor}
	if f.inicio == nil {
		f.inicio = no
	} else {
		f.fim.prox = no
	}
	f.fim = no
	f.tamanho++
}

// Desenfileirar um elemento da fila
func (f *Fila) desenfileirar() (int, error) {
	if f.inicio == nil {
		return 0, fmt.Errorf("Fila vazia!")
	}

	valorRemovido := f.inicio.valor
	f.inicio = f.inicio.prox
	f.tamanho--

	if f.inicio == nil {
		f.fim = nil
	}

	return valorRemovido, nil
}

func main() {
	var fila Fila
	fila.percorrer()

	fila.enfileirar(10)
	fila.enfileirar(20)
	fila.enfileirar(30)

	fila.percorrer()

	valorRemovido, _ := fila.desenfileirar()
	fmt.Println("Valor removido:", valorRemovido)

	fila.percorrer()

	fila.enfileirar(40)
	fila.enfileirar(50)

	fila.percorrer()
}
