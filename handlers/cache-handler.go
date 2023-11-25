package handlers

import (
	"fmt"
	"net/http"
	"time"
)

const (
    ContentTypeJSON     = "application/json"
    ContentTypeXML      = "application/xml"
    ContentTypeHTML     = "text/html"
    ContentTypePlain    = "text/plain"
    ContentTypePDF      = "application/pdf"
    ContentTypeCSS      = "text/css"
    ContentTypeJavaScript = "application/javascript"
    ContentTypeSVG      = "image/svg+xml"
    ContentTypeGIF      = "image/gif"
    ContentTypeBMP      = "image/bmp"
    ContentTypeICO      = "image/x-icon"
    ContentTypeMP4      = "video/mp4"
    ContentTypeMP3      = "audio/mpeg"
    ContentTypeJPG      = "image/jpeg"
	ContentTypeFontOpen = "font/opentype"
)



func SetCacheHeaders(w http.ResponseWriter, contentType string, cacheDuration time.Duration, key string) {
    w.Header().Set("Content-Type", contentType)
    w.Header().Set("Cache-Control", fmt.Sprintf("max-age=%d", int(cacheDuration.Seconds())))
    w.Header().Set("Expires", time.Now().Add(cacheDuration).Format(http.TimeFormat))
    w.Header().Set("ETag", fmt.Sprintf("sysd-cache-%s", key))
}
