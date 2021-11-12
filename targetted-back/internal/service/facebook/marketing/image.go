package marketing

type Image interface {
	Add()
	GetByHash()
	GetByAccount()
}

// https://developers.facebook.com/docs/marketing-api/reference/ad-image

type FacebookImage struct {
}

// curl \
//   -F 'filename=@<IMAGE_PATH>' \
//   -F 'access_token=<ACCESS_TOKEN>' \
//   https://graph.facebook.com/v2.11/act_<AD_ACCOUNT_ID>/adimages
func (img *FacebookImage) Add() {

}

// curl \
//   -F 'bytes=iVBORw0KGgoAAAANSUhEUgAAABEAAAARCAMAAAAMs7fIAAAAOVBMVEX///87WZg7WZg7WZg7WZg7WZg7WZg7WZg7WZg7WZg7WZhMeMJEaa5Xi9tKdb0+Xp5Wi9tXjNxThNH+wk/7AAAACnRSTlMAsHIoaM7g/fx9Zr/g5QAAAGlJREFUeNplkFsOwCAIBPGJrtbX/Q/bqm1qwnxuJrBAE6OVD15pQy/WYePsDiIjp9FGyuC4DK7l6pOrVH4s41D6R4EzpJGXsa0MTQqp/yQo8hhHMuApoB1JQ5COnCN3yT6ys7xL3i7/cwMYsAveYa+MxAAAAABJRU5ErkJggg=='
//   -F 'access_token=<ACCESS_TOKEN>' \
// "https://graph.facebook.com/<API_VERSION>/act_<ACCOUNT_ID>/adimages"
func (img *FacebookImage) AddBytes() {

}

// curl -G \
//   -d 'access_token=<ACCESS_TOKEN>' \
//   https://graph.facebook.com/v2.11/act_<AD_ACCOUNT_ID>/adimages
func (img *FacebookImage) GetByHash() {

}

// curl -G \
//   -d 'hashes=["<IMAGE_1_HASH>","<IMAGE_2_HASH>"]' \
//   -d 'access_token=<ACCESS_TOKEN>' \
//   https://graph.facebook.com/v2.11/act_<AD_ACCOUNT_ID>/adimages
func (img *FacebookImage) GetByAddAccount() {

}
