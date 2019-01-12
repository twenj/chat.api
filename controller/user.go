package controller

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "home")
}

/**
  注册
 */
func Register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  // 解析提交的表单
	fmt.Fprintln(w, r.Form)
}
