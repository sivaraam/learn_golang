package main
import "fmt"

type MyClass struct {
	  name string
}

func (self MyClass) String() string {
	return "Name: " + self.name
}

func (self MyClass) ICannotModifyMe(selfagain MyClass) {
  self.name= "Unmodifiable name"
  fmt.Println(self, "; ", selfagain)
  selfagain.name = "Unmodifiable selfagain name"
  fmt.Println(self, "; ", selfagain)
}

func (self *MyClass) ICanModifyMe(selfagain *MyClass) {
  self.name= "Unmodifiable name"
  fmt.Println(self, "; ", selfagain)
  selfagain.name = "Unmodifiable selfagain name"
  fmt.Println(self, "; ", selfagain)
}

func main() {
	  myobj0 := &MyClass{}
    myobj0.name = "Main name"
    myobj0.ICannotModifyMe(*myobj0)
    fmt.Println("myobj0: ", myobj0)
    myobj0.ICanModifyMe(myobj0)
    fmt.Println("myobj0: ", myobj0)
}
