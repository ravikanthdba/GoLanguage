package dog

import (
    "testing"
)

func TestYears(t *testing.T) {
    var human_years int = 40;
    var dog_years int;

    dog_years = human_years * 7;
    if dog_years != 280 {
        t.Errorf("The test is not right, expected and the returned output is not correct. Passed human_years: %d, Expected dog_years: 280, Returned:%d", human_years, dog_years);
    }
}

