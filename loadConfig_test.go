package go_util

import "testing"

func TestCase001(t *testing.T) {

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			ret := addInt(i,j)
			if ret != (i+j) {
				t.Errorf("addInt(%d,%d)",i,j)
			}
		}
	} 

}

