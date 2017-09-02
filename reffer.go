package reffer

type Reffer struct {
}

type IDer interface {
	ID() string
}

type Object interface {
	IDer
	Refferer() IDer
	Refferrals() []IDer

	SetRefferer(Object)
	AddRefferal(refferal Object) error
}

type Store interface {
	Get(id string) (Object, error)
	Set(Object) error

	GetRefferer(Object, level int) (Object, error)
	SetRefferer(o Object, refferer Object) error

	RefferralsCount(o Object, level int) (int, error)
}

type defaultObject struct {
	id         string
	refferer   string
	refferrals []string
}

func (do defaultObject) ID() string {
	return do.id
}

func (do defaultObject) Refferer() IDer {
	return defaultObject{id: do.refferer}
}

func (do defaultObject) Refferrals() (res []IDer) {
	for _, v := range do.refferrals {
		res = append(res, defaultObject{id: v})
	}
	return
}

func (do *defaultObject) AddRefferal(o Object) error {
	do.refferrals = append(do.refferrals, o.ID())
	return nil
}

func (do *defaultObject) SetRefferer(o Object) {
	do.refferer = o.ID()
	return
}

type defaultStore struct {
	objects []Object
}

func (ds *defaultStore) Get(id string) (Object, error) {
	for _, v := range ds.objects {
		if v.ID() == id {
			return v, nil
		}
	}
	return nil, ErrNotExists
}

func (ds *defaultStore) Set(o Object) error {
	for i, v := range ds.objects {
		if v.ID() == o.ID() {
			ds.objects[i] = o
			return nil
		}
	}
	ds.objects = append(ds.objects, o)
	return nil
}

func (ds *defaultStore) RefferralsCount(o Object, level int) (count int, err error) {
	count += len(o.Refferrals())
	for _, v := range o.Refferrals() {
		var oro Object
		oro, err = ds.Get(v.ID())
		if err != nil {
			return
		}
		var cnt int
		cnt, err = ds.RefferralsCount(oro, level-1)
		if err != nil {
			return
		}
		count += cnt
	}
	return
}

func (ds *defaultStore) GetRefferer(o Object, level int) (reso Object, err error) {
	id := o.Refferer().ID()
	for i := 0; i != level; i++ {
		reso, err = ds.Get(id)
		if err != nil {
			return
		}
		id = reso.Refferer().ID()
	}
	return
}

func (ds *defaultStore) SetRefferer(o Object, ref Object) (err error) {
	o.SetRefferer(ref)
	ref.AddRefferal(o)
	err = ds.Set(o)
	if err != nil {
		return
	}
	err = ds.Set(ref)
	if err != nil {
		return
	}
	return
}
