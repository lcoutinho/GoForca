package main

import(
	"fmt"
	"strings"
	"syscall"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {

	errorsImage := map[int]string{
		0: `
    +---+
    |   |
    	|
    	|
    	|
        |
=========`,
		1: `

    +---+
    |   |
    O   |
	|
	|
	|
=========`,
		2:   `	
    +---+
    |   |
    O   |
    |   |
	|
	|
=========`,
		3: `
   +---+
   |    |
   O    |
  /|    |
	|
	|
=========`,
		4: `
    +---+
    |   |
    O   |
   /|\  |
	|
	|
=========`,
		5: `
    +---+
    |   |
    O   |
   /|\  |
   /    |
	|
=========`,
		6: `
    +---+
    |   |
    O   |
   /|\  |
   / \  |
	|
=========`,
	}
	

var palavra string
var tentativa string
var digitadas []string
var acertos []string
erros := 0

fmt.Println("Digite sua palavra:")
bytesPalavra, _ := terminal.ReadPassword(int(syscall.Stdin)) 
palavra =   strings.Replace(string(bytesPalavra)," ","_",-1)

for{
	senha := ""
	for _, letra := range palavra{
		senha += func() string{
			if inArray(string(letra),acertos) {
				return string(letra)
			}else{
				return "_"
			}
		}()
	}
	fmt.Println(senha)

	if senha == palavra{
		fmt.Println("Você acertou!!")
		break
	}

	fmt.Println("Digite uma letra:")
	fmt.Scanln(&tentativa)

	if inArray(tentativa,digitadas){
		fmt.Println("Você já digitou essa letra!")
		continue
	}else{
		digitadas = append(digitadas,tentativa)
		if strings.Index(palavra,tentativa) != -1{
			acertos = append(acertos,tentativa)
		}else{
			fmt.Println("Você errou!")
			erros+=1
		}
	}
	fmt.Println(errorsImage[erros])

	if erros == 6{
		fmt.Println("Enforcado!")
		break
	}
}
}


func inArray(val string, array []string) bool{
	for i := range array{
		if ok := array[i] == val; ok{
			return true
		}
	}
	return false
}
