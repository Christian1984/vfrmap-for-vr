package drm

type DrmData struct{}

func New() DrmData {
	return DrmData{}
}

func (dd *DrmData) Valid() bool {
	return false
}

func (dd *DrmData) Serial() string {
	return "unlicensed"
}

func (dd *DrmData) Email() string {
	return "unlicensed"
}