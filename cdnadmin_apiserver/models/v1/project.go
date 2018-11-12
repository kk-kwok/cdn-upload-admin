package models

import (
	"cdn-upload-admin/cdnadmin_apiserver/config"
	"os"

	"github.com/json-iterator/go"
)

// GetProjectList : GET PROJECT LIST
func GetProjectList(pageNum int, pageSize int, pj *Project, platName string) (content []byte, err error) {
	var p Project
	// query db
	db := config.GetDB()
	sql := "SELECT a.`id`, a.`name`, a.`platform_id`, b.`name`, IFNULL(a.`cdn_id`, 0), a.`path`, a.`domain`, a.`status`, a.`create_time`, a.`update_time`" +
		"FROM t_project a INNER JOIN t_cdn_platform b WHERE a.`platform_id` = b.`id`"
	if pj.Name != "" {
		sql = sql + " AND a.`name` like '%" + pj.Name + "%'"
	}
	if pj.Status != "" {
		sql = sql + " AND a.`status` like '%" + pj.Status + "%'"
	}
	if platName != "" {
		sql = sql + " AND b.`name` like '%" + platName + "%'"
	}
	sql = sql + " AND a.`is_delete`=0 ORDER BY a.`id`;"
	rows, err := db.Query(sql)
	if err != nil {
		return
	}
	defer rows.Close()

	projects := make([]Project, 0)
	for rows.Next() {
		rows.Scan(&p.ID, &p.Name, &p.PlatformID, &p.PlatformName, &p.CdnID,
			&p.Path, &p.Domain, &p.Status, &p.CreateTime, &p.UpdateTime)
		projects = append(projects, p)
	}
	total := len(projects)
	if total < 1 {
		content, err = jsoniter.Marshal(Response{Code: 4000, Msg: "record not found", Total: total})
	} else {
		pageList := make([]interface{}, len(projects))
		for i, v := range projects {
			pageList[i] = v
		}
		newProjects := Pages(pageNum, pageSize, pageList)
		content, err = jsoniter.Marshal(Response{Code: 200, Msg: "get success", Data: newProjects, Total: total})
	}
	if err != nil {
		return
	}
	return
}

// AddProject : CREATE PROJECT
func AddProject(pj *Project) (content []byte, err error) {
	var p Project
	// query db
	db := config.GetDB()
	sql := "SELECT `id`, `is_delete`, `name` FROM t_project WHERE `name`=?;"
	err = db.QueryRow(sql, pj.Name).Scan(&p.ID, &p.IsDelete, &p.Name)

	// if err is nil and it's means record is exist
	if err == nil {
		if p.IsDelete == 1 {
			sql = "UPDATE t_project SET `is_delete`=0, `update_time`=now()" +
				"WHERE `name`=?;"
			stmt, _ := db.Prepare(sql)
			_, err = stmt.Exec(pj.Name)
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
		sql = "INSERT INTO t_project (`id`, `name`, `platform_id`, `cdn_id`, `path`, `domain`, `status`)" +
			" VALUES (?, ?, (SELECT `id` FROM t_cdn_platform WHERE `name`=? LiMIT 1), ?, ?, ?, ?);"
		stmt, _ := db.Prepare(sql)
		_, err = stmt.Exec(pj.ID, pj.Name, pj.PlatformName,
			pj.CdnID, pj.Path, pj.Domain, pj.Status)
		defer stmt.Close()
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
		fullPath := rootPath + path + pj.Path
		err = os.MkdirAll(fullPath, os.ModePerm)
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

// UpdateProject : UPDATE PROJECT
func UpdateProject(pj *Project) (content []byte, err error) {
	var p Project
	// query db
	db := config.GetDB()
	sql := "SELECT `id`, `name`, `path` FROM t_project WHERE `id`=? AND `name`=?;"
	err = db.QueryRow(sql, pj.ID, pj.Name).Scan(&p.ID, &p.Name, &p.Path)

	// if err is not nil and it's means record not found
	if err != nil {
		content, err = jsoniter.Marshal(Response{Code: 4000,
			Msg: "record not found"})
	} else {
		sql = "UPDATE t_project SET `id`=?, `name`=?, " +
			"`platform_id`=(SELECT `id` FROM t_cdn_platform WHERE `name`=? LiMIT 1), " +
			" `cdn_id`=?, `path`=?, `domain`=?, `status`=?, " +
			"`update_time`=now() WHERE `id`=?"
		stmt, _ := db.Prepare(sql)
		rs, execErr := stmt.Exec(pj.ID, pj.Name, pj.PlatformName,
			pj.CdnID, pj.Path, pj.Domain, pj.Status, pj.ID)
		defer stmt.Close()
		if execErr != nil {
			return
		}
		// get exec sql affect rows
		affect, _ := rs.RowsAffected()
		if affect != 0 {
			rootPath := config.Config().HTTP.RootPath
			var path string
			if os.IsPathSeparator('\\') {
				path = "\\"
			} else {
				path = "/"
			}
			oldPath := rootPath + path + p.Path
			newPath := rootPath + path + pj.Path
			if oldPath != newPath {
				err = os.Rename(oldPath, newPath)
				if err != nil {
					return
				}
			}
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

// DeleteProject : DELETE PROJECT
func DeleteProject(pj *Project) (content []byte, err error) {
	var p Project
	// query db
	db := config.GetDB()
	sql := "SELECT `id`, `name`, `path` FROM t_project WHERE `id`=? AND `name`=?;"
	err = db.QueryRow(sql, pj.ID, pj.Name).Scan(&p.ID, &p.Name, &p.Path)

	// if err is not nil and it's means record not found
	if err != nil {
		content, err = jsoniter.Marshal(Response{Code: 4000,
			Msg: "record not found"})
	} else {
		sql = "UPDATE t_project SET `is_delete`=1, `status`='已删除', `update_time`=now() WHERE `id`=? AND `name`=?;"
		stmt, _ := db.Prepare(sql)
		rs, execErr := stmt.Exec(pj.ID, pj.Name)
		defer stmt.Close()
		if execErr != nil {
			return
		}
		// get exec sql affect rows
		affect, _ := rs.RowsAffected()
		if affect != 0 {
			// update t_cdn_file
			sql = "UPDATE t_cdn_file SET `is_delete`=1, `update_time`=now() WHERE `project_id`=?;"
			stmt, _ := db.Prepare(sql)
			_, err = stmt.Exec(pj.ID)
			defer stmt.Close()
			if err != nil {
				return
			}
			// delete project path
			rootPath := config.Config().HTTP.RootPath
			var path string
			if os.IsPathSeparator('\\') {
				path = "\\"
			} else {
				path = "/"
			}
			oldPath := rootPath + path + p.Path
			err = os.RemoveAll(oldPath)
			if err != nil {
				return
			}
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
