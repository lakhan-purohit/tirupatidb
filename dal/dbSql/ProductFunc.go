package dbSql

import (
	"fmt"
	"tirupatiBE/dal/dbModel"
)

func (d *dbAccess) AddProduct(product dbModel.Product) (int64, error) {

	query := "insert into " + dbModel.PRODUCTLIST + " (color,nack,size,brand_name,hand_slive,pocket,cloth_type,role,price) values(?,?,?,?,?,?,?,?,?)"

	res, err := d.db.Exec(query, product.Color, product.Nack, product.Size, product.BrandName, product.HandSlive, product.Pocket, product.ClothType, product.Role, product.Price)

	if err != nil {
		return 0, err
	}

	ID, errID := res.LastInsertId()

	if errID != nil {
		return 0, nil
	}
	return ID, nil
}

func (d *dbAccess) UpdateProduct(product dbModel.Product) (bool, error) {

	query := "update " + dbModel.PRODUCTLIST + " set color =? ,  nack =? ,   size =? ,  brand_name =? ,  hand_slive =? , pocket =? ,  cloth_type =? ,  role =?, price  where id =?"

	_, err := d.db.Exec(query, product.Color, product.Nack, product.Size, product.BrandName, product.HandSlive, product.Pocket, product.ClothType, product.Role, product.Price, product.ID)

	if err != nil {

		fmt.Println(err)
		return false, err
	}

	return true, nil
}

func (d *dbAccess) SaveImagesProduct(ID int64, imageName string) (bool, error) {

	queryPhoto := "insert into " + dbModel.CLOUTHPHOTO + " (photo,photo_thumb,clothTable_id) values(?,?,?)"

	_, errPhoto := d.db.Exec(queryPhoto, imageName, imageName, ID)

	if errPhoto != nil {
		return false, errPhoto
	}

	return true, nil

}

func (d *dbAccess) DeleteImageProduct(ID int64) (bool, error) {

	queryPhoto := "delete from " + dbModel.CLOUTHPHOTO + " where id =?"

	_, errPhoto := d.db.Exec(queryPhoto, ID)

	if errPhoto != nil {
		return false, errPhoto
	}

	return true, nil

}

func (d *dbAccess) DeleteProduct(ID int64) (bool, error) {

	queryPhoto := "delete from " + dbModel.CLOUTHPHOTO + " where clothTable_id = ?"

	_, errPhoto := d.db.Exec(queryPhoto, ID)

	if errPhoto != nil {
		return false, errPhoto
	}

	queryProduct := "delete from " + dbModel.PRODUCTLIST + " where id = ?"

	_, errProduct := d.db.Exec(queryProduct, ID)

	if errProduct != nil {
		fmt.Println(errProduct)
		return false, errProduct
	}

	return true, nil

}

func (d *dbAccess) GetProductList(role int64) ([]dbModel.Product, error) {

	query := "SELECT * from " + dbModel.PRODUCTLIST

	query += " WHERE `role`=?"

	var data []dbModel.Product

	rows, err := d.db.Queryx(query, role)
	if err != nil {

		return nil, err
	}

	for rows.Next() {
		var t dbModel.Product
		if err := rows.StructScan(&t); err != nil {

			return []dbModel.Product{}, err
		}

		r1, _ := d.db.Queryx("select photo,photo_thumb,id from cloth_photo where `clothTable_id` = ?", t.ID)
		for r1.Next() {

			var z dbModel.ProductPhoto
			if err1 := r1.StructScan(&z); err1 != nil {
				fmt.Println(err1)
			}

			t.Photos = append(t.Photos, z)
		}

		data = append(data, t)

	}

	return data, nil
}

func (d *dbAccess) SaveUnSaveProduct(productID int64, userID int64) (bool, error) {

	var saveProduct dbModel.SaveProduct

	querySelect := "select * from " + dbModel.SAVEPRODUCT + " where user_id = ? and product_id = ?"
	insert := "insert into " + dbModel.SAVEPRODUCT + " (user_id , product_id) values(?,?) "
	delete := "delete from " + dbModel.SAVEPRODUCT + " where id  = ?"

	if errFetch := d.db.QueryRowx(querySelect, userID, productID).StructScan(&saveProduct); errFetch != nil {

		if _, err := d.db.Exec(insert, userID, productID); err != nil {
			return false, err
		}
		return true, nil
	}

	if _, errDelete := d.db.Exec(delete, saveProduct.ID); errDelete != nil {
		return false, nil
	}

	return true, nil

}
