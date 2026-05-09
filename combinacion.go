// Derek Ramon Tirado Barajas 23760343
package main

import (
	"fmt"
	"sort"
)

type Producto struct {
	Nombre      string
	Precio      float64
	Categoria   string
	Descripcion string
}

var catalogo = []Producto{
	{Nombre: "Agua mineral", Precio: 20.00, Categoria: "Bebida", Descripcion: "Botella 600ml sin gas"},
	{Nombre: "Té verde", Precio: 30.00, Categoria: "Bebida", Descripcion: "Infusión caliente"},
	{Nombre: "Café Americano", Precio: 35.00, Categoria: "Bebida", Descripcion: "Café negro doble shot"},
	{Nombre: "Jugo de naranja", Precio: 45.00, Categoria: "Bebida", Descripcion: "Natural exprimido 355ml"},
	{Nombre: "Pastel de chocolate", Precio: 55.00, Categoria: "Postre", Descripcion: "Rebanada individual 120g"},
	{Nombre: "Burrito de res", Precio: 75.00, Categoria: "Comida", Descripcion: "Tortilla, carne, frijoles"},
	{Nombre: "Sandwich de pollo", Precio: 85.00, Categoria: "Comida", Descripcion: "Pan integral, pollo, verduras"},
	{Nombre: "Ensalada César", Precio: 95.00, Categoria: "Comida", Descripcion: "Lechuga romana, crutones"},
	{Nombre: "Pizza personal", Precio: 110.00, Categoria: "Comida", Descripcion: "4 rebanadas, queso y jitomate"},
	{Nombre: "Combo del día", Precio: 130.00, Categoria: "Combo", Descripcion: "Plato fuerte + bebida + postre"},
}

func encontrarCombinaciones(productos []Producto, presupuesto float64) [][]Producto {
	var resultado [][]Producto
	var actual []Producto

	var backtrack func(inicio int, gastoActual float64)
	backtrack = func(inicio int, gastoActual float64) {
		if len(actual) > 0 {
			copia := make([]Producto, len(actual))
			copy(copia, actual)
			resultado = append(resultado, copia)
		}

		for i := inicio; i < len(productos); i++ {
			nuevoCosto := gastoActual + productos[i].Precio

			if nuevoCosto > presupuesto {
				continue
			}

			actual = append(actual, productos[i])
			backtrack(i+1, nuevoCosto)

			actual = actual[:len(actual)-1]
		}
	}

	backtrack(0, 0)
	return resultado
}

func totalCombinacion(combo []Producto) float64 {
	total := 0.0
	for _, p := range combo {
		total += p.Precio
	}
	return total
}

func imprimirResultados(combis [][]Producto, presupuesto float64) {
	fmt.Printf("\nPresupuesto: $%.2f\n", presupuesto)

	if len(combis) == 0 {
		fmt.Println("No se encontraron combinaciones dentro del presupuesto.")
		return
	}

	fmt.Printf("Total de combinaciones: %d\n", len(combis))

	grupoPorTamaño := make(map[int][]int)
	for i, combo := range combis {
		n := len(combo)
		grupoPorTamaño[n] = append(grupoPorTamaño[n], i)
	}

	tamaños := make([]int, 0, len(grupoPorTamaño))
	for t := range grupoPorTamaño {
		tamaños = append(tamaños, t)
	}
	sort.Ints(tamaños)

	fmt.Println("\nPor cantidad de productos:")
	for _, t := range tamaños {
		cantidad := len(grupoPorTamaño[t])
		if t == 1 {
			fmt.Printf("  1 producto:  %d combinación(es)\n", cantidad)
		} else {
			fmt.Printf("  %d productos: %d combinación(es)\n", t, cantidad)
		}
	}

	fmt.Println()
	for i, combo := range combis {
		total := totalCombinacion(combo)
		cambio := presupuesto - total
		n := len(combo)

		if n == 1 {
			fmt.Printf("[%d] 1 producto — Total: $%.2f — Cambio: $%.2f\n", i+1, total, cambio)
		} else {
			fmt.Printf("[%d] %d productos — Total: $%.2f — Cambio: $%.2f\n", i+1, n, total, cambio)
		}

		for _, p := range combo {
			fmt.Printf("     • %-22s $%.2f\n", p.Nombre, p.Precio)
		}
	}

	mejor := combis[0]
	for _, combo := range combis[1:] {
		if totalCombinacion(combo) > totalCombinacion(mejor) {
			mejor = combo
		}
	}

	totalMejor := totalCombinacion(mejor)
	cambioMejor := presupuesto - totalMejor
	fmt.Println("\nMejor combinación (mayor gasto):")
	for _, p := range mejor {
		fmt.Printf("     • %-22s $%.2f\n", p.Nombre, p.Precio)
	}
	fmt.Printf("     Total: $%.2f  Cambio: $%.2f\n", totalMejor, cambioMejor)
}

func main() {
	var presupuesto float64
	fmt.Print("Ingresa tu presupuesto: $")
	_, err := fmt.Scan(&presupuesto)
	if err != nil || presupuesto <= 0 {
		fmt.Println("Error: ingresa un número mayor a cero.")
		return
	}

	combis := encontrarCombinaciones(catalogo, presupuesto)
	imprimirResultados(combis, presupuesto)
}
