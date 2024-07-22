package main
type request struct {

	Password strin
	Address string
	Email string
	Username string

}
 func main() {

	// good function a fucntion that contaiants maximum 3 arguments

	// func validate(usnername, email, password, address string) bool {

	// 	return true
	

	func validate(request Request) bool {
		inValid := request.email != "" && strings.Contains(request.Email, "@") && request.Password != && 
		return isValid
	}
	

 }