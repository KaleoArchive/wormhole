package api

import "github.com/segmentio/ksuid"

// Trans ...
type Trans struct {
	ID            string `json:"id"`
	SrcRegistryID string `json:"srcId"`
	DstRegistryID string `json:"dstId"`
}

var transMap map[string]*Trans

func init() {
	transMap = make(map[string]*Trans)
}

// NewTrans ...
func NewTrans(t *Trans) (string, error) {
	rs := GetRegistry(t.SrcRegistryID)
	rd := GetRegistry(t.DstRegistryID)

	if rs == nil || rd == nil {
		return "", errInvalidRegistry
	}

	ok, err := TransExist(t)
	if err != nil {
		return "", err
	}
	if ok {
		return t.ID, nil
	}

	id := ksuid.New()
	t.ID = id.String()
	AddTrans(t)
	return t.ID, nil
}

// AddTrans ...
func AddTrans(t *Trans) {
	transMap[t.ID] = t
}

// TransExist ...
func TransExist(t *Trans) (bool, error) {
	return false, nil
}
