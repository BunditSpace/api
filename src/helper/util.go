package helper

//Check check with parameter
func Check(err error) {
	if err != nil {
		panic(err)
	}
}
