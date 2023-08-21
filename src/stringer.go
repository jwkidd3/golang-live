package main
 
import "fmt"

// START OMIT
type Person struct {
    Name, AlmaMater string
}

func (p Person) String() string {
    return fmt.Sprintf("%v (Alma Mater: %v)", p.Name, p.AlmaMater)
}

func main() {
    rs := Person{"The Hamburglar", "McDonalds University"}
    ab := Person{"Colonel Sanders", "KFC Technical College"}
    fmt.Printf("%s\n%s\n", rs, ab)
}
// END OMIT
