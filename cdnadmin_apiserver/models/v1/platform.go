package models

import (
	"cdn-upload-admin/cdnadmin_apiserver/config"

	"github.com/json-iterator/go"
)

// GetPlatformList : GET PLATFORM LIST
func GetPlatformList(pageNum int, pageSize int, pl *Platform) (content []byte, err error) {
	var plat Platform
	// query db
	db := config.GetDB()
	sql := "SELECT `id`, `name`, IFNULL(`secret_id`, '-'), IFNULL(`secret_key`, '-'), " +
		"IFNULL(`api_url`, '-'), IFNULL(`action`, '-'), `status` FROM t_cdn_platform WHERE 1=1 AND `is_delete`=0"
	if pl.Name != "" {
		sql = sql + " AND `name` like '%" + pl.Name + "%'"
	}
	if pl.Status != "" {
		sql = sql + " AND `status` like '%" + pl.Status + "%'"
	}
	sql = sql + " ORDER BY `id`;"
	rows, err := db.Query(sql)
	if err != nil {
		return
	}
	defer rows.Close()

	platforms := make([]Platform, 0)
	for rows.Next() {
		rows.Scan(&plat.ID, &plat.Name, &plat.SecretID, &plat.SecretKey,
			&plat.APIUrl, &plat.Action, &plat.Status)
		platforms = append(platforms, plat)
	}
	total := len(platforms)
	if total < 1 {
		content, err = jsoniter.Marshal(Response{Code: 4000, Msg: "record not found", Total: total})
	} else {
		pageList := make([]interface{}, len(platforms))
		for i, v := range platforms {
			pageList[i] = v
		}
		newPlatforms := Pages(pageNum, pageSize, pageList)
		content, err = jsoniter.Marshal(Response{Code: 200, Msg: "get success", Data: newPlatforms, Total: total})
	}
	if err != nil {
		return
	}
	return
}

// AddPlatform : CREATE PLATFORM
func AddPlatform(pl *Platform) (content []byte, err error) {
	var plat Platform
	// query db
	db := config.GetDB()
	sql := "SELECT `id`, `is_delete`, `name` FROM t_cdn_platform WHERE `name`=?;"
	err = db.QueryRow(sql, pl.Name).Scan(&plat.ID, &plat.IsDelete, &plat.Name)

	// if err is nil and it's means record is exist
	if err == nil {
		if plat.IsDelete == 1 {
			sql = "UPDATE t_cdn_platform SET `is_delete`=0, `secret_id`=?, " +
				"`secret_key`=?, `api_url`=?, `action`=?, `status`=?, " +
				"`update_time`=now() WHERE `name`=?;"
			stmt, _ := db.Prepare(sql)
			_, execErr := stmt.Exec(pl.SecretID, pl.SecretKey,
				pl.APIUrl, pl.Action, pl.Status, pl.Name)
			defer stmt.Close()
			if execErr != nil {
				return
			} else {
				content, err = jsoniter.Marshal(Response{Code: 200,
					Msg: "create success"})
			}
		} else {
			content, err = jsoniter.Marshal(Response{Code: 4000,
				Msg: "record already exist"})
		}
	} else {
		sql = "INSERT INTO t_cdn_platform (`name`, `secret_id`, `secret_key`, " +
			"`api_url`, `action`, `status`) VALUES (?, ?, ?, ?, ?, ?);"
		stmt, _ := db.Prepare(sql)
		_, err = stmt.Exec(pl.Name, pl.SecretID, pl.SecretKey,
			pl.APIUrl, pl.Action, pl.Status)
		defer stmt.Close()
		if err != nil {
			return
		}
		content, err = jsoniter.Marshal(Response{Code: 200,
			Msg: "create success"})
	}

	if err != nil {
		return
	}
	return
}

// UpdatePlatform : UPDATE PLATFORM
func UpdatePlatform(pl *Platform) (content []byte, err error) {
	var plat Platform
	// query db
	db := config.GetDB()
	sql := "SELECT `id`, `name` FROM t_cdn_platform WHERE `id`=? AND `name`=?;"
	err = db.QueryRow(sql, pl.ID, pl.Name).Scan(&plat.ID, &plat.Name)

	// if err is not nil and it's means record not found
	if err != nil {
		content, err = jsoniter.Marshal(Response{Code: 4000,
			Msg: "record not found"})
	} else {
		sql = "UPDATE t_cdn_platform SET `name`=?, `secret_id`=?, " +
			"`secret_key`=?, `api_url`=?, `action`=?, `status`=?, " +
			"`update_time`=now() WHERE `id`=?"
		stmt, _ := db.Prepare(sql)
		rs, execErr := stmt.Exec(pl.Name, pl.SecretID, pl.SecretKey,
			pl.APIUrl, pl.Action, pl.Status, pl.ID)
		defer stmt.Close()
		if execErr != nil {
			return
		}
		// get exec sql affect rows
		affect, _ := rs.RowsAffected()
		if affect != 0 {
			content, err = jsoniter.Marshal(Response{Code: 200,
				Msg: "update success"})
		} else {
			content, err = jsoniter.Marshal(Response{Code: 4000,
				Msg: "update failed"})
		}

	}

	if err != nil {
		return
	}
	return
}

// DeletePlatform : DELETE PLATFORM
func DeletePlatform(pl *Platform) (content []byte, err error) {
	var plat Platform
	// query db
	db := config.GetDB()
	sql := "SELECT `id`, `name` FROM t_cdn_platform WHERE `id`=? AND `name`=?;"
	err = db.QueryRow(sql, pl.ID, pl.Name).Scan(&plat.ID, &plat.Name)

	// if err is not nil and it's means record not found
	if err != nil {
		content, err = jsoniter.Marshal(Response{Code: 4000,
			Msg: "record not found"})
	} else {
		sql = "SELECT `id` FROM t_project WHERE `platform_id`=? AND `is_delete`=0 LIMIT 1;"
		err = db.QueryRow(sql, pl.ID).Scan(&plat.ID)
		if err == nil {
			content, err = jsoniter.Marshal(Response{Code: 4000,
				Msg: "project using this platform"})
		} else {

			sql = "UPDATE t_cdn_platform SET `is_delete`=1, `status`='已删除', `update_time`=now() WHERE `id`=? AND `name`=?;"
			stmt, _ := db.Prepare(sql)
			rs, execErr := stmt.Exec(pl.ID, pl.Name)
			defer stmt.Close()
			if execErr != nil {
				return
			}
			// get exec sql affect rows
			affect, _ := rs.RowsAffected()
			if affect != 0 {
				content, err = jsoniter.Marshal(Response{Code: 200,
					Msg: "delete success"})
			} else {
				content, err = jsoniter.Marshal(Response{Code: 4000,
					Msg: "delete failed"})
			}
		}
	}

	if err != nil {
		return
	}
	return
}
