package main

import (
	fface "facefinder/findface/facerec"
)

func b2vector(jb []byte) (vector []float32) {
	rec, err := fface.NewRecognizer(RECOGNIZER_DIR)
	if err != nil {
		return vector
	}
	defer rec.Close()

	recs, err := rec.Recognize(jb)
	if len(recs) < 1 {
		return vector
	}
	for _, point := range recs[0].Descriptor {
		vector = append(vector, point)
	}
	return vector
}
