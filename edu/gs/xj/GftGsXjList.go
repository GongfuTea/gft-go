package xj

type GftGsXjList []*GftGsXj

func (l GftGsXjList) Includes(it *GftGsXj) bool {
	for _, x := range l {
		if x.Xh == it.Xh {
			return true
		}
	}

	return false
}

func (l GftGsXjList) FindByXh(xh string) GftGsXjList {
	ls := make(GftGsXjList, 0)

	for _, x := range l {
		if x.Xh == xh {
			ls = append(ls, x)
		}
	}
	return ls
}

func (l GftGsXjList) FindByXhLastVersion(xh string) (s *GftGsXj) {
	for _, x := range l {
		if x.Xh == xh {
			if s == nil || s.TlVersion < x.TlVersion {
				s = x
			}
		}
	}
	return s
}

func (l GftGsXjList) FindAdded(oldList GftGsXjList) (added GftGsXjList) {
	added = make(GftGsXjList, 0)

	for _, x := range l {
		if !oldList.Includes(x) {
			added = append(added, x)
		}
	}

	return added
}

func (l GftGsXjList) FindRemoved(oldList GftGsXjList) GftGsXjList {
	return oldList.FindAdded(l)
}

func (l GftGsXjList) FilterByNj(nj int) GftGsXjList {
	ls := make(GftGsXjList, 0)
	for _, x := range l {
		if x.Nj == nj {
			ls = append(ls, x)
		}
	}
	return ls
}

func (l GftGsXjList) FindUpdated(oldList GftGsXjList) (updated GftGsXjList) {
	updated = make(GftGsXjList, 0)

	for _, x := range l {
		if old := oldList.FindByXhLastVersion(x.Xh); old != nil {
			diff := x.Diff(old)
			if len(diff) > 0 {
				x.TlDiff = diff
				// x.TlVersion = old.TlVersion + 1
				// x.Id = fmt.Sprintf("%s-%d", x.Xh, x.TlVersion)
				// x.CreatedAt = time.Now()
				updated = append(updated, x)
			}
		}
	}

	return updated
}
