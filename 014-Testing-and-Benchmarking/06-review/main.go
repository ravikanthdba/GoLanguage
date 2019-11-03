package main

import (
    "fmt"
    "strings"
    "github.com/ravikanthdba/GoLanguage/014-Testing-and-Benchmarking/06-review/splitJoin"
)

const sentence = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Risus in hendrerit gravida rutrum. Turpis massa sed elementum tempus egestas sed. Enim ut sem viverra aliquet eget sit. Justo nec ultrices dui sapien. Eu sem integer vitae justo. Consectetur libero id faucibus nisl tincidunt eget. Condimentum vitae sapien pellentesque habitant morbi. Vel turpis nunc eget lorem dolor sed viverra. Consectetur a erat nam at lectus urna duis convallis convallis. Quam vulputate dignissim suspendisse in est ante."
const anotherSentence = "Here I come!!!"

func main() {
    wordsList := strings.Split(sentence, " ");
    fmt.Println("After Splitting:");
    fmt.Println(wordsList);
    fmt.Println();
    fmt.Println("After Joining with StringJoin function:");
    fmt.Println(splitJoin.StringJoin(wordsList, " "));
    fmt.Println();
    fmt.Println("After Joining with strings.Join function:");
    fmt.Println(splitJoin.StringJoinWithBuiltinFunction(wordsList, " "));



    wordsList = strings.Split(anotherSentence, " ");
    fmt.Println("After Splitting:");
    fmt.Println(wordsList);
    fmt.Println();
    fmt.Println("After Joining with StringJoin function:");
    fmt.Println(splitJoin.StringJoin(wordsList, " "));
    fmt.Println();
    fmt.Println("After Joining with strings.Join function:");
    fmt.Println(splitJoin.StringJoinWithBuiltinFunction(wordsList, " "));    
}
