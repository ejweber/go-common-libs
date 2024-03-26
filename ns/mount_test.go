// TODO: Remove this file or use it to implement correct tests.

package ns

import "testing"

func TestNSMounterList(t *testing.T) {
	nsMounter := NewNSMounter("")
	result, err := nsMounter.List()
	t.Logf("%v", err)
	t.Errorf("%v", result)
}
