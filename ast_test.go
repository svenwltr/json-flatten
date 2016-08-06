package main

import (
	"fmt"
	"testing"
)

func assertEquals(t *testing.T, expect string, object interface{}) {
	obtained := fmt.Sprintf("%v", object)
	if obtained != expect {
		t.Errorf("Assertion failed.")
		t.Errorf("\tInput:     %#v", object)
		t.Errorf("\tGenerated: %#v", obtained)
		t.Errorf("\tExpexted:  %#v", expect)
	}
}

func TestStrings(t *testing.T) {
	assertEquals(t, "a=b",
		Assignment{Path{ObjectName("a")}, "b"})

	assertEquals(t, ".=a",
		Assignment{Path{}, "a"})

	assertEquals(t, "a.b=c",
		Assignment{[]Name{ObjectName("a"), ObjectName("b")}, "c"})

	assertEquals(t, "a=b\nc=d",
		Assignments{
			&Assignment{Path{ObjectName("a")}, "b"},
			&Assignment{Path{ObjectName("c")}, "d"},
		})

	assertEquals(t, "a=b\nc=d",
		Assignments{
			&Assignment{Path{ObjectName("a")}, "b"},
			&Assignment{Path{ObjectName("c")}, "d"},
		})

	assertEquals(t, "a[0]=b\na[1]=b",
		Assignments{
			&Assignment{Path{ObjectName("a"), ArrayIndex(0)}, "b"},
			&Assignment{Path{ObjectName("a"), ArrayIndex(1)}, "b"},
		})
	assertEquals(t, `a.b\.c=d`,
		Assignment{[]Name{ObjectName("a"), ObjectName("b.c")}, "d"})
	assertEquals(t, `a.b\\c=d`,
		Assignment{[]Name{ObjectName("a"), ObjectName("b\\c")}, "d"})

}
