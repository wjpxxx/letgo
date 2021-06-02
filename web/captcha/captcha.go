package captcha

import (
	"io/ioutil"
	"image/color"
	"image"
	"image/draw"
	"math"

	"github.com/golang/freetype/truetype"
	"github.com/golang/freetype"
	"github.com/wjpxxx/letgo/log"
	"github.com/wjpxxx/letgo/lib"
)

//常量
const (
	NUM =iota  //number
	LCHAR	//Lowercase letters
	UCHAR	//Upper case letters
	ALL	//number+letters
)
var captchaType =map[int]string{
	NUM:"number",
	LCHAR:"Lowercase letters",
	UCHAR:"Upper case letters",
	ALL:"number+letters",
}

//Size
type Size struct{
	Width int
	Height int
}

//Captcha
type Captcha struct {
	fonts []*truetype.Font
	fontColors []color.Color
	backColors []color.Color
	size Size
}

//AddFonts
func(c *Captcha)AddFonts(paths ...string){
	for _,p:=range paths{
		fontData,err:=ioutil.ReadFile(p)
		if err!=nil{
			log.DebugPrint("captcha load font error %v",err)
			continue
		}
		font,err:=freetype.ParseFont(fontData)
		if err!=nil{
			log.DebugPrint("captcha parse font error %v",err)
			continue
		}
		c.fonts=append(c.fonts, font)
	}
}
//AddColors
func (c *Captcha)AddColors(colors ...color.Color){
	if len(c.fontColors)>0{
		c.fontColors=c.fontColors[:0]
	}
	for _,v:=range colors{
		c.fontColors=append(c.fontColors, v)
	}
}

//AddBackColors
func (c *Captcha)AddBackColors(colors ...color.Color){
	if len(c.backColors)>0{
		c.backColors=c.backColors[:0]
	}
	for _,v:=range colors{
		c.backColors=append(c.backColors, v)
	}
}

//SetSize
func (c *Captcha)SetSize(width, height int) {
	c.size.Width=width
	c.size.Height=height
}

//Create
func (c *Captcha)Create(num int, option int) (*Image,string){
	if c.fonts == nil {
		log.DebugPrint("please call AddFonts to Add font")
		panic("please call AddFonts to Add font")
	}
	if num<=0{
		num=4
	}
	dst:=NewImage(c.size.Width,c.size.Height)
	c.drawBackground(dst)
	text:=c.rand(num,option)
	c.drawString(dst,text)
	return dst,text
}
//getFontSize
func (c *Captcha)getFontSize()int{
	if c.size.Width>c.size.Height{
		return int(float64(c.size.Height)*0.7)
	} else {
		return int(float64(c.size.Width)*0.7)
	}
}
//getPadding
func (c *Captcha)getPadding(fontSize,n int)int{
	return int(math.Ceil(float64(c.size.Width-n*fontSize)/float64(n+1)))
}
//randColor
func (c *Captcha)randColor(k int)color.Color{
	n:=len(c.fontColors)
	i:=lib.Rand(0,n-1,lib.Time()+k)
	return c.fontColors[i]
}

//randBackColor
func (c *Captcha)randBackColor()color.Color{
	n:=len(c.backColors)
	i:=lib.Rand(0,n-1,lib.Time())
	return c.backColors[i]
}
//randFont
func (c *Captcha)randFont(k int)*truetype.Font{
	n:=len(c.fonts)
	i:=lib.Rand(0,n-1,lib.Time()+k)
	return c.fonts[i]
}
//drawString
func (c *Captcha)drawString(img *Image,text string) {
	n:=len(text)
	fontSize:=c.getFontSize()
	padding:=c.getPadding(fontSize,n)
	//fmt.Println("allwidth:",c.size.Width,"allheight:",c.size.Height,"fontSize:",fontSize,"padding:",padding)
	for i,char:=range text{
		textImage:=NewImage(fontSize,fontSize)
		textImage.DrawString(c.randFont(i),c.randColor(i),string(char),float64(fontSize))
		s:=textImage.Bounds().Size()
		left:=(i+1)*padding+i*s.X
		top:=(c.size.Height-s.Y)/2
		x:=left+s.X
		y:=top+s.Y
		//fmt.Println("left:",left,"top:",top,"right:",x,"bottom:",y,"width:",s.X,"height:",s.Y)
		draw.Draw(img,image.Rect(left,top,x,y),textImage,image.ZP,draw.Over)
	}
}

//rand
func (c *Captcha)rand(num int,option int)string{
	switch option{
	case NUM:
		return lib.RandNumber(num)
	case LCHAR:
		return lib.RandLChar(num)
	case UCHAR:
		return lib.RandUChar(num)
	case ALL:
		return lib.RandChar(num)
	default:
		return lib.RandChar(num)
	}
}
//drawBackground
func (c *Captcha)drawBackground(img *Image){
	bg:=image.NewUniform(c.randBackColor())
	img.Fill(bg)
}

//NewCaptcha
func NewCaptcha()*Captcha{
	return &Captcha{
		size: Size{140,40},
		fontColors: []color.Color{color.Black},
		backColors: []color.Color{color.White},
	}
}