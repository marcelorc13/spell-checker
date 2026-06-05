package trie

import "testing"

func TestInsertSearch(t *testing.T) {
	tr := New()
	tr.Insert("hello", 1500)
	tr.Insert("help", 1200)
	tr.Insert("world", 1000)

	freq, found := tr.Search("hello")
	if !found || freq != 1500 {
		t.Fatalf("Search(hello): got freq=%d found=%v, want 1500 true", freq, found)
	}

	freq, found = tr.Search("help")
	if !found || freq != 1200 {
		t.Fatalf("Search(help): got freq=%d found=%v, want 1200 true", freq, found)
	}

	_, found = tr.Search("hell")
	if found {
		t.Fatal("Search(hell): expected not found")
	}

	_, found = tr.Search("xyz")
	if found {
		t.Fatal("Search(xyz): expected not found")
	}
}

func TestInsertEmpty(t *testing.T) {
	tr := New()
	tr.Insert("", 100)
	freq, found := tr.Search("")
	if !found || freq != 100 {
		t.Fatalf("Search(''): got freq=%d found=%v, want 100 true", freq, found)
	}
}
