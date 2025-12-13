<p align="right">
⭐ &nbsp;&nbsp;<strong>the project to show your appreciation.</strong> :arrow_upper_right:
</p>

<p align="right">
<a href="http://godoc.org/github.com/romance-dev/ascii-art"><img src="http://godoc.org/github.com/romance-dev/ascii-art?status.svg" /></a>
<a href="https://goreportcard.com/report/github.com/romance-dev/ascii-art"><img src="https://goreportcard.com/badge/github.com/romance-dev/ascii-art" /></a>
</p>

# ASCII ART for Go

## Description
Prints beautiful ASCII art from text.
It supports [FIGlet](http://www.figlet.org/) files, and most of its features.

This package is a fork of [go-figure](https://github.com/common-nighthawk/go-figure).
**UPDATES:** `go.mod` support was added, the API was cleaned up, and font-loading was done as plugins to save memory.

## Installation
`go get github.com/romance-dev/ascii-art`

## Basic Example
```go
package main

import(
    "github.com/romance-dev/ascii-art"
    _ "github.com/romance-dev/ascii-art/fonts" // load all fonts
    // _ "github.com/romance-dev/ascii-art/fonts/alphabet" // load a specific font
)

func main() {
  myFigure := asciiart.NewFigure("Hello World", "", true)
  myFigure.Print()
}
```

```txt
  _   _          _   _          __        __                 _       _ 
 | | | |   ___  | | | |   ___   \ \      / /   ___    _ __  | |   __| |
 | |_| |  / _ \ | | | |  / _ \   \ \ /\ / /   / _ \  | '__| | |  / _` |
 |  _  | |  __/ | | | | | (_) |   \ V  V /   | (_) | | |    | | | (_| |
 |_| |_|  \___| |_| |_|  \___/     \_/\_/     \___/  |_|    |_|  \__,_|
```

You can also make colorful figures:

```go
func main() {
  myFigure := asciiart.NewColorFigure("Hello World", "", asciiart.Green, true)
  myFigure.Print()
}
```


## More Examples
`asciiart.NewFigure("Go-Figure", "isometric1", true).Print()`

```
      ___           ___           ___                       ___           ___           ___           ___     
     /\  \         /\  \         /\  \          ___        /\  \         /\__\         /\  \         /\  \    
    /::\  \       /::\  \       /::\  \        /\  \      /::\  \       /:/  /        /::\  \       /::\  \   
   /:/\:\  \     /:/\:\  \     /:/\:\  \       \:\  \    /:/\:\  \     /:/  /        /:/\:\  \     /:/\:\  \  
  /:/  \:\  \   /:/  \:\  \   /::\~\:\  \      /::\__\  /:/  \:\  \   /:/  /  ___   /::\~\:\  \   /::\~\:\  \ 
 /:/__/_\:\__\ /:/__/ \:\__\ /:/\:\ \:\__\  __/:/\/__/ /:/__/_\:\__\ /:/__/  /\__\ /:/\:\ \:\__\ /:/\:\ \:\__\
 \:\  /\ \/__/ \:\  \ /:/  / \/__\:\ \/__/ /\/:/  /    \:\  /\ \/__/ \:\  \ /:/  / \/_|::\/:/  / \:\~\:\ \/__/
  \:\ \:\__\    \:\  /:/  /       \:\__\   \::/__/      \:\ \:\__\    \:\  /:/  /     |:|::/  /   \:\ \:\__\  
   \:\/:/  /     \:\/:/  /         \/__/    \:\__\       \:\/:/  /     \:\/:/  /      |:|\/__/     \:\ \/__/  
    \::/  /       \::/  /                    \/__/        \::/  /       \::/  /       |:|  |        \:\__\    
     \/__/         \/__/                                   \/__/         \/__/         \|__|         \/__/    
```

`asciiart.NewFigure("Foo Bar Pop", "smkeyboard", true).Print()`

```
 ____  ____  ____  ____  ____  ____  ____  ____  ____ 
||F ||||o ||||o ||||B ||||a ||||r ||||P ||||o ||||p ||
||__||||__||||__||||__||||__||||__||||__||||__||||__||
|/__\||/__\||/__\||/__\||/__\||/__\||/__\||/__\||/__\|
```

`asciiart.NewFigure("Keep Your Eyes On Me", "rectangles", true).Print()`

```
                                                                                          
 _____                 __ __                 _____                 _____       _____      
|  |  | ___  ___  ___ |  |  | ___  _ _  ___ |   __| _ _  ___  ___ |     | ___ |     | ___ 
|    -|| -_|| -_|| . ||_   _|| . || | ||  _||   __|| | || -_||_ -||  |  ||   || | | || -_|
|__|__||___||___||  _|  |_|  |___||___||_|  |_____||_  ||___||___||_____||_|_||_|_|_||___|
                 |_|                               |___|                                  
```

`asciiart.NewFigure("ABCDEFGHIJ", "eftichess", true).Print()`

```
#########         #########   ___   #########         #########                           
##[`'`']#  \`~'/  ##'\v/`##  /\*/\  ##|`+'|##  '\v/`  ##\`~'/##  [`'`']   '\v/`    \`~'/  
###|  |##  (o o)  ##(o 0)## /(o o)\ ##(o o)##  (o 0)  ##(o o)##   |  |    (o 0)    (o o)  
###|__|##   \ / \ ###(_)###   (_)   ###(_)###   (_)   ###\ / \#   |__|     (_)      \ / \ 
#########    "    #########         #########         ####"####                      "    
```

## Supported Fonts
* 3-d
* 3x5
* 5lineoblique
* acrobatic
* alligator
* alligator2
* alphabet
* avatar
* banner
* banner3-D
* banner3
* banner4
* barbwire
* basic
* bell
* big
* bigchief
* binary
* block
* bubble
* bulbhead
* calgphy2
* calligraphy
* catwalk
* chunky
* coinstak
* colossal
* computer
* contessa
* contrast
* cosmic
* cosmike
* cricket
* cursive
* cyberlarge
* cybermedium
* cybersmall
* diamond
* digital
* doh
* doom
* dotmatrix
* drpepper
* eftichess
* eftifont
* eftipiti
* eftirobot
* eftitalic
* eftiwall
* eftiwater
* epic
* fender
* fourtops
* fuzzy
* goofy
* gothic
* graffiti
* hollywood
* invita
* isometric1
* isometric2
* isometric3
* isometric4
* italic
* ivrit
* jazmine
* jerusalem
* katakana
* kban
* larry3d
* lcd
* lean
* letters
* linux
* lockergnome
* madrid
* marquee
* maxfour
* mike
* mini
* mirror
* mnemonic
* morse
* moscow
* nancyj-fancy
* nancyj-underlined
* nancyj
* nipples
* ntgreek
* o8
* ogre
* pawp
* peaks
* pebbles
* pepper
* poison
* puffy
* pyramid
* rectangles
* relief
* relief2
* rev
* roman
* rot13
* rounded
* rowancap
* rozzo
* runic
* runyc
* sblood
* script
* serifcap
* shadow
* short
* slant
* slide
* slscript
* small
* smisome1
* smkeyboard
* smscript
* smshadow
* smslant
* smtengwar
* speed
* stampatello
* standard
* starwars
* stellar
* stop
* straight
* tanja
* tengwar
* term
* thick
* thin
* threepoint
* ticks
* ticksslant
* tinker-toy
* tombstone
* trek
* tsalagi
* twopoint
* univers
* usaflag
* wavy
* weird


