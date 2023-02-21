package djikstra

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

func TestMeuOvo(t *testing.T) {

	valor, _ := decimal.NewFromString("10.00")

	div, _ := decimal.NewFromString("3")

	divResult := valor.Div(div)


	var no1 byte = 'A'
	//var no2 byte = 'B'

	fmt.Printf("%c\n", no1)


	if otherValue := divResult.Mul(div); otherValue == valor {
		t.Errorf("Division not precise. Result %s where expected was %s.", otherValue, valor)
	}

}
