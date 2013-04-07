package main
import(
	"fmt"
	"net/http"
	"strings"
    "log"
)

func index(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
    fmt.Println(r.URL.Path)
    fmt.Println(r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k,v := range r.Form{
        fmt.Println("key:",k)
        fmt.Println("value:",strings.Join(v," "))
    }
	fmt.Fprintln(w,"index page")
}

func main(){
	http.HandleFunc("/",index)
	err := http.ListenAndServe(":8088",nil)
	if err != nil{
		log.Fatal("ListenAndServe:",err)
	}
}

