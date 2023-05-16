import "encoding/json"

func UnmarshalWelcome(data []byte) (Welcome, error) {
	var r Welcome
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Welcome) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Welcome struct {
	Country                      string        `json:"country"`
	AppSourceCopMandatoryAffixes []string      `json:"appSourceCopMandatoryAffixes"`
	AppFolders                   []interface{} `json:"appFolders"`
	TestFolders                  []interface{} `json:"testFolders"`
	BcptTestFolders              []interface{} `json:"bcptTestFolders"`
}
