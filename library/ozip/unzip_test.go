package ozip

import (
	`fmt`
	`testing`
)

func TestDeCompressZip(t *testing.T) {
	err := DeCompressZip("test.zip", "./tmp", "123456", nil, 0)
	if err != nil {
		fmt.Println(err)
	}
	return
}
