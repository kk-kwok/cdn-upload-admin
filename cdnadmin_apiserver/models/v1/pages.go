package models

// get new list base on pagenum pagesize
func Pages(pageNum int, pageSize int, pageList []interface{}) []interface{} {
	var pages int
	newList := make([]interface{}, 0)
	if pageNum == 0 || pageSize == 0 {
		newList = pageList
	} else {
		// get pages
		total := len(pageList)
		if total > pageSize {
			pages = total / pageSize
			remainder := total % pageSize
			if remainder > 0 {
				pages = pages + 1
			}
		} else {
			pages = 1
		}
		// get lower and uper
		if pageNum <= pages {
			pageLow := (pageNum - 1) * pageSize
			pageUp := pageLow + pageSize
			if pageUp+1 > total {
				pageUp = total
			}
			newList = pageList[pageLow:pageUp]
		} else {
			newList = make([]interface{}, 0)
		}
	}
	return newList
}
