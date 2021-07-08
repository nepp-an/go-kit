package main

type Database interface {
    Fetch(key string) (int, error)
}

var DB Database

func isOver9000(db Database) bool {
    i, err := db.Fetch("powerlevel")
    if err != nil {
        // if this were a real program, this should return an error
        // but this is an example, so it's ok.
        return false
    }

    if i > 9000 {
        return true
    }

    return false
}
