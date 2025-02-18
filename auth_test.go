package mangodex

import (
	"testing"
)

func TestDexClient(t *testing.T) {
	t.Log("Testing Login")
	err := testClient.Login(user, pwd)
	if err != nil {
		t.Error("Failed:", err)
	} else {
		t.Logf("Passed: %+v", *testClient)
	}

	t.Log("Testing CheckToken")
	res, err := testClient.CheckToken()
	if err != nil {
		t.Error("Failed:", err)
	} else {
		t.Logf("Passed: %+v", res)
	}

	t.Log("Testing RefreshSessionToken")
	t.Logf("Before: %+v", *testClient)
	err = testClient.RefreshSessionToken()
	if err != nil {
		t.Error("Failed:", err)
	} else {
		t.Logf("After: %+v", *testClient)
	}

	t.Log("Testing Logout")
	err = testClient.Logout()
	if err != nil {
		t.Error("Failed:", err)
	} else {
		t.Logf("Passed: %+v", *testClient)
	}
}
