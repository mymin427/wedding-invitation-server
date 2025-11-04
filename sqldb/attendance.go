package sqldb

import (
	"fmt"
	"time"
	"github.com/mymin427/wedding-invitation-server/types"
)

func initializeAttendanceTable() error {
	_, err := sqlDb.Exec(`
		CREATE TABLE IF NOT EXISTS attendance (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			side VARCHAR(10),
			name VARCHER(20),
			meal VARCHAR(20),
			count INTEGER,
			timestamp INTEGER
		)
	`)
	return err
}

func CreateAttendance(side, name, meal string, count int) error {
	_, err := sqlDb.Exec(`
		INSERT INTO attendance (side, name, meal, count, timestamp)
		VALUES (?, ?, ?, ?, ?)
	`, side, name, meal, count, time.Now().Unix())
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func GetAttendance(offset, limit int) ([]types.AttendanceItem, error) {
	rows, err := sqlDb.Query(`
		SELECT id, side, name, meal, count, timestamp
		FROM attendance
		ORDER BY timestamp DESC
		LIMIT ? OFFSET ?
	`, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []types.AttendanceItem{}
	for rows.Next() {
		var it types.AttendanceItem
		if err := rows.Scan(&it.Id, &it.Side, &it.Name, &it.Meal, &it.Count, &it.Timestamp); err != nil {
			return nil, err
		}
		items = append(items, it)
	}
	return items, nil
}

func CountAttendance() (int, error) {
	row := sqlDb.QueryRow(`
		SELECT COUNT(*) FROM attendance
	`)
	var total int
	if err := row.Scan(&total); err != nil {
		return 0, err
	}
	return total, nil
}
