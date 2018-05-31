package service

type QRCodeService struct {
	Base *BaseService
}

func (ctx *QRCodeService) Encode(source string) (string, error) {
	return "", nil
}

func (ctx *QRCodeService) Decode(source string) (string, error) {
	return "", nil
}
