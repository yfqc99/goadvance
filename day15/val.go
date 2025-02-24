package main
import (
	"fmt"
	"reflect"
)
func main() {
	var f float32
	f=6.6
	rval := reflect.ValueOf(f)
	// f = rval.Float()
	/* cannot use rval.Float() 
	(value of type float64) as float32 value in assignment */
	f = float32(rval.Float())
	fmt.Println(f)
}