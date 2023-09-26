package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()
	require.True(t, pila.EstaVacia())
	pila.Apilar("hola")
	require.False(t, pila.EstaVacia())
	pila.Desapilar()
	require.True(t, pila.EstaVacia())
}

func TestComportamiento(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(10)
	require.Equal(t, 10, pila.VerTope())
	pila.Apilar(11)
	require.Equal(t, 11, pila.VerTope())
	pila.Apilar(12)
	require.Equal(t, 12, pila.VerTope())
	require.Equal(t, 12, pila.Desapilar())
	require.Equal(t, 11, pila.Desapilar())
	require.Equal(t, 10, pila.Desapilar())
}

func TestVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 1; i <= 10000; i++ {
		pila.Apilar(i)

	}
	for j := 0; j < 10000; j++ {
		require.Equal(t, 10000-j, pila.VerTope())
		require.Equal(t, 10000-j, pila.Desapilar())
	}
	require.True(t, pila.EstaVacia())
}

func TestComportamientosInvalidos(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[float64]()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	pila.Apilar(1.34)
	pila.Desapilar()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}
