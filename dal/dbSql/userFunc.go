package dbSql

import (
	"tirupatiBE/config"
	"tirupatiBE/dal/dbModel"
)

func (d *dbAccess) CreateUser(userInfo dbModel.UserInfo) (dbModel.UserInfo, error) {

	var userInfoValue dbModel.UserInfo
	selectQuery := "select * from " + dbModel.USERINFOTABLE + " where id = ?"

	query := "insert into " + dbModel.USERINFOTABLE + " (user_mobile_number,user_name,password,email_id,firm_name,blocked,role) values(?,?,?,?,?,?,?)"
	password, errPass := config.GeneratePasswordFunc(userInfo.Password)
	if errPass != nil {
		return dbModel.UserInfo{}, errPass
	}
	res, err := d.db.Exec(query, userInfo.UserMobileNumber, userInfo.UserName, password, userInfo.EmailID, userInfo.FirmName, userInfo.Blocked, userInfo.Role)
	if err != nil {

		return dbModel.UserInfo{}, err
	}
	ID, errID := res.LastInsertId()
	if errID != nil {

		return dbModel.UserInfo{}, errID
	}

	if errFatch := d.db.QueryRowx(selectQuery, ID).StructScan(&userInfoValue); errFatch != nil {
		return dbModel.UserInfo{}, errFatch
	}
	return userInfoValue, nil

}

func (d *dbAccess) UserDetailsUpdate(userInfo dbModel.UserInfo) (dbModel.UserInfo, error) {

	var userInfoValue dbModel.UserInfo
	selectQuery := "select * from " + dbModel.USERINFOTABLE + " where id = ?"

	query := "update " + dbModel.USERINFOTABLE + " set user_mobile_number = ?,user_name = ?,email_id = ?,firm_name = ?,blocked = ?,role = ? where id = ?"

	_, err := d.db.Exec(query, userInfo.UserMobileNumber, userInfo.UserName, userInfo.EmailID, userInfo.FirmName, userInfo.Blocked, userInfo.Role, userInfo.ID)
	if err != nil {

		return dbModel.UserInfo{}, err
	}

	if errFatch := d.db.QueryRowx(selectQuery, userInfo.ID).StructScan(&userInfoValue); errFatch != nil {
		return dbModel.UserInfo{}, errFatch
	}
	return userInfoValue, nil

}

func (d *dbAccess) LoginUser(userInfo dbModel.UserInfo) (dbModel.UserInfo, error) {

	var userInfoValue dbModel.UserInfo
	selectQuery := "select * from " + dbModel.USERINFOTABLE + " where email_id = ? and password = ?"

	password, errPass := config.GeneratePasswordFunc(userInfo.Password)
	if errPass != nil {
		return dbModel.UserInfo{}, errPass
	}

	if errFatch := d.db.QueryRowx(selectQuery, userInfo.EmailID, password).StructScan(&userInfoValue); errFatch != nil {
		return dbModel.UserInfo{}, errFatch
	}
	return userInfoValue, nil

}

func (d *dbAccess) UserList() ([]dbModel.UserInfo, error) {

	var userInfoValue []dbModel.UserInfo
	selectQuery := "select * from " + dbModel.USERINFOTABLE + " where role <> 1"

	rawValue, err := d.db.Queryx(selectQuery)

	if err != nil {
		return userInfoValue, err
	}

	for rawValue.Next() {

		var val dbModel.UserInfo

		if errFatch := rawValue.StructScan(&val); errFatch != nil {
			return []dbModel.UserInfo{}, errFatch
		}

		userInfoValue = append(userInfoValue, val)
	}

	return userInfoValue, nil

}
