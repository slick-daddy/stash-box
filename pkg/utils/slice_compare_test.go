// nolint: revive
package utils

import (
	"reflect"
	"testing"
)

type strSliceTest struct {
	subject         []string
	against         []string
	expectedAdded   []string
	expectedMissing []string
}

var strSliceTests = []strSliceTest{
	{[]string{"a"}, []string{"b"}, []string{"a"}, []string{"b"}},
	{[]string{"a", "b"}, []string{"b"}, []string{"a"}, nil},
	{[]string{"a"}, []string{"a", "b"}, nil, []string{"b"}},
	{[]string{"a"}, nil, []string{"a"}, nil},
	{[]string{"a"}, []string{}, []string{"a"}, nil},
	{nil, []string{"b"}, nil, []string{"b"}},
	{[]string{}, []string{"b"}, nil, []string{"b"}},
	{[]string{"a", "a"}, []string{"b", "b"}, []string{"a"}, []string{"b"}},
	{[]string{"a"}, []string{"a"}, nil, nil},
}

func TestStrSliceCompare(t *testing.T) {
	for _, test := range strSliceTests {
		added, missing := SliceCompare(test.subject, test.against)

		if !reflect.DeepEqual(added, test.expectedAdded) {
			t.Errorf("added(%v,%v) = %v; want %v", test.subject, test.against, added, test.expectedAdded)
		}

		if !reflect.DeepEqual(missing, test.expectedMissing) {
			t.Errorf("missing(%v,%v) = %v; want %v", test.subject, test.against, missing, test.expectedMissing)
		}
	}
}

func BenchmarkSliceCompare(b *testing.B) {
	subject := make([]string, 0, 1000)
	against := make([]string, 0, 1000)

	for i := 0; i < 1000; i++ {
		value := string(rune('a'+(i%26))) + string(rune('a'+((i/26)%26))) + string(rune('a'+((i/676)%26)))
		subject = append(subject, value)
	}
	for i := 500; i < 1500; i++ {
		value := string(rune('a'+(i%26))) + string(rune('a'+((i/26)%26))) + string(rune('a'+((i/676)%26)))
		against = append(against, value)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SliceCompare(subject, against)
	}
}
