package calculator

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
)

type ResponseDTO[T any] struct {
	code    string
	message string
	data    T
}

type User struct {
	name string
}

func TestAdd(t *testing.T) {
	expected := 15

	actual := Add[int](5, 5, 5)

	if expected != actual {
		t.Errorf("Get %d but expected %d", actual, expected)
	}
}

func TestPrint(t *testing.T) {
	expected := "1 is one"

	actual := Print[int, string](1, "one")

	if expected != actual {
		t.Errorf("Get %v but expected %v", actual, expected)
	}
}

func TestGenericClass(t *testing.T) {
	user := User{name: "cris"}
	responseDTO := ResponseDTO[User]{code: "00", message: "Success", data: user}

	fmt.Println(responseDTO)
}

func Test(t *testing.T) {
	// Automatically read environment variables
	viper.AutomaticEnv()

	// Accessing an environment variable
	value := viper.GetString("HOME")
	fmt.Printf("HOME = %s\n", value)
}
