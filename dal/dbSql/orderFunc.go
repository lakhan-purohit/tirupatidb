package dbSql

import (
	"fmt"
	"tirupatiBE/dal/dbModel"

	"github.com/jmoiron/sqlx"
)

func (d *dbAccess) MakeOrderList(order dbModel.Order) (bool, error) {

	var productData dbModel.Product

	query := "insert into " + dbModel.ORDERLIST + " (user_id, product_id, qty, total, status,price) values (?,?,?,?,?,?)"

	selectQuery := "select * from " + dbModel.PRODUCTLIST + " where id = ?"

	for _, val := range order.Order {

		d.db.QueryRowx(selectQuery, val.ProductID).StructScan(&productData)

		total := int64(val.Qty) * productData.Price

		d.db.Exec(query, val.UserID, val.ProductID, val.Qty, total, val.Status, productData.Price)

	}

	return true, nil
}

func (d *dbAccess) UpdateOrder(ID int64, price int64, qty int64) (bool, error) {

	query := "update " + dbModel.ORDERLIST + " set  qty = ?, total = ? where id = ?"
	total := qty * price

	if _, err := d.db.Exec(query, qty, total, ID); err != nil {
		return false, err
	}

	return true, nil
}

func (d *dbAccess) ChangeOrderStatus(ID int64, status int64) (bool, error) {

	query := "update " + dbModel.ORDERLIST + " set  status = ? where id = ?"

	if _, err := d.db.Exec(query, status, ID); err != nil {
		return false, err
	}

	return true, nil
}

func (d *dbAccess) DeleteOrder(ID int64) (bool, error) {

	query := "delete from " + dbModel.ORDERLIST + " where id = ?"

	if _, err := d.db.Exec(query, ID); err != nil {
		return false, err
	}

	return true, nil
}

func (d *dbAccess) GetOrder(ID int64) ([]dbModel.OrderDetails, error) {

	var orderData []dbModel.OrderDetails

	selectQuery := "select * from " + dbModel.ORDERLIST

	selectProduct := "select * from " + dbModel.PRODUCTLIST + " where id = ?"
	selectUser := "select * from " + dbModel.USERINFOTABLE + " where id = ?"

	if ID != 0 {
		selectQuery += " where id = ?"
	}
	var val *sqlx.Rows
	var err error

	if ID == 0 {
		val, err = d.db.Queryx(selectQuery)
		if err != nil {
			fmt.Println(err)
			return orderData, nil
		}
	} else {
		val, err = d.db.Queryx(selectQuery, ID)
		if err != nil {
			fmt.Println(err)
			return orderData, nil
		}
	}

	for val.Next() {
		var t dbModel.MakeOrder
		var z dbModel.OrderDetails

		if errFetch := val.StructScan(&t); errFetch != nil {
			fmt.Println(errFetch)
			return orderData, nil
		}

		if errP := d.db.QueryRowx(selectProduct, t.ProductID).StructScan(&z.Product); errP != nil {
			fmt.Println(errP)
			return orderData, nil
		}

		if errU := d.db.QueryRowx(selectUser, t.UserID).StructScan(&z.User); errU != nil {
			return orderData, nil
		}

		z.CreatedAt = t.CreatedAt
		z.ID = t.ID
		z.Paid = t.Paid
		z.Total = t.Total
		z.UpdateAt = t.UpdateAt
		z.Status = t.Status
		z.Qty = t.Qty
		z.Price = t.Price

		orderData = append(orderData, z)

	}

	return orderData, nil

}
