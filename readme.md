# GO LCD Fonts

A font for use with TinyGo and a relevant display (LCD/OLED etc). Use the [TinyFont](tinygo.org/x/tinyfont) library.

Using the provided fonts in the TinyFont library as a template [](https://pkg.go.dev/tinygo.org/x/tinyfont#section-directories), I've hand coded the dotmatrix font by marking out each glyph in the Excel file and copying the encoded bytes and width information.

I haven't provided every glyph, and most of the info has been entered through trial and error.

## Example usage
The [TinyFont](tinygo.org/x/tinyfont) library will write to any pointer to the screen, using a pointer to the font.

```go
import (
  ...
  "github.com/chilledoj/golcdfonts/dotmatrix"
)

...

tinyfont.WriteLine(&oled,&dotmatrix.Regular, 0, 7, "Dot Matrix",ssd1305.WHITE)

tinyfont.WriteLine(&oled,&dotmatrix.BOLD, 0, 15, "Bold !",ssd1305.WHITE)
```

## Notes
In my usage, I noted that when using the provided fonts in the __TinyFont__ libary, I had to set the `y` position as the baseline of the font. These fonts have been encoded in the same way, but might be a little off. In general, They seemed to be ok on a 128x32 OLED screen when using the `y` positions of `7`,`15`,`23`,`31`.

