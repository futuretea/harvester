package volume

type ExportVolumeInput struct {
	DisplayName  string `json:"displayName"`
	Namespace    string `json:"namespace"`
	DiskSelector string `json:"diskSelector"`
	NodeSelector string `json:"nodeSelector"`
}
