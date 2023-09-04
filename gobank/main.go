package main

import "log"

// docker run --name gobank -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=gobank -p 5432:5432 postgres
// https://youtu.be/f3mwdyxxMfg?list=PL0xRBLFXXsP6nudFDqMXzrvQCZrxSOm-2&t=5

func main() {
	store, err := NewPostgresStorage()

	if err != nil {
		log.Fatal("Database error: ", err)
	}

	if err := store.Init(); err != nil {
		log.Fatal("Database error: ", err)
	}


	server := NewAPIServer(":3000", store)

	server.Run()

}
