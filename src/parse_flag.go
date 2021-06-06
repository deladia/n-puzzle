package main

func parse_flag(hueristic_name string) bool {
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