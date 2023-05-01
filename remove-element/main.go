package main

func removeLeftRight(arr []any, position string) []any {
	if position == "left" {
		return arr[1:]
	} else {
		return arr[:len(arr)-1]
	}
}

func main() {
	data1 := []any{1, 2, 3, 4, 5}
	fmt.Println(removeLeftRight(data1, "left"))
	fmt.Println(removeLeftRight(data1, "right"))

	data2 := []any{"Edo", "Budi", "Joko", "Tono"}
	fmt.Println(removeLeftRight(data2, "left"))
	fmt.Println(removeLeftRight(data2, "right"))
}
