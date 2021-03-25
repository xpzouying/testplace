package main

import "log"

type RenderType = uint8

const (
	RenderTypeHeader RenderType = iota
	RenderTypeFoot

	RenderTypeCount
)

type Render interface {
	Rend() error

	Type() RenderType
}

type HeaderRender struct {
	Title string
	Logo  string
}


func (r HeaderRender) Rend() error {
	log.Printf("heder_render: title:%s logo:%s", r.Title, r.Logo)
	return nil
}

func (HeaderRender) Type() RenderType {
	return RenderTypeHeader
}

type FootRender struct {
	Statement string
	Year      int
}

func (r FootRender) Rend() error {
	log.Printf("foot_render: statement:%s year:%d", r.Statement, r.Year)
	return nil
}

func (FootRender) Type() RenderType {
	return RenderTypeFoot
}
