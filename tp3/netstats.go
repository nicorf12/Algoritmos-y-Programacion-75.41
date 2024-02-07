package main

import (
	"bufio"
	"fmt"
	"netstats/errores"
	"netstats/utilidades"
	"os"
	"strings"
	TDADiccionario "tdas/diccionario"
	TDAGrafo "tdas/grafo"
)

const (
	PARAMETROS_INICIALES_ESPERADOS = 2
)

func main() {
	var args = os.Args

	if len(args) != PARAMETROS_INICIALES_ESPERADOS {
		fmt.Println(errores.ErrorParametros{}.Error())
		return
	}

	archivo, err := os.Open(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer archivo.Close()

	grafo := utilidades.MoldearEnGrafo(archivo)

	dict_page_rank := make(map[string]float32)
	dict_cfc := make(map[string]int)
	cfcs := make([][]string, 0)
	label := make(map[string]string)
	comunidades := make(map[string][]string)
	dict_CC := make(map[string]float32)

	comandos := []utilidades.Comando{
		{Nombre: "camino", Funcion: utilidades.Camino},
		{Nombre: "diametro", Funcion: utilidades.Diametro},
		{Nombre: "rango", Funcion: utilidades.Rango},
		{Nombre: "listar_operaciones", Funcion: utilidades.Listar_operaciones},
		{Nombre: "ciclo", Funcion: utilidades.Ciclo},
		{Nombre: "lectura", Funcion: utilidades.Lectura},
		{Nombre: "navegacion", Funcion: utilidades.Navegacion},
		{Nombre: "clustering", Funcion: utilidades.Clustering},
		{Nombre: "mas_importantes", Funcion: utilidades.Mas_importantes},
		{Nombre: "conectados", Funcion: utilidades.Conectados},
		{Nombre: "comunidad", Funcion: utilidades.Comunidad},
	}

	dict_comandos := TDADiccionario.CrearHash[string, func([]string, TDAGrafo.Grafo[string], []utilidades.Comando, map[string]float32, map[string]int, *[][]string, map[string]string, map[string][]string, map[string]float32)]()

	inicializarDiccionarioComandos(comandos, dict_comandos)

	input := bufio.NewScanner(os.Stdin)
	var fin bool
	for !fin {
		input.Scan()
		str_comando := input.Text()
		str_comando = strings.TrimSpace(str_comando)
		partes_comando := strings.Fields(str_comando)
		var comando string
		var parametros []string

		if len(partes_comando) > 0 {
			comando = partes_comando[0]
			parametros = partes_comando[1:]
		}

		if !dict_comandos.Pertenece(comando) {
			fin = true
		} else {
			dict_comandos.Obtener(comando)(parametros, grafo, comandos, dict_page_rank, dict_cfc, &cfcs, label, comunidades, dict_CC)
		}
	}
}

func inicializarDiccionarioComandos(comandos []utilidades.Comando, dict_comandos TDADiccionario.Diccionario[string, func([]string, TDAGrafo.Grafo[string], []utilidades.Comando, map[string]float32, map[string]int, *[][]string, map[string]string, map[string][]string, map[string]float32)]) {
	for _, c := range comandos {
		dict_comandos.Guardar(c.Nombre, c.Funcion)
	}
}
