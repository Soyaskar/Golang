package auth
import( "fmt")

func Login(name string)string{
	res:=fmt.Sprintf("Successfully logged in by %s",name)
	return res
}