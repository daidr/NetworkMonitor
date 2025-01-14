package fetch

import (
	"testing"
)

func TestSftpWithPassword_GetData(t *testing.T) {
	fetcher, err := Spawn["sftp"](map[string]any{
		"host":     "172.16.4.6",
		"port":     22,
		"user":     "ubuntu",
		"password": "111111",
		"filepath": "/var/log/bird/{{- (Now.Add (Second -60)).Format \"01-02-2006-15-04\"}}.mrt",
	})
	if err != nil {
		t.Error(err)
		return
	}
	data, err := fetcher.GetData()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(data)
}
