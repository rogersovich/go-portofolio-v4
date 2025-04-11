package upload

import (
	"fmt"

	"github.com/rogersovich/go-portofolio-v4/utils"
)

func BuildMinioURL(endpoint, bucket, fileName string) string {
	return fmt.Sprintf("%s://%s/%s/%s", utils.GetProtocol(), endpoint, bucket, fileName)
}
