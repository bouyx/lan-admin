package handler

import "testing"

func TestSum(t *testing.T) {
    allUsers := getAllUsers()
    if len(allUsers) == 0 {
       t.Errorf("We need at least one user")
    }
}