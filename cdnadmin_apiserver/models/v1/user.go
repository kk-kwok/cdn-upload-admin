package models

import (
	"cdn-upload-admin/cdnadmin_apiserver/config"

	"github.com/json-iterator/go"
)

// GetUserList : get user list
func GetUserList(pageNum int, pageSize int, user *User) (content []byte, err error) {
	var u User
	// query db
	db := config.GetDB()
	sql := "SELECT `id`, `username`, `name`, `is_admin`, `status`, `create_time`, `update_time` FROM t_user WHERE `is_delete`=0"
	if user.UserName != "" {
		sql = sql + " AND `username` like '%" + user.UserName + "%'"
	}
	sql = sql + " ORDER BY `id`;"
	rows, err := db.Query(sql)
	if err != nil {
		return
	}
	defer rows.Close()

	uList := make([]User, 0)
	for rows.Next() {
		rows.Scan(&u.ID, &u.UserName, &u.Name, &u.IsAdmin, &u.Status, &u.CreateTime, &u.UpdateTime)
		uList = append(uList, u)
	}
	total := len(uList)
	if total < 1 {
		content, err = jsoniter.Marshal(Response{Code: 4000, Msg: "record not found", Total: total})
	} else {
		pageList := make([]interface{}, len(uList))
		for i, v := range uList {
			pageList[i] = v
		}
		newUList := Pages(pageNum, pageSize, pageList)
		content, err = jsoniter.Marshal(Response{Code: 200, Msg: "get success", Data: newUList, Total: total})
	}
	if err != nil {
		return
	}
	return
}

// AddUser : create new user
func AddUser(user *User) (content []byte, err error) {
	var u User
	// query db
	db := config.GetDB()
	sql := "SELECT `id`, `is_delete`, `username` FROM t_user WHERE `username`=?"
	err = db.QueryRow(sql, user.UserName).Scan(&u.ID, &u.IsDelete, &u.UserName)

	// if err is nil and it's means record is exist
	if err == nil {
		if u.IsDelete == 1 {
			sql = "UPDATE t_user SET `is_delete`=0, `password`=?, `is_admin`=1," +
				" `name`=?, `status`=?, `update_time`=now() WHERE `username`=?;"
			stmt, _ := db.Prepare(sql)
			_, err = stmt.Exec(user.Password, user.Name, user.Status, user.UserName)
			defer stmt.Close()
			if err != nil {
				return
			}
			content, err = jsoniter.Marshal(Response{Code: 200, Msg: "create success"})
		} else {
			content, err = jsoniter.Marshal(Response{Code: 4000, Msg: "record already exist"})
		}
	} else {
		sql = "INSERT INTO t_user (`username`, `password`, `is_admin`, `name`, `status`) VALUES (?, ?, 1, ?, ?);"
		stmt, _ := db.Prepare(sql)
		_, err = stmt.Exec(user.UserName, user.Password, user.Name, user.Status)
		defer stmt.Close()
		if err != nil {
			return
		}
		content, err = jsoniter.Marshal(Response{Code: 200, Msg: "create success"})
	}
	if err != nil {
		return
	}
	return
}

// UpdateUser : update user info
func UpdateUser(user *User) (content []byte, err error) {
	var u User
	// query db
	db := config.GetDB()
	sql := "SELECT `id`, `is_delete`, `username` FROM t_user WHERE `id`=?"
	err = db.QueryRow(sql, user.ID).Scan(&u.ID, &u.IsDelete, &u.UserName)

	if err != nil {
		return
	} else if u.ID < 1 {
		content, err = jsoniter.Marshal(Response{Code: 4000,
			Msg: "record not found"})
	} else {
		sql = "UPDATE t_user SET `name`=?, `status`=?, `update_time`=now() WHERE `id`=?;"
		stmt, _ := db.Prepare(sql)
		rs, execErr := stmt.Exec(user.Name, user.Status, user.ID)
		defer stmt.Close()
		if execErr != nil {
			return
		}
		// get exec sql affect rows
		affect, _ := rs.RowsAffected()
		if affect != 0 {
			content, err = jsoniter.Marshal(Response{Code: 200, Msg: "update success"})
		} else {
			content, err = jsoniter.Marshal(Response{Code: 4000, Msg: "update failed"})
		}

	}
	if err != nil {
		return
	}
	return
}

// UpdateUserPwd : update user password
func UpdateUserPwd(user *User) (content []byte, err error) {
	var u User
	// query db
	db := config.GetDB()
	sql := "SELECT `id`, `password`, `username` FROM t_user WHERE `id`=?"
	err = db.QueryRow(sql, user.ID).Scan(&u.ID, &u.Password, &u.UserName)

	if err != nil {
		return
	} else if u.ID < 1 {
		content, err = jsoniter.Marshal(Response{Code: 4000,
			Msg: "record not found"})
	} else if user.OldPassword != u.Password {
		content, err = jsoniter.Marshal(Response{Code: 4000,
			Msg: "current password not match"})
	} else {
		sql = "UPDATE t_user SET `password`=?, `update_time`=now() WHERE `id`=?;"
		stmt, _ := db.Prepare(sql)
		rs, execErr := stmt.Exec(user.Password, user.ID)
		defer stmt.Close()
		if execErr != nil {
			return
		}
		// get exec sql affect rows
		affect, _ := rs.RowsAffected()
		if affect != 0 {
			content, err = jsoniter.Marshal(Response{Code: 200, Msg: "update password success"})
		} else {
			content, err = jsoniter.Marshal(Response{Code: 4000, Msg: "update password failed"})
		}

	}
	if err != nil {
		return
	}
	return
}

// ResetUserPwd : reset user password
func ResetUserPwd(user *User) (content []byte, err error) {
	var u User
	// query db
	db := config.GetDB()
	sql := "SELECT `id` FROM t_user WHERE `id`=?"
	err = db.QueryRow(sql, user.ID).Scan(&u.ID)

	if err != nil {
		return
	} else if u.ID < 1 {
		content, err = jsoniter.Marshal(Response{Code: 4000,
			Msg: "record not found"})
	} else {
		sql = "UPDATE t_user SET `password`=md5('Aa@1234%'), `update_time`=now() WHERE `id`=?;"
		stmt, _ := db.Prepare(sql)
		rs, execErr := stmt.Exec(user.ID)
		defer stmt.Close()
		if execErr != nil {
			return
		}
		// get exec sql affect rows
		affect, _ := rs.RowsAffected()
		if affect != 0 {
			content, err = jsoniter.Marshal(Response{Code: 200, Msg: "reset password success"})
		} else {
			content, err = jsoniter.Marshal(Response{Code: 4000, Msg: "reset password failed"})
		}

	}
	if err != nil {
		return
	}
	return
}

// SetUserAdmin : set user is admin or not
func SetUserAdmin(user *User) (content []byte, err error) {
	var u User
	// query db
	db := config.GetDB()
	sql := "SELECT `id` FROM t_user WHERE `id`=?"
	err = db.QueryRow(sql, user.ID).Scan(&u.ID)

	if err != nil {
		return
	} else if u.ID < 1 {
		content, err = jsoniter.Marshal(Response{Code: 4000,
			Msg: "record not found"})
	} else {
		sql = "UPDATE t_user SET `is_admin`=? , `update_time`=now() WHERE `id`=?;"
		stmt, _ := db.Prepare(sql)
		rs, execErr := stmt.Exec(user.IsAdmin, user.ID)
		defer stmt.Close()
		if execErr != nil {
			return
		}
		// get exec sql affect rows
		affect, _ := rs.RowsAffected()
		if affect != 0 {
			content, err = jsoniter.Marshal(Response{Code: 200, Msg: "update success"})
		} else {
			content, err = jsoniter.Marshal(Response{Code: 4000, Msg: "update failed"})
		}

	}
	if err != nil {
		return
	}
	return
}

// DeleteUser : delete user
func DeleteUser(user *User) (content []byte, err error) {
	var u User
	// query db
	db := config.GetDB()
	sql := "SELECT `id`, `is_delete`, `username` FROM t_user WHERE `id`=?"
	err = db.QueryRow(sql, user.ID).Scan(&u.ID, &u.IsDelete, &u.UserName)

	if err != nil {
		return
	} else if u.ID < 1 {
		content, err = jsoniter.Marshal(Response{Code: 4000,
			Msg: "record not found"})
	} else {
		sql = "UPDATE t_user SET `is_delete`=1, `status`='已删除', `update_time`=now() WHERE `id`=?;"
		stmt, _ := db.Prepare(sql)
		rs, execErr := stmt.Exec(user.ID)
		defer stmt.Close()
		if execErr != nil {
			return
		}
		// get exec sql affect rows
		affect, _ := rs.RowsAffected()
		if affect != 0 {
			content, err = jsoniter.Marshal(Response{Code: 200, Msg: "delete success"})
		} else {
			content, err = jsoniter.Marshal(Response{Code: 4000, Msg: "delete failed"})
		}

	}
	if err != nil {
		return
	}
	return
}
