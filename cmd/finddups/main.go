package main

import (
	"crypto/sha256"
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/johnllao/macgo/pkg/queue"
)

var (
	path    string
	pathq   *queue.StringQueue
	fileset [][]string
)

func init() {
	pathq = queue.NewStringQueue()
	fileset = make([][]string, 0)

	flag.StringVar(&path, "p", "", "path to search for duplicate files")
	flag.Parse()
}

func main() {
	var err error

	log.Printf("searching for duplicates at %s", path)

	var filesm = make(map[[32]byte][]string)

	pathq.Push(path)

	for pathq.Len() > 0 {

		var pathval = pathq.Pop()

		var fileinfos []os.FileInfo
		fileinfos, err = ioutil.ReadDir(pathval)
		if err != nil {
			log.Print("WARN: ", err)
		}

		for _, fileinfo := range fileinfos {
			var pp = filepath.Join(pathval, fileinfo.Name())
			if fileinfo.IsDir() {
				pathq.Push(pp)
			}

			if !fileinfo.IsDir() {
				var cksum [sha256.Size]byte
				cksum, err = FileHash(pp)
				if err != nil {
					log.Print("WARN: ", err)
					continue
				}
				if _, ok := filesm[cksum]; ok {
					var l = filesm[cksum]
					l = append(l, pp)
					filesm[cksum] = l
				} else {
					var l = make([]string, 1)
					l[0] = pp
					filesm[cksum] = l
				}
			}
		}
	}

	for _, i := range filesm {
		if len(i) > 1 {
			fileset = append(fileset, i)
		}
	}

	var t *template.Template
	t, err = template.New("duplicates").Parse(templ)
	if err != nil {
		log.Fatal("ERR: ", err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, fileset)
	})
	log.Print("open http://localhost:8080 to view the report")
	http.ListenAndServe("localhost:8080", nil)
}

var templ = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		<title>Duplicate Files</title>
		<style type="text/css">
		body { background-color: #000; color: #ddd; font-family: Tahoma; font-size: 10pt; }
		</style>
	</head>

	<body>
		<div>
			{{ range $i, $e := . }}<tr>
				<div>
					<ol>
						{{ range $e}}<li> {{ . }} </li>
						{{ end }}
					</ol>
				</div>
			</tr>{{ end }}
		</div>
	</body>

</html>

`
