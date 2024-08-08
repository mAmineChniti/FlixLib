package utils

import (
	"fmt"
	"github.com/a-h/templ"
	"net/http"
)

func Render(component templ.Component) {
	component.Render(context.Background(), os.Stdout)
}
