package utilidades

import (
	"math/rand"
	TDACola "tdas/cola"
	TDAGrafo "tdas/grafo"
	TDAPila "tdas/pila"
	"time"
)

const (
	D           = 0.85
	K_comunidad = 40
	K_pagerank  = 20
)

func BFS[V comparable](grafo TDAGrafo.Grafo[V], inicio V, destino *V) (map[V]V, map[V]int) {

	if !grafo.ExisteVertice(inicio) || (destino != nil && !grafo.ExisteVertice(*destino)) {
		panic("No existen los vertices")
	}

	visitado := make(map[V]bool)
	cola := TDACola.CrearColaEnlazada[V]()
	padres := make(map[V]V)
	orden := make(map[V]int)

	cola.Encolar(inicio)
	visitado[inicio] = true
	orden[inicio] = 0

	for !cola.EstaVacia() {

		v := cola.Desencolar()
		if destino != nil && v == *destino {
			return padres, orden
		}
		for _, w := range grafo.ObtenerAdyacentes(v) {
			if !visitado[w] {
				cola.Encolar(w)
				visitado[w] = true
				padres[w] = v
				orden[w] = orden[v] + 1
			}
		}
	}
	return padres, orden
}

func ReconstruirCamino[V comparable](padres map[V]V, inicio V, destino V) ([]V, int) {
	camino := []V{}

	actual := destino
	costo := 0

	for actual != inicio {
		costo++
		camino = append(camino, actual)
		actual = padres[actual]
		if costo == 1000 {
			return []V{}, -1
		}
	}
	camino = append(camino, inicio)
	return camino, costo
}

func _obtenerCiclo[V comparable](grafo TDAGrafo.Grafo[V], v V, n int, visitados map[V]bool, orden map[V]int, ciclo *[]V) bool {
	visitados[v] = true
	if orden[v] > n {
		visitados[v] = false
		return false
	}

	*ciclo = append(*ciclo, v)

	for _, w := range grafo.ObtenerAdyacentes(v) {
		if !visitados[w] {
			orden[w] = orden[v] + 1
			if _obtenerCiclo(grafo, w, n, visitados, orden, ciclo) {
				return true
			}
		} else if orden[w] == 0 && len(*ciclo) == n {
			*ciclo = append(*ciclo, w)
			return true
		}
	}

	visitados[v] = false
	*ciclo = (*ciclo)[:len(*ciclo)-1]
	return false
}

func ObtenerCiclo[V comparable](grafo TDAGrafo.Grafo[V], v V, n int) []V {
	visitados := make(map[V]bool)
	orden := make(map[V]int)
	ciclo := make([]V, 0, n)
	orden[v] = 0

	_obtenerCiclo(grafo, v, n, visitados, orden, &ciclo)

	return ciclo
}

func OrdenTopologico[V comparable](grafo TDAGrafo.Grafo[V]) []V {
	gr_e := gradosEntrada(grafo)
	c := TDACola.CrearColaEnlazada[V]()
	orden := make([]V, 0)
	for key, value := range gr_e {
		if value == 0 {
			c.Encolar(key)
		}
	}

	for !c.EstaVacia() {
		v := c.Desencolar()
		orden = append(orden, v)
		for _, w := range grafo.ObtenerAdyacentes(v) {
			_, esta := gr_e[w]
			if esta && w != v {
				gr_e[w]--
				if gr_e[w] == 0 {
					c.Encolar(w)
				}
			}
		}
	}
	return orden
}

func gradosEntrada[V comparable](grafo TDAGrafo.Grafo[V]) map[V]int {
	gr_e := make(map[V]int)
	vertices := grafo.ObtenerVertices()
	for _, v := range vertices {
		gr_e[v] = 0
	}

	for _, v := range vertices {
		for _, w := range grafo.ObtenerAdyacentes(v) {
			gr_e[w]++
		}
	}
	return gr_e
}

func vertices_entrada[V comparable](grafo TDAGrafo.Grafo[V]) map[V][]V {
	resultado := make(map[V][]V)
	for _, v := range grafo.ObtenerVertices() {
		for _, w := range grafo.ObtenerAdyacentes(v) {
			if v != w {
				resultado[w] = append(resultado[w], v)
			}
		}
	}
	return resultado
}

func TarjanCFCS[V comparable](grafo TDAGrafo.Grafo[V], cfcs *[][]V) {
	visitados := make(map[V]bool)
	for _, v := range grafo.ObtenerVertices() {
		if !visitados[v] {
			var contador int
			_tarjanCFCS(v, grafo, cfcs, visitados, make(map[V]int), make(map[V]int), TDAPila.CrearPilaDinamica[V](), make(map[V]bool), &contador)
		}
	}
}

func _tarjanCFCS[V comparable](v V, grafo TDAGrafo.Grafo[V], cfcs *[][]V, visitados map[V]bool, orden map[V]int, mb map[V]int, p TDAPila.Pila[V], apilados map[V]bool, contador *int) {
	mb[v] = *contador
	orden[v] = *contador
	(*contador)++
	visitados[v] = true
	p.Apilar(v)
	apilados[v] = true

	for _, w := range grafo.ObtenerAdyacentes(v) {
		if !visitados[w] {
			_tarjanCFCS(w, grafo, cfcs, visitados, orden, mb, p, apilados, contador)
		}
		if apilados[w] {
			mb[v] = min(mb[v], mb[w])
		}
	}

	if orden[v] == mb[v] {
		compo := make([]V, 0)
		for true {
			w := p.Desapilar()
			apilados[w] = false
			compo = append(compo, w)
			if w == v {
				break
			}
		}
		*cfcs = append(*cfcs, compo)
	}
}

func EncontrarMaximo[V comparable](m map[V]int) (int, V) {
	var max_int int
	var max_v V

	for clave, valor := range m {
		if valor > max_int {
			max_int = valor
			max_v = clave
		}
	}

	return max_int, max_v
}

func min(n1, n2 int) int {
	if n1 > n2 {
		return n2
	}
	return n1
}

func CalcularPageRank[V comparable](dict_page_rank map[V]float32, grafo TDAGrafo.Grafo[V]) {
	vertices := grafo.ObtenerVertices()
	vertices_e := vertices_entrada(grafo)
	for i := 0; i < K_pagerank; i++ {
		for _, v := range vertices {
			pr(v, dict_page_rank, len(vertices), vertices_e, grafo)
		}
	}
}

func pr[V comparable](v V, dict_page_rank map[V]float32, n int, vertices_e map[V][]V, grafo TDAGrafo.Grafo[V]) {
	var suma float32
	for _, w := range vertices_e[v] {
		adyacentes := grafo.ObtenerAdyacentes(w)
		if len(adyacentes) > 0 {
			suma += (dict_page_rank[w] / float32(len(adyacentes)))
		}
	}
	dict_page_rank[v] = ((1 - D) / float32(n)) + D*suma
}

func LabelPropagation[V comparable](grafo TDAGrafo.Grafo[V], label map[V]V, raiz V) []V {
	vertices := grafo.ObtenerVertices()
	vertices_e := vertices_entrada(grafo)
	comunidad := make([]V, 0)
	for _, v := range vertices {
		label[v] = v
	}
	shuffleArray(&vertices)
	for i := 0; i < K_comunidad; i++ {
		for _, v := range vertices {
			label[v] = max_freq(vertices_e[v], label, v)
		}
	}

	for _, v := range grafo.ObtenerVertices() {
		if label[v] == label[raiz] {
			comunidad = append(comunidad, v)
		}
	}
	return comunidad
}

/* func LabelPropagation[V comparable](grafo TDAGrafo.Grafo[V], label map[V]V) {
	vertices := grafo.ObtenerVertices()
	vertices_e := vertices_entrada(grafo)
	for _, v := range vertices {
		label[v] = v
	}
	shuffleArray(&vertices)
	for i := 0; i < K_comunidad; i++ {
		for _, v := range vertices {
			label[v] = max_freq(vertices_e[v], label, v)
		}
	}
}
*/
func max_freq[V comparable](v_entrada []V, label map[V]V, actual V) V {
	if len(v_entrada) == 0 {
		return label[actual]
	}

	aux := make(map[V]int)
	for _, v := range v_entrada {
		aux[label[v]]++
	}

	rand.Seed(time.Now().UnixNano())
	i_a := rand.Intn(len(v_entrada))
	var maxFrecuencia int = aux[label[v_entrada[i_a]]]
	var maxComunidad V = label[v_entrada[i_a]]

	for comunidad, frecuencia := range aux {
		if frecuencia > maxFrecuencia {
			maxComunidad = comunidad
			maxFrecuencia = frecuencia
		}
	}
	return maxComunidad
}

func shuffleArray[V comparable](arr *[]V) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(*arr), func(i, j int) {
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	})
}

func CoeficienteClusteringGenerico[V comparable](grafo TDAGrafo.Grafo[V], dict_CC map[V]float32) {
	for _, v := range grafo.ObtenerVertices() {
		dict_CC[v] = CoeficienteClustering(grafo, v)
	}
}

func CoeficienteClustering[V comparable](grafo TDAGrafo.Grafo[V], v V) float32 {
	ady := grafo.ObtenerAdyacentes(v)
	var ady_cont int
	var contador int
	for _, d := range ady {
		if d == v {
			continue
		}
		for _, w := range ady {
			if w == v {
				continue
			}
			if d != w && grafo.ExisteArista(d, w) {
				contador++
			}
		}
		ady_cont++
	}

	if ady_cont < 2 {
		return 0
	}

	return float32(contador) / float32(ady_cont*(ady_cont-1))
}

func Invertir[V comparable](arr []V) {
	length := len(arr)
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
