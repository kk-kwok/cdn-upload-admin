package models

import (
	"cdn-upload-admin/cdnadmin_apiserver/config"

	"github.com/json-iterator/go"
)

// GetFileSuffixList : get file suffix list
func GetFileSuffixList(pageNum int, pageSize int) (content []byte, err error) {
	var fsuffix FileSuffix
	// query db
	db := config.GetDB()
	sql := "SELECT `id`, `file_suffix` FROM t_file_suffix WHERE `is_delete`=0 ORDER BY `id`;"
	rows, err := db.Query(sql)
	if err != nil {
		return
	}
	defer rows.Close()

	fsList := make([]FileSuffix, 0)
	for rows.Next() {
		rows.Scan(&fsuffix.ID, &fsuffix.FileSuffix)
		fsList = append(fsList, fsuffix)
	}
	total := len(fsList)
	if total < 1 {
		content, err = jsoniter.Marshal(Response{Code: 4000, Msg: "record not found", Total: total})
	} else {
		pageList := make([]interface{}, len(fsList))
		for i, v := range fsList {
			pageList[i] = v
		}
		newfsList := Pages(pageNum, pageSize, pageList)
		content, err = jsoniter.Marshal(Response{Code: 200, Msg: "get success", Data: newfsList, Total: total})
	}
	if err != nil {
		return
	}
	return
}

// AddFileSuffix : add file suffix
func AddFileSuffix(fs *FileSuffix) (content []byte, err error) {
	var fsuffix FileSuffix
	// query db
	db := config.GetDB()
	sql := "SELECT `id`, `is_delete`, `file_suffix` FROM t_file_suffix WHERE `file_suffix`=?;"
	err = db.QueryRow(sql, fs.FileSuffix).Scan(&fsuffix.ID,
		&fsuffix.IsDelete, &fsuffix.FileSuffix)

	// if err is nil and it's means record is exist
	if err == nil {
		if fsuffix.IsDelete == 1 {
			sql = "UPDATE t_file_suffix SET `is_delete`=0, `update_time`=now()" +
				"WHERE `file_suffix`=?;"
			stmt, _ := db.Prepare(sql)
			_, err = stmt.Exec(fs.FileSuffix)
			defer stmt.Close()
			if err != nil {
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
		sql = "INSERT INTO t_file_suffix (`file_suffix`) VALUES (?);"
		stmt, _ := db.Prepare(sql)
		_, err = stmt.Exec(fs.FileSuffix)
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

// UpdateFileSuffix : update file suffix
func UpdateFileSuffix(fs *FileSuffix) (content []byte, err error) {
	var fsuffix FileSuffix
	// query db
	db := config.GetDB()
	sql := "SELECT `id`, `file_suffix` FROM t_file_suffix WHERE `id`=?;"
	err = db.QueryRow(sql, fs.ID).Scan(&fsuffix.ID, &fsuffix.FileSuffix)

	// if err is not nil and it's means record not found
	if err != nil {
		content, err = jsoniter.Marshal(Response{Code: 4000,
			Msg: "record not found"})
	} else {
		sql = "UPDATE t_file_suffix SET `file_suffix`=? WHERE `id`=?"
		stmt, _ := db.Prepare(sql)
		rs, execErr := stmt.Exec(fs.FileSuffix, fs.ID)
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

// DeleteFileSuffix : delete file suffix
func DeleteFileSuffix(fs *FileSuffix) (content []byte, err error) {
	var fsuffix FileSuffix
	// query db
	db := config.GetDB()
	sql := "SELECT `id`, `file_suffix` FROM t_file_suffix WHERE `id`=? AND `file_suffix`=?;"
	err = db.QueryRow(sql, fs.ID, fs.FileSuffix).Scan(&fsuffix.ID, &fsuffix.FileSuffix)

	// if err is not nil and it's means record not found
	if err != nil {
		content, err = jsoniter.Marshal(Response{Code: 4000,
			Msg: "record not found"})
	} else {

		sql = "UPDATE t_file_suffix SET `is_delete`=1, `update_time`=now() WHERE `id`=? AND `file_suffix`=?;"
		stmt, _ := db.Prepare(sql)
		rs, execErr := stmt.Exec(fs.ID, fs.FileSuffix)
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

	if err != nil {
		return
	}
	return
}
