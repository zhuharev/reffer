package reffer

import "testing"

func TestGetleveledRefferals(t *testing.T) {
	obj1 := &defaultObject{id: "1"}
	obj2 := &defaultObject{id: "2"}

	ds := &defaultStore{}
	ds.Set(obj1)
	ds.Set(obj2)

	_, err := ds.GetRefferer(obj1, 1)
	if err != ErrNotExists {
		t.Fatalf("expected error, got %s", err)
	}

	ds.SetRefferer(obj1, obj2)
	refferer, err := ds.GetRefferer(obj1, 1)
	if err != nil {
		t.Fatalf("expected error, got %s", err)
	}
	if refferer.ID() != obj2.ID() {
		t.Fatalf("expected error, got %s", err)
	}

	obj3 := &defaultObject{id: "3"}

	ds.SetRefferer(obj2, obj3)
	refferer, err = ds.GetRefferer(obj2, 1)
	if err != nil {
		t.Fatalf("expected error, got %s", err)
	}
	if refferer.ID() != obj3.ID() {
		t.Fatalf("expected error, got %s", err)
	}

	refferer, err = ds.GetRefferer(obj1, 2)
	if err != nil {
		t.Fatalf("expected error, got %s", err)
	}
	if refferer.ID() != obj3.ID() {
		t.Fatalf("not expected refferer %s", err)
	}

	cnt, err := ds.RefferralsCount(obj3, 2)
	if err != nil {
		t.Fatalf("r cnt error %s", err)
	}
	if cnt != 2 {
		t.Fatalf("not expected refferer count %d", cnt)
	}

}
