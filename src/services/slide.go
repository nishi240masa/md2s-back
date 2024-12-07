package services

import "md2s/utils/slide"

func GetSlide(content []byte) (marp string, err error) {
	marp, err = slide.SlideConverter(content)
	return marp, err
}
