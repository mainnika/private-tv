package backend

import (
	"fmt"
	"os"

	"github.com/lidouf/gst"
)

const (
	srcName       = "src"
	decoderName   = "decoder"
	converterName = "converter"
	sinkName      = "sink"
)

// Pipeline Pipeline
type Pipeline struct {
	*gst.Pipeline
	src       *gst.Element
	decoder   *gst.Element
	converter *gst.Element
	sink      *gst.Element
}

// NewPipeline NewPipeline
func NewPipeline() (error, *Pipeline) {

	p := &Pipeline{}

	p.Pipeline = gst.NewPipeline("main")
	if p.Pipeline == nil {
		return fmt.Errorf("can not create pipeline"), nil
	}

	p.src = gst.ElementFactoryMake("filesrc", srcName)
	if p.src == nil {
		return fmt.Errorf("can not create filesrc"), nil
	}

	p.decoder = gst.ElementFactoryMake("decodebin", decoderName)
	if p.decoder == nil {
		return fmt.Errorf("can not create decodebin"), nil
	}

	p.converter = gst.ElementFactoryMake("videoconvert", converterName)
	if p.converter == nil {
		return fmt.Errorf("can not create videoconvert"), nil
	}

	p.sink = gst.ElementFactoryMake("aasink", sinkName)
	if p.sink == nil {
		return fmt.Errorf("can not create aasink"), nil
	}

	p.Pipeline.Add(p.src, p.decoder, p.converter, p.sink)

	p.src.Link(p.decoder)
	p.converter.Link(p.sink)
	p.decoder.ConnectNoi("pad-added", p.connectDynPad, p.converter.GetStaticPad("sink"))

	return nil, p
}

// SetSource SetSource
func (p *Pipeline) SetSource(file string) {

	p.src.SetProperty("location", file)
}

// connectDynPad function for "pad-added" event
func (p *Pipeline) connectDynPad(targetPad, createdPad *gst.Pad) {

	if !createdPad.CanLink(targetPad) {
		fmt.Fprintln(os.Stderr, "can't link:", createdPad.GetName(), targetPad.GetName())
		return
	}

	if createdPad.Link(targetPad) != gst.PAD_LINK_OK {
		fmt.Fprintln(os.Stderr, "link error:", createdPad.GetName(), targetPad.GetName())
	}
}
