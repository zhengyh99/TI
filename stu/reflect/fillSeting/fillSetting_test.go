package fillSeting

import "testing"

type stu struct {
	Name string
	Age  int
}
type epl struct {
	Name string
	Age  int
}

func TestFillSeting(t *testing.T) {
	settings := map[string]interface{}{"Name": "xiaohua", "Age": 20}
	s := stu{}
	e := new(epl)
	if err := FillSeting(&s, settings); err == nil {
		t.Log(s)
	} else {
		t.Log(err)
	}
	if err := FillSeting(e, settings); err == nil {
		t.Log(e)
	} else {
		t.Log(err)
	}
}
