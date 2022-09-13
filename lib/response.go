package lib

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)


func NewResponse(raw_res Response) []byte {
	extra_options := 0
	output := []byte(raw_res.Data)
    
    string_resp := fmt.Sprintf("RS~STAT=%d~SIZE=%d~ITEM=%d~OPTS=%d\r\n", raw_res.Status, len(output), raw_res.NumPost, extra_options)
	resp := []byte(string_resp)
    
    log.Debug("RES HEADER " , string_resp)
	
	resp = append(resp, output...)

	return resp

}
