package dogapm

import "database/sql"

type dBUtil struct {
}

var DBUtil = &dBUtil{}

// Query 查询
func (d *dBUtil) Query(rows *sql.Rows, err interface{}) []map[string]interface{} {
	if err != nil {
		return nil
	}
	if rows == nil {
		return []map[string]interface{}{}
	}
	defer rows.Close()

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}
	res := make([]map[string]interface{}, 0, 5)
	for rows.Next() {
		record := make(map[string]interface{})
		rows.Scan(scanArgs...)
		for i, col := range values {
			if col != nil {
				switch col.(type) {
				case []byte:
					record[columns[i]] = string(col.([]byte))
				default:
					record[columns[i]] = col
				}

			}
		}
		res = append(res, record)
	}

	return res
}

// QueryFirst 查询第一条
func (d *dBUtil) QueryFirst(rows *sql.Rows, err interface{}) map[string]interface{} {
	res := d.Query(rows, err)
	if len(res) > 0 {
		return res[0]
	}

	return nil
}
