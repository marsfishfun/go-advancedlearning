package dao

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

const (
	userName = "root"
	password = "689zcb*ZCB"
	ip       = "10.21.16.201"
	port     = "3306"
	dbName   = "sdp-admin"
)

var db *sql.DB

func OpenDB() error {
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	var err error
	db, err = sql.Open("mysql", path)
	// if err != nil {
	// 	return errors.Wrap(err, "open db failed")
	// }

	return errors.Wrap(err, "open db failed")
}

func Close() {
	db.Close()
}

func FindUserByIP(ip string) (string, error) {
	query := `SELECT DISTINCT(acct.PRACCT_NAME) FROM tr_pracct_vip vip, td_acct_primary acct WHERE vip.virtual_vip = ? AND vip.pracct_id = acct.PRACCT_ID AND acct.DATA_SOURCE = 1`
	queryStr := `SELECT DISTINCT(acct.PRACCT_NAME) FROM tr_pracct_vip vip, td_acct_primary acct WHERE vip.virtual_vip = %v AND vip.pracct_id = acct.PRACCT_ID AND acct.DATA_SOURCE = 1`

	var username string
	err := db.QueryRow(query, ip).Scan(&username)

	queryStr = fmt.Sprintf(queryStr, ip)
	return username, errors.Wrap(err, queryStr)
	// return fmt.Errorf(queryStr, err, ip)
}
