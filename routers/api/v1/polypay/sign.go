package polypay

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
)

func sign(bytes []byte, t time.Time, token string) (s string) {
	m := md5.New()
	m.Write([]byte(string(bytes) + strconv.FormatInt(t.Unix(), 10) + token))
	return hex.EncodeToString(m.Sum(nil))
}
