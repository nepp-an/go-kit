package main

import (
    "fmt"
    pigo "github.com/esimov/pigo/core"
    _ "image/jpeg"
    "io/ioutil"
    "log"
    "os"
    "time"
)
func main() {
    start := time.Now()
    cascadeFile, err := ioutil.ReadFile("/Users/anpuqiang/Documents/code/gitlab/src/github.com/esimov/pigo/cascade/facefinder")
    if err != nil {
        log.Fatalf("Error reading the cascade file: %v", err)
    }
    dir := "/Users/anpuqiang/Documents/pictures/8da4864389d3dff31bba4fa75d37ce17.jpeg"
    file, err := os.Open(dir)
    if err != nil {
        log.Fatalf("1 Cannot open the image file: %v", err)
    }
    defer file.Close()
    src, err := pigo.GetImage(dir)
    if err != nil {
        log.Fatalf("Cannot open the image file: %v", err)
    }

    pixels := pigo.RgbToGrayscale(src)
    cols, rows := src.Bounds().Max.X, src.Bounds().Max.Y

    cParams := pigo.CascadeParams{
        MinSize:     30,
        MaxSize:     1000,
        ShiftFactor: 0.1,
        ScaleFactor: 1.1,

        ImageParams: pigo.ImageParams{
            Pixels: pixels,
            Rows:   rows,
            Cols:   cols,
            Dim:    cols,
        },
    }

    pigo2 := pigo.NewPigo()
    // Unpack the binary file. This will return the number of cascade trees,
    // the tree depth, the threshold and the prediction from tree's leaf nodes.
    classifier, err := pigo2.Unpack(cascadeFile)
    if err != nil {
        log.Fatalf("Error reading the cascade file: %s", err)
    }
    fmt.Println()

    angle := 0.0 // cascade rotation angle. 0.0 is 0 radians and 1.0 is 2*pi radians

    // Run the classifier over the obtained leaf nodes and return the detection results.
    // The result contains quadruplets representing the row, column, scale and detection score.
    var trueDets []pigo.Detection
    dets := classifier.RunCascade(cParams, angle)
    fmt.Printf("%v", dets)
    for _, v := range dets {
        if v.Q < 5.0 {
            continue
        }
        trueDets = append(trueDets, v)
    }
    fmt.Println()
    fmt.Printf("%v", trueDets)

    fmt.Println()
    // Calculate the intersection over union (IoU) of two clusters.
    dets = classifier.ClusterDetections(trueDets, 0.1)
    fmt.Printf("%v", dets)
    var height, width, weight float32
    for _, v := range dets {
        height += float32(v.Row) * v.Q
        width += float32(v.Col) * v.Q

        weight += v.Q
    }
    fmt.Println()
    fmt.Println(width/weight, height/weight)

    fmt.Printf(time.Since(start).String())
}
