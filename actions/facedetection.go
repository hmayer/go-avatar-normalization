package actions

import (
	"fmt"
	"image"
	"log"
	"math"

	"gocv.io/x/gocv"
)

func findFace(processed gocv.Mat) image.Rectangle {
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()
	loaded := classifier.Load("./resources/haarcascade_frontalface_default.xml")

	if !loaded {
		err := gocv.LastExceptionError()
		log.Fatal("Error loading classifier", err)
	}

	faces := classifier.DetectMultiScaleWithParams(
		processed,
		1.5,
		5,
		0,
		image.Point{X: 25, Y: 25},
		image.Point{X: 400, Y: 400},
	)
	fmt.Println("Detected", len(faces), "faces in the image, getting the first one")
	return faces[0]
}

func takeFaceFromOriginal(original gocv.Mat, face image.Rectangle) (cropped gocv.Mat) {
	height := face.Max.X - face.Min.X
	width := face.Max.Y - face.Min.Y
	avatarSize := math.Max(float64(height), float64(width))
	padding := int(math.Floor(avatarSize / 4))

	avatarRectangle := image.Rectangle{
		image.Point{X: face.Min.X - padding, Y: face.Min.Y - padding},
		image.Point{X: face.Max.X + padding, Y: face.Max.Y + padding},
	}

	// ensure we dont overflow image border adding padding
	if avatarRectangle.Min.X < 0 {
		avatarRectangle.Min.X = 0
	}
	if avatarRectangle.Min.Y < 0 {
		avatarRectangle.Min.Y = 0
	}
	if avatarRectangle.Max.X > original.Cols() {
		avatarRectangle.Max.X = original.Cols()
	}
	if avatarRectangle.Max.Y > original.Rows() {
		avatarRectangle.Max.Y = original.Rows()
	}

	cropped = original.Region(avatarRectangle)
	return
}

func FaceDetection(filename string) image.Image {
	original := gocv.IMRead(filename, gocv.IMReadColor)
	processed := gocv.NewMat()
	defer original.Close()

	err := gocv.CvtColor(original, &processed, gocv.ColorRGBToGray)
	if err != nil {
		log.Fatal("Unable to convert to color", err)
	}
	face := findFace(processed)
	croppedFace := takeFaceFromOriginal(original, face)
	avatar, _ := croppedFace.ToImage()
	return avatar
}
