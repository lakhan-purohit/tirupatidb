package router

import (
	"encoding/json"
	"fmt"
	"strconv"
	"tirupatiBE/dal/dbModel"

	"github.com/valyala/fasthttp"
)

func (r *Rtr) makeOrder(ctx *fasthttp.RequestCtx) {

	var order dbModel.Order

	if errFetch := json.Unmarshal(ctx.PostBody(), &order); errFetch != nil {
		fmt.Println(errFetch)
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong fetch data\"}")
		return

	}

	val, err := r.db.MakeOrderList(order)
	if err != nil {
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong \"}")
		return
	}

	b, errMar := json.Marshal(val)
	if errMar != nil {
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong marchal \"}")
		return
	}

	ctx.WriteString(string(b))

}

func (r *Rtr) getOrder(ctx *fasthttp.RequestCtx) {

	ID, err := strconv.Atoi(ctx.UserValue("id").(string))

	if err != nil {
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong  fetch data\"}")
		return
	}

	val, err := r.db.GetOrder(int64(ID))
	if err != nil {
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong \"}")
		return
	}

	b, errMar := json.Marshal(val)
	if errMar != nil {
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong marchal \"}")
		return
	}

	ctx.WriteString(string(b))

}

func (r *Rtr) upateOrder(ctx *fasthttp.RequestCtx) {

	ID, err := strconv.Atoi(ctx.UserValue("id").(string))

	if err != nil {
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong  fetch data\"}")
		return
	}

	price, errP := strconv.Atoi(ctx.UserValue("price").(string))

	if errP != nil {
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong  fetch data\"}")
		return
	}
	qty, errQ := strconv.Atoi(ctx.UserValue("qty").(string))

	if errQ != nil {
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong  fetch data\"}")
		return
	}

	val, err := r.db.UpdateOrder(int64(ID), int64(price), int64(qty))
	if err != nil {
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong \"}")
		return
	}

	b, errMar := json.Marshal(val)
	if errMar != nil {
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong marchal \"}")
		return
	}

	ctx.WriteString(string(b))

}

func (r *Rtr) upateOrderStatus(ctx *fasthttp.RequestCtx) {

	ID, err := strconv.Atoi(ctx.UserValue("id").(string))

	if err != nil {
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong  fetch data\"}")
		return
	}

	status, errP := strconv.Atoi(ctx.UserValue("status").(string))

	if errP != nil {
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong  fetch data\"}")
		return
	}

	val, err := r.db.ChangeOrderStatus(int64(ID), int64(status))
	if err != nil {
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong \"}")
		return
	}

	b, errMar := json.Marshal(val)
	if errMar != nil {
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong marchal \"}")
		return
	}

	ctx.WriteString(string(b))

}

func (r *Rtr) deleteOrder(ctx *fasthttp.RequestCtx) {

	ID, err := strconv.Atoi(ctx.UserValue("id").(string))

	if err != nil {
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong  fetch data\"}")
		return
	}

	val, err := r.db.DeleteOrder(int64(ID))
	if err != nil {
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong \"}")
		return
	}

	b, errMar := json.Marshal(val)
	if errMar != nil {
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong marchal \"}")
		return
	}

	ctx.WriteString(string(b))

}
