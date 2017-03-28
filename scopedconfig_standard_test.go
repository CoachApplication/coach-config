package configuration

import "testing"

func TestStandardScopedConfig_Get(t *testing.T) {
	sc := NewStandardScopedConfig().ScopedConfig()

	sc.Set("one", NewTestStringConfig(t, "one","one").Config())

	if c, err := sc.Get("one"); err != nil {
		t.Error("ScopedConfig didn't retrieve the assigned Config")
	} else {
		s := "zero"
		c.Get(&s)
		if s == "zero" {
			t.Error("ScopedConfig retrieved incorrect Config")
		}
	}
}

func TestStandardScopedConfig_Set(t *testing.T) {
	sc := NewStandardScopedConfig().ScopedConfig()

	sc.Set("one", NewTestStringConfig(t, "one","one").Config())

	if c, err := sc.Get("one"); err != nil {
		t.Error("ScopedConfig didn't retrieve the assigned Config")
	} else {
		s := "zero"
		c.Get(&s)
		if s == "zero" {
			t.Error("ScopedConfig retrieved incorrect Config")
		}
	}

	sc.Set("two", NewTestStringConfig(t, "two","two").Config())
	sc.Set("two", NewTestStringConfig(t, "nexttwo","nexttwo").Config())

	if c, err := sc.Get("two"); err != nil {
		t.Error("ScopedConfig didn't retrieve the assigned Config")
	} else {
		s := "zero"
		c.Get(&s)
		if s == "zero" {
			t.Error("ScopedConfig retrieved incorrect Config")
		} else if s == "two" {
			t.Error("ScopedConfig did not overwrite Config on second assign")
		}
	}
}



func TestStandardScopedConfig_List(t *testing.T) {
	sc := NewStandardScopedConfig().ScopedConfig()

	sc.Set("one", NewTestStringConfig(t, "one","one").Config())
	sc.Set("two", NewTestStringConfig(t, "two", "two").Config()) // overwrite the previous one
	sc.Set("three", NewTestStringConfig(t, "three", "three").Config())

	list := sc.Order()

	if len(list) != 3 {
		t.Error("Incorrect number of configs returned")
	} else if list[0] != "one" || list[1] != "two" || list[2] != "three" {
		t.Error("Incorrect config list order returned.")
	}
}
