package twopointers

import (
	"math"
)

// Idea de por qué funciona mover el puntero de la menor columna. Llamemos a este puntero j.
/* Hay que mover un puntero hacia adentro, pero esto reduce el ancho del contenedor.
   La idea es encontrar una columna más alta para compensar ese ancho.
   Al mover j, si encontramos una columna más alta que i y j, habrá una ganancia de volumen entre i y la nueva columna.
   Si este nuevo volumen no es mayor que el anterior, moveremos entonces i y si la nueva columna es mayor, habrá una nueva
   ganancia de volumen. Entonces, en general, al mover el puntero más pequeño (j), uno está descartando la pérdida de volumen de haber elegido en su lugar j,
   ya que i provee la mayor ganancia al ser mayor.
*/
func MaxArea(height []int) int {
	front, back := 0, len(height)-1
	mostWater := 0
	currWater := 0

	for front < back {
		smallestHeight := int(math.Min(float64(height[front]), float64(height[back])))
		currWater = smallestHeight * (back - front)
		mostWater = int(math.Max(float64(mostWater), float64(currWater)))
		if height[front] < height[back] {
			front++
		} else {
			back--
		}
	}
	return mostWater
}
