package main

type S3Service struct {
    s3 string
}

type ImageDBService struct {
    imageDB string
}

type CacheService struct {
    cache string
}


type GetService interface {
    GetImage(string)  ([]byte, error)
}

type token struct {
    fcm string
}

func main() {
    var _ Shape = (*Square)(nil)

}


type Shape interface {
    Sides() int
    Area() int
}
type Square struct {
    len int
}

func (s Square) Area() int {
    panic("implement me")
}

func (s *Square) Sides() int {
    return 4
}
//func main() {
//    s := Square{len: 5}
//    fmt.Printf("%d\n",s.Sides())
//}





