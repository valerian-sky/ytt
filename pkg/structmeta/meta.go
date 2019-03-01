package structmeta

import (
	"fmt"
	"strings"
)

type AnnotationName string
type AnnotationNs string

const (
	AnnotationNameComment AnnotationName = "comment"
)

type Meta struct {
	Annotations []*Annotation
}

type Annotation struct {
	Name    AnnotationName // eg template/code
	Content string         // eg if True:
}

// Supported formats:
//   "! comment"
//   "@comment content"
//   "@ if True:"
//   "@template/code"
//   "@template/code if True:"
//   "@text/trim-left,text/trim-right,template/code if True:"

func NewMetaFromString(data string) (Meta, error) {
	meta := Meta{}

	// TODO better error messages?
	switch {
	case len(data) == 0:
		return Meta{}, fmt.Errorf("Expected non-empty metadata")

	case data[0] == '!':
		meta.Annotations = []*Annotation{{
			Name:    AnnotationNameComment,
			Content: data[1:],
		}}

	case data[0] == '@':
		pieces := strings.SplitN(data[1:], " ", 2)
		for _, name := range strings.Split(pieces[0], ",") {
			meta.Annotations = append(meta.Annotations, &Annotation{
				Name: AnnotationName(name),
			})
		}
		if len(pieces) == 2 {
			meta.Annotations[len(meta.Annotations)-1].Content = pieces[1]
		}

	default:
		return Meta{}, fmt.Errorf("Unknown metadata format (use '#@' or '#!')")
	}

	return meta, nil
}