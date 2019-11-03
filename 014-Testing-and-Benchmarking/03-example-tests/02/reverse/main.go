// This package is used for reversing a particular variable.
package reverse 

import (
    "math"
)

/* This function takes a string and reverses the string */
func ReverseString(s string) string {
    reversed_string := "";

    for alphabet := len(s) - 1; alphabet >=0; alphabet -- {
        reversed_string += string(s[alphabet]);
    }

    return reversed_string;
}

/* This function takes a integer and reverses the integer */
func ReverseInteger(number int) int {
    reversed_number := 0;
    count := countInt(number)
    i := int(math.Pow(10, count - 1));
    
    for number != 0 {
        reversed_number += ((number % 10) * i);
        number = number / 10;
        i /= 10;
    }

    return reversed_number;
}

func countInt(number int) float64 {
    counter := 0;

    for number != 0 {
        counter ++;
        number /= 10;
    }

    return float64(counter);
}
