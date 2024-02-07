package utilidades

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDAColaPrioridad "tdas/cola_prioridad"
	TDAGrafo "tdas/grafo"
)

const (
	NO_IMPORTA = 0
)

type Comando struct {
	Nombre  string
	Funcion func([]string, TDAGrafo.Grafo[string], []Comando, map[string]float32, map[string]int, *[][]string, map[string]string, map[string][]string, map[string]float32)
}

func MoldearEnGrafo(archivo *os.File) TDAGrafo.Grafo[string] {
	grafo := TDAGrafo.CrearGrafo[string](true, false)
	s := bufio.NewScanner((archivo))
	for s.Scan() {
		resultado := strings.Split(s.Text(), "	")
		if !grafo.ExisteVertice(resultado[0]) {
			grafo.AgregarVertice(resultado[0])
		}

		for i := 1; i < len(resultado); i++ {
			if !grafo.ExisteVertice(resultado[i]) {
				grafo.AgregarVertice(resultado[i])
			}
			grafo.AgregarArista(resultado[0], resultado[i], NO_IMPORTA)
		}
	}
	return grafo
}

func Listar_operaciones(parametros []string, grafo TDAGrafo.Grafo[string], comandos []Comando, dict_page_rank map[string]float32, dict_cfc map[string]int, cfcs *[][]string, label map[string]string, comunidades map[string][]string, dict_CC map[string]float32) {
	for _, c := range comandos {
		if c.Nombre == "listar_operaciones" {
			continue
		}
		fmt.Println(c.Nombre)
	}
}

func Camino(parametros []string, grafo TDAGrafo.Grafo[string], comandos []Comando, dict_page_rank map[string]float32, dict_cfc map[string]int, cfcs *[][]string, label map[string]string, comunidades map[string][]string, dict_CC map[string]float32) {
	aux := strings.Join(parametros, " ")
	parametros = strings.Split(aux, ",")

	if len(parametros) != 2 {
		fmt.Println("PARAMETROS INCORRECTOS")
		return
	}

	inicio := parametros[0]
	destino := parametros[1]

	padres, _ := BFS(grafo, inicio, &destino)
	camino, costo := ReconstruirCamino(padres, inicio, destino)
	if len(camino) == 0 {
		fmt.Println("No se encontro recorrido")
		return
	}
	Invertir(camino)
	var salida string
	for _, vertice := range camino {
		salida += vertice + " -> "
	}
	fmt.Println(salida[:len(salida)-4])
	fmt.Println("Costo:", costo)
}

func Diametro(parametros []string, grafo TDAGrafo.Grafo[string], comandos []Comando, dict_page_rank map[string]float32, dict_cfc map[string]int, cfcs *[][]string, label map[string]string, comunidades map[string][]string, dict_CC map[string]float32) {
	if len(parametros) != 0 {
		fmt.Println("PARAMETROS INCORRECTOS")
		return
	}

	var padres_max map[string]string
	var inicio_max string
	var destino_max string

	diametro := 0
	for _, vertice := range grafo.ObtenerVertices() {
		padres, orden := BFS(grafo, vertice, nil)
		aux, destino := EncontrarMaximo(orden)
		if aux > diametro {
			diametro = aux
			inicio_max = vertice
			destino_max = destino
			padres_max = padres
		}
	}

	camino, _ := ReconstruirCamino(padres_max, inicio_max, destino_max)
	Invertir(camino)
	var salida string
	for _, vertice := range camino {
		salida += vertice + " -> "
	}

	fmt.Println(salida[:len(salida)-4])
	fmt.Println("Costo:", diametro)
}

func Rango(parametros []string, grafo TDAGrafo.Grafo[string], comandos []Comando, dict_page_rank map[string]float32, dict_cfc map[string]int, cfcs *[][]string, label map[string]string, comunidades map[string][]string, dict_CC map[string]float32) {
	separado := strings.Split(strings.Join(parametros, " "), ",")
	if len(separado) != 2 {
		return
	}

	inicio := separado[0]
	n, err := strconv.Atoi(separado[1])
	if err != nil {
		return
	}

	_, orden := BFS(grafo, inicio, nil)
	contador := 0
	for _, v := range grafo.ObtenerVertices() {
		if orden[v] == n {
			contador++
		}
	}
	fmt.Println(contador)
}

func Ciclo(parametros []string, grafo TDAGrafo.Grafo[string], comandos []Comando, dict_page_rank map[string]float32, dict_cfc map[string]int, cfcs *[][]string, label map[string]string, comunidades map[string][]string, dict_CC map[string]float32) {
	separado := strings.Split(strings.Join(parametros, " "), ",")
	if len(separado) != 2 {
		fmt.Println("PARAMETROS INCORRECTOS")
		return
	}
	v := separado[0]
	n, err := strconv.Atoi(separado[1])
	if err != nil {
		fmt.Println("PARAMETROS INCORRECTOS")
		return
	}

	ciclo := ObtenerCiclo(grafo, v, n)

	if len(ciclo) == 0 {
		fmt.Println("No se encontro recorrido")
		return
	}

	var salida string
	for i := 0; i < len(ciclo); i++ {
		salida += ciclo[i] + " -> "
	}
	fmt.Println(salida[:len(salida)-4])
}

func Lectura(parametros []string, grafo TDAGrafo.Grafo[string], comandos []Comando, dict_page_rank map[string]float32, dict_cfc map[string]int, cfcs *[][]string, label map[string]string, comunidades map[string][]string, dict_CC map[string]float32) {
	if len(parametros) == 0 {
		fmt.Println("PARAMETROS INCORRECTOS")
		return
	}
	aux := strings.Join(parametros, " ")
	vertices := strings.Split(aux, ",")
	grafo_aux := TDAGrafo.CrearGrafo[string](true, false)
	for _, v := range vertices {
		grafo_aux.AgregarVertice(v)
	}
	for _, v := range vertices {
		for _, w := range vertices {
			if grafo.ExisteArista(v, w) {
				grafo_aux.AgregarArista(v, w, NO_IMPORTA)
			}
		}
	}

	orden := OrdenTopologico(grafo_aux)
	if len(orden) != len(vertices) {
		fmt.Println("No existe forma de leer las paginas en orden")
		return
	}

	Invertir(orden)

	var salida string
	for _, w := range orden {
		salida += w + ", "
	}
	fmt.Println(salida[:len(salida)-2])
}

func Conectados(parametros []string, grafo TDAGrafo.Grafo[string], comandos []Comando, dict_page_rank map[string]float32, dict_cfc map[string]int, cfcs *[][]string, label map[string]string, comunidades map[string][]string, dict_CC map[string]float32) {
	v := strings.Join(parametros, " ")
	if !grafo.ExisteVertice(v) {
		panic("NO EXISTE TAL VERTICE")
	}

	if len(*cfcs) == 0 {
		TarjanCFCS(grafo, cfcs)
		for i := 0; i < len(*cfcs); i++ {
			for _, v := range (*cfcs)[i] {
				dict_cfc[v] = i
			}
		}
	}

	var salida string
	for _, w := range (*cfcs)[dict_cfc[v]] {
		salida += w + ", "
	}
	fmt.Println(salida[:len(salida)-2])
}

func Navegacion(parametros []string, grafo TDAGrafo.Grafo[string], comandos []Comando, dict_page_rank map[string]float32, dict_cfc map[string]int, cfcs *[][]string, label map[string]string, comunidades map[string][]string, dict_CC map[string]float32) {
	v := strings.Join(parametros, " ")
	if !grafo.ExisteVertice(v) {
		panic("NO EXISTE TAL VERTICE")
	}
	var salida string = v
	for i := 0; i < 20; i++ {
		ady := grafo.ObtenerAdyacentes(v)
		if len(ady) == 0 {
			break
		}
		v = ady[0]
		salida += " -> " + v
	}
	fmt.Println(salida)
}

func Mas_importantes(parametros []string, grafo TDAGrafo.Grafo[string], comandos []Comando, dict_page_rank map[string]float32, dict_cfc map[string]int, cfcs *[][]string, label map[string]string, comunidades map[string][]string, dict_CC map[string]float32) {
	if len(parametros) != 1 {
		fmt.Println("PARAMETROS INCORRECTOS")
		return
	}
	n, err := strconv.Atoi(parametros[0])
	if err != nil {
		fmt.Println("PARAMETROS INCORRECTOS")
		return
	}

	if len(dict_page_rank) == 0 {
		CalcularPageRank(dict_page_rank, grafo)
	}

	heap_page_rank := TDAColaPrioridad.CrearHeap[string](func(s1, s2 string) int {
		if dict_page_rank[s1] > dict_page_rank[s2] {
			return -1
		} else if dict_page_rank[s2] > dict_page_rank[s1] {
			return 1
		}
		return 0
	})

	vertices := grafo.ObtenerVertices()

	for i := 0; i < n; i++ {
		heap_page_rank.Encolar(vertices[i])
	}

	for i := n; i < len(vertices); i++ {
		if dict_page_rank[heap_page_rank.VerMax()] < dict_page_rank[vertices[i]] {
			heap_page_rank.Desencolar()
			heap_page_rank.Encolar(vertices[i])
		}
	}

	var salida string
	vertices_top := make([]string, n)
	for i := 0; i < n; i++ {
		b := heap_page_rank.Desencolar()
		vertices_top[i] = b
		salida += b + ", "
	}
	fmt.Println(salida[:len(salida)-2])
}

func Comunidad(parametros []string, grafo TDAGrafo.Grafo[string], comandos []Comando, dict_page_rank map[string]float32, dict_cfc map[string]int, cfcs *[][]string, label map[string]string, comunidades map[string][]string, dict_CC map[string]float32) {
	v := strings.Join(parametros, " ")
	if !grafo.ExisteVertice(v) {
		panic("NO EXISTE TAL VERTICE")
	}

	/* if len(comunidades) == 0 {
		LabelPropagation(grafo, label)
		for v, comunidad := range label {
			comunidades[comunidad] = append(comunidades[comunidad], v)
		}
	} */

	comunidad := LabelPropagation[string](grafo, label, v)

	/* var salida string
	for i := 0; i < len(comunidades[label[v]]); i++ {
		salida += comunidades[label[v]][i] + ", "
	} */

	var salida string
	for i := 0; i < len(comunidad); i++ {
		salida += comunidad[i] + ", "
	}

	fmt.Println(salida[:len(salida)-2])
}

func Clustering(parametros []string, grafo TDAGrafo.Grafo[string], comandos []Comando, dict_page_rank map[string]float32, dict_cfc map[string]int, cfcs *[][]string, label map[string]string, comunidades map[string][]string, dict_CC map[string]float32) {
	if len(parametros) != 0 {
		v := strings.Join(parametros, " ")
		if !grafo.ExisteVertice(v) {
			panic("NO EXISTE TAL VERTICE")
		}
		valor, esta := dict_CC[v]
		if esta {
			fmt.Printf("%.3f\n", valor)
		} else {
			fmt.Printf("%.3f\n", CoeficienteClustering(grafo, v))
		}
		return
	}

	if len(dict_CC) == 0 {
		CoeficienteClusteringGenerico(grafo, dict_CC)
	}

	var suma float32
	for _, valor := range dict_CC {
		suma += valor
	}

	fmt.Printf("%.3f\n", 1/float32(len(dict_CC))*suma)
}
