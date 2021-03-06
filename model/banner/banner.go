package banner

type BannerAd struct {
	Banner []*BannerDetail `json:"banner"`
	Ad     []*BannerDetail `json:"ad"`
}

type BannerDetail struct {
	Id      uint64     `json:"id"`
	EleInfo []*EleInfo `json:"ele_info"`
	TagName string     `json:"tag_name"`
}

type EleInfo struct {
	ImageUrl    string `json:"image_url"`
	RedirectUrl string `json:"redirect_url"`
	Sort        uint32 `json:"sort"`
}
