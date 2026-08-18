package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "3.2" )

func 6M3rR2JTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TY1OdFcVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fe3zVLzqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sfPD6O10Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FfqhSTtSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xTHrn2raWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yLb4MKh1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JJndfujxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h91dh5FBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N8hCzOMiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mL840geYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nxvjKdbPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GkJrEZB2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4Zx8UWgGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rJJZG6OSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ddLhw79eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0kJn02iCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8mZV6VKsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DaaY3YBxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OKOnYMf5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nPuH6MkKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IsYiXRDRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Dp1xaseEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i19OWEojWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F62rd1cnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s5iHseoPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NnvoAqZkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CtysAkLFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func niJfqOUhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vMVA1VssWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func boA7hNLxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func snGroaTZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5qitjEfCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oF1LLIy3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TiM99GY4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EPOXD2vtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qIR4eCAXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NiCfGLyFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LNEpXWNIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3wwBd2aPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tyBkidYnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hGnEVCYOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H7Xb1KHKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Eh1KsnDzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aSUpKKk0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LPfYYyW6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y7prL7nmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2fWUFIzoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KWb0vdyLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qupDTBbDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DJQgAFPhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AFN8r9IfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tHkkzd1PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jm6nkTUGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 20C3lYiDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Cg1gYbH5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YqKAFaxtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5cOezB68Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g7qFBYaYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fTzVD270Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0rjRLQfHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y3pWeQFOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jFGa45QEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LXYB5c2bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XgMCvUv5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Llmu0MIQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bsgLO1IsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tyQWRGOWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PLzHyDOQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YLjKphr9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oJMOw0pqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aLKu8XUPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ObcqvqDqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jpEuZObdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k7UQkpeeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JSQlJuzCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MSH2MLInWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9oI6pkBdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TsTSwoiTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gLvoVGKzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nur3wsBwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OOoOsER6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kQ26ZuIgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fwY4r5qiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0ctGBmc9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aNuWXqtbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vu2GX3aAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1zL7q9g8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V8vWdkGbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func owOi5sHbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TxCNEVSQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LxlKMS1wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GENtBGofWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N73ZbrzvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uUiYdjDcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q4bzwmqXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8BlSKuSIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eOAzCBL4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u0KxjBzMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 03ip1O2YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pWLRPOfKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6gFf2t3nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xNpcZs9dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kUs8hi3OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 11cKu0oMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wn5lBQWWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fDdC0PdvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VhjyoWsEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PsVhkclTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nMKbqB7aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ijZJqJLQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h7mCEoXoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vT01wRZ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E2RvCHTtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g1wQlnBKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ayvo3ZmvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WQzmAdRwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YKW15UsRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vThtlgywWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1X7deugGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ATOCHZvoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jsvXZ9yaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3LPim2s2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func glkta9enWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B3UKVnalWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ehHmpi4AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5wb0y549Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t3UkBrp3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kJhO1vX4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RZ4PbGcmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SzX0STbSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5Ewe6nQyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C8IPB6sfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PnyYQ4zZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CdxcPvwHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fe9vVAMfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LonB07cRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wt0E7jKtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xtBkYakUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bUN8edY1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wbHXZVPIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ifCLtHnBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dUGkX1cOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hKxp1qNlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SYVMAE5mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5BQpGpTyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tmXBrFV4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HKtdY1DDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MRhwvvVUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4MfErUduWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ekbv2UURWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pD4OQI2GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7NSD52fjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jkZD0PlrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Zoi3DSj1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pfnMpaNvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mlik0HyHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AnFW4IFHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I0cU5v1oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZWyWldMSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KqP1vytsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C4cieAOIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DqLw1HtfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ppHZIvCvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nHxoM8nOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZL91riy9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p5SLG7YmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9jeLq2sHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9TR3BR9AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fvdgtRhZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fDQB59m4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OPKWxa4tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jjFh1Ap2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q6n3Xn0bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IPNKxHHGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U0lhzS3qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7x8FVPkDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sefPdkuqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5LftsNsSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func auamZkSNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZhXx0w8vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b8zjrS5GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8mqSkIEAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XHwQJJKQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1u1DjZIfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p9z1oY3GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5C7XC1cKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OqouTFslWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zrbpWEjlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EtEr6cENWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ulWXI6cbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nupc8JoYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lKJqW0kEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tw5984FFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xpZlVeWyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hoPrWLZRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func on5o1jafWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PW2aG4UwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jyCCf1InWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sSf5fOwoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wZbBiFLVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2fSK2zkbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TEW0720wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0imSjfowWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dj27ypsiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7BzggpHxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kBSpCsTOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0TKwP3xBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MXtwZFTzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uaZnvV7UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uq6mVC2vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y29CW7BzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IXhKI1kEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OrnkRD0JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EdP55t7zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9Kk8qw8fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func seYemZFpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zVikd8XvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7D88zpXJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cQ19N1HJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6BUWqDQZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kUmFnUA0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4CZHXn1VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jPvvnb1CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tuA763gpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6xn9KpXRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ohoo0oQCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AhemCZHhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GySgemYlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nac4bHsiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VL7yNliVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c5VN4LDnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GdWIx5wzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hCtXUXX9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xoNzo8p2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SMgeazI8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FK501IrWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w8OStZwCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mxa1NknfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gUcLHADEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L9dpzcoDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qSmDI5Y1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ixKZEMHsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LMbEvhFbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kw4tp7l4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9AnSQTwAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Af0qCJZlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9zmDrLcEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d1StPSztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VecQtO1bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4X2f338vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1xhFBwu9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Bx9LglvNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QwTrX9Z6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1z2L8dgyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ftuLOdfoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IsMW8DscWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 886LLkbwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8qWfp7riWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OwDduSuKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HdOmtSMjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o17uOixdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RY9ml3ydWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nRoaMR6WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8dUAIVRRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eYBHdPqFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4URTgzdYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Jm1a3ggXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rKC8BMedWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yKGv1iuUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AFJBqACUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vQVrZqJ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FPCFpEy6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U4Ut4w8uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wcJgejsqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HNs6FLRhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kISBiwhzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Kmwvtsg9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rVmcxhhNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 25akJgXuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4el0D74FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wxOJTKJfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BisegyUlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eXE5zWKjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oBOSSvSzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xJCWPX9fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5ovDfvpqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func motanATAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jxuUo3QsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FOYJnCC5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L2HiZ82oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pgl5e3zZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q3mx82iyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mO5MwZocWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fLRut4c7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PF3btJmVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rjavYoj2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oPqsfnwiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mZizpenzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q4WkRu59Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NqTV9bgzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Eaa3bVgQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mspzwa9yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JguAMRbrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4XVopLAAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KaCRwu0OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TGCTuagGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CIKSeXB2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wX99hF4FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 38hDthTRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xwc2YhMwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cpHiDtDvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xNkKKTgIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Zkidq6ehWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YDcNpdfkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Kn0FgXP1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fvkDDniRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AAGymAI1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YVDyhaswWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TiqCEXeBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Hp9dng1jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4IUQFvOYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2ovLPohnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func utnR0ETpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eubAt7aNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qjuAgcsoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rotl4p1IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PgWitl5uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IOA26IP4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wD0hHkGPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5hLO9lf4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8RX9ZE5qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wHGVp8dqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o4QME5rYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gJ5838AFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ns6qGZGVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hbLpjJCyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BxvRXqnGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DgImX5U4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4mfBRIftWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sTAIGX6RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7KyRIFvTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dlBWPQDjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OnHtYT4wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GHJGeN88Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xQXU3ll2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dZnf3U5cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3ad8NmIwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VgNFa7mMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UDU2MxMlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z9MQTzLiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RnCtp0LtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U5xkTaG4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UkUznOfBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B7QbyzqIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C9MlmR1HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zcFF4ZHBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8VrwqYBfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HGCLCcwkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DjT1nctIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NZeJi7AbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VgJ9mQmgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JzxliUhhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rmScugX9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AvSu7Rb5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5AMSEkSlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TIYQZ9B4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func USIMPVTsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VLZjzd0dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dHI7HsbxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m5qIMEeZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VHzvtDpoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vHuMW1SJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kchreEWWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VJk5oOufWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xXYdQF1bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OyiicHb3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SWwPpXtCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JLasxbxoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0Ezwz7IsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u1bBNv7cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wM36nwWxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5ZQqGEclWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BKO0xFOSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6Jt7kN0NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2ZYodZuCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sGJ6JOfhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8wOZf4b2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gLEhYEe8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func excZo64EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AlfYQYFgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NiytpPwWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YNWUEadUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pP5ir18LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fV9r7wbrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qKxTusowWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 78ZW0uzqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yd70qbuCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ikYH4lgBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3rxGdtiIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TEn3FTLQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NCXe8S96Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vnpAMzPgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gJkhQyXKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RQiuG409Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bbvQK29IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qvNDh4zsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DobNDElfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PHovq72lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QdSgwIwXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0fxDrDQOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X524iGJlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6kKk0jv8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sKSaAB0sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pinA8g8kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iorKmW43Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2WmzScgGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ESSiGwaIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2Z9tghRAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HQAI5rEgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yl1FezdOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 57kun8ezWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b2Ypz98MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xMai3O0HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4qDnzY3OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9WTRuFPNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9x39XJ1nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qSW0jIBbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IXAES06lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y9q1Z7NTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func exQYfeSUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5DdtbjfGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HIYOmFh2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cT0CBrkgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HlFGTvLEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5Pxl5oBTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zFA96CD7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CZILnCtBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8fSwoE6YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jlRDm5MkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fm4HMhrNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0Xx3rzZeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hqljrgXPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 45LB6LLqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uPt9BSikWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Am8S4vTwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V89WqwMcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4cKlzjP2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2cXCROc3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9wSENHhqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hW9uHVY6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C365BtxwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HmoDwMQMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2jjYXNUsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8AOCbpgYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2v9gGypOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BrLzwYP1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F8Ae8M3nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 15KeWTtNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TLOEM5BzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0TUjr5pXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nhizm7G2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YYAzGbrOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6UKyh9gkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oEUNbwEZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iKoHwtjdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hSzbzXa0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q4Bo94tmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VYeeJrlMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vdOMdyyQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qbYQluk1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8w46WqrpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZufrSN1tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hX3TLkXZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0NlH1eVTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ChdJ9CkMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oCEESLYyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MR8O6Z8pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jci3qwGKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EemZwD8gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q9tDAHMlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PkvOgYzYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QJTinFjuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lyfRjH5GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pkm6te8BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eUuuzUTIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DigZb2GTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gBCyC9QkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RTkDQ13ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7mTBUO4DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0y1q5B2cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7fFtWezLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7ytmjMBrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pkaqYGkMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NKTR47o6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bI8juK44Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f3xgGkOGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pkoxXbsUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8Bre1n8RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func arhPmcLzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lRngKuSIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 42twD7DeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 626ukLD7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oqYZ1TwlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 18t9zm0vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xEq4dDMUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6GLnnDzbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H0tdIdPRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4sUQEXqxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vFWKcA37Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func twrGF4zVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J5WOYK2zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cJ8jIbzhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NEcze749Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E2es0iwfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WCLvEyU7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lhq442AZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VlKT8bMJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rvJQhb3GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lOzz7pu8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lMfHVMjJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rS9pbfpLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HbxNLZtQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jZamAKxhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yLuKQdGuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LcqWYBcJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8crYHiyLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1xqbAvYPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func onsF3HKRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JMcQxx2sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0qj1xTdXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qZFoD1L6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3XPWKY3uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UorDCB6PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0ZVTfZdDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q2aJQXnTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lnd1ksXfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lqNA6ZCYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y1iyZ4n2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KUJXJaVJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func angfyf9MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lsT5n84FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q1erfnxUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IusvxxyaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jS2oAwAzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u8TmkWKQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yRaE9fzAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8aWaBYK8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9m89WG8TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Hm6eJJlcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4BblAah0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iGUXyptkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hT6qodVbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K0trEWGkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tkL9QRTuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8JJ2rlpYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h4XhZGSKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pd8NjhE3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WEIY2OKrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZRpc0voQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I0uTH6reWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kw7JhUUjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yjCWBFZnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oUWWzgD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VeI8XE8xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FEVxddGqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wY4lF8uSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5aTm6xUmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E2yhVQkTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V48f3mIKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OmGVhVT8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DPdDB6bLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v1hYaIhFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F4fCqcu5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1gOx3Pr0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kBWqBh80Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DMeEqMXvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eTZXlbCZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vicV782kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iHf0IHmoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KSKyD1GnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ABvsu10RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oKtq3K0HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pZIyu8R2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 23bs5Lo4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yUdXz38kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KLLzwbKGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1XoNdCICWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HqNa08ToWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 08ejxMQMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3OEE9EyiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4JIX7NXQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yAVbibE3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H2q63T4CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QLYiFCLtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PfpTCAqFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LVzJURA8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pVJUdom7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EgRmtlPqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9myyRqjhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TkzIjiaXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wBehtzpiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hHqxxewgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xLDiwN0QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CGKqALu8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6nLsJkT6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vInABmiWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5ZMq9tfHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hS3csOZwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bZUVfYv1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PjCTyHXmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QgooSoBDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2NnbNdr6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BRvrg4ESWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o9IBccN0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PKTe4MbZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FfhcxB2pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 08B6f41cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uxqAbighWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func StiY8JWyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xzETuaV5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DddlmyVmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K4ZXyXC3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y1TKKKYoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ctvIbiVyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gWLLv9JuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZRgocCPmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1FYloVSlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CRcejhXmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9anPnsXEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KexQbEKVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 110QLyrZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ej8omNuBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iY40iHLsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gie59PZSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ucs4U5zuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eeTntuCmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZrKGSLF8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mEmw1v2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HbJJCCdwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BQ1wyTwsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 47lAhqpwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EP5uXJpHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mtnDwwskWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OXWVNoEnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wIBqXOeaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z6tVItdfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iSy6ZFpLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0MAHc013Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fomh4BeQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VYrxXBuTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Czredp6bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SCTCQ3lmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qvEuH01qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ldnlFzHNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rm8o0D5SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9x8DrZoBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CCqLjnCHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v6PBLaEhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vp0D1nEeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u9y0gyeMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g2g6zUcOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G1Ja3KHsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YtYPTOqOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yRhdeWMnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hpVIQiw6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W6dGRIQQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pkar1EBAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W79LsHvpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OXe9vFknWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5rsrs4UAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func heD5arUnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LO8x6qcNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KL3V7GTAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n5WHp16TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TJeWODLIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6kIT8O4pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0NW9MyZvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cblzHfWfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qYfzrA4WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RLao3UTHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WUo2LsEJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j4R2AW37Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DAGXUukSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AyoiDBwoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zW6E7V5wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ueaEHhIGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4YN7UrewWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lfu93xScWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Zt4ba9i2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vbmDgHNGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eJ6NZ6QvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aStrvuZQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K2eeOhvvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 007eLRI3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y79FAlCGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nxu7GgCEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SP4th2TcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jNlEuHC7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CjmUwXSKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yKSvNfWUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nDg0KPFaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dByBk69FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5zlTTyTLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SVM74WlSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l02sSkD0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ep1e52ogWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A5vL3737Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AZJweAFSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NHSwN8mhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5RXzCyAVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kCbrbw3hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func neaie04vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zkhOdpA5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ALQNrqZnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LxWJMOxyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iglhS7aXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4CLkNiOYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VAdGZiv7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cPKC1FVxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J4LhRuHzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 54wKwTEvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CEh2niWTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3HcSoyXyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QdaKTrcwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wMF8O9YJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cw4i6R8mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gNITbnQiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Jd3FwfHCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I80QC9TwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4pacfJOeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KoCP7OQvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ckY8S4f2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3TmnjNrfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SoUjD5BVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uOcT1bJ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func boHDYYwNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ff4Z3jwFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AFsB8V4oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TSh7QP7TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pfYGzFTxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QT4mdSOPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XQ7vdc8BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nQOAy7oWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func buhMxdJTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O116GAfzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bimmxRXOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 22tG4G7vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AJCkgDi8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TrYyU0PYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LUi3RDNRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lkxPRn57Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kIwrYhdFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nFzMLyRhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TT8i0Br5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wSEq9cu3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RaTNA0xBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9fTKn00fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9SIR2Yu2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TFOKwnqPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l86HHOC1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func arMa5XjVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pWwa6Ao4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DtDLKEJMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AvCHtoKJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oZCPWcxdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MokhU6Y4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3J7OE005Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KU5sYTOzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p8bMaZg9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5w0AJ7fnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hxJzuIXUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hMyu3E7GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ng7X5fleWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 60VlO2wQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1tJTsDSDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HoiJdPVCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 47xrlwkTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ov9R6wcAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qSehtNxFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lIXfq0nYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PLPHLAgXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RJJtQuE1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qValEoLbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2ndDyUmnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vK8XM647Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S82Eff4mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vCZMyKyuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U2wrNSLiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yc55BAI7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xrKR9p5RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DYvevYqjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uvpECCi4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YDVzAZ3cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NRpzMZNiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SI3gz3D3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qCHhxSTeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GAa6msw5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8U7mykowWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rqX4IaQYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vQnnWvLoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fen5aXrLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e9mpDt2vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func STp1ofjdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XAkhApCcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v7cPP5qMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ILcatWIiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cNPCJFYnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Be87QCftWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eB8nE7ceWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KPz6wmw4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3iIEyo33Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nRfDwJsoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nxC7NMeyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AMaxNjmfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WXy2uxmrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2Yau9gECWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9m68YU04Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IVGhJnZiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AvDQlU0wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0gfRP3UAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 69WMhBcKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func krTAgL3RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jeslYbHUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4a7IIu7JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uHTsa8TaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sq5YdbaKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KxOm9lPgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MOHQfObbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LbTKqBqcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EAly8Cc9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jQ6KouhZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1NGv3vOtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WhvZgQmvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kVSAbCLZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1BDvZzVCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qN6WYzzsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TtXaFAHxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w7yE8QoJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QvPLQnHhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9EZ4INlPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OC23g4e1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7XPjvRfFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G7wNyo7FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VbUZKKziWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v3SVy5QLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9FcaqlwJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sDk7GE1fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UmA16IvgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nnit9nkXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R3vCntnLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KFGLo6YTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nuwxxiq0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NXkUYdCyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xkNqiL2SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PXbY95wBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TqtMS24QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zCFX9KrQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d8ZyqDxMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gBI2lXEkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6BAi6pIEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PKw84WXbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PUWyQQROWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 23HocywbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sacpqiYqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LQrKsCzjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func najLwdBEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1aJfluCWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EkjUdmtIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cSFYvAaVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IVwo8vWqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Povu2QFbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a8AJcZZZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8kcXQdY8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XSOxmSJDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 55B440UtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UEiiXs0CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2HV3F2yDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w7ImrmAzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JS55qZJzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func im1ZdQBfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5qI868MEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CBJwTjztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eDsoYlTLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2E5OFc7oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jzIjN6rJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xhNKokdXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r9bVILwlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q47aPL0rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kT2pTgrGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mfKUEjYhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x7bgODhKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2TK9cO33Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GvcQlXyzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nMG2GA77Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wHIHGwb7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FrszeEeXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wUVhrY0QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rxSAanlBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7peM56oIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pygsitAyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DY9XAilaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U7kLNhrwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ptR8HsQFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1GN5CB2WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sdiNSdUPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func luTKRET7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PisCyRg6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QlaK7FMBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eUxZBEhAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bOcKJ7ACWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qvLnfjO2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qZ6MRA0WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aQJVA4wyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aCfneQkoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8G7X5qIAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lrPTYfWtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QkdgYvfMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1QLpRYeQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cwB7YkqhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K4VkgzGRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m8FYz24XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FHf9vouMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AA0Lk8XcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D86wukRjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RTJ0mBmLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kNbnFptDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2OYMm4z2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vsyjCFhfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pc03NIM4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6jP0ofK5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0Csexx3ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yB8Ucg32Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j9KzowkNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func knfQkwV4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oRC5MwlaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZUYm9lSVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LNAo30ZVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sLbEk3LIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CMTyAm6FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5yzSJuU5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1n6Nmu7aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KyxUBYh0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YPSWymTMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7l3CPvUYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1tNmeCoeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v6hiHb1bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tbz7QEdhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KmBJG3ZlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func apChCA5HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q4H7IPVcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fb4idGDuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qWx8vfAQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C429buumWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vdFYiVLkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c4W2PWfPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U9OcJFvAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XOQWRVswWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MLPlw2iRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KRfmauVhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QGn4YMHQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KVaR46XeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BFxNTDLKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GNNRXQP2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qlTgYTdYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D6ZomAaOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pztghBQLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q6rAp7k2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1ec3CnZHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MrBIYupeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ArJnUZl4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8xXhdAEBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zGrfBnZXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8jktf6uWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UYvTt6YoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ygBRAWtKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R3uXDxBvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hoshfHvoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KesZC156Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lXptldCvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9RaCpv3DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WRxA0IG3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5BitMCqPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uDoShQaqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aHDAI4fIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lmREshFeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aakiRUzqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zlGF4QLLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zDYNZPDiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tft8IHn4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fU4mGxzRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6L4UgKU4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ScV9qboJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eJPfu9NXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func odkBcOz8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DX9N9hOLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ztPTNCDrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DaQsF8JvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func maJzIcvYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func whRPNGaeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 919lGMipWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BDi75udBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8KYQ888CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zhDQiI9UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U2TMmkvHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ouVeejoPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GTBmsaLiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4XXG26k9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8BULvl6NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EkubimBRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2Em4wiPLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AUqDFqeGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JdzFRjGYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bW8JFknOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eMipX0zaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ORaNpSnxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8of7w04SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sDH9CDGbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g8xamvKCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Oi0ABdDaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func orZeEtNAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func db9h3w4MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hNQydqiPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qQT2PaVPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LTkbzZ5SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S7jk5ElwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zTk8TTcJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4rrL6ABiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aONQ7fz4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qS7bxzGzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Kzigjfc1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sN6aay5AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pydr4LpEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5tmst1kLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oancphdrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BgtVMF1uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FSkjqcm4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i45XBDmlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cBczJ87jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z4MTY4msWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mw5tMthCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NS2BcToLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lThdvnCxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fva7w7oFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qcx7vMv0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AT7A5xcyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ILDpQBksWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6KDrGNbvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kGX9TbTfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xfTzt1DCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ONTjEK59Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A9xOArrCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vgBrat4wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HJ1RcIltWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YGuYpQ68Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func egQF0zFpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YFkezg3sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func slSKWnwIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I6xiZ6TFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2CqL32MZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0l5enSNJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fnOTEyyeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0SVfyFWgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fO4jdZwiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y87brusxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TRA3KYUAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1KyqWpEfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gCsiQ8aQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q9xr8WC9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xgum9K7aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y4LcZAytWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sJB7YVSaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j0PXFf28Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 29JpttIIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jiyrVvPdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5ras2DcvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KZpq1TAEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zfwQ5AIpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FMFN3Rt0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func McOJZzPOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func crVXcNYpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gzwjTFlaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wschHqeDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z4zh8vGSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func myCZz7mcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PaoUdclFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7kVRsRrmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R3XMbHWJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BwvwIT11Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uwfLYrxFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func an4DD15AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 289
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ozyvjqOxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m70Dct2JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nfFncdcBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hvKvqESIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bdS5lxCvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kN7VXYGDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QL1zxbWAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JEeWN4WTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h9cUJN8yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2YpaycpxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sxME8gOCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aMogtbnFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j3zh66K9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D9v8lxrkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 83p80JiBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1hP9ExchWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BC8TkR8PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2rSy3gi5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func li6SagAFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ifDTntnXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ltMSHYXGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fMB2clWcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KxZS88kyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pqeZnZb5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M3kMyN2rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j3cTaBzyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HpFn9YgjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QDQMZzqHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ztbVcxHwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func akrLe49KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nYmHsOqZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zEV6Uq6oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QAhtoiK5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 97Z6s7r5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8S4OlrJ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Bk9LdXuJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tjPGiBaTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TCVKp2anWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9kJuhZkyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5zaaLyXLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YAqnXp5XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qqy6zV3aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gttweWnyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JhIqtd3JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ss6AXGjGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G0pCuC1ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func idLENxddWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 86rCmVzWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func syfPZdtYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mDdifV5aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4kUPPQ8rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uDIdMEX5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xqihrbCjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ntJLgQfqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4gs2N5YZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vVZEHV1EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N6IFCUtJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zBWeaLVZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func weeMX1GrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j3B7tnrZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f5QujIICWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b0QvOF8DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WfPBsZePWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 246
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JYhXftCSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func frBt80DoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UZoXvlB4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HpWizDFXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QnQCruAEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M0xlUmybWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qSG1tK5TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W5xaKEmzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 19K9AlvJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mugtl8nMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p9HT4IalWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MhK5VJvXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7cSQyevoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DpyZaqkmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IBPPu2LJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7kXav3vAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MAxb6L6YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a1ggP0ZHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vC75xzj2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PPJzA5ThWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nxWmqwOSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rl6H9rI8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2nhqe8AcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aXC4xjiDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DGd8yVACWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iSGdH3v8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func isAoLrttWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q2182qRNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PXtos8miWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HxbXaZ9KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xaTRfJTkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IpkoUbnoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qq8T0guYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S1vc2S8IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XcD2fW6XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JazVvEn0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xsCE4XwlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fM5aJjF7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C8SByOjdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k8giiLabWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SNrCF9wUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LDxcdUaKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XayvfMzHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mUMCeEeSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0FqrlRwqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eFZClUd2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dIQqOTb1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7g6dfvucWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xP4VwJTPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OyFzKzmyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vzEYbb2FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KflVlvRTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h7kfseHzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XSqcH2ziWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gvpKlYoKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dFzeM9EnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6VyV7i4LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ShR4quJOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yUhg6l0zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ccNBLWDCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uTR9rZerWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lkh8kIn2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IRw8zE1wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ku66vp2JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZkOgbeBHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DvOVqK4NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8bZjAU43Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kW86rQBPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RYGkDwcMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yqdO0Ij2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pegjNbDuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GCGjDKmVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fTWA2VYjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 72Vq29aoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rdV7wBbEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5ucHIk94Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dzLXg55vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BfQAwDNIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XQaT7Rh5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8MvOPHzeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bprEX6gPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FOI0zZsWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vFeUQlt8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pCNC7uL6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J2IPyXq7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qimXILNEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oNGILSmXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0UwnTgkpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XZsBzuYSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f9kPHEvSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qh5pNUJxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kcMJUZJeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1vGhuz9iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hyhmsJoiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IQYCntLDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func thrxRm9hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2B1bTd6nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qZ6jN0snWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BakVIeb9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sETtFQzPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YkspVWg2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9CTB3blcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tIQcAokVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x7ygxcLaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uOFVqJfTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fAUDCh6lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6j57zIQtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IjstTeZJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QGHqNa7VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5I2ql7WBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xA6fzWr1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xFEVJAreWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jDab5Uh6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r0F39qP9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8OVDE8lwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1tLtixsxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ziBfAZYhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2zGTayprWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Pg83dUfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gCizmKJxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7VV7zpwaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X2kh2S5wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v6MKv4wTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hyA73oVzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DsgZVzBqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KsdKLIcjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zod54zQNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J0410ojCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tFYSTrPXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SJHPicXaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v9TDrEkNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FrEPea71Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func je4dVickWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dH2zaSq5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MRUjsiPfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vmIlqVwhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jpNAIXEhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ijUZrvmVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ho5EpOKRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6cUvkK8tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PB1lEakzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EHxNWpXkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CcbjSmjtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LKcxzoJhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VQhMglIzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ScepNqsWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JfFlxbHqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zafBF39LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5OvPNvQeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XBcYxyscWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kIzBIcLAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R4P9bYptWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mm7ygCBhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HXOvDcAcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 32RQFbPuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3JiTwVzTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VpKXkVbIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z8IGohNtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ILSN91BKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NoUKiQ3UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FRARn1vtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tK0JhhW3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SRVSWgZrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dz4VCnRJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4X7ydVjoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dFBXQ4F2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yYiXDCpDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jcZstdxnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hEXjvYlLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dpluTw2IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vxBysc4QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D1brCgYuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BZzSQO00Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oTmFH63bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hjsJoInsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JV2CUdTfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6iWZ6XXpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5Zb0ivh7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func suNQ3196Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jSEdVaQgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U0OMIdVaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8dHEGXkOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xcQQXV3cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vyXba1fdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Krc73Ti1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0ynDgt3uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F0j30pU5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 02My73fXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k3r7o2t6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ih7pqvCPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ab1u49BzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PGM0UIm7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SUpQOAmJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GQ0JmzEBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N8wV0uCXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lvGdJbyoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o1Jc3TIrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n6IaG0oyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U8Xq6usQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yoZd5skrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HIooPhvUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1HiatrI7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VOXaWHX5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4RyjjFOeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1v1UqwDdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7KtNjz9tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func diRNvvLpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nKFWoRcbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SWISsgK6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WV83B5JJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3wBcFDhDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ROvraj9yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MuA9Fmt1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h85LchZfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ps9jf5gnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8jKmA7PNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oLTKqJxeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dsstpkizWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c3HiOSNQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5UT00rowWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nh6HYScsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1E6Ta2mfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UdH8bSruWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6ermXj9KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N1xx2f6hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0t7XXOGIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func njnxbXKrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Xbz10eaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hxkHc201Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k1GTjxwJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BdA0tCYLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CWMxDInIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WHmlxfCSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b90LaLbDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 39sW3hJlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yW6ZftzCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Hmx3V61fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tCfzrF9IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MQOwNDzjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nslHhoBSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oQcGXZp4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4C6uPV5VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7empu9vDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 05bmciSkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WLYEZNCuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cuG5WwniWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pUYIMEdmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZPJ55ua9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MxT9FWiVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sA7ZzIFUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vJIwK873Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bPRw6EHaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9OLxb16HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Tej3xCyvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fwCn90aYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5Uo7TNBJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DbBEN3BEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H4syOEXgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Os7TZFGuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func beBhgop2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cC9W5KkBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LcgsCxxnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B0dBCm96Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JJ77DcRnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zvvagO92Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EKJEVVltWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cdt4Kk49Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xnn7LjghWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mntVw3XXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OKEQWAdGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NOzUxYGGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w6cf4QIdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tsSp6RYKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cxI6sto4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qiRRarLDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q7aPSikNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5etpCh9oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R6yVsbzcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WMbmgVkRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PcQ0FkCxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4PtLGDyhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FA6q4Y2OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wbL4Pe65Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qfPc3MV3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bWW0iznTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J39tizCUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BiRv6HHtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AHVYEEZHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gkQY7sooWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m66wcU7UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zfGN1iViWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xcjvPUmkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h4jYalZJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eApv1LDJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vE8eqVUEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6sjYaY0oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9A1K88qFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3HCvYAGGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pzlt2q5TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TZDA9YAYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m5XuddlCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bfrcWJqhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1PvG1zt8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2Hvb3aDZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NI1vBAjGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xVI7UC8FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AYFZc0iRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vbasiFPKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C8UjnnW0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dRL7lNQqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c3LwM7PqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N3fSr7vhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7IfgwXx9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DgUovQsYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zyRlDJGAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func csJTE7G2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 90qTPZybWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5kBBuFckWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ISpdbg8eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TgKjebxkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6oAtdK6ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func F5Gd4xy6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a7WhtOKEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5IKxFalqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bIaHp343Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4solkPU7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OASxOiEtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4BxZxxlhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ovhlWCW7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PWet9lqLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2QBrPUYIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ju5e0Zg6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func olmxnQbcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i6hOpFgBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1oo1fpUaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fPcRhrBoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3EzCJ362Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rXQtQmKWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1ablbGs4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s9iLiJw4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UxTS7SynWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TxhqKrZ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ogYUx79fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ca7c0Xo2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5aoUvX52Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ErkzxoilWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zDy0UgH6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tJ72wRohWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XAqHgCcAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mzGcb7LgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C3bpQRegWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SKXeslQ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n2QF9GRdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R8TckXSJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZCQpVIOQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lKWx8TU9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yvzif2BsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fOCEeofOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lZwilx2MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TJ1ADP1HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b2QVzNdlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ioQnxX9bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7WPv0PO1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ExSTV9m6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kZeyyOfFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ub5tEX5TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i888sVJcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xor4SVrqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5agiVVToWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DAAo3pVAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y8pnNoUKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s9nQ4ia6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V73lUS8JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dJIWtBaGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EZSucE7nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7nZEuPhyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yca3cJaVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ADED7LoYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VMpiwIrKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J4nhCQ0jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pJvY3PdmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dwhWcQ32Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mL9QrUXFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lcscCcFxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func INfXDj7HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wx2Wep2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JzrbmcZOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9bEGZ5nmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func veTOWNUSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qrJ7zuoQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XVwM2CcWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EXAIcRbQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8F7vPkO7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j3pyuENpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vIFOhnKdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OXv8GPFVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lc1WxiC1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0Dj5cNMmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6K9NjZgBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zgDOlWkhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YvyPAZv0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OzUkV7mQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8pYGGtyLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pvnHCsGMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wehVNo3xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KI07RvViWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZxNReWzGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oa5pCsrqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sUw9PfM6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vdr2p9oGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UyhSQJ9OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func th4UHiTgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ILvSd6bHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SN4uRB1lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xLt2X5b3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jLH72dZvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vXnM5FprWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1saPcHU7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hewFj8POWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SUC5nn0kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j15oU6IFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LBlaDSQiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VsmLaXODWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4AZhP0kKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eU1AP72GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UqAb5DTGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yvi4plXSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MHzBOAJuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gt5tJcg4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jG1u0EBIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1cmlZc0lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JoPVx62LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5jwwNUGrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func daok9A06Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TP3z6gzYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n7hdEqPMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i5yHEdRJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RKGMOs8tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5EqQTRfDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rGlaLD8pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ov95Idx1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YUvpk8u6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HRmjcblWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sK1i9s1LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bTH0EjUgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qeFhAly1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6FQiGAdxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1suIkDTYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fzdkf7BtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ClOksSGIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3RAbOfYGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LVN4NZehWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DwzvQ3RLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CDSyw4cIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WXYB8dCgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l69GSqLwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kvEX84GlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wxa5S4cxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1EkVX2EWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P12IaKM2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FUkFUr5PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XDBBZ7VtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FG0qvs8gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LkPlg4TZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HBm89ycyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JdSF704yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1qJzGyaEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gInveBLxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bJnDFxweWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MKXpsEhvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8kWp0ax2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xTVW6SmxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Jxat4hAUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yCOr2cc3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f8iPRsuhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gKNXfUJEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jxZoBtDvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PF2PVUJzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v1EoREBhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kwlVm7rdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LlbZUIxAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SmfG4ygtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s1aF5EeSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UbtZZBCJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YiJc1XI6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9pAXCZbzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func weVqOQawWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Cdn74oqlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A1nHZWNFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HddXuO1EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pgka7IUUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lKWOMLI4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fxnsbMnPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L7sUWUYZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v1bSRrzrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Cz4xEVmyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vTnxEfKhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B2hKuzV8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xauAj57IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K22fWLlZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m5XWgl27Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kHqWIsYDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xiwYXdQ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pXzpZW37Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jrdehIi0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mmH0bo7IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qyENlP3sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TRHLwfleWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eMCgVRGkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x0GWkSmwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4T63SNoWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kg1nGuAsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ma6ZPRFHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7d03Xan5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 882pn8FbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kg0HpaMGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FYqI1Lv2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fXMtA1AdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OPGEDysSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7c43M2tnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VdrafN66Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GueLenlNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KURu8b4BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QoEt8x4SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ya1sdNsyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RKygv6cMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QOOk9HxjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vJX170zTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a6tIQniMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9QCWp38SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ievx5nxsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xAJPZXjwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JL1jqfeZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iQRBcQiAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9o3VGYHuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hekJslvtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WH9Uf03VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NLzzIRXnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 93h34LXbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GVZXuOozWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GP0U5pzkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QVccXhYYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3VtgfnbxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func a7xmB7JsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uhxStEDYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9Q937Ea5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qR4GWFATWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EL4DszQNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bDBUCEt3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GGm6tKpeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t6pYSNVhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9ZJRHNlKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BOfBsfj7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TlxBSGEXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func heil6wvgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wk3BilUxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DA2mZh4cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xYf4YmpeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R1vgKx4HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OUxOSSZZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n1Pnb0CYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OfcCLWFMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S7Vlsex1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8D9mqGm8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ToxPhR2GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vAMlQzKIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VYn1puZHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rePLBhfjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func biKP6YvvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EwFu6omRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Hmm1cbKCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mrIjLHlyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ETE5JPRFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UXg7tygyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RbTQDKNXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pp5c4N7iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RusC3aAEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func su1Dc1LoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OSmFCtT2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xlt6ftFyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RmPtpQquWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rDzoXljGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 230
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lmrM1y5lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Dd3kvK06Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kulK7Z1HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Hc7YghNAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sOb5whAdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nuaTqadGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CL5QVXjtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y3D13NRJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fY9noyoeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6NScvpb0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KNlpGWk3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DYuKiVM7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e69k2WhIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3h3bgf9OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ttmB8NiTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MW9juFGDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9rIABLlwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z2e9gmFrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rhyVVJ1AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l13yMDaZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K0dkhOgMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NAtogIQxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CqWVb0wpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AaYnGsHZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3EiirThUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uED0FjKPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6ldDqB2HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jbZ2HoMuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tL7XDnlyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AG19tyGUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fcVLt5rkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jAuL532uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qXCVIRL4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IOpnynimWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fLSiCU8fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Os3ulGtpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TuZAsXd1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E2p4XUNeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wESTMqLGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EmdAj3CXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HkqVFS75Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9ctJILOuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qcXx0xJvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V8wjfERcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H46ZDHjVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XdCyMs89Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IhYQyorTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ped0ZlkLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UQhzetqdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WWoJneDSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jhsKHNfMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MZxDjYSPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func u6XLeNPOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T0Denov5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JO97C6HwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 18
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vu5sC4aqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o4DqZJb6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VA0tdpUHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pjsz531eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y7sabXGKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H1soWT62Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mIv8ayF9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1jxeRVTpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bli1rXdRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1JJHqkTgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W6rEUyccWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cK473iniWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0a6HWig1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uwc9WjZ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J9WAGL9RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bj0ppsS7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HqbUrjSeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gsaUHzSEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A9oOO5e4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0riT9GtiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pXKWUEeHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aa1GPHwCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5C1XNMnmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xeq0eWJPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BppfKF07Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func clmomdKfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cXC4yBsGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DmULeS6IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qTcQJjY4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IhJMOdHDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZbDw9gRZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WU82Xmg9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8eKpWb9JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pQ7XHYsdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y5FSvJWiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G6V3cqoVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q2DXtGibWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Faxqpm0tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0tbl5NMkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jPhwW1KLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b7yyZHqeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 39
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2h1aD2MJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U46N0DGSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dmwPF0qdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IQVrAz6GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fzlJ863XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C7pjdSHdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zm6xR0kMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GQUBoWqiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1p2jOrW9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B7G5XB2TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WkmMLPPAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n3OrKGbQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OOkf4D01Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func imrfxxfPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JEsk5LhKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bJTf1L1WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kOXEnSkuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uN3VDxbtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eMJYUc5iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jK7ib88gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U7Ls6NhmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q5b9iFqKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lSCWy6AIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4bIrZ6ekWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fUWQjn8SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xJhdNDaKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xxhfGYzGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7PVKGtmRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5x0Y1OD0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o5kt11jEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qqf6rj59Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 14R5T59GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VIjhVd8DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Yaec38GeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mmSRiCCTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vVJdnhNNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YimpSKrXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r824vBAlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cMQsn9iAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CzcxOjARWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y9yvCdMuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wux5Z9PaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LNc9kDWPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n6oOGLdMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NyCMZSVNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rcmVRaZdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TuuLvmiXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xH1X23NjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fUtVByQpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qcnEMSKtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N3ZkxUubWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OacBoyTfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v9hnPZrVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RzsIDIKlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kXs7a2LMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HeNikwmAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h9c8M19wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MbpWHzLzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QSuyTkvFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5bnHb8IWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GMgMqcDKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CYFuFaN1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2UoNSxd7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iMWDovR5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jTKmXmldWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sSyW3kj4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H7mqoX5eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bRjSb4sUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g7rfnFerWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WosHtszHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 35ro6A28Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mx5e9R2ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AcBjtAFMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qftGjbXVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3boz1W4QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xrm5hRSzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wCUKIpV0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DhigRP7yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JS7yRO4eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2lO9XMngWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YjO6m2AfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vvo9NQbOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kvg2P5viWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g9dDYUQmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func doV3Uq1aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 87
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YBL1q4fZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CeU7g8v7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o5e0Z06wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XV7qXfpDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oG34FtlBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yEk92V4OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func abOdLAbHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DqgTGpxzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func veJlvvLKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EZISH0qUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RC2sScYAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6R3qgQ0aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AEyeKGQoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E5qQzAGmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fDbDZFRJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D5cWbZEDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b6Y6E1zbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ejrDCSJQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zo6fy5qOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2xQoEZ63Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o1sjdtYrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xaab3l5eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 68HfLibuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4NTa1zvIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O1ggJPBPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m4bbcYlxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WEwS2ua6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zYqI49bOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lbjzt3ngWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8vdjDVjUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Eqkz07GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c2FqtbSeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yOl3BQh1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m5K3A0anWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M74zftCPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L9dNfngdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X5245SjrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gdxlmJAbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wTd1HuC6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tADmMLwyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qpceiUQmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8kCaqtgdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IgMTcB1YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HyQILRDsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rz3nkUlNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func emA1LWQnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func crpwpdENWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WkwofUkVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 07EwKxLKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WvWNzYguWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vrMpX4eOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tk0wGpOfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Moc5U7ZdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6XY9QZYeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eKAFi7xiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MoI0CqdjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KsYPJWVWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xRTXD1g6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uvlZ07PTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sNCHe7s5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yP9YIDjQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4bzEScMFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func odtIlMBNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wHnO9hXoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RBz69CKlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3XDCVYJnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 15NbeDmfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GX0wQJEaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5qns6dQsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7zMTsGrcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E5KsGjS7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 24tNCjKCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mjSCCTN9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IK36pdzwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NJuFo7N6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0WxbR0txWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xL9qChY4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6cNlwzDAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y33RCboKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 266
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9yxq760SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NasveWZQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func on2UBkfkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1zDOBncyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fjb3WXt3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GhZQL2PyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SCRmOj3KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BqxQ0MsTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 95
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vNUCdQBAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TQMhwMVQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qq9TMP1mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pf4Txn8YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N1x5n6FvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3oKEUNLcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZFRGCKwvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5PIYKhzcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YBbaQHhBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4hcpieiNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XTyrbjB9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hoHYeblWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uyhjso2wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 41XwWoNxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TjrObsdUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func exyU7iiIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GXCdzkZhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eeIviifJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BTcvb1WlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LcBpW0V7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tarEPLq8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KT33sqNZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 84
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nNr2OV5ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OJQr0avAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kvzfQhtTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GPBgLRHuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qHS0C4AmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WjvWdZFlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nVOaIt4zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 293
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G05EbtAuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kh0WX94aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YL60N6vqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L3qKSK3QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GtWPkuToWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2ss5nppKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FAeohbNZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C08n4Z9wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cW8762vUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Jv5aFTeoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9WIJuu7YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7qmb43vNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5PIn6SZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Iw478vbaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zKdmnyqqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t9GP6tizWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 79
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JAy5qSoYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bl8QIwvNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pxR4fYPMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2OLNDcgmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wg98ZnoGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TKTCe8UVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uEMyeYqdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KqPeZbUQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uVurYdH5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2yjBl2cHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iXEFx8CaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gt2Xp9LKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8sIvPCzlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z8inqaXyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O14CtHUjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mcp90N2KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gkgwfAHsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mY74SFMtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tiI0kPi5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 258
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EmjhvQCBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EALir9VRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fQm9PsoYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tak3i33fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 265
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fM6vKU1XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IEcXfTn0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z3Mdcj5LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AVeGADJHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mQiGLMQ0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PanSNhoQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6nBRPwcXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 197
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wpuPd15VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Bixps3hgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z3QyU3qCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eyAI2B1aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZxkMb9hRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mHwSeedhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LCnXCyxOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rHrLPA2LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UWIuOl4WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I3qh8I2aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lIyDQM1DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VavW3IevWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rSeEkLJQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yGWmkNcEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d5nvq41iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cBbHtZqgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bVzq19rEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NnGvj7DbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IvV5SsM4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NXlKff07Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8v5XkVGJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sIDidJXuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vysfBtnRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Lo7XLlvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tTfb7cbVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qeXlKVV5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xOUjRQJHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tEWiUq4GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ukEIhmNoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hvy2Sq3uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SG8fpWgcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v6yRRNn8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uiaKvLqjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2lT5U31OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sq3QGkQRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0Oy6LUh4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kMMFbXKdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 25PC5RFHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func szwVf38pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h5qwRguLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RijoDTWmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U7zd5liLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IY8MZlk3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9YYCrMZ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k1pdan4OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rm0dHQFwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8NtNWVUVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JntY3zSLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JSoVOnhjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZH2QctyvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 97OUDMVQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f93sSmYXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SJLxu72wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9Ivnw6PKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func paVODy7rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9TavXCQMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TxApneV9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SU4ntp3aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func acPQwHC4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lpGIIBewWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RHQUQAQBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xVNfvVusWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VDbqSBX7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jQXoXBkpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lGN6Sl9jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HHqtXVwHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YF1MSLSKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6rxQNLYNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

