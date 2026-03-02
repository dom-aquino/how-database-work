package main

func main() {
	dbName := "app.db"
	_, err := PageManagerCreator(dbName)
	if err != nil {
		panic(err)
	}
}
