package grafo

type grafo[V comparable] struct {
	dic         map[V][]arista[V]
	cantidad    int
	es_dirigido bool
	es_pesado   bool
}

type arista[V comparable] struct {
	vertice V
	peso    int
}

func CrearGrafo[V comparable](es_dirigido bool, es_pesado bool) Grafo[V] {
	return &grafo[V]{make(map[V][]arista[V]), 0, es_dirigido, es_pesado}
}

func (graf *grafo[V]) AgregarVertice(v V) {
	if graf.ExisteVertice(v) {
		panic("YA EXISTIA ESTE VERTICE")
	}
	graf.dic[v] = make([]arista[V], 0)
	graf.cantidad++
}

func (graf *grafo[V]) SacarVertice(v V) {
	if !graf.ExisteVertice(v) {
		panic("NO HABIA TAL VERTICE")
	}
	delete(graf.dic, v)
	graf.cantidad--
	for _, value := range graf.dic {
		for i := 0; i < len(value); i++ {
			if value[i].vertice == v {
				value[i], value[0] = value[0], value[i]
				value = value[1:]
			}
		}
	}
}

func (graf *grafo[V]) AgregarArista(v, w V, peso int) {
	esta_v := graf.ExisteVertice(v)
	esta_w := graf.ExisteVertice(w)
	if !esta_v || !esta_w {
		panic("NO HABIAN TAL VERTICES")
	}
	if !graf.es_pesado {
		peso = 1
	}

	graf.dic[v] = append(graf.dic[v], arista[V]{vertice: w, peso: peso})

	if !graf.es_dirigido {
		graf.dic[w] = append(graf.dic[w], arista[V]{vertice: v, peso: peso})
	}
}

func (graf *grafo[V]) SacarArista(v, w V) {
	esta_v := graf.ExisteVertice(v)
	esta_w := graf.ExisteVertice(w)
	if !esta_v || !esta_w {
		panic("NO HABIAN TAL VERTICES")
	}
	var habia bool
	for i := 0; i < len(graf.dic[v]); i++ {
		if graf.dic[v][i].vertice == w {
			graf.dic[v][i], graf.dic[v][0] = graf.dic[v][0], graf.dic[v][i]
			graf.dic[v] = graf.dic[v][1:]
			habia = true
		}
	}
	if !graf.es_dirigido {
		for i := 0; i < len(graf.dic[w]); i++ {
			if graf.dic[w][i].vertice == v {
				graf.dic[w][i], graf.dic[w][0] = graf.dic[w][0], graf.dic[w][i]
				graf.dic[w] = graf.dic[w][1:]
				habia = true
			}
		}
	}

	if !habia {
		panic("NO HABIA TAL ARISTA")
	}
}

func (graf *grafo[V]) ExisteArista(v, w V) bool {
	esta_v := graf.ExisteVertice(v)
	esta_w := graf.ExisteVertice(w)
	if !esta_v || !esta_w {
		panic("NO HABIAN TAL VERTICES")
	}

	for i := 0; i < len(graf.dic[v]); i++ {
		if graf.dic[v][i].vertice == w {
			return true
		}
	}

	return false
}

func (graf *grafo[V]) ExisteVertice(v V) bool {
	_, esta_v := graf.dic[v]
	return esta_v
}

func (graf *grafo[V]) ObtenerVertices() []V {
	vertices := make([]V, graf.cantidad)
	i := 0
	for vertice := range graf.dic {
		vertices[i] = vertice
		i++
	}
	return vertices
}

func (graf *grafo[V]) ObtenerAdyacentes(v V) []V {
	if !graf.ExisteVertice(v) {
		panic("NO HABIAN TAL VERTICES")
	}
	adyacentes := make([]V, len(graf.dic[v]))
	for i := 0; i < len(graf.dic[v]); i++ {
		adyacentes[i] = graf.dic[v][i].vertice
	}
	return adyacentes
}

func (graf *grafo[V]) ObtenerPeso(v, w V) int {
	esta_v := graf.ExisteVertice(v)
	esta_w := graf.ExisteVertice(w)
	if !esta_v || !esta_w {
		panic("NO HABIAN TAL VERTICES")
	}
	if !graf.es_pesado {
		panic("EL GRAFO ES NO PESADO")
	}
	var peso int
	for i := 0; i < len(graf.dic[v]); i++ {
		if graf.dic[v][i].vertice == w {
			peso = graf.dic[v][i].peso
		}
	}
	return peso
}

func (graf *grafo[V]) EsPesado() bool {
	return graf.es_pesado
}
func (graf *grafo[V]) EsDirigido() bool {
	return graf.es_dirigido
}
