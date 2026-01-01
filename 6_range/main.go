package main

func main() {
	nums := []int{10, 20, 30, 40, 50}

	for _, v := range nums {
		println(v)
	}

	m := map[string]string{
		"one":   "First",
		"two":   "Second",
		"three": "Third",
	}

	for k, v := range m {
		println(k, ":", v)
	}
}
