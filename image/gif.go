package image

import (
	"fmt"
	"image/gif"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/malivvan/mate/view"
)

// GifView is a box which displays animated gifs via Omnikron13's pixelview
// dynamic color rendering.  It automatically draws the right frame based on
// time elapsed since creation.  You can trigger re-drawing by executing
// Animate(view.Application) in a goroutine.
type GifView struct {
	sync.Mutex
	*view.Box

	// Timing for the frames
	delay         []time.Duration
	frames        []string
	startTime     time.Time
	totalDuration time.Duration
}

// NewGifView returns a new GifView.
func NewGifView() *GifView {
	return &GifView{
		Box:       view.NewBox(),
		startTime: time.Now(),
	}
}

// GifFromImage creates a new GifView from a GIF Image
func GifFromImage(img *gif.GIF) (*GifView, error) {
	g := NewGifView()
	return g.SetImage(img)
}

// GifFromImagePath creates a new GifView from a file on disk
func GifFromImagePath(imgPath string) (*GifView, error) {
	g := NewGifView()
	return g.SetImagePath(imgPath)
}

// SetImagePath sets the image to a given GIF path
func (g *GifView) SetImagePath(imgPath string) (*GifView, error) {
	file, err := os.Open(imgPath)
	if err != nil {
		return g, fmt.Errorf("Unable to open file: %v", err)
	}
	defer file.Close()

	image, err := gif.DecodeAll(file)
	if err != nil {
		return g, fmt.Errorf("Unable to decode GIF: %v", err)
	}

	return g.SetImage(image)
}

// SetImage sets the content to a given gif.GIF
func (g *GifView) SetImage(img *gif.GIF) (*GifView, error) {
	g.Lock()
	defer g.Unlock()

	// Store delay in milliseconds
	g.totalDuration = time.Duration(0)
	for _, i := range img.Delay {
		d := time.Duration(i*10) * time.Millisecond
		g.delay = append(g.delay, d)
		g.totalDuration += d
	}

	// Set height,width of the box
	g.SetRect(0, 0, img.Config.Width, img.Config.Height)

	// Convert images to text
	frames := []string{}
	for i, elem := range img.Image {
		parsed, err := FromImage(elem)
		if err != nil {
			return g, fmt.Errorf("Unable to convert frame %d: %v", i, err)
		}
		frames = append(frames, parsed)
	}

	// Store the output
	g.frames = frames

	return g, nil
}

// GetCurrentFrame returns the current frame the GIF is on
func (g *GifView) GetCurrentFrame() int {
	// Always at frame 0
	if g.totalDuration == 0 {
		return 0
	}

	dur := time.Since(g.startTime) % g.totalDuration
	for i, d := range g.delay {
		dur -= d
		if dur < 0 {
			return i
		}
	}
	return 0
}

// Draw renders the current frame of the GIF
func (g *GifView) Draw(screen tcell.Screen) {
	g.Lock()
	defer g.Unlock()

	currentFrame := g.GetCurrentFrame()

	frame := strings.Split(g.frames[currentFrame], "\n")
	x, y, w, _ := g.GetInnerRect()

	for i, line := range frame {
		view.Print(screen, []byte(line), x, y+i, w, view.AlignLeft, tcell.ColorWhite)
	}
}

var globalAnimationMutex = &sync.Mutex{}

// Animate triggers the application to redraw every 50ms
func Animate(app *view.Application) {
	globalAnimationMutex.Lock()
	defer globalAnimationMutex.Unlock()

	for {
		app.QueueUpdateDraw(func() {})
		time.Sleep(50 * time.Millisecond)
	}
}
