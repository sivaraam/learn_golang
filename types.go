package main
import "fmt"

func main() {
  var i int = 10;
  var i8_negative int8 = -128; // 8 bit integer
  var i8_positive int8 = 127; // 8 bit integer
  var ch int = 'a';
  var str string = "Hello, world!"

  fmt.Println("integer i = ", i);
  fmt.Println("integer (8 bit, low) = ", i8_negative);
  fmt.Println("integer (8 bit, high) = ", i8_positive);
  fmt.Println("char ch = ", ch);
  fmt.Println("string str = ", str);
}
