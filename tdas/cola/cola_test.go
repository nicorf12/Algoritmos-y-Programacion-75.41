package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRecienCreada(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()

	//apenas creada
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

	//la usamos
	cola.Encolar("prueba1")
	cola.Encolar("prueba2")
	require.EqualValues(t, "prueba1", cola.Desencolar())
	require.EqualValues(t, "prueba2", cola.Desencolar())

	//luego de usarse
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestComportamiento(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[float64]()
	require.True(t, cola.EstaVacia())

	cola.Encolar(10.8)
	cola.Encolar(11.2)
	cola.Encolar(12.7)
	require.False(t, cola.EstaVacia())

	require.Equal(t, 10.8, cola.VerPrimero())
	require.Equal(t, 10.8, cola.Desencolar())
	require.Equal(t, 11.2, cola.VerPrimero())
	require.Equal(t, 11.2, cola.Desencolar())
	require.Equal(t, 12.7, cola.VerPrimero())
	require.Equal(t, 12.7, cola.Desencolar())

	require.True(t, cola.EstaVacia())
}

func TestVolumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 1; i <= 10000; i++ {
		cola.Encolar(i)

	}
	for j := 1; j <= 10000; j++ {
		require.Equal(t, j, cola.VerPrimero())
		require.Equal(t, j, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
}
