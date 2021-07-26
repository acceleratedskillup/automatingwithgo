package main

import "fmt"

type Artwork struct {
	title  string
	artist string
	epoch  string
}

type Gallery struct {
	artworks []Artwork
}

func (g *Gallery) addArtwork(a Artwork) {
	g.artworks = append(g.artworks, a)
}

func (g *Gallery) displayArtworks() {
	for _, artwork := range g.artworks {
		fmt.Println("Title:", artwork.title)
		fmt.Println("Artist:", artwork.artist)
		fmt.Println("Epoch:", artwork.epoch)
		fmt.Println("-----")
	}
}

func (g *Gallery) searchByArtist(artistName string) {
	for _, artwork := range g.artworks {
		if artwork.artist == artistName {
			fmt.Println("Title:", artwork.title)
			fmt.Println("Epoch:", artwork.epoch)
			fmt.Println("-----")
		}
	}
}

func main() {
	monaLisa := Artwork{title: "Mona Lisa", artist: "Leonardo da Vinci", epoch: "Renaissance"}
	starryNight := Artwork{title: "Starry Night", artist: "Vincent van Gogh", epoch: "Post-Impressionism"}

	virtualGallery := &Gallery{}
	virtualGallery.addArtwork(monaLisa)
	virtualGallery.addArtwork(starryNight)

	virtualGallery.displayArtworks()
	virtualGallery.searchByArtist("Vincent van Gogh")
}
