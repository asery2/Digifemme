package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] == "" {
		// Si l'argument est manquant ou une chaîne vide, le programme ne produit aucune sortie.
		return
	}
	// Lecture du contenu du fichier "standard.txt" dans la variable "template".
	template, err := ioutil.ReadFile("standard.txt")
	if err != nil {
		// Affiche une erreur si la lecture du fichier échoue.
		fmt.Println(err)
		return
	}
	// Traitement de l'entrée en remplaçant les caractères "\\n" par des sauts de ligne.
	input := strings.ReplaceAll(os.Args[1], "\\n", "\n")
	if input == "\n" {
		// Si l'entrée est une nouvelle ligne, le programme affiche simplement une nouvelle ligne.
		fmt.Println()
		return
	}
	// Vérification des caractères de l'entrée pour s'assurer qu'ils sont affichables en ASCII.
	for i := 0; i < len(input); i++ {
		if (input[i] < 32 || input[i] > 127) && input[i] != 10 {
			fmt.Println("Entrée incorrecte.")
			return
		}
	}
	// Séparation du modèle de texte en blocs représentant chaque caractère ASCII.
	splitted := strings.Split(string(template)[1:], "\n\n")
	if len(splitted) != 95 {
		fmt.Println("Modèle incorrect.")
		return
	}
	// Création de la représentation ASCII pour chaque ligne de l'entrée.
	lines := strings.Split(input, "\n")
	res := ""
	for _, line := range lines {
		if line == "" && res != "" {
			res += string('\n')
			continue
		}
		for row := 0; row < 8; row++ {
			for i := 0; i < len(line); i++ {
				temp := strings.Split(splitted[line[i]-32], "\n")[row]
				res += temp
			}
			res += string('\n')
		}
	}
	// Affichage du résultat final.
	fmt.Print(res)
}
