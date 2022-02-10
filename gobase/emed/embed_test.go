/*
 * @Date: 2021-11-01 16:49:21
 * @LastEditors: seven sun
 * @LastEditTime: 2021-11-01 16:51:25
 * @FilePath: /interview/gobase/emed/embed_test.go
 */

package emed_test

import (
	_ "embed"
	"testing"
)

//go:embed embed_test.go
var src string

func TestBed(t *testing.T) {
	t.Log(src)
	t.Log("dfdf")
}
