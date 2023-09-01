package main
import ( 
    "fmt" 
    "html" 
    "log" 
    "net/http"
    ) 
func main () {
	http.HandleFunc("/Carpetas", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Bienvenidos al compartimiento")
    })
    log.Fatal(http.ListenAndServe(":5001", nil))
}