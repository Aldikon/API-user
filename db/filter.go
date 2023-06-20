package db

import (
	"fmt"
	"strings"
)

type filter string

func newFilter(base string) filter {
	return filter(base)
}

func (f filter) end() string {
	return string(f) + ";"
}

func (f filter) trim(l int) filter {
	if l > 0 {
		w := []rune(f)
		if len(w) >= l {
			w = w[:len(w)-1-l]
		}
		w = append(w, ' ')
		return filter(w)
	}
	return f
}

type orderBy struct {
	asc  []string
	desc []string
}

func (o orderBy) notZero() bool {
	if len(o.asc) != 0 || len(o.desc) != 0 {
		return true
	}
	return false
}

func (f filter) withOrderBy(o orderBy) filter {
	if o.notZero() {
		b := strings.Builder{}
		b.WriteString(string(f))
		b.WriteString("ORDER BY ")
		for _, v := range o.asc {
			if v != "" {
				b.WriteString(fmt.Sprintf("%s ASC, ", v))
			}
		}
		for _, v := range o.desc {
			if v != "" {
				b.WriteString(fmt.Sprintf("%s DESC, ", v))
			}
		}
		return filter(b.String()).trim(2)
	}
	return f
}

type where map[string]string

func (w where) notZero() bool {
	for _, v := range w {
		if len(v) != 0 {
			return true
		}
	}
	return false
}

func (f filter) withWhere(w where) filter {
	if w.notZero() {
		b := strings.Builder{}
		b.WriteString(string(f))
		b.WriteString("WHERE ")
		for k, v := range w {
			if v != "" {
				b.WriteString(fmt.Sprintf("%s = '%s' AND ", k, v))
			}
		}
		return filter(b.String()).trim(4)
	}
	return f
}

func (f filter) withOffset(d uint) filter {
	if d != 0 {
		return filter(fmt.Sprintf("%s OFFSET %d", f, d))
	}
	return f
}

func (f filter) withLimit(d uint) filter {
	if d != 0 {
		return filter(fmt.Sprintf("%s LIMIT %d", f, d))
	}
	return f
}
