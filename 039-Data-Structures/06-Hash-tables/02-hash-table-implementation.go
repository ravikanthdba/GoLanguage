package main

import (
    "fmt"
    "net/http"
    "log"
    "bufio"
)

func main() {
    res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
    if err != nil {
        log.Fatalln("error reading the file");
    }

    scanner := bufio.NewScanner(res.Body);
    defer res.Body.Close();

    scanner.Split(bufio.ScanWords);
    buckets := make([]int, 200);
    fmt.Println(buckets);

    for scanner.Scan() {
        n := HashBucket(scanner.Text(), 12);
        buckets[n]++;
    }
    fmt.Println(buckets[0:12]);
}

func HashBucket(word string, buckets int) int {
    var sum int;
    for _, v := range word {
        sum += int(v);
    }
    return sum % buckets;
}