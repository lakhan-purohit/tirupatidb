package router

import (
	"encoding/json"
	"fmt"
	"tirupatiBE/dal/dbModel"

	"github.com/valyala/fasthttp"
)

func (r *Rtr) createUser(ctx *fasthttp.RequestCtx) {
	var userInfo dbModel.UserInfo

	e := json.Unmarshal(ctx.PostBody(), &userInfo)
	if e != nil {
		fmt.Println("err", e)
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong unmarshal\"}")
		return
	}

	fmt.Println(userInfo.UserMobileNumber)

	res, err := r.db.CreateUser(userInfo)
	if err != nil {
		fmt.Println(err)
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong fetch data\"}")
		return
	}

	val, errVal := json.Marshal(res)
	if errVal != nil {
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong unmarshal\"}")
		return

	}

	ctx.WriteString(string(val))
	return

}

func (r *Rtr) loginUser(ctx *fasthttp.RequestCtx) {
	var userInfo dbModel.UserInfo

	e := json.Unmarshal(ctx.PostBody(), &userInfo)
	if e != nil {
		fmt.Println("err", e)
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong unmarshal\"}")
		return
	}

	fmt.Println(userInfo.UserMobileNumber)

	res, err := r.db.LoginUser(userInfo)
	if err != nil {
		fmt.Println(err)
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong fetch data\"}")
		return
	}

	val, errVal := json.Marshal(res)
	if errVal != nil {
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong unmarshal\"}")
		return

	}

	ctx.WriteString(string(val))
	return

}

func (r *Rtr) userList(ctx *fasthttp.RequestCtx) {
	res, err := r.db.UserList()
	if err != nil {
		fmt.Println(err)
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong fetch data\"}")
		return
	}

	val, errVal := json.Marshal(res)
	if errVal != nil {
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong unmarshal\"}")
		return

	}

	ctx.WriteString(string(val))
	return

}

func (r *Rtr) userDetailsUpdate(ctx *fasthttp.RequestCtx) {

	var userDetails dbModel.UserInfo

	err := json.Unmarshal(ctx.PostBody(), &userDetails)
	if err != nil {
		fmt.Println(err)
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong fetch data\"}")
		return
	}

	res, err := r.db.UserDetailsUpdate(userDetails)
	if err != nil {
		fmt.Println(err)
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong \"}")
		return
	}

	val, errVal := json.Marshal(res)
	if errVal != nil {
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong unmarshal\"}")
		return

	}

	ctx.WriteString(string(val))
	return

}
