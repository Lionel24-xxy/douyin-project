package repository

import "testing"

func TestInitRedisClient(t *testing.T) {
	if err := InitRedisClient(); err != nil {
		t.Errorf(err.Error())
	}
}