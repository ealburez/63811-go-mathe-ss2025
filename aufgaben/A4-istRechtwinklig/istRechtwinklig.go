package main

// istRechtwinklig prüft, ob das durch die Seitenlängen x, y, z gebildete
// Dreieck rechtwinklig ist.
func istRechtwinklig(x, y, z int) bool {
	
	if x > y && x > z{
		return pythagorasCheck(x,y,z)
	} else if y > x && y > z {
		return pythagorasCheck(y,x,z)
	}else{
		return pythagorasCheck(z,y,x)
	}

}

func pythagorasCheck(a,b,c int) bool {
	result:= false
	if (a*a) == (b*b + c*c) {
		result = true
	}
	return result
}
