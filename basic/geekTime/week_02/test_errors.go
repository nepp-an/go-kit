package main
import(
    "database/sql"
    "github.com/pkg/errors"

    log "github.com/Sirupsen/logrus"
)

func main() {
    var id = "week_2"
    err := service(id)
    if err != nil {
        log.Errorf("service failed, err is %+v", err)
    }
    return
}

func service(id string) error {
    return biz(id)
}

func dao(id string) error {
    var err = sql.ErrNoRows
    if err == sql.ErrNoRows {
        return errors.Wrap(err, "dao failed, id is "+ id)
    }
    return nil
}

func biz(id string) error {
    return dao(id)
}