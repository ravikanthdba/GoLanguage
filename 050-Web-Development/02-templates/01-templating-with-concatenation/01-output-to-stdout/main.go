package _1_output_to_stdout

import "fmt"

func main() {
	name := "Ravikanth Garimella"

	text := `
		<!DOCTYPE html> 
		<html>
			<title> Example </title>
			<body>
				<p1>` + name + `</p> 
			</body>
		</html>
	`
	fmt.Println(text)
}
