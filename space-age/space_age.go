package space

type Planet string

// Calculate a year based in seconds
func Age(seconds float64, planet Planet) float64 {
	var years float64
	const (
		earthSeconds = 31557600
		mercuryPeriod float64 = 0.2408467
		venusPeriod   float64 = 0.61519726
		marsPeriod    float64 = 1.8808158
		jupiterPeriod float64 = 11.862615
		saturnPeriod  float64 = 29.447498
		uranusPeriod  float64 = 84.016846
		neptunePeriod float64 = 164.79132
	)

	switch planet {
	case "Earth":
		years = seconds / earthSeconds
	case "Mercury":
		years = seconds / (mercuryPeriod*earthSeconds)
	case "Venus":
		years = seconds / (venusPeriod*earthSeconds)
	case "Mars":
		years = seconds / (marsPeriod*earthSeconds)
	case "Jupiter":
		years = seconds / (jupiterPeriod*earthSeconds)
	case "Saturn":
		years = seconds / (saturnPeriod*earthSeconds)
	case "Uranus":
		years = seconds / (uranusPeriod*earthSeconds)
	case "Neptune":
		years = seconds / (neptunePeriod*earthSeconds)
	}

	return years
}
