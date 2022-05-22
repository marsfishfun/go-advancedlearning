package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"

	"error/dao"
)

type Response struct {
}

type UserRsp struct {
	Code     int    `json:"code, int"`
	Msg      string `json:"msg,string"`
	Username string `json:"username,string"`
}

func FindUserByIP(w http.ResponseWriter, r *http.Request) {
	var rsp *UserRsp
	ip := r.FormValue("ip")

	username, err := dao.FindUserByIP(ip)
	switch err {
	case nil:
		rsp = &UserRsp{Code: 1, Msg: "成功", Username: username}
	case sql.ErrNoRows:
		fmt.Printf("original error: %T, %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace: \n%+v\n", err)
		rsp = &UserRsp{Code: 0, Msg: "该IP没有对应的用户"}
	default:
		fmt.Printf("original error: %T, %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace: \n%+v\n", err)
		rsp = &UserRsp{Code: 0, Msg: err.Error()}
	}
	if err != nil {
		fmt.Printf("original error: %T, %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace: \n%+v\n", err)
		rsp = &UserRsp{Code: 0, Msg: "该IP没有对应的用户"}
	} else {

	}

	jrsp, _ := json.Marshal(rsp)

	fmt.Fprintf(w, string(jrsp))
}
