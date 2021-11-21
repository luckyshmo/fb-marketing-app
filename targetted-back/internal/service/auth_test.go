package service

import "testing"

// Test only to generate constant hashes for passwords for demo users
func Test_generatePass(t *testing.T) {
	t.Skip()
	pass := generatePasswordHash("123123")
	t.Error(pass)
}
