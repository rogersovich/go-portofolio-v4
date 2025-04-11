package upload

import (
	"fmt"
	"path/filepath"
	"slices"
	"strings"

	"github.com/rogersovich/go-portofolio-v4/utils"
)

func ValidateSize(size int64) (err []utils.FieldError) {
	// Optional: File size validation (e.g. max 2MB)
	if size > 2*1024*1024 {
		errors := utils.GenerateFieldErrorResponse("avatar_file", "File size exceeds 2MB")
		return errors
	}

	return nil
}

func ValidateExtension(fileName string, allowedExtensions []string) (err []utils.FieldError) {
	// Jika nil atau kosong, pakai default
	if len(allowedExtensions) == 0 {
		allowedExtensions = []string{".jpg", ".jpeg", ".png", ".webp"}
	}
	ext := strings.ToLower(filepath.Ext(fileName))
	if !slices.Contains(allowedExtensions, ext) {
		message := fmt.Sprintf("File must be %s", FormatAllowedExtensions(allowedExtensions))
		errors := utils.GenerateFieldErrorResponse("avatar_file", message)
		return errors
	}

	return nil
}

func FormatAllowedExtensions(exts []string) string {
	n := len(exts)
	if n == 0 {
		return ""
	} else if n == 1 {
		return exts[0]
	} else if n == 2 {
		return fmt.Sprintf("%s or %s", exts[0], exts[1])
	}

	// Multiple values
	return fmt.Sprintf("%s or %s",
		strings.Join(exts[:n-1], ", "),
		exts[n-1],
	)
}
