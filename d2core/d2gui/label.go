package d2gui

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

// Label is renderable text
type Label struct {
	widgetBase

	renderer d2interface.Renderer
	text     string
	font     *d2asset.Font
	surface  d2interface.Surface
}

func createLabel(renderer d2interface.Renderer, text string, font *d2asset.Font) *Label {
	label := &Label{
		font:     font,
		renderer: renderer,
	}

	err := label.setText(text)
	if err != nil {
		log.Print(err)
		return nil
	}

	label.SetVisible(true)

	return label
}

func (l *Label) render(target d2interface.Surface) error {
	return target.Render(l.surface)
}

func (l *Label) getSize() (width, height int) {
	return l.surface.GetSize()
}

// GetText returns the label text
func (l *Label) GetText() string {
	return l.text
}

// SetText sets the label text
func (l *Label) SetText(text string) error {
	if text == l.text {
		return nil
	}

	return l.setText(text)
}

func (l *Label) setText(text string) error {
	width, height := l.font.GetTextMetrics(text)

	surface, err := l.renderer.NewSurface(width, height, d2enum.FilterNearest)
	if err != nil {
		return err
	}

	if err := l.font.RenderText(text, surface); err != nil {
		return err
	}

	l.surface = surface
	l.text = text

	return nil
}
