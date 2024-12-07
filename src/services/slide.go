package services

import (
	"md2s/dto"
	"md2s/utils/slide"
)

func GetSlide(input dto.RequestBody) (marp string, err error) {
	marp, err = slide.SlideConverter(input)
	return marp, err
}
