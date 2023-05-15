package makeString

import (
	"fmt"
	"strconv"
	"strings"
)

func MakeString() {
	s := "hello World"

	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(s[0])        // ascii value of the character
	fmt.Printf("%c\n", s[0]) // this should print out the h
	fmt.Println(s[0:6])
	fmt.Println(s[:6])
	fmt.Println(s[6:])

	s = s + " Again"
	s += " Again"

	fmt.Println(s)

	fmt.Println("Hello \nWorld")
	fmt.Println("Hello \tWorld")
	fmt.Println("Hello \bWorld")

	abc := "abc"
	b := []byte(abc)

	fmt.Printf("%s, %s", abc, b)
	abc2 := string(b)
	fmt.Println(abc2)

	fmt.Println("안녕")
	fmt.Println("안녕"[0])
	fmt.Println(len("안녕"))

	fmt.Println(strings.ToUpper("hello World"))
	fmt.Println(strings.ToLower("HELLo World"))

	fmt.Println(strings.HasPrefix("Hello world", "Hello"))
	fmt.Println(strings.HasSuffix("Hello world", "world"))

	fmt.Println(strings.Replace("Hellow world world", "world", "Hello", 5))
}

func StringBuilder() {
	var sb strings.Builder
	fmt.Println("this is the string builder:", sb.String())
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())

	sb.WriteString("Hellooooo")
	fmt.Println("this is the string builder:", sb.String())
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())

	sb.Grow(100)
	fmt.Println(sb.Cap())

	sb.Reset()

	fmt.Println(sb.String())

	sb.Write([]byte{65, 65, 65})
	fmt.Println(sb.String())

	x := 123
	y := strconv.Itoa(x)

	fmt.Println(y)
	fmt.Printf("%T\n", y)

	z, _ := strconv.Atoi(y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
