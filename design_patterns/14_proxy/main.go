package main

import "fmt"

type Image interface {
	DisplayImage()
}

type RealImage struct {
	filename string
}

type ProxyImage struct {
	filename string
	image Image
}

func NewRealImage(filename string) *RealImage {
	image := &RealImage{filename:filename}
	image.loadImageFromDisk()
	return image
}

func (image *RealImage) loadImageFromDisk() {
	fmt.Println("Load from disk: ", image.filename)
}

func (image *RealImage) DisplayImage() {
	fmt.Println("Displaying ", image.filename)
}


func NewProxyImage(filename string) *ProxyImage {
	return &ProxyImage{filename:filename}
}

func (pi *ProxyImage) DisplayImage() {
	if pi.image == nil {
		pi.image = NewRealImage(pi.filename)
	}
	pi.image.DisplayImage()
}


func main() {
	image := NewRealImage("image")
	proxyImage := NewProxyImage("proxyImage")
	// render image immediately
	image.DisplayImage()
	//first load the image and then display it
	proxyImage.DisplayImage()
}