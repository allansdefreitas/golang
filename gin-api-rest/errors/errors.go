package errors

func HandleError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
