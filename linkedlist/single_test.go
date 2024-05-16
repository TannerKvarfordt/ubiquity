package linkedlist_test

import (
	"testing"

	"github.com/TannerKvarfordt/ubiquity/linkedlist"
)

func TestDefaults(t *testing.T) {
	l := linkedlist.SinglyLinkedList[int]{}
	if l.Len() != 0 {
		t.Error()
	}
}

func TestAppend(t *testing.T) {
	l := linkedlist.SinglyLinkedList[int]{}
	l.Append(0)
	if l.Len() != 1 {
		t.Error()
	}
	l.ForEach(func(i int) {
		if i != 0 {
			t.Error()
		}
	})

	l.Append(1)
	if l.Len() != 2 {
		t.Error()
	}
	itr := 0
	l.ForEach(func(i int) {
		if itr != i {
			t.Error()
		}
		itr++
	})
}

func TestPrepend(t *testing.T) {
	l := linkedlist.SinglyLinkedList[int]{}
	l.Prepend(2)
	if l.Len() != 1 {
		t.Error()
	}
	l.ForEach(func(i int) {
		if i != 2 {
			t.Error()
		}
	})

	l.Prepend(1)
	l.Prepend(0)
	if l.Len() != 3 {
		t.Error()
	}
	itr := 0
	l.ForEach(func(i int) {
		if itr != i {
			t.Error()
		}
		itr++
	})
}

func TestGetData(t *testing.T) {
	l := &linkedlist.SinglyLinkedList[int]{}
	_, err := l.GetData(0)
	if err == nil {
		t.Error("Expected non-nil error")
	}

	l = linkedlist.NewSinglyLinkedList(0, 1, 2)
	for i := 0; i < l.Len(); i++ {
		val, err := l.GetData(i)
		if err != nil {
			t.Error(err)
		}
		if i != val {
			t.Error()
		}
	}
}

func TestSetData(t *testing.T) {
	l := linkedlist.SinglyLinkedList[string]{}
	if err := l.SetData(0, "foo"); err == nil {
		t.Error()
		return
	}
	l.Append("bar")
	if err := l.SetData(0, "foo"); err != nil {
		t.Error(err)
		return
	}
	if data, err := l.GetData(0); err != nil {
		t.Error(err)
		return
	} else if data != "foo" {
		t.Errorf("expected: foo, actual: %s", data)
	}
}

func TestInsert(t *testing.T) {
	l := &linkedlist.SinglyLinkedList[string]{}
	if err := l.Insert(-1, "b"); err == nil {
		t.Error("Expected error")
		return
	}
	if err := l.Insert(1, "b"); err == nil {
		t.Error("Expected error")
		return
	}

	if err := l.Insert(0, "b"); err != nil {
		t.Error(err)
		return
	}
	if val, _ := l.GetData(0); val != "b" {
		t.Error()
		return
	}

	if err := l.Insert(1, "d"); err != nil {
		t.Error(err)
		return
	}
	if val, _ := l.GetData(0); val != "b" {
		t.Error()
		return
	}
	if val, _ := l.GetData(1); val != "d" {
		t.Error()
		return
	}

	if err := l.Insert(0, "a"); err != nil {
		t.Error(err)
		return
	}
	if val, _ := l.GetData(0); val != "a" {
		t.Error()
		return
	}
	if val, _ := l.GetData(1); val != "b" {
		t.Error()
		return
	}
	if val, _ := l.GetData(2); val != "d" {
		t.Error()
		return
	}

	if err := l.Insert(2, "c"); err != nil {
		t.Error(err)
		return
	}
	if val, _ := l.GetData(0); val != "a" {
		t.Error()
		return
	}
	if val, _ := l.GetData(1); val != "b" {
		t.Error()
		return
	}
	if val, _ := l.GetData(2); val != "c" {
		t.Error()
		return
	}
	if val, _ := l.GetData(3); val != "d" {
		t.Error()
		return
	}

	if err := l.Insert(4, "e"); err != nil {
		t.Error(err)
		return
	}
	if val, _ := l.GetData(0); val != "a" {
		t.Error()
		return
	}
	if val, _ := l.GetData(1); val != "b" {
		t.Error()
		return
	}
	if val, _ := l.GetData(2); val != "c" {
		t.Error()
		return
	}
	if val, _ := l.GetData(3); val != "d" {
		t.Error()
		return
	}
	if val, _ := l.GetData(4); val != "e" {
		t.Error()
		return
	}
}

func TestDelete(t *testing.T) {
	l := linkedlist.NewSinglyLinkedList(0, 1, 2, 3, 4, 5)
	if _, err := l.Delete(-1); err == nil {
		t.Error()
	}
	if _, err := l.Delete(l.Len()); err == nil {
		t.Error()
	}
	if _, err := l.Delete(l.Len() + 1); err == nil {
		t.Error()
	}

	if v, err := l.Delete(0); err != nil {
		t.Error(err)
	} else {
		if v != 0 {
			t.Error()
		}
		if l.Len() != 5 {
			t.Error()
		}
		if d, err := l.GetData(0); err != nil {
			t.Error(err)
		} else if d != 1 {
			t.Error()
		}
	}

	if v, err := l.Delete(4); err != nil {
		t.Error(err)
	} else {
		if v != 5 {
			t.Error()
		}
		if l.Len() != 4 {
			t.Error()
		}
		if d, err := l.GetData(3); err != nil {
			t.Error(err)
		} else if d != 4 {
			t.Error()
		}
	}

	if v, err := l.Delete(1); err != nil {
		t.Error(err)
	} else {
		if v != 2 {
			t.Error()
		}
		if l.Len() != 3 {
			t.Error()
		}
		if d, err := l.GetData(1); err != nil {
			t.Error(err)
		} else if d != 3 {
			t.Error()
		}
	}
}
