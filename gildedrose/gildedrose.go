package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

const (
	agedBrie      = "Aged Brie"
	backstagePass = "Backstage passes to a TAFKAL80ETC concert"
	sulfuras      = "Sulfuras, Hand of Ragnaros"
	minQuality    = 0
	maxQuality    = 50
)

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		switch items[i].Name {
		case agedBrie:
			updateAgedBrie(items[i])
		case backstagePass:
			updateBackstagePass(items[i])
		case sulfuras:
			continue
		default:
			updateDefaultItem(items[i])
		}
	}
	//for i := 0; i < len(items); i++ {

	//	if items[i].Name != agedBrie && items[i].Name != backstagePass {
	//		if items[i].Quality > 0 {
	//			if items[i].Name != sulfuras {
	//				items[i].Quality = items[i].Quality - 1
	//			}
	//		}
	//	} else {
	//		if items[i].Quality < 50 {
	//			items[i].Quality = items[i].Quality + 1
	//			if items[i].Name == backstagePass {
	//				if items[i].SellIn < 11 {
	//					if items[i].Quality < 50 {
	//						items[i].Quality = items[i].Quality + 1
	//					}
	//				}
	//				if items[i].SellIn < 6 {
	//					if items[i].Quality < 50 {
	//						items[i].Quality = items[i].Quality + 1
	//					}
	//				}
	//			}
	//		}
	//	}
	//
	//	if items[i].Name != sulfuras {
	//		items[i].SellIn = items[i].SellIn - 1
	//	}
	//
	//	if items[i].SellIn < 0 {
	//		if items[i].Name != agedBrie {
	//			if items[i].Name != backstagePass {
	//				if items[i].Quality > 0 {
	//					if items[i].Name != sulfuras {
	//						items[i].Quality = items[i].Quality - 1
	//					}
	//				}
	//			} else {
	//				items[i].Quality = items[i].Quality - items[i].Quality
	//			}
	//		} else {
	//			if items[i].Quality < 50 {
	//				items[i].Quality = items[i].Quality + 1
	//			}
	//		}
	//	}
	//}

}

func updateAgedBrie(item *Item) {
	item.Quality = agedBrieQuality(item)
	item.SellIn -= 1
}

func agedBrieQuality(item *Item) int {
	if item.Quality == maxQuality {
		return item.Quality
	}

	if item.SellIn > minQuality {
		return item.Quality + 1
	} else {
		return item.Quality + 2
	}
}

func updateBackstagePass(item *Item) {

}

func updateDefaultItem(item *Item) {

}
