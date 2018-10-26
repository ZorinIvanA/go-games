package fias

//Envelope 12
type Envelope struct {
	Body struct {
		GetAllDownloadFileInfoResponse struct {
			GetAllDownloadFileInfoResult struct {
				DownloadFileInfo []struct {
					VersionID          int
					TextVersion        string
					FiasCompleteDbfURL string
					FiasCompleteXMLURL string
					FiasDeltaDbfURL    string
					FiasDeltaXMLURL    string
					Kladr4ArjURL       string
					Kladr47ZURL        string
				}
			}
		}
	}
}
