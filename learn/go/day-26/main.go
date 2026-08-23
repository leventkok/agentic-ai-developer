package main


import(
	"net/http"
	"log"
	"fmt"
)



func homeHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/plain");
	w.WriteHeader(http.StatusOK);
	fmt.Fprintln(w, "Welcome to MasterFabric cafe")
}


func healthHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK);
	fmt.Fprintln(w, "ok")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusNotFound);
	fmt.Fprintln(w, "404 Not Found")
}

func logRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s from %s | User-Agent: %s\n",
			r.Method, r.URL.Path, r.RemoteAddr, r.Header.Get("User-Agent"))
		next(w, r)
	}
}



func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", logRequest(homeHandler))
	mux.HandleFunc("/health", logRequest(healthHandler))

	fmt.Println("Server is running http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
