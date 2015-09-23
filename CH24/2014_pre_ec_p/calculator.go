package main

import (
    "flag"
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
    "math"
    "time"
)

func main() {
    t0 := time.Now()

    const (
        defaultFile = "film_logistics_input"
        usage       = "Input file"
    )

    var sourceFile string
    var movieCount int
    var firstlat float64
    var firstlon float64
    var prewlat float64
    var prewlon float64

    flag.StringVar(&sourceFile, "file", defaultFile, usage)
    flag.Parse()

    inputFile, err := os.Open("inputs/" + sourceFile + ".in");
    defer inputFile.Close()

    if err != nil {
        panic("Input file unreadable")
    }

    scanner  := bufio.NewScanner(bufio.NewReader(inputFile))
    i        := 0
    distance := 0.0

    for scanner.Scan() {

        if i == 0 {
            movieCount, _ = strconv.Atoi(scanner.Text())
        } else {
            coordinates := strings.Split(scanner.Text(), " ")
            lat, _      := strconv.ParseFloat(coordinates[0], 64)
            lon, _      := strconv.ParseFloat(coordinates[1], 64)

            if i == 1 {
                firstlat = lat
                firstlon = lon
            } else {
                distance = distance + getDistance(prewlat, prewlon, lat, lon)
            }

            if i == movieCount {
                distance = distance + getDistance(lat, lon, firstlat, firstlon)
            }

            prewlat = lat
            prewlon = lon
        }

        i = i + 1
    }

    writeToFile(sourceFile, distance)

    t1 := time.Now()
    fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
    fmt.Println(distance / 1000)
}

func writeToFile(src string, dis float64) {
    source, _ := os.OpenFile("outputs/p" + src + ".out", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0777)
    source.Write([]byte(strconv.FormatFloat(dis, 'f', 6, 64)))
    source.Close()
}

func getDistance(lat1 float64, lon1 float64, lat2 float64, lon2 float64) (distance float64) {
    R      := 6371000.0
    sinLat := math.Sin(getRadian(lat1)) * math.Sin(getRadian(lat2))
    cosLat := math.Cos(getRadian(lat1)) * math.Cos(getRadian(lat2)) * math.Cos(getRadian(lon2-lon1))

    distance = math.Acos(sinLat + cosLat)  * R

    return distance
}

func getRadian(num float64) (rad float64) {
    rad = num * math.Pi / 180
    return  rad
}
