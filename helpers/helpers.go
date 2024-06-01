package helpers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
	"time"

	"gitlab.com/rodrigoodhin/gocure/models"
)

func HumanizeExecution(execution int) string {
	duration := time.Duration(execution)

	if duration.Milliseconds() < 1000.0 {
		return fmt.Sprintf("<span class='timer-bold'>%d</span>ms", int64(duration.Milliseconds()))
	}

	if duration.Seconds() < 60.0 {
		remainingMilliseconds := math.Mod(float64(duration.Milliseconds()), 1000)
		return fmt.Sprintf("<span class='timer-bold'>%d</span>s <span class='timer-bold'>%d</span>ms", int64(duration.Seconds()), int64(remainingMilliseconds))
	}

	if duration.Minutes() < 60.0 {
		remainingMilliseconds := math.Mod(float64(duration.Milliseconds()), 1000)
		remainingSeconds := math.Mod(duration.Seconds(), 60)
		return fmt.Sprintf("<span class='timer-bold'>%d</span>m <span class='timer-bold'>%d</span>s <span class='timer-bold'>%d</span>ms", int64(duration.Minutes()), int64(remainingSeconds), int64(remainingMilliseconds))
	}

	if duration.Hours() < 24.0 {
		remainingMilliseconds := math.Mod(float64(duration.Milliseconds()), 1000)
		remainingMinutes := math.Mod(duration.Minutes(), 60)
		remainingSeconds := math.Mod(duration.Seconds(), 60)
		return fmt.Sprintf("<span class='timer-bold'>%d</span>h <span class='timer-bold'>%d</span>m <span class='timer-bold'>%d</span>s <span class='timer-bold'>%d</span>ms",
			int64(duration.Hours()), int64(remainingMinutes), int64(remainingSeconds), int64(remainingMilliseconds))
	}

	remainingMilliseconds := math.Mod(float64(duration.Milliseconds()), 1000)
	remainingHours := math.Mod(duration.Hours(), 24)
	remainingMinutes := math.Mod(duration.Minutes(), 60)
	remainingSeconds := math.Mod(duration.Seconds(), 60)
	return fmt.Sprintf("<span class='timer-bold'>%d</span>d <span class='timer-bold'>%d</span>h <span class='timer-bold'>%d</span>m <span class='timer-bold'>%d</span>s <span class='timer-bold'>%d</span>ms",
		int64(duration.Hours()/24), int64(remainingHours),
		int64(remainingMinutes), int64(remainingSeconds), int64(remainingMilliseconds))
}

func ReportParse(ignoreBadFiles bool, content []byte) (features []models.Feature, err error) {

	err = json.Unmarshal(content, &features)
	if err != nil && !ignoreBadFiles {
		return nil, fmt.Errorf("error trying to unmarshal json file : %v", err)
	}

	return
}

func ToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func ReadReport(ignoreBadFiles bool, report string) (features []models.Feature, err error) {
	content, err := ioutil.ReadFile(report)
	if err != nil {
		return nil, fmt.Errorf("error reading report file: %v", err)
	}

	features, err = ReportParse(ignoreBadFiles, content)
	if err != nil {
		return nil, fmt.Errorf("error parsing report: %v", err)
	}

	return
}

func ReadFolders(folder string) (fileList []string, err error) {

	dirContent, err := os.ReadDir(folder)
	if err != nil {
		return nil, fmt.Errorf("error reading directory: %v", err)
	}

	folder = strings.TrimSuffix(folder, "/")

	for _, file := range dirContent {
		if file.IsDir() {
			files, err := ReadFolders(folder + "/" + file.Name())
			if err != nil {
				return nil, fmt.Errorf("error reading directory: %v", err)
			}
			fileList = append(fileList, files...)
		} else {
			fileList = append(fileList, folder+"/"+file.Name())
		}
	}

	return
}

func GetFileExtension(mimeTyoe string) (fileExtension string) {

	m := make(map[string]string)
	m["application/envoy"] = "evy"
	m["application/fractals"] = "fif"
	m["application/futuresplash"] = "spl"
	m["application/hta"] = "hta"
	m["application/internet-property-stream"] = "acx"
	m["application/mac-binhex40"] = "hqx"
	m["application/msword"] = "doc"
	m["application/octet-stream"] = "exe"
	m["application/oda"] = "oda"
	m["application/olescript"] = "axs"
	m["application/pdf"] = "pdf"
	m["application/pics-rules"] = "prf"
	m["application/pkcs10"] = "p10"
	m["application/pkix-crl"] = "crl"
	m["application/postscript"] = "ai"
	m["application/postscript"] = "eps"
	m["application/postscript"] = "ps"
	m["application/rtf"] = "rtf"
	m["application/set-payment-initiation"] = "setpay"
	m["application/set-registration-initiation"] = "setreg"
	m["application/vnd.ms-excel"] = "xls"
	m["application/vnd.ms-outlook"] = "msg"
	m["application/vnd.ms-pkicertstore"] = "sst"
	m["application/vnd.ms-pkiseccat"] = "cat"
	m["application/vnd.ms-pkistl"] = "stl"
	m["application/vnd.ms-powerpoint"] = "pps"
	m["application/vnd.ms-project"] = "mpp"
	m["application/vnd.ms-works"] = "wks"
	m["application/winhlp"] = "hlp"
	m["application/x-bcpio"] = "bcpio"
	m["application/x-cdf"] = "cdf"
	m["application/x-compress"] = "z"
	m["application/x-compressed"] = "tgz"
	m["application/x-cpio"] = "cpio"
	m["application/x-csh"] = "csh"
	m["application/x-director"] = "dcr"
	m["application/x-dvi"] = "dvi"
	m["application/x-gtar"] = "gtar"
	m["application/x-gzip"] = "gz"
	m["application/x-hdf"] = "hdf"
	m["application/x-internet-signup"] = "ins"
	m["application/x-iphone"] = "iii"
	m["application/x-javascript"] = "js"
	m["application/x-latex"] = "latex"
	m["application/x-msaccess"] = "mdb"
	m["application/x-mscardfile"] = "crd"
	m["application/x-msclip"] = "clp"
	m["application/x-msdownload"] = "dll"
	m["application/x-msmediaview"] = "mvb"
	m["application/x-msmetafile"] = "wmf"
	m["application/x-msmoney"] = "mny"
	m["application/x-mspublisher"] = "pub"
	m["application/x-msschedule"] = "scd"
	m["application/x-msterminal"] = "trm"
	m["application/x-mswrite"] = "wri"
	m["application/x-netcdf"] = "cdf"
	m["application/x-perfmon"] = "pmc"
	m["application/x-pkcs12"] = "pfx"
	m["application/x-pkcs7-certificates"] = "spc"
	m["application/x-pkcs7-certreqresp"] = "p7r"
	m["application/x-pkcs7-mime"] = "p7c"
	m["application/x-pkcs7-signature"] = "p7s"
	m["application/x-sh"] = "sh"
	m["application/x-shar"] = "shar"
	m["application/x-shockwave-flash"] = "swf"
	m["application/x-stuffit"] = "sit"
	m["application/x-sv4cpio"] = "sv4cpio"
	m["application/x-sv4crc"] = "sv4crc"
	m["application/x-tar"] = "tar"
	m["application/x-tcl"] = "tcl"
	m["application/x-tex"] = "tex"
	m["application/x-texinfo"] = "texinfo"
	m["application/x-troff"] = "roff"
	m["application/x-troff-man"] = "man"
	m["application/x-troff-me"] = "me"
	m["application/x-troff-ms"] = "ms"
	m["application/x-ustar"] = "ustar"
	m["application/x-wais-source"] = "src"
	m["application/x-x509-ca-cert"] = "crt"
	m["application/ynd.ms-pkipko"] = "pko"
	m["application/zip"] = "zip"
	m["audio/basic"] = "au"
	m["audio/mid"] = "mid"
	m["audio/mpeg"] = "mp3"
	m["audio/x-aiff"] = "aiff"
	m["audio/x-mpegurl"] = "m3u"
	m["audio/x-pn-realaudio"] = "ra"
	m["audio/x-wav"] = "wav"
	m["image/bmp"] = "bmp"
	m["image/cis-cod"] = "cod"
	m["image/gif"] = "gif"
	m["image/webp"] = "webp"
	m["image/png"] = "png"
	m["image/ief"] = "ief"
	m["image/jpeg"] = "jpg"
	m["image/pipeg"] = "jfif"
	m["image/svg+xml"] = "svg"
	m["image/tiff"] = "tiff"
	m["image/x-cmu-raster"] = "ras"
	m["image/x-cmx"] = "cmx"
	m["image/x-icon"] = "ico"
	m["image/x-portable-anymap"] = "pnm"
	m["image/x-portable-bitmap"] = "pbm"
	m["image/x-portable-graymap"] = "pgm"
	m["image/x-portable-pixmap"] = "ppm"
	m["image/x-rgb"] = "rgb"
	m["image/x-xbitmap"] = "xbm"
	m["image/x-xpixmap"] = "xpm"
	m["image/x-xwindowdump"] = "xwd"
	m["text/css"] = "css"
	m["text/h323"] = "323"
	m["text/html"] = "html"
	m["text/iuls"] = "uls"
	m["text/plain"] = "txt"
	m["text/richtext"] = "rtx"
	m["text/scriptlet"] = "sct"
	m["text/tab-separated-values"] = "tsv"
	m["text/webviewhtml"] = "htt"
	m["text/x-component"] = "htc"
	m["text/x-setext"] = "etx"
	m["text/x-vcard"] = "vcf"
	m["video/mpeg"] = "mpeg"
	m["video/mp4"] = "mp4"
	m["video/quicktime"] = "mov"
	m["video/x-la-asf"] = "lsf"
	m["video/x-ms-asf"] = "asf"
	m["video/x-msvideo"] = "avi"
	m["video/x-sgi-movie"] = "movie"

	fileExtension = m[mimeTyoe]

	return
}

func ByteCountSI(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

func GetBase64Size(b string) int64 {
	return int64(len(b)*3/4 - strings.Count(b, "="))
}
