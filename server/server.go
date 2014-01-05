package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

var (
	homeTemplate = template.Must(template.ParseFiles("layouts/index.html"))
)

func StartServer(port int) {
	log.Println("Starting server on port :", port)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/", handleHomePage)
	http.HandleFunc("/dashboard/", handleDashboard)

	addr := fmt.Sprintf(":%d", port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}

func handleHomePage(c http.ResponseWriter, req *http.Request) {
	files := getAllFiles("./dashboards")
	data := make(map[string]interface{})
	data["host"] = req.Host
	data["dashboards"] = files
	homeTemplate.Execute(c, data)
}

func handleDashboard(c http.ResponseWriter, r *http.Request) {

	dashboard := r.URL.Path[len("/dashboard/"):]
	_, err := template.ParseFiles("./dashboards/" + dashboard + ".html")
	if err != nil {
		fmt.Println(err)
		http.Error(c, "Dashboard not found", 404)
		return
	}

	dashTemplate := template.New(dashboard)

	dashTemplate.ParseFiles("./dashboards/" + dashboard + ".html")

	pattern := filepath.Join("./widgets/", "*.widget")
	dashTemplate.ParseGlob(pattern)
	data := make(map[string]interface{})
	data["host"] = r.Host
	dashTemplate.ExecuteTemplate(c, dashboard, data)

}

func getAllFiles(path string) []string {
	dir, err := os.Open(path)
	checkErr(err)
	defer dir.Close()
	fi, err := dir.Stat()
	checkErr(err)

	filenames := make([]string, 0)
	if fi.IsDir() {
		fis, err := dir.Readdir(-1) // -1 means return all the FileInfos
		checkErr(err)
		for _, fileinfo := range fis {
			if !fileinfo.IsDir() {
				filenames = append(filenames, removeExtension(fileinfo.Name()))
			}
		}
	}
	return filenames
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func removeExtension(filename string) string {
	var extension = filepath.Ext(filename)
	return filename[0 : len(filename)-len(extension)]
}
