package mapobject

import (
	"mapserver/coords"
	"mapserver/mapobjectdb"

	"github.com/minetest-go/mapparser"
)

type PoiBlock struct {
	Color string
}

func (this *PoiBlock) onMapObject(mbpos *coords.MapBlockCoords, x, y, z int, block *mapparser.MapBlock) *mapobjectdb.MapObject {
	md := block.Metadata.GetMetadata(x, y, z)

	o := mapobjectdb.NewMapObject(mbpos, x, y, z, "poi")
	o.Attributes["name"] = md["name"]
	o.Attributes["category"] = md["category"]
	o.Attributes["url"] = md["url"]
	o.Attributes["owner"] = md["owner"]
	o.Attributes["icon"] = md["icon"]
	o.Attributes["color"] = this.Color

	return o
}
