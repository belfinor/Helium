package boltdb

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2018-11-06

func Close() {
	if dbh != nil {
		dbh.Close()
		dbh = nil
	}
}

func Get(bucket, key []byte) []byte {

	if dbh != nil {
		return dbh.Get(bucket, key)
	}

	return nil
}

func Set(bucket, key, value []byte) {

	if dbh != nil {
		dbh.Set(bucket, key, value)
	}

}

func Delete(bucket, key []byte) {
	if dbh != nil {
		dbh.Delete(bucket, key)
	}
}

func NextId(bucket []byte) int64 {
	if dbh != nil {
		return dbh.NextId(bucket)
	}

	return 0
}

func Prefix(bucket, prefix []byte, f FETCH_CALLBACK) {
	if dbh != nil {
		dbh.Prefix(bucket, prefix, f)
	}
}

func Range(bucket, from, to []byte, f FETCH_CALLBACK) {
	if dbh != nil {
		dbh.Range(bucket, from, to, f)
	}
}

func ForEach(bucket []byte, f FETCH_CALLBACK) {
	if dbh != nil {
		dbh.ForEach(bucket, f)
	}
}

func Transaction(f TRANSACTION_CALLBACK) {
	if dbh != nil {
		dbh.Transaction(f)
	}
}
