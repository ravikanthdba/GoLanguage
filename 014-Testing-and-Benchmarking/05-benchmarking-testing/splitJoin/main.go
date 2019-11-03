package splitJoin

import (
    "strings"
)

func StringJoin(words []string) string {
    var outputString string;

    for _, word := range words {
        outputString = outputString + (word + " ");
    }

    return outputString;
}

func StringJoinWithStringsFunction(words []string, separator string) string {
    return strings.Join(words, separator)
}