# Go Image Search
[![Build Status](https://travis-ci.org/herlon214/go-imgsearch.svg?branch=master)](https://travis-ci.org/herlon214/go-imgsearch)
[![codecov](https://codecov.io/gh/herlon214/go-imgsearch/branch/master/graph/badge.svg)](https://codecov.io/gh/herlon214/go-imgsearch)

Search if one image is inside other image and where it is.

```
$ go get github.com/herlon214/go-imgsearch
```

In your project:
```
package main

func main() {
    imageA := // your image.Image var
    imageB := // your image.Image var

    // This will return imagesearch.BestMatch struct that contains `X int`, `Y int`, `Confidence float64` (between 0 and 1)
    result := imgsearch.SearchImage(imageA, imageB)
}
```

![Match result](./testdata/match_result.png "Result")
