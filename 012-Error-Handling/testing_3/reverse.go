package main

func reverse(s string) string {
    var reversed_string string = "";

    for str := len(s) - 1; str >= 0; str -- {
        reversed_string += string(s[str]);
    }

    return reversed_string;
}
