package kata

import (
	"reflect"
	"sort"
	"testing"
)

func TestWordFreq(t *testing.T) {
	got := WordFreq("Hi, hi!! go-go?")
	want := map[string]int{"hi": 2, "go": 2}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got=%v want=%v", got, want)
	}
}

func TestUniqueTrimmed(t *testing.T) {
	got := UniqueTrimmed([]string{" a ", "", "b", "a", "  b  ", "c"})
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got=%v want=%v", got, want)
	}
}

func TestGroupAnagrams(t *testing.T) {
	got := GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})

	// normalize: sort each group and compare keys we expect exist
	for k := range got {
		sort.Strings(got[k])
	}

	if _, ok := got["aet"]; !ok || !reflect.DeepEqual(got["aet"], []string{"ate", "eat", "tea"}) {
		t.Fatalf("expected key aet => [ate eat tea], got=%v", got["aet"])
	}
	if _, ok := got["ant"]; !ok || !reflect.DeepEqual(got["ant"], []string{"nat", "tan"}) {
		t.Fatalf("expected key ant => [nat tan], got=%v", got["ant"])
	}
	if _, ok := got["abt"]; !ok || !reflect.DeepEqual(got["abt"], []string{"bat"}) {
		t.Fatalf("expected key abt => [bat], got=%v", got["abt"])
	}
}

func TestRLE(t *testing.T) {
	got := RLE("aaabbc")
	want := "a3b2c"
	if got != want {
		t.Fatalf("got=%q want=%q", got, want)
	}
}

func TestTwoSum(t *testing.T) {
	i, j, ok := TwoSum([]int{2, 7, 11, 15}, 9)
	if !ok {
		t.Fatalf("expected ok=true")
	}
	if !(i == 0 && j == 1) {
		t.Fatalf("got indices (%d,%d), want (0,1)", i, j)
	}
}

func TestMergeIntervals(t *testing.T) {
	got := MergeIntervals([]Interval{{1, 3}, {2, 6}, {8, 10}, {9, 12}})
	want := []Interval{{1, 6}, {8, 12}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got=%v want=%v", got, want)
	}
}

func TestNetBalances(t *testing.T) {
	got := NetBalances([]string{"alice->bob:30", "bob->carl:10", "alice->carl:5"})
	want := map[string]int{"alice": -35, "bob": 20, "carl": 15}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got=%v want=%v", got, want)
	}
}

func TestTopK(t *testing.T) {
	got := TopK([]int{1, 1, 1, 2, 2, 3, 3, 3, 3}, 2)
	want := []int{3, 1}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got=%v want=%v", got, want)
	}
}

func TestIsBalanced(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"([{}])", true},
		{"(]", false},
		{"a+(b*[c])", true},
		{"(", false},
		{"", true},
	}
	for _, tc := range cases {
		if got := IsBalanced(tc.in); got != tc.want {
			t.Fatalf("IsBalanced(%q)=%v want=%v", tc.in, got, tc.want)
		}
	}
}

// //
// // Harder 11) LRU cache tests
// //

// func TestLRU(t *testing.T) {
// 	c, err := NewLRU(2)
// 	if err != nil {
// 		t.Fatalf("unexpected err: %v", err)
// 	}

// 	c.Put("a", 1)
// 	c.Put("b", 2)

// 	if v, ok := c.Get("a"); !ok || v != 1 {
// 		t.Fatalf("expected Get(a)=1,true got=%v,%v", v, ok)
// 	}

// 	// a is now most recent; inserting c should evict b
// 	c.Put("c", 3)

// 	if _, ok := c.Get("b"); ok {
// 		t.Fatalf("expected b to be evicted")
// 	}
// 	if v, ok := c.Get("a"); !ok || v != 1 {
// 		t.Fatalf("expected a to remain")
// 	}
// 	if v, ok := c.Get("c"); !ok || v != 3 {
// 		t.Fatalf("expected c to exist")
// 	}
// 	if c.Len() != 2 {
// 		t.Fatalf("expected len=2 got=%d", c.Len())
// 	}
// }

// //
// // Harder 12) Tokenizer tests
// //

// func TestTokenize(t *testing.T) {
// 	in := `sum(x1, 12.5, "hi\nthere") + y_2`
// 	got, err := Tokenize(in)
// 	if err != nil {
// 		t.Fatalf("unexpected err: %v", err)
// 	}

// 	// Spot-check a few tokens (not all) to keep test robust
// 	if len(got) < 8 {
// 		t.Fatalf("expected >= 8 tokens, got=%d", len(got))
// 	}
// 	if got[0].Type != TokIdent || got[0].Text != "sum" {
// 		t.Fatalf("expected first token ident sum, got=%v", got[0])
// 	}
// 	// string token should contain an actual newline due to \n escape
// 	foundString := false
// 	for _, tk := range got {
// 		if tk.Type == TokString {
// 			foundString = true
// 			if tk.Text != "hi\nthere" {
// 				t.Fatalf("expected string token hi\\nthere => actual newline, got=%q", tk.Text)
// 			}
// 		}
// 	}
// 	if !foundString {
// 		t.Fatalf("expected a string token")
// 	}
// }

// //
// // Harder 13) JSON-lite parser tests
// //

// func TestParseJSONLite(t *testing.T) {
// 	in := `{"a": 1, "b": [true, null, "x"], "c": {"d": -2.5}}`
// 	got, err := ParseJSONLite(in)
// 	if err != nil {
// 		t.Fatalf("unexpected err: %v", err)
// 	}

// 	obj, ok := got.(map[string]any)
// 	if !ok {
// 		t.Fatalf("expected object, got %T", got)
// 	}

// 	if obj["a"].(float64) != 1 {
// 		t.Fatalf("expected a=1 got=%v", obj["a"])
// 	}

// 	barr := obj["b"].([]any)
// 	if barr[0].(bool) != true {
// 		t.Fatalf("expected b[0]=true")
// 	}
// 	if barr[1] != nil {
// 		t.Fatalf("expected b[1]=nil got=%v", barr[1])
// 	}
// 	if barr[2].(string) != "x" {
// 		t.Fatalf("expected b[2]=x")
// 	}

// 	cobj := obj["c"].(map[string]any)
// 	if cobj["d"].(float64) != -2.5 {
// 		t.Fatalf("expected c.d=-2.5 got=%v", cobj["d"])
// 	}
// }
