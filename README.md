# go-httplogger
A simple lite http logger for golang. equivalent to morgan in nodeJs 

## Usage
` go get github.com/jesseokeya/go-httplogger `
<br/> 		or  <br/>
` go get -u github.com/jesseokeya/go-httplogger `

```go
import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	httplogger "github.com/jesseokeya/go-httplogger"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.ListenAndServe(port("3000"), httplogger.Golog(r))
}

// HomeHandler function
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome To A Test Server")
}

func port(p string) string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = p
	}
	return ":" + port
}
```
## Terminal Snippet
![alt text](https://raw.githubusercontent.com/jesseokeya/go-httplogger/master/example/src/resources/images/terminal.png)
