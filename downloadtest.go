package main


import (
  "net/http"
  "fmt"
  "os"  
  "os/exec"
)

var (
  genScript = "./generate-data.sh"
)

func genData() ([]byte, error){
  out, err := exec.Command(genScript).CombinedOutput()
  return out, err
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
  welcomepage := `
    <html>
      <p>Welcome to getter which allows you to download generated zero byte files.  When the app starts up it runs generate-data.sh which will generate 1,2,5,10,20,50,100, and 1024mb files for you to download</p>
      <p>You can download these fils using the following example</p>
      <pre>
wget http://app.domain/downloads/1mb
wget http://app.domain/downloads/2mb
wget http://app.domain/downloads/5mb
wget http://app.domain/downloads/10mb
wget http://app.domain/downloads/20mb
wget http://app.domain/downloads/50mb
wget http://app.domain/downloads/100mb
wget http://app.domain/downloads/1024mb
</pre>
    </html>
  `
  w.Write([]byte(welcomepage))
}

func main(){
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }
  
  out, err := genData()
  if err != nil {
    fmt.Printf("Failed to generate downloadable data: %s\n%s\n", out, err)
    os.Exit(100)
  }
  
  http.HandleFunc("/", rootHandler)
  http.Handle("/downloads/", http.FileServer(http.Dir("")))
  fmt.Printf("Starting Serve on port %s\n", port)
  err = http.ListenAndServe(":"+port, nil)
  if err != nil {
    fmt.Printf("Failed to start http server: %s\n", err)
    os.Exit(200)
  }
  
}
