package model_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gorm-demo/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	user := model.User{
		FirstName: "test",
	}
	assert.NotEqual(t, user.FirstName, "")
	pretty([]byte(fmt.Sprintf("%v", user)), t)
}

func pretty(obj []byte, t *testing.T) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, obj, "", "\t")
	if err != nil {
		t.Error(err)
	}
	t.Log(string(prettyJSON.Bytes()))
}
