package progress

import (
	"flag"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"math/rand"
	"os"
	"testing"
	"time"
)

func Test_Example(t *testing.T) {
	b := NewInt(50)

	for i := 0; i <= 20; i++ {
		b.ValueInt(i)
		b.Text(fmt.Sprintf("iteration %d", i))
		b.WriteTo(os.Stdout)
		time.Sleep(time.Millisecond * 100)
	}
}

var preview = flag.Bool("preview", false, "Preview output rendering.")

func TestBar_previewWidth(t *testing.T) {
	if !*preview {
		t.SkipNow()
	}

	b := NewInt(10)
	b.Width = 25
	b.Empty = " "

	for i := 0; i <= 10; i++ {
		b.ValueInt(i)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		b.WriteTo(os.Stdout)
	}
}

func TestBar_previewDefaults(t *testing.T) {
	if !*preview {
		t.SkipNow()
	}

	b := NewInt(20)

	for i := 0; i <= 20; i++ {
		b.ValueInt(i)
		b.Text(fmt.Sprintf("iteration %d", i))
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		b.WriteTo(os.Stdout)
	}
}

func TestBarString(t *testing.T) {
	Convey("test", t, func() {
		b := NewInt(1000)
		So(`  0% |░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░| `, ShouldEqual, b.String())

		b.ValueInt(250)
		So(` 25% |███████████████░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░| `, ShouldEqual, b.String())

		b.ValueInt(750)
		So(` 75% |█████████████████████████████████████████████░░░░░░░░░░░░░░░| `, ShouldEqual, b.String())

		b.ValueInt(1000)
		So(`100% |████████████████████████████████████████████████████████████| `, ShouldEqual, b.String())
	})

}

func TestBarText(t *testing.T) {
	Convey("test", t, func() {
		b := NewInt(1000)

		b.Text("Building")
		So(`  0% |░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░| Building`, ShouldEqual, b.String())

		b.Text("Installing")
		b.ValueInt(250)
		So(` 25% |███████████████░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░| Installing`, ShouldEqual, b.String())
	})

}

func TestBarTemplate(t *testing.T) {
	Convey("test", t, func() {
		b := NewInt(1000)
		b.Template(`{{.Bar}} {{.Percent}}%`)
		So(`|░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░| 0%`, ShouldEqual, b.String())

		b.ValueInt(250)
		So(`|███████████████░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░| 25%`, ShouldEqual, b.String())

		b.ValueInt(750)
		So(`|█████████████████████████████████████████████░░░░░░░░░░░░░░░| 75%`, ShouldEqual, b.String())

		b.ValueInt(1000)
		So(`|████████████████████████████████████████████████████████████| 100%`, ShouldEqual, b.String())
	})
}

func TestBarDelimiters(t *testing.T) {
	Convey("test", t, func() {
		b := NewInt(1000)
		b.StartDelimiter = "["
		b.EndDelimiter = "]"
		So(`  0% [░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░] `, ShouldEqual, b.String())
	})

}

func TestNewSimple(t *testing.T) {
	Convey("test", t, func() {
		b := NewSimple(10)
		So("  0%", ShouldEqual, b.String())
	})
}

func TestBar_AutoString(t *testing.T) {
	Convey("test", t, func() {
		b := NewSimple(10)
		for i := 0; i < 10; i++ {
			fmt.Println(b.AutoString())
		}
	})
}
