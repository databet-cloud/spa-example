package apierror

import (
	"fmt"
	"testing"
)

func TestE(t *testing.T) {
	e := Error{
		error: fmt.Errorf("hello world"),
		code:  "code",
		level: "level",
		data:  map[string]any{"1": 2},
	}

	fmt.Println(e.String())
}
