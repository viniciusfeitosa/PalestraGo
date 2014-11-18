package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Correios struct {
	Bairro, Logradouro, Cep, Uf, Localidade string
}

func (c *Correios) ServeHTTP(rw http.ResponseWriter, req *http.Request) {

	caminho := fmt.Sprintf("http://cep.correiocontrol.com.br/%s.json", req.URL.Path[len("/correios/"):])
	log.Println(caminho)
	resp, _ := http.Get(caminho)

	defer resp.Body.Close()
	var correios Correios

	dec := json.NewDecoder(resp.Body)
	dec.Decode(&correios)

	fmt.Fprintf(rw, "%s \n\n", correios.Bairro)
	fmt.Fprintf(rw, "%s \n\n", correios.Logradouro)
	fmt.Fprintf(rw, "%s \n\n", correios.Cep)
	fmt.Fprintf(rw, "%s \n\n", correios.Uf)
	fmt.Fprintf(rw, "%s \n\n", correios.Localidade)
}

func main() {
	correiosHandler := new(Correios)

	http.Handle("/correios/", correiosHandler)
	log.Println("Start app on :8888")
	http.ListenAndServe(":8888", nil)
}
