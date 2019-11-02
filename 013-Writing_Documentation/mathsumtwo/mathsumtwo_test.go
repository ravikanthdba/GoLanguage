package mathsumtwo

import (
    "testing"
)

func Test_MathSumTwo(t *testing.T) {
    var variables = []int{10,20,30,40,50};
    var sum int;

    sum = MathSumTwo(variables);
    if sum != 150 {
        t.Errorf("Expected is 150 and the sum after the function is: %d", sum);
    }

    variables = []int{100,200,300,400,500};
    sum = 0;
    sum = MathSumTwo(variables);
    if sum != 1500 {
        t.Errorf("Expected is 1500 and the sum after the function is: %d", sum);
    }
    
}
