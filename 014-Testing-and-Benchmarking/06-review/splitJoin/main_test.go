package splitJoin

import (
    "testing"
    "strings"
    "fmt"
)

const initialSentence = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Risus in hendrerit gravida rutrum. Turpis massa sed elementum tempus egestas sed. Enim ut sem viverra aliquet eget sit. Justo nec ultrices dui sapien. Eu sem integer vitae justo. Consectetur libero id faucibus nisl tincidunt eget. Condimentum vitae sapien pellentesque habitant morbi. Vel turpis nunc eget lorem dolor sed viverra. Consectetur a erat nam at lectus urna duis convallis convallis. Quam vulputate dignissim suspendisse in est ante."

func TestStringJoin(t *testing.T) {
    sentence1 := "Here I come!!!";
    wordsList := strings.Split(sentence1, " ");
    if StringJoin(wordsList, " ") != sentence1 {
        t.Errorf("Expected: %v, but got %v", sentence1, StringJoin(wordsList, " "))
    }
}

func TestStringJoin_another_check(t *testing.T) {
    wordsList := strings.Split(initialSentence, " ");
    if StringJoin(wordsList, " ") != initialSentence {
        t.Errorf("Expected: %v, but got %v", initialSentence, StringJoin(wordsList, " "))
    }
}


func TestStringJoinWithBuiltinFunction(t *testing.T) {
    sentence1 := "Here I come!!!";
    wordsList := strings.Split(sentence1, " ");
    if StringJoinWithBuiltinFunction(wordsList, " ") != sentence1 {
        t.Errorf("Expected: %v, but got %v", sentence1, StringJoinWithBuiltinFunction(wordsList, " "))
    }
}

func TestStringJoinWithBuiltinFunction_another_check(t *testing.T) {
    wordsList := strings.Split(initialSentence, " ");
    if StringJoinWithBuiltinFunction(wordsList, " ") != initialSentence {
        t.Errorf("Expected: %v, but got %v", initialSentence, StringJoinWithBuiltinFunction(wordsList, " "))
    }
}

func ExampleStringJoin() {
    sentence1 := "Here I come!!!";
    wordsList := strings.Split(sentence1, " ");
    outputString := StringJoin(wordsList, " ");
    fmt.Println(outputString);
    // Output:
    // Here I come!!! 
}

func ExampleStringJoin_another_example() {
    wordsList := strings.Split(initialSentence, " ");
    outputString := StringJoin(wordsList, " ");
    fmt.Println(outputString);
    // Output:
    // Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Risus in hendrerit gravida rutrum. Turpis massa sed elementum tempus egestas sed. Enim ut sem viverra aliquet eget sit. Justo nec ultrices dui sapien. Eu sem integer vitae justo. Consectetur libero id faucibus nisl tincidunt eget. Condimentum vitae sapien pellentesque habitant morbi. Vel turpis nunc eget lorem dolor sed viverra. Consectetur a erat nam at lectus urna duis convallis convallis. Quam vulputate dignissim suspendisse in est ante.
}


func ExampleStringJoinWithBuiltinFunction() {
    sentence1 := "Here I come!!!";
    wordsList := strings.Split(sentence1, " ");
    outputString := StringJoinWithBuiltinFunction(wordsList, " ");
    fmt.Println(outputString);
    // Output:
    // Here I come!!!
}

func ExampleStringJoinWithBuiltinFunction_another_example() {
    wordsList := strings.Split(initialSentence, " ");
    outputString := StringJoinWithBuiltinFunction(wordsList, " ");
    fmt.Println(outputString);
    // Output:
    // Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Risus in hendrerit gravida rutrum. Turpis massa sed elementum tempus egestas sed. Enim ut sem viverra aliquet eget sit. Justo nec ultrices dui sapien. Eu sem integer vitae justo. Consectetur libero id faucibus nisl tincidunt eget. Condimentum vitae sapien pellentesque habitant morbi. Vel turpis nunc eget lorem dolor sed viverra. Consectetur a erat nam at lectus urna duis convallis convallis. Quam vulputate dignissim suspendisse in est ante.
}

func BenchmarkStringJoin(b *testing.B) {
    sentence1 := "Here I come!!!";
    wordsList := strings.Split(sentence1, " ");
    for i := 0; i < b.N; i ++ {
        StringJoin(wordsList, " ");
    }
}

func BenchmarkStringJoin_another_example(b *testing.B) {
    wordsList := strings.Split(initialSentence, " ");
    for i := 0; i < b.N; i ++ {
        StringJoin(wordsList, " ");
    }
}

func BenchmarkStringJoinWithBuiltinFunction(b *testing.B) {
    sentence1 := "Here I come!!!";
    wordsList := strings.Split(sentence1, " ");    
    for i := 0; i < b.N; i ++ {
       StringJoinWithBuiltinFunction(wordsList, " "); 
    }
}

func BenchmarkStringJoinWithBuiltinFunction_another_example(b *testing.B) {
    wordsList := strings.Split(initialSentence, " ");    
    for i := 0; i < b.N; i ++ {
       StringJoinWithBuiltinFunction(wordsList, " "); 
    }
}