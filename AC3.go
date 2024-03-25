//package main

//import "fmt"

// Definição do struct Contato
type Contato struct {
    Nome  string
    Email string
}

// Função para adicionar um contato ao array de contatos
func adicionarContato(contato Contato, arrayContatos []Contato) []Contato {
    for i := range arrayContatos {
        if arrayContatos[i].Nome == "" && arrayContatos[i].Email == "" {
            arrayContatos[i] = contato
            return arrayContatos
        }
    }
    fmt.Println("Não foi possível adicionar o contato. O array está cheio.")
    return arrayContatos
}

func main() {
    // Array de contatos com capacidade para até 5 elementos
    arrayContatos := make([]Contato, 2)

    // Criando alguns contatos
    contato1 := Contato{"Joao", "Joaopedromenezes890@gmail.com"}
    contato2 := Contato{"Victor", "victor@gmail.com"}

    // Adicionando os contatos ao array
    arrayContatos = adicionarContato(contato1, arrayContatos)
    arrayContatos = adicionarContato(contato2, arrayContatos)

    // Exibindo o array de contatos
    fmt.Println("Array de Contatos:")
    for _, contato := range arrayContatos {
        fmt.Printf("Nome: %s, Email: %s\n", contato.Nome, contato.Email)
    }
}
