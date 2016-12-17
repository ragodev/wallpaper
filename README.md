# wallpaper

Cross platform Go library for setting wallpaper

## Installation

```
$ go get github.com/redpois0n/wallpaper
```

Import in your go source file

```go
import "github.com/redpois0n/wallpaper"
```

Set wallpaper

```go
func set() {
    path := "/home/user/image.jpg"
    wallpaper.SetWallpaper(path)
}
```