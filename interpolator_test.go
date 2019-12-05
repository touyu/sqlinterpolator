package sqlinterpolator

import (
	"testing"
	"time"
)

func TestInterpolate(t *testing.T) {
	fullQuery := "select * from users where id = 1 and name <> 'jack' and is_checked = 1 and created_at < '1997-02-22 22:22:22'"
	placeholderQuery := "select * from users where id = ? and name <> ? and is_checked = ? and created_at < ?"
	createdAt := time.Date(1997, 2, 22, 22, 22, 22, 0, time.UTC)
	args := []interface{}{1, "jack", true, createdAt}

	res, err := Interpolate(placeholderQuery, args)
	if err != nil {
		t.Errorf("Can't interpolate")
	}

	if res != fullQuery {
		t.Errorf("Not equal fullQuery")
	}
}
