package datasources

type DataSource interface {
	Schema()
	Scan()
}
