package sqlinterpolator

import "testing"

func TestInterpolate(t *testing.T) {
	fullQuery := "select * from users where id = 1 and name <> 'jack' and is_checked = 1"
	placeholderQuery := "select * from users where id = ? and name <> ? and is_checked = ?"
	args := []interface{}{1, "jack", true}

	res, err := Interpolate(placeholderQuery, args)
	if err != nil {
		t.Errorf("Can't interpolate")
	}

	if res != fullQuery {
		t.Errorf("Not equal fullQuery")
	}
}
