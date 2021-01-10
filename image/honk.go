package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
)

func handler(w http.ResponseWriter, r *http.Request) {
  var name, _ = os.Hostname()

  fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>Honk!</title>
</head>
<body>
<pre>
                                   ___
                               ,-""   ` + "`" + ` .     < Honk from %s !>
                             ,'  _   e )` + "`" + `-._ /
                            /  ,' ` + "`" + `-._<.===-'
                           /  /
                          /  ;
              _          /   ;
 (` + "`" + `._    _.-"" ""--..__,'    |
 <_  ` + "`" + `-""                     \
  <` + "`" + `-                          :
   (__   <__.                  ;
     ` + "`" + `-.   '-.__.      _.'    /
        \      ` + "`" + `-.__,-'    _,'
         ` + "`" + `._    ,    /__,-'
            ""._\__,'< <____
                 | |  ` + "`" + `----.` + "`" + `.
                 | |        \ ` + "`" + `.
                 ; |___      \-` + "`" + `` + "`" + `
                 \   --<
                  ` + "`" + `.` + "`" + `.<
                    ` + "`" + `-'

</pre>
</body>
</html>
`, name)
log.Print("Served request ",r,"\n")
}

func main() {
  log.SetOutput(os.Stderr)
  log.Println("Starting server ...")
  http.HandleFunc("/", handler)
  err := http.ListenAndServe(":8080",nil)
  if err != nil {
    log.Fatal("ListenAndServer: ", err)
  }
}
