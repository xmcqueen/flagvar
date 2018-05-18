package flagvar

import (
	"fmt"
	"strings"
)

// KV is a key-value pair of strings
type KV struct {
	Key   string
	Value string
}

// Assignment is a `flag.Value` for `KEY=VALUE` arguments.
// The value of the `Separator` field is used instead  of `"="` when set.
type Assignment struct {
	Separator string

	Value KV
	Text  string
}

// Set is flag.Value.Set
func (fv *Assignment) Set(v string) error {
	separator := "="
	if fv.Separator != "" {
		separator = fv.Separator
	}
	i := strings.Index(v, separator)
	if i < 0 {
		return fmt.Errorf(`"%s" must have the form KEY%sVALUE`, v, separator)
	}
	fv.Text = v
	fv.Value = KV{
		Key:   v[:i],
		Value: v[i+len(separator):],
	}
	return nil
}

func (fv *Assignment) String() string {
	return fv.Text
}

// Assignments is a `flag.Value` for `KEY=VALUE` arguments.
// The value of the `Separator` field is used instead  of `"="` when set.
type Assignments struct {
	Separator string

	Values []KV
	Texts  []string
}

// Set is flag.Value.Set
func (fv *Assignments) Set(v string) error {
	separator := "="
	if fv.Separator != "" {
		separator = fv.Separator
	}
	i := strings.Index(v, separator)
	if i < 0 {
		return fmt.Errorf(`"%s" must have the form KEY%sVALUE`, v, separator)
	}
	fv.Texts = append(fv.Texts, v)
	fv.Values = append(fv.Values, KV{
		Key:   v[:i],
		Value: v[i+len(separator):],
	})
	return nil
}

func (fv *Assignments) String() string {
	return strings.Join(fv.Texts, ", ")
}
