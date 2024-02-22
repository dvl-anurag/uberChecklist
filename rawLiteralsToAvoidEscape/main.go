package main

import "fmt"

func main() {
	// Without raw string literal - need to escape backslashes
	notRaw := "C:\\Users\\User\\Documents\\file.txt"
	fmt.Println(notRaw)

	// With raw string literal - backslashes are treated as literal
	raw := `C:\Users\User\Documents\file.txt`
	fmt.Println(raw)
}
