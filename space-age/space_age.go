package space

//Planet is of type string
type Planet string

//Age takes age in seconds and a platenmt as varialbes and outputs the number of years a person of age s would be on planet p
func Age(s float64, p Planet) float64 {

	var y float64

	// 31557600 = 60 * 60 *24 *365.35
	y = s / 31557600
	switch p {

	case "Mercury":
		y = y / 0.2408467
	case "Venus":
		y = y / 0.61519726
	case "Earth":
		y = y
	case "Mars":
		y = y / 1.8808158
	case "Jupiter":
		y = y / 11.862615
	case "Saturn":
		y = y / 29.447498
	case "Uranus":
		y = y / 84.016846
	case "Neptune":
		y = y / 164.79132
	}
	return y
}
