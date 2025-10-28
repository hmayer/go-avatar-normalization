package handlers

import (
	"go-avatar-normalization/actions"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"github.com/google/uuid"
)

func saveToDisc(file *multipart.File, extension string) (filename string) {
	randomUUID, _ := uuid.NewRandom()
	filename = filepath.Join("./resources/images", randomUUID.String()+"-avatar"+extension)
	destiny, _ := os.Create(filename)
	defer destiny.Close()
	io.Copy(destiny, *file)
	return
}

func AvatarUploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, header, err := r.FormFile("avatar")
	defer file.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	mimeType := header.Header["Content-Type"][0]
	mimeRegexp := regexp.MustCompile("(image/png|image/jpeg|image/webp)")
	if !mimeRegexp.MatchString(mimeType) {
		w.WriteHeader(http.StatusBadRequest)
	}

	extension := filepath.Ext(header.Filename)
	filename := saveToDisc(&file, extension)
	avatar := actions.FaceDetection(filename)

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("Pragma", "no-cache")
	png.Encode(w, avatar)
}
