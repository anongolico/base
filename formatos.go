package base

import (
	"fmt"
	"strings"
)

var (
	Formats = make(map[string][]string, 0)
)

func ScanFormats(r *Rouz) string {
	comments := len(r.Comentarios)
	commentsWithAttachment := 0
	for _, v := range r.Comentarios {
		data := v.Media.Url
		if data == "" {
			continue
		}

		// TODO: if attachment is an external video (e.g. Youtube)
		if strings.Contains(data, "https://") {
			continue
		}

		_, format, _ := strings.Cut(v.Media.Url, ".")
		Formats[format] = append(Formats[format], BaseMediaUrl+v.Media.Url)
		commentsWithAttachment++
	}

	return fmt.Sprintf("%d comentarios encontrados, %d con archivo adjunto", comments, commentsWithAttachment)
}
