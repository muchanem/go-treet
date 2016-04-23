package treet_test

import (
	//"fmt"
	"github.com/fadb/go-treet"
)

func ExamplePrint() {
	treeish := map[string]interface{}{
		"blah": map[string]string{"another": "finally"},
	}
	treet.Print(treeish)

	// Output:
	// blah
	// └── another
	//     └── finally
}
