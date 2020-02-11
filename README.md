# komkujson

start the server and autoreloading:

     make watch


Edit the `main.go` file (or any other `.go` file) and the server will automatically reload if the build succeeded. 
  


test it with curl

    curl -X POST --header "Content-Type: application/json" --data '{"name":"Hank"}' http://localhost:8080/jsonbodyrequest
