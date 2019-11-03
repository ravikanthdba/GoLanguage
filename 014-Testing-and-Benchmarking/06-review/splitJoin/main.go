// This package is for joining the words
package splitJoin

import (
    "strings"
)

// This function takes the words list, which is a slice of string, and joins them with a separator " ". This is a fucntion created by us.
func StringJoin(wordsList []string, separator string) string {
    sentence := "";
    sentence += wordsList[0];

    for _, word := range wordsList[1:] {
        sentence += separator;
        sentence += word;
    }

    return sentence;
}

// This function takes the words list, which is a slice of string and joins them using the strings.Join() function.
func StringJoinWithBuiltinFunction(wordsList []string, separator string) string {
    return strings.Join(wordsList, separator)
}