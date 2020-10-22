package main

import(
	"fmt"
	"os"
	"sort"
)
//Entradas para el programa
type Entradas struct{
	Input string
}

//Ascendente para ordenar en ese orden
type Ascendente []Entradas

func (a Ascendente) Len() int           { return len(a) }
func (a Ascendente) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Ascendente) Less(i, j int) bool { return a[i].Input < a[j].Input }

//Descendente para ordenar en ese orden
type Descendente []Entradas

func (a Descendente) Len() int           { return len(a) }
func (a Descendente) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Descendente) Less(i, j int) bool { return a[i].Input < a[j].Input }

func main(){
	n := 0
	aux := ""
	In := []Entradas{}
	fmt.Println("Indique el numero de strings a insertar: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++{
		fmt.Println("Ingrese el string ", i+1, ": ")
		fmt.Scan(&aux)
		In = append(In , Entradas{aux})
	}

	file, err := os.Create("ascendente.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer file.Close()
	sort.Sort(Ascendente(In))

	for j := 0; j < n; j++{
		file.WriteString(In[j].Input)
		file.WriteString("\n")
	}
	fmt.Println(In)

	desce, err := os.Create("descendente.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer desce.Close()
	sort.Sort(sort.Reverse(Descendente(In)))

	for k := 0; k < n; k++{
		desce.WriteString(In[k].Input)
		desce.WriteString("\n")
	}
	fmt.Println(In)

}