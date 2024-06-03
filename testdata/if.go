package testdata

func zero() int {
	return 0
}

func zeroAndOne() (int, int) {
	return 0, 1
}

func main() {
	if z := zero(); z > 0 { // want "variable z can be removed and use assignee directly"
		println("Zero is greater than 0")
	}

	if z := zero(); z > 0 {
		println(z)
	}

	if z := zero(); z > 0 { // want "variable z can be removed and use assignee directly"
		println("Zero is greater than 0")
	} else {
		println("Zero is not greater than 0")
	}

	if z := zero(); z > 0 {
		println("Zero is greater than 0")
	} else {
		println(z)
	}

	if z := zero(); z > 0 && z < 10 {
		println("Zero is greater than 0 and less than 10")
	}

	if z, _ := zeroAndOne(); z > 0 {
		println("Zero is greater than 0")
	}
}
