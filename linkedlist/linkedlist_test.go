package linkedlist

import (
	"bytes"
	"os"
	"testing"
)

func TestSize(t *testing.T) {
	l := New()
	if l.Size() != 0 {
		t.Errorf("expected size 0, got %d", l.Size())
	}
	l.Append(12)
	if l.Size() != 1 {
		t.Errorf("expected size 1, got %d", l.Size())
	}
	for i := 0; i < 300; i++ {
		l.Append(i)
	}
	if l.Size() != 301 {
		t.Errorf("expected size 301, got %d", l.Size())
	}
}

func TestPrint(t *testing.T) {
	oldStdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	l := New()
	l.Append(5)
	l.Append(6)
	l.Append(7)
	l.Append(8)
	l.Print()

	w.Close()
	os.Stdout = oldStdOut
	var buf bytes.Buffer
	buf.ReadFrom(r)

	got := buf.String()
	expected := "Linked List: {5, 6, 7, 8}\n"
	if got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

func TestAppend(t *testing.T) {
	new := New()
	new.Append(5)
	if new.Size() != 1 {
		t.Errorf("expected 1, got %d", new.Size())
	}
	new.Append(1)
	new.Append(1)
	new.Append(1)
	new.Append(1)
	new.Append(1)
	new.Append(1)
	if new.Size() != 7 {
		t.Errorf("expected 7, got %d", new.Size())
	}
}

func TestDeleteNode(t *testing.T) {
	new := New()

	new.Append(5)
	new.Append(7)
	new.Append(13)
	got := new.DeleteNode(3)
	if got.Data != 13 {
		t.Errorf("expected 13, got %d", got.Data)
	}
}
