package api

import "github.com/segmentio/ksuid"

// Trans ...
type Trans struct {
	ID            string `json:"id"`
	SrcRegistryID string `json:"src_id"`
	DstRegistryID string `json:"dst_id"`
}

var transMap map[string]*Trans

// NewTrans ...
func NewTrans(srcID, dstID string) (string, error) {
	rs := GetRegistry(srcID)
	rd := GetRegistry(dstID)

	if rs == nil && rd == nil {
		return "", errInvalidRegistry
	}

	t := &Trans{
		SrcRegistryID: srcID,
		DstRegistryID: dstID,
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
	return true, nil
}
