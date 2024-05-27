package gildedrose

import "strings"

type Item struct {
	Name            string
	SellIn, Quality int
}

const (
	agedBrie      = "Aged Brie"
	backstagePass = "Backstage passes"
	sulfuras      = "Sulfuras"
	conjured      = "Conjured"
	minQuality    = 0
	maxQuality    = 50
)

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		switch {
		case strings.Contains(items[i].Name, agedBrie):
			updateAgedBrie(items[i])
		case strings.Contains(items[i].Name, backstagePass):
			updateBackstagePass(items[i])
		case strings.Contains(items[i].Name, sulfuras):
			continue
		case strings.Contains(items[i].Name, conjured):
			updateConjured(items[i])
		default:
			updateDefaultItem(items[i])
		}
	}
}

func updateAgedBrie(item *Item) {
	item.Quality = agedBrieQuality(item)
	item.SellIn -= 1
}

func agedBrieQuality(item *Item) int {
	if item.Quality == maxQuality {
		return item.Quality
	}

	if item.SellIn > 0 {
		return sumQuality(item, 1)
	} else {
		return sumQuality(item, 2)
	}
}

func updateBackstagePass(item *Item) {
	item.Quality = backstagePassQuality(item)
	item.SellIn -= 1
}

func backstagePassQuality(item *Item) int {
	if item.Quality == maxQuality && item.SellIn > 0 {
		return item.Quality
	}

	if item.SellIn > 10 {
		return sumQuality(item, 1)
	}

	if item.SellIn > 5 {
		return sumQuality(item, 2)
	}

	if item.SellIn > 0 {
		return sumQuality(item, 3)
	}

	return 0
}

func updateDefaultItem(item *Item) {
	item.Quality = defaultItemQuality(item)
	item.SellIn -= 1
}

func defaultItemQuality(item *Item) int {
	if item.Quality == minQuality {
		return item.Quality
	}

	if item.SellIn > 0 {
		return sumQuality(item, -1)
	} else {
		return sumQuality(item, -2)
	}
}

func updateConjured(item *Item) {
	item.Quality = conjuredQuality(item)
	item.SellIn -= 1
}

func conjuredQuality(item *Item) int {
	if item.Quality == minQuality {
		return item.Quality
	}

	if item.SellIn > 0 {
		return sumQuality(item, -2)
	} else {
		return sumQuality(item, -4)
	}
}

func sumQuality(item *Item, quality int) int {
	sum := item.Quality + quality

	if sum > maxQuality {
		return maxQuality
	}

	return sum
}
