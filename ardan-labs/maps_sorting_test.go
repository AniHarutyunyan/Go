package ardanlabs

import (
	"fmt"
	"sort"
	"testing"
)

func TestSortingAlpha(t *testing.T){
	type user struct{
		name string
		surname string
	}
	users := map[string]user{

		"R": {"Roy","Royson"},
		"B": {"Poll","Royson"},
		"D": {"Ann","Royson"},
		"A": {"Gab","Royson"},
	}
	var keys []string
	for key := range users{
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _,key := range keys{
		fmt.Println(key, users[key])
	}
}