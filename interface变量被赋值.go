type Speaker interface {
	Hello()
}

type User struct {
	Name string
	Age  int
}

func (this *User) Hello() {
	fmt.Println("hello my name is", this.Name)
}
func FuncT() Speaker {

	return &User{"wss", 10}
}
func main() {

	s := FuncT()
	s.Hello()

}