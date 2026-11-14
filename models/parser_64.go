package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "4.5" )

func b6jyWtWNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lGiaUryFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yneo5iNqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IKJuo2LuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oFLxwweRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UTw4aBp3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8z8r5mw0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QflHReTSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8CPFrfOUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jmsRZsTtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GyW9P6wrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Auh92KMmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D3b6HXQBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eEecwfvoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w8alnzhnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iyzLjMwRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uL1VBUZfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func prw5osvCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bTpd5Oo7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gblkmALKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QZWO3OLjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4qx0z7plWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vV63X2gjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7evxhwueWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kgu7EN8nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jCt5dwQPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qToo3CytWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UjUzrpGcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6pkxWDiuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ikTGNObYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mEIvB3WiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PFwTftfxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0K7izwNIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NB7tiGoGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vs1OT5c8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SNoGqKJ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NV9X8V6QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EagGejg3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VF2iCH5oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ggEddEJsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6QEvVSTdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jyFJIuPhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ddA0VgtIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e3ptv8lNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4nZJrvohWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pM99d0s8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ig3BAdjAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RjzgIREyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func icFhLdPYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6wJsDmlBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hHedULm1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VWGCn131Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func na2r4iBRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YuQfeNohWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JjqmqRAyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EcIslRbrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PIx6qCbaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IvnZvD6hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZYEqXReJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7mFbdOGkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lU7cMFKGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U999Wr0yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UE01xDBtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 10EZAslgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YvKBlVY1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 39OWwVHBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lw819A0DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rC9F5dsOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZnL8gglBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3yhMAapcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yp07Mvj4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0tR0nBM6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ns1GBGzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D3Ii9Wy1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LDyVAZNxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qRzwsRcMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IUi0ep1OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tJ6s7z1xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dvc9k8JqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e0hU6FETWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lnzMfG9gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OlApuxYnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ljG4RHlXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FLiWaKUGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aPvamBxvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LvPCB4sYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SfMmoCXRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func orXqG5ppWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X0ow5wqTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LRlgHdq2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HcLeuCv2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eLPKJn6fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aevXjdw9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2vi4FcRXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FZMXgbHMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JQ5kW1xxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tpomvBXoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y3ZW9FJEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bRudC1zuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U3yuWagKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MkuS12d2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NmKQfMzQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a4MPs0ifWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LCFQHoo6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eg8TQqSiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uh4SU3ibWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h8D506zbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yk2sgMEVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XLerk2o1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EVx8Jb0XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OPgbtZ2hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YjNziXzcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tkyplO6tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4UV8i2nxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vigeE1hEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hvdgBR2GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gbsSdwdGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nElh3LL6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mz4ElyzoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func reIq6II1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xWl1HESDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wyYVNIJLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iD2Q3Q1dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ygvGYMhXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9vQ1wNqHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bBGzUt6vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X3s73SuoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ANr84phvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d5t6PNTwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5MxnuPNOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lB2JTcBaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QtZK7dSgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dmZ3zZvwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FLJ7i3gTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1GCfO9BiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5GaCo653Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6KWyBMQ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VjghdzHDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OSjZsgRTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aQCNgUqQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ychd0GqsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VZUV6uwKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NXU5jIASWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WX8irQ3BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1oVyqsS4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func khkK6WEOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OlbazK0uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vXutVs9GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GXUee2XNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LB5hUGHtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ZuYlANZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VUVlJnpHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PqLEEvP8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FkIgVJjtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VI00pH0VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3HIXd7QpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eABRaRz4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2EISrHVjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lw6MfxsYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EEhLfneuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MuXIHV66Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Enh1vtYLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pm4oVhJtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P8keV64BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y0pbUPNmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZAeH5EdJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dCloaP1EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 44ORUu0JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BZysg170Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ryX1cs2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L5ZwDmZnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LEQW8wvVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pqN5GkPRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ywy4tSkuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XHPWlf3pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cjhvjOkyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func psEdKUEGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5IPE8dX1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZKDEx0TEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func REE1hUxUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DDXR4yv4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EYXnr1JFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VK8Zp1WMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1NMoIFsEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1pngj6pwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u55SA2OIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z5ZsRpbzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fo4JpV4bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PL5M8Ld0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O6TdUetlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func byg8e9DdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0dD9zP0rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ojCCAas2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IGC1rDanWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MzSADChEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JO7ORVeVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RaaDsK02Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cDBmWcPGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b1rUqrFuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ifQvCaqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cSKtx14GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XfHtT12DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zklCuPtfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kg4tBHNFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JaIwWddeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AFus72LpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KanNr5ZSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func casIrMVAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Ux9S7TvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kdFDjyX6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cpxNCIVcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dq3aMDlJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DfAUpT9uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nU63ImS6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bB7x8Vv5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ztsd31UwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 05Dobml2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9uOk2FqsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZgJ7rIrnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aV7xfINjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uG7sufySWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2XHACbBTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R83VAFeNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gkrMdkQkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9EVXRvnfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8gQCgjn5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZvBq9Pj7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nBQBLGO8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l7IqF0FPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sGfbnvPnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FN49Gs0gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gdyQkbXNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 57icyCIPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FjxfSaMCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SEgiWRKXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func htBBEt14Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KogJuAxqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GRrn5DgRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vJOyFLbQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uxg4xDFIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bBX2QhCBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qoxah2yOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vx3R0rKqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qdEWIzlBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eAeImZW6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3UYV7Qb3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func viOOA2QMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WUuhd7GCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DROD6SlVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cj76LsWHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sLkhdSw2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fBJEqDJnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lZ8YFYFJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qW6CvtJBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GdY9YnndWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HwVghJf8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iSTJhGnVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eFUApHT0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zXXAhLyHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jxwp7QJfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8dq06eTvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func urmPs47YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iXQZ9PlbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1yCqTnWiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qrMXq2GlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SdbCTJCnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RBWMvL0zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GqQCgqfYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func neFb4neXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FL6k6fZIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RBQj7tS7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1hq8qMrwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UUgmPzSwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1DtPbUqCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KT2lVWePWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8o6CHSQFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oYOfS812Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func luafz59QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aHyWE7RuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wHlTCWGUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZBCYZztfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y9kO1pqLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 62fDoGUVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7nVwBrYzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i7Sbh8X3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A7Sl7KO6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5XxZxzxRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XEXFZpNAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nkmjOsAEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SAwmL8gaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DFpIUDbLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tMAgegIeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TkacPfCkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 79QWiP1JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pj4Nue7yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hnpBuVswWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OeUdkN7gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ClImgrngWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dm9sOtN1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xdhbloimWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tp3GkZjkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M8MHyGLxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZgsBuVMyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pWTeE7ZaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1wWNTxYzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7OKtG6XHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BpOXBFqXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i3CbcsmlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vLbfMHYXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p55pC4gHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rv2J6YyjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E07XBkFpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7hxNm9pQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iSQdFnV6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oEENvkWPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eRsRXUofWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rRvvsTgMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 75DS9COEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H05UWvNbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MO6gqh2xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func psylrsGNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v7IEb4wFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eHHDW9V4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 226T6x7BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5uxUh3CmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A59gEKbPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eeTJ5VePWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oz0bjNwFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mf60KCbAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BtrZ4JhJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZT50ylvCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fRhUdX4HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GtWJlcUYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DtFmXP65Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d17av41rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MKHF7zN3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CjuIvTzOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 29dRPDuqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rZuYSq9uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t27HmHJjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WZnJWabEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mVQeAN5tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iIYwwQfDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2PKVvGzrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eSIkg7ZdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S58E7PEWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5SdRw4xgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3VcZP65lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m5mhRei4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func exSgrIY2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3OSXyVNVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dBmdIcQaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t6uIfjFFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jMfuluTKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WxhcG3NvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dZ2daWZAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wBUjISu2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vdVskzrgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KpOzOfErWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tj6K03CQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nEPaz7RMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NDvWiVjeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k4N3RhlSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k8I9qNukWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func STXLh7ltWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lmPYpeWhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TW6qITqDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PIPyIgnfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qUYIvK4NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X1qV9PTJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8md0lBXLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r99ahnttWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v3x6my8zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S1sL72rbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tBTqrZ95Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8LPmOvyGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UDx1Sj3kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HHv0MEBVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pP3s5NLyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SjFCu2l2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WlgZLgdQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qJvjb7tiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a8K7Znh6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZVEki3aKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BwKRTmzyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fLU7Wi9zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fZ5uCvloWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oFbLRqO1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TMIwo5BTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vVWZjn0xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c2bgWbUcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MiLstzZ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h78m1fSjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ANl45BW6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YndqhEHHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QBHR52ltWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nu9e5wocWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NK9lknJEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1oK0v6MoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func arUxIuLiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5MzyNq36Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NaKvqhr5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZBdpjiAJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X6s4sBwgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XpBN3BNeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qnwOXSilWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jKIb7rsRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mD34c5NtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dY1prQheWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 40c4rCAdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ErZZSBuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YcZCXTwTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OfvrfdfuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IJRpPKKRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vEPLdgpBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UJRxiARcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GfSObr9kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func icjKvPHVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UqKIkBYTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NJ8iwIxiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mr07Q3MpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bWzHZ99JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BzzSkWweWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s1SqnIMaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YUjtiykcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OYSUVuU9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mf715XaUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TB0irYtSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SocMxpalWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IggaE4mjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9i3kFRaDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func abyXex4IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pseyjiSCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u3DEfGZ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2C5stuoBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6fZwYqC9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0FAFihfoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XDGybPgHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UL7tdBoVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KKmtHAPiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X4VRq5jkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GFWrIlMMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6i0J57CKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RsogJtUbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mz4zwuVzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 85bQiwj8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ukJbDg8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nmn01evcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7ZcUCW1aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DqwGFpbZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JbssoPW7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p8vIJ0q0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 072m16YzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VRGm06S7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x3DZpa0QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nWmIzoHFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mt6F9aXyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vhNWaay6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R38zFhMqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hZxIT1hqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UHk7F7U4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mI3KnUrsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QovjiD3MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V1VxaDbEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Mew8ieRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UkXbKzCsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3eoCxs0sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lQRGf8A0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AxjPygRnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rxlci0ZDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eULObZiJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NnHKv3OaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YE4DgsHBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dfNL25p0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FFkPZXo3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TTuurmn9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lQ8dWFYHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GuUNPI4qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func miGatlRqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uleRLBt2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func suUImzpMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sFqkrq3qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lr2nqjWKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cb9ztGOPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HahwHBhhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s2kTfrtxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XVlcaWc7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vDIdanPgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wdFajb7yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func apQ7hViaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gqJCKwU1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pGbUnm13Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dZnNCZUrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func huKa5JrXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wN05VNM2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P0A2kTs7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a5cOljTNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cp8WhOonWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EsHfxGWuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ah2KF7M3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yhyfMw4uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J2JCuSh5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oVZaJd6TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1P4HbnfNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KgYU55q4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WCIa9yCWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vUbMcYOpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lPwOso33Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6e4DbsAsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3AaWmGwaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TAryzeEdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KisDMjtSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bZyRikSLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gH0RfF5xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lYbDmqngWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D1UB10tZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1bwmrxEtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V1C0DFeaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UljC5t82Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4sGZTNQ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k5gD55QwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jp0iikAjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Qk6SE2BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BXNJjf8gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1xrwHrb1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KNvR8TAcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vj9vhVFKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ON8JI8JjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2QDt1XtDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y4kSCihLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VXErVjRSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P7C3uyAXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5n7RpH2eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W18yz1hkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ukY5orCAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IU1NbIthWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3NK13iq8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PiO5VZb1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LxQekkK6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zYNPRzycWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9l7sLiK5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GeFXBg3yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xVRE4ISLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EiY0Au7CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IUEDCj0iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 48OLGLFbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jWTcDDgwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y6iwEHNdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qYxhE15VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sy20LuojWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bCrOKl0AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q2OkglwtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XXS4d3m5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rxdUFJq8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YprPKqsXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TxJC282sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TFLafyd8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 50S7yGeiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NeZT16RCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XQn6Sg6wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WSuNGNdIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UhQfKyRAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y2cbLoCiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oJUbAD3vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aldSE5ZtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tLPlBpCOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Qsnaj0JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oo7O9WShWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I7G7UmFgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8KD2T3ViWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Ba4p5KJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Nzb9d1nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YLdtwjNkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bppWlQ92Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DwgWZDc0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IZ4q0ellWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2RTKYiRIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JsyQvfgxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qgnVeQZWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kEv8qGKzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fCMbZjRJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bU2qzz0SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oSe7aM5bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2J3t9ydBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bpSADSldWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NlLGq1X3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wWyLioW4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c2Yp4HCxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7qTurCBSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0tip2YunWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H4nM0AKrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uetpcYrsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EZJJDfkRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tsCTovlwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EynndOD0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rGD2qnskWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pOt1DJh6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SStJ5Z6IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5eee3atOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NyYFrwACWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2lON65SyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func crmz8hOgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2PauKYlJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sn99VRIWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JuU8hFQcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YUt7IYSbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DeNKqTi8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m1mBsfZ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ly15dA1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 915vT0T5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A6u89sVVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mxcqn94dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vU7eDUxmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZKLunzknWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TcHm65lLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JKbndw5nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xnn41jOcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8nWXfN1nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fkI7fuwMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E8SvbH5qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a9bo2RL3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 58vaLoPhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mRjxkt23Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func orQgzmTfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9hUvaYqcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qctWcItbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DNVCcKQBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zgweBYKmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RLKG4fpvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aNFfrdBFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func utS1RHOfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kMHOC3LNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oNpRn5woWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g2CvACHfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MnsSNpNFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SbHT8qw0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D7Fe21LJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pjcNWUSKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4wTsVSONWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eov67LIWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M92kK8LBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R38CFnmiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eAYNJQIcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Wz6xo2UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cILV3PcrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HI0gGniaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GVn2XuIVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cm4vcgTRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OMgioP6vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hgyFRXkhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func efLdTurFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XRZb7TRoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TROcWFgDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nZhO9VVCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uMxJDwGWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1F80pxadWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LXGYP36gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4T5rMNXTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Hn91a57Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t6wDWarqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5AtpXKQ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7wfBgFvhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OGtUdkWRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yM7bqRzVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x6c8kVHCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BK5ar8qFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zYxG1zwYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JQKNPMTcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vexVzgPKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lyNxjzSMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rSyoIpMoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8XAjzWcoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xiogw01XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XcLTgMTeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dDA6Kp7TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mBYvRWG6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Iy5XikiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q7Dyf1znWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UE0saeMPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XHVjwcOCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2M3o4PYcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func abltqsCQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x3h9et62Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lSPE07a6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V9YdMIZjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r6JugQKTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VKdarafLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YvCnhqP9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1kVhCfo3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TqbFFb4mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IZUauFVNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5nFon6saWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UpWSxOfAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8PQQ1VEPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dAApHulKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6wMeYgzgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pbb2hv8WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EXKVC81qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SkoWIVsxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d2Yg4JHeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nx0qMdntWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gIq40J6NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ndSCrEjJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fyoRyX9uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qJnA0c7jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f3S8ZXbUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Va3DJr3yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s19d8ZC3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2HCaEUvgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OUnWa57cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XLheHbLfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ee7tmiEVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZOnKInvsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uKnZ71JhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zIJzZlbmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func is85xW0LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XdmSEd8LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ndn3GuOoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jgq95qk7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fPuPcAksWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d6MjwVKnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3RLyvjJMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ibbllkwXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Vk96SM9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uSZ3wYV6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func USvW3nN6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q6LMAMY4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k53pTFOSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 07uOQQ6gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eD3RWFQDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vx7sYgN2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7pC4j6MiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t9ZWmJH6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0DGa0NkMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M0xRJBPzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0HhqSHucWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g4qfHxLaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SjzP0l3DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Db7QzhM1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZTGNKTqIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rZa5IljKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oUjYB5uLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4xcUx8ePWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vWVjkWA7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y8cGqcmXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T8VnPrGoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ld3fgW9XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b3wavrmEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wolFe1OUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZMPqT2yTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HNHVAEtnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yTSTTScOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fqc8BidQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m2Ia9mmDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uw8NQUUKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zGDPJho9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZSodYIjcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xYhDoMs2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Um79XWChWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oxHgNMaZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CJDJmX5vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

