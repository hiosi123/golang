package testing

import (
	"os"
	"testing"
	"time"
)

func TestAbs(t *testing.T) {
	if Abs(-1) < 0 {
		t.Error("negative value found in abs() with", -1)
	}
	if Abs(0) < 0 {
		t.Error("Negative value foujnd in abs() with", 0)
	}
	if Abs(1) < 0 {
		t.Error("Negative value foujnd in abs() with", 1)
	}
}

func TestAbsSUb(t *testing.T) {
	t.Run("Positive", func(t *testing.T) {
		if Abs(1) < 0 {
			t.Error("negative value found in abs()")

		}
	})
}

func TestSkip(t *testing.T) {
	if len(os.Getenv("GOPATH")) == 0 {
		t.Skip("Skpping test because GOPATH isn't set")
	}
	t.Log("testing with GOPATH:", os.Getenv("GOPATH"))
}

func TestCleanup(t *testing.T) {
	t.Cleanup(func() {
		t.Log("cleanup")
	}) // cuncurrently safe
	t.Log("Running some test")
}

func TestParallel(t *testing.T) {
	t.Parallel()
}

func TestParallelOne(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
}
func TestParallelTwo(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
}
func TestParallelThree(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
}

// t.Log("Similar to fmt.Println() and concurrently safe")
// t.Fail()    //will show a test case has failed in the results
// t.FailNow() // t.Fail() + safely exit without continuing
// t.Error("t.Log() + t.Fail()")
// t.Fatal("t.Log() + t.FailNow()")
