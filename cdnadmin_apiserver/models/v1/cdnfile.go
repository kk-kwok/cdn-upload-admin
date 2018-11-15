package models

import (
	"cdn-upload-admin/cdnadmin_apiserver/config"
	"os"
	"strings"

	"github.com/json-iterator/go"
)

// GetCDNFileList : GET CDN FILE LIST
func GetCDNFileList(pageNum int, pageSize int, cf *CDNFile, suffix string) (content []byte, err error) {
	var cfile CDNFile
	// query db
	db := config.GetDB()
	sql := "SELECT b.`id`, a.`name`, a.`domain`, a.`path`, b.`file_name`, b.`file_size`," +
		" b.`file_md5`, b.`comment`,  b.`user_id`, c.`username`, b.`create_time`, b.`update_time`" +
		" FROM (t_project a INNER JOIN t_cdn_file b ON a.`id`=b.`project_id` AND b.`is_delete`=0)" +
		" INNER JOIN t_user c WHERE b.`user_id`=c.`id`"
	if cf.ProjectName != "" {
		sql = sql + " AND a.`name` like '%" + cf.ProjectName + "%'"
	}
	if suffix != "" {
		sql = sql + " AND b.`file_name` like '%" + suffix + "%'"
	}
	sql = sql + " ORDER BY b.`update_time` DESC;"
	rows, err := db.Query(sql)
	if err != nil {
		return
	}
	defer rows.Close()

	cfList := make([]CDNFile, 0)
	var dm = "http://"
	for rows.Next() {
		rows.Scan(&cfile.ID, &cfile.ProjectName, &cfile.Domain, &cfile.Path,
			&cfile.FileName, &cfile.FileSize, &cfile.FileMD5, &cfile.Comment,
			&cfile.UserID, &cfile.UserName, &cfile.CreateTime, &cfile.UpdateTime)
		if strings.HasSuffix(cfile.FileName, "plist") {
			dm = "https://"
		} else {
			dm = "http://"
		}
		if !strings.HasPrefix(cfile.Domain, "http") {
			cfile.Domain = dm + cfile.Domain
		}
		cfList = append(cfList, cfile)
	}
	total := len(cfList)
	if total < 1 {
		content, err = jsoniter.Marshal(Response{Code: 4000, Msg: "record not found", Total: total})
	} else {
		pageList := make([]interface{}, len(cfList))
		for i, v := range cfList {
			pageList[i] = v
		}
		newCFList := Pages(pageNum, pageSize, pageList)
		content, err = jsoniter.Marshal(Response{Code: 200, Msg: "get success", Data: newCFList, Total: total})
	}
	if err != nil {
		return
	}
	return
}

// AddCDNFile : CREATE CDNFILE
func AddCDNFile(tmpFile string, cf *CDNFile) (content []byte, err error) {
	var cfile CDNFile
	// query db
	db := config.GetDB()
	sql := "SELECT `id`, `path` FROM t_project WHERE `name`=?;"
	err = db.QueryRow(sql, cf.ProjectName).Scan(&cfile.ProjectID, &cfile.Path)
	if err != nil {
		return
	}
	if cfile.Path == "" {
		content, err = jsoniter.Marshal(Response{Code: 4000, Msg: "project path undefined"})
	} else {
		rootPath := config.Config().HTTP.RootPath
		var path string
		if os.IsPathSeparator('\\') {
			path = "\\"
		} else {
			path = "/"
		}
		fullPath := rootPath + path + cfile.Path
		os.MkdirAll(fullPath, os.ModePerm)
		err = os.Rename(tmpFile, fullPath+path+cf.FileName)
		if err != nil {
			return
		}
		// query it's already exist
		sql = "SELECT `id` FROM t_cdn_file WHERE `project_id`=? AND `file_name`=?;"
		err = db.QueryRow(sql, cfile.ProjectID, cf.FileName).Scan(&cfile.ID)
		if err != nil && cfile.ID < 1 {
			sql = "INSERT INTO t_cdn_file (`project_id`, `file_name`, `file_md5`, `file_size`, `comment`, `user_id`) VALUES (?, ?, ?, ?, ?, ?)"
			stmt, _ := db.Prepare(sql)
			res, execErr := stmt.Exec(cfile.ProjectID, cf.FileName, cf.FileMD5, cf.FileSize, cf.Comment, cf.UserID)
			if execErr != nil {
				return
			}
			id, _ := res.LastInsertId()
			content, err = CDNFilePush(int(id))
			content, err = jsoniter.Marshal(Response{Code: 200, Msg: "create success"})
		} else {
			sql = "UPDATE t_cdn_file SET `is_delete`=0, `file_name`=?, `file_md5`=?, `file_size`=?, `comment`=?, `user_id`=?, `update_time`=NOW() WHERE `id`=?;"
			stmt, _ := db.Prepare(sql)
			_, err = stmt.Exec(cf.FileName, cf.FileMD5, cf.FileSize, cf.Comment, cf.UserID, cfile.ID)
			if err != nil {
				return
			}
			content, err = CDNFilePush(cfile.ID)
			content, err = jsoniter.Marshal(Response{Code: 200, Msg: "update success"})
		}
	}
	return
}

// UpdateCDNFile : UPDATE CDN FILE
func UpdateCDNFile(cf *CDNFile) (content []byte, err error) {
	var cfile CDNFile
	// query db
	db := config.GetDB()
	sql := "SELECT a.`id`, b.`path`, a.`file_name`, a.`comment` FROM t_cdn_file a INNER JOIN t_project b WHERE a.`id`=? AND a.`project_id`=b.`id`;"
	err = db.QueryRow(sql, cf.ID).Scan(&cfile.ID, &cfile.Path, &cfile.FileName, &cfile.Comment)
	if err != nil {
		return
	}
	if cf.FileName != cfile.FileName {
		rootPath := config.Config().HTTP.RootPath
		var path string
		if os.IsPathSeparator('\\') {
			path = "\\"
		} else {
			path = "/"
		}
		oldPath := rootPath + path + cfile.Path + path + cfile.FileName
		newPath := rootPath + path + cfile.Path + path + cf.FileName
		err = os.Rename(oldPath, newPath)
		if err != nil {
			return
		}
		content, err = CDNFilePush(cfile.ID)
	}
	if cf.Comment == "" {
		cf.Comment = cfile.Comment
	}
	sql = "UPDATE t_cdn_file SET `file_name`=?, `comment`=?, `update_time`=NOW() WHERE `id`=?;"
	stmt, _ := db.Prepare(sql)
	_, err = stmt.Exec(cf.FileName, cf.Comment, cfile.ID)
	if err != nil {
		return
	}
	content, err = jsoniter.Marshal(Response{Code: 200, Msg: "update success"})

	return
}

// DeleteCDNFile : DELETE CDN FILE
func DeleteCDNFile(cf *CDNFile) (content []byte, err error) {
	var cfile CDNFile
	// query db
	db := config.GetDB()
	sql := "SELECT a.`id`, b.`path`, a.`file_name` FROM t_cdn_file a INNER JOIN t_project b WHERE a.`id`=? AND a.`project_id`=b.`id`;"
	err = db.QueryRow(sql, cf.ID).Scan(&cfile.ID, &cfile.Path, &cfile.FileName)
	if err != nil {
		return
	}
	sql = "UPDATE t_cdn_file SET `is_delete`=1, `update_time`=NOW() WHERE `id`=?;"
	stmt, _ := db.Prepare(sql)
	_, err = stmt.Exec(cfile.ID)
	if err != nil {
		return
	}
	rootPath := config.Config().HTTP.RootPath
	var path string
	if os.IsPathSeparator('\\') {
		path = "\\"
	} else {
		path = "/"
	}
	oldPath := rootPath + path + cfile.Path + path + cfile.FileName
	f, err := os.Stat(oldPath)
	if err != nil {
		return
	}
	if !f.IsDir() {
		err = os.Remove(oldPath)
		if err != nil {
			return
		}
	}
	content, err = jsoniter.Marshal(Response{Code: 200, Msg: "delete success"})

	return
}
