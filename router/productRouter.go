package router

import (
	"encoding/json"
	"fmt"
	"strconv"
	"tirupatiBE/dal/dbModel"

	"github.com/valyala/fasthttp"
)

func (r *Rtr) addProduct(ctx *fasthttp.RequestCtx) {

	var product dbModel.Product

	if err := json.Unmarshal(ctx.PostBody(), &product); err != nil {

		fmt.Println(err)
		ctx.WriteString("Something went wrong unmarshal")
		return
	}

	res, errRes := r.db.AddProduct(product)

	if errRes != nil {
		ctx.WriteString("something went wrong post data")
		return
	}

	val, errMarshal := json.Marshal(res)
	if errMarshal != nil {
		ctx.WriteString("something went wrong marshal")
		return
	}

	ctx.WriteString(string(val))

}

func (r *Rtr) updateProduct(ctx *fasthttp.RequestCtx) {

	var product dbModel.Product

	if err := json.Unmarshal(ctx.PostBody(), &product); err != nil {

		fmt.Println(err)
		ctx.WriteString("Something went wrong unmarshal")
		return
	}

	res, errRes := r.db.UpdateProduct(product)

	if errRes != nil {
		ctx.WriteString("something went wrong post data")
		return
	}

	val, errMarshal := json.Marshal(res)
	if errMarshal != nil {
		ctx.WriteString("something went wrong marshal")
		return
	}

	ctx.WriteString(string(val))

}

func (r *Rtr) deleteProduct(ctx *fasthttp.RequestCtx) {

	ID, errID := strconv.Atoi(ctx.UserValue("id").(string))

	if errID != nil {
		ctx.WriteString("something went wrong")
	}

	res, errRes := r.db.DeleteProduct(int64(ID))

	if errRes != nil {
		ctx.WriteString("something went wrong post data")
		return
	}

	val, errMarshal := json.Marshal(res)
	if errMarshal != nil {
		ctx.WriteString("something went wrong marshal")
		return
	}

	ctx.WriteString(string(val))
}

func (r *Rtr) saveImages(ctx *fasthttp.RequestCtx) {

	ID, errID := strconv.Atoi(string(ctx.UserValue("id").(string)))

	if errID != nil {
		ctx.WriteString("Something went to wrong")
	}

	file, err := ctx.FormFile("file")

	if err != nil {
		ctx.WriteString("something went wrong")
	}

	errMultipart := fasthttp.SaveMultipartFile(file, "image/"+file.Filename)

	if errMultipart != nil {
		ctx.WriteString("something went wrong")
	}

	res, errFile := r.db.SaveImagesProduct(int64(ID), file.Filename)

	if errFile != nil {
		ctx.WriteString("something went wrong")
	}

	b, errJson := json.Marshal(res)

	if errJson != nil {
		ctx.WriteString("something went wrong")
	}

	ctx.WriteString(string(b))

}

func (r *Rtr) deleteImages(ctx *fasthttp.RequestCtx) {

	ID, errID := strconv.Atoi(string(ctx.UserValue("id").(string)))

	if errID != nil {
		ctx.WriteString("Something went to wrong")
	}

	res, errFile := r.db.DeleteImageProduct(int64(ID))

	if errFile != nil {
		ctx.WriteString("something went wrong")
	}

	b, errJson := json.Marshal(res)

	if errJson != nil {
		ctx.WriteString("something went wrong")
	}

	ctx.WriteString(string(b))

}

func (r *Rtr) getProduct(ctx *fasthttp.RequestCtx) {

	appValue, errID := strconv.Atoi(string(ctx.UserValue("role").(string)))
	if errID != nil {
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong\"}")
		return

	}

	res, err := r.db.GetProductList(int64(appValue))
	if err != nil {
		fmt.Println(err.Error())
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong\"}")
		return
	}
	b, err := json.Marshal(res)
	if err != nil {
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong\"}")
		return

	}
	if string(b) == "null" {
		b = []byte("[]")
	}

	ctx.WriteString("{\"s\":1,\"m\":\"success\",\"rs\":" + string(b) + "}")
}

func (r Rtr) serveFile(ctx *fasthttp.RequestCtx) {

	filName := string(ctx.UserValue("fileName").(string))

	ctx.SendFile("image/" + filName)

}

func (r Rtr) saveUnSaveProduct(ctx *fasthttp.RequestCtx) {

	var saveProduct dbModel.SaveProduct

	errFetch := json.Unmarshal(ctx.PostBody(), &saveProduct)
	if errFetch != nil {
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong fetch data\"}")
		return
	}

	res, err := r.db.SaveUnSaveProduct(saveProduct.ProductID, saveProduct.UserID)
	if err != nil {
		fmt.Println(err.Error())
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong\"}")
		return
	}
	b, err := json.Marshal(res)
	if err != nil {
		ctx.WriteString("{\"s\":0,\"m\":\"something went wrong\"}")
		return

	}

	ctx.WriteString("{\"s\":1,\"m\":\"success\",\"rs\":" + string(b) + "}")

}
