package main

func parse_flag_h(hueristic_name string) bool {
	switch hueristic_name {
	case "hamming":
		return true
	case "manhattan":
		return true
	case "conflict":
		return true
	default:
		return false
	}
}

func parse_flag_s(algorithm_name string) bool {
	switch algorithm_name {
	case "a_star":
		return true
	case "greedy":
		return true
	case "uniform":
		return true
	default:
		return false
	}
}