package error

import (
	"strconv"
)

// func main() {
// 	err := returnError()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

type specialError struct {
	errorMessage string
	errorCode    int
}

func (s specialError) Error() string {
	return s.errorMessage + ":" + strconv.Itoa(s.errorCode)
}

func returnError(returnError bool) (string, error) {
	if returnError {
		return "", specialError{"specialError", 123}
	}
	return "good", nil
}

func errorHandling()
