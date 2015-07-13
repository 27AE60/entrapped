package entrapped

import (
	"crypto/rand"
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"os"
)

var logger = log.New(os.Stdout, "[entrapped]", log.Ldate|log.Ltime|log.Llongfile)

func randomInt(max int) int {
	if max <= 0 {
		max = 1
	}

	num, numErr := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if numErr != nil {
		return 0
	}

	return int(num.Int64())
}

func decodeJSON(req *http.Request, dst interface{}) error {
	return json.NewDecoder(req.Body).Decode(dst)
}

func error500(rw http.ResponseWriter, err error) {
	rw.WriteHeader(500)
	rw.Write([]byte(err.Error()))
}
