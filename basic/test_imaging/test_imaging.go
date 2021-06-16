package main

import (
    "bytes"
    "errors"
    "fmt"
    "io"
    "os"
)

func main() {
    file, _ := os.Create("new.png")
    defer file.Close()
    //mark, _ := imaging.Open("/Users/anpuqiang/Documents/code/gitlab/src/image-server/resource/watermark.png")
    //mark = imaging.Resize(mark, 3*mark.Bounds().Max.X/4, 3*mark.Bounds().Max.Y/4, imaging.Lanczos)
    //_ = imaging.Encode(file, mark, 1)


    err := Dexif("/Users/anpuqiang/Documents/pictures/big.jpeg", "/Users/anpuqiang/Documents/pictures/big.jpeg")
    if err != nil {
        fmt.Println(err)
    }

}



func Dexif(filepath string, destpath string) error {
    f, err := os.Open(filepath)
    defer func() {
        _ = f.Close()
    }()
    if err != nil {
        return err
    }

    b := make([]byte, 2)
    _, err = f.Seek(2, 0)
    if err != nil {
        return err
    }
    _, err = f.Read(b)
    if err != nil {
        return err
    }

    res := bytes.Compare(b, []byte{0xFF, 0xE1})
    if res == 0 {
        for i := 0; i < 32768; i++ {
            _, err = f.Read(b)
            if err != nil {
                return err
            }
            res := bytes.Compare(b, []byte{0xFF, 0xD9})
            if res == 0 {
                break
            }
        }
    } else {
        return errors.New("no exif data found")
    }

    fout, err := os.Create(destpath)
    if err != nil {
        return err
    }
    defer func() {
        _ = fout.Close()
    }()

    _, err = fout.Write([]byte{0xFF, 0xD8})
    if err != nil {
        return err
    }

    _, err = io.Copy(fout, f)
    if err != nil {
        return err
    }

    return nil
}