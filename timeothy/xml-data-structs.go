package timeothy

import "encoding/xml"

// Map was generated 2022-07-09 15:23:45 by sparksb on sparksb10u.
type TiledMap struct {
	XMLName      xml.Name `xml:"map"`
	Text         string   `xml:",chardata"`
	Version      string   `xml:"version,attr"`
	Tiledversion string   `xml:"tiledversion,attr"`
	Orientation  string   `xml:"orientation,attr"`
	Renderorder  string   `xml:"renderorder,attr"`
	Width        int      `xml:"width,attr"`
	Height       int      `xml:"height,attr"`
	Tilewidth    int      `xml:"tilewidth,attr"`
	Tileheight   int      `xml:"tileheight,attr"`
	Infinite     int      `xml:"infinite,attr"`
	Nextlayerid  int      `xml:"nextlayerid,attr"`
	Nextobjectid int      `xml:"nextobjectid,attr"`
	Tileset      struct {
		Text     string `xml:",chardata"`
		Firstgid int    `xml:"firstgid,attr"`
		Source   string `xml:"source,attr"`
	} `xml:"tileset"`
	Layer struct {
		Text   string `xml:",chardata"`
		ID     int    `xml:"id,attr"`
		Name   string `xml:"name,attr"`
		Width  int    `xml:"width,attr"`
		Height int    `xml:"height,attr"`
		Data   struct {
			Text     string `xml:",chardata"`
			Encoding string `xml:"encoding,attr"`
		} `xml:"data"`
	} `xml:"layer"`
}

// Tileset was generated 2022-07-09 15:24:06 by sparksb on sparksb10u.
type TiledTileset struct {
	XMLName      xml.Name `xml:"tileset"`
	Text         string   `xml:",chardata"`
	Version      string   `xml:"version,attr"`
	Tiledversion string   `xml:"tiledversion,attr"`
	Name         string   `xml:"name,attr"`
	Tilewidth    int      `xml:"tilewidth,attr"`
	Tileheight   int      `xml:"tileheight,attr"`
	Spacing      int      `xml:"spacing,attr"`
	Tilecount    int      `xml:"tilecount,attr"`
	Columns      int      `xml:"columns,attr"`
	Image        struct {
		Text   string `xml:",chardata"`
		Source string `xml:"source,attr"`
		Trans  string `xml:"trans,attr"`
		Width  int    `xml:"width,attr"`
		Height int    `xml:"height,attr"`
	} `xml:"image"`
}
