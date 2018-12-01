package main

import (
	"image"
	"image/color"
)

 func main() {

	  }
func drawrect(nr *image.NRGBA,dx,dy,w,l int)(xa,ya,xb,yb int){
	for x:=0;x<=w;x++{
		nr.Set(x+dx,dy,color.RGBA{0,0,0,255})
		nr.Set(x+dx,l+dy,color.RGBA{0,0,0,255})

	}

	for x:=0;x<=l;x++{
		nr.Set(dx,x+dy,color.RGBA{0,0,0,255})
		nr.Set(w+dx,x+dy,color.RGBA{0,0,0,255})

	}
	return dx+w,dy,dx,dy+l
}