package database

var connection string

func init() {
	connection = "MySQL"
}

func GetConnection() string {
	return connection
}

//kalau init akan otomatis dieksekusi ketika package diimport
