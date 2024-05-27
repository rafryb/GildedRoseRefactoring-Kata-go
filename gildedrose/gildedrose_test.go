package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
	"github.com/stretchr/testify/assert"
)

func Test_Foo(t *testing.T) {
	var items = []*gildedrose.Item{
		{"foo", 0, 0},
	}

	gildedrose.UpdateQuality(items)

	if items[0].Name != "fixme" {
		t.Errorf("Name: Expected %s but got %s ", "fixme", items[0].Name)
	}
}

func Test_First_Day(t *testing.T) {
	var items = []*gildedrose.Item{
		{"+5 Dexterity Vest", 10, 20},
		{"Aged Brie", 2, 0},
		{"Elixir of the Mongoose", 5, 7},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
		{"Conjured Mana Cake", 3, 6}, // <-- :O
	}

	gildedrose.UpdateQuality(items)

	assert.Equal(t, items[0].SellIn, 9)
	assert.Equal(t, items[0].Quality, 19)

	assert.Equal(t, items[1].SellIn, 1)
	assert.Equal(t, items[1].Quality, 1)

	assert.Equal(t, items[2].SellIn, 4)
	assert.Equal(t, items[2].Quality, 6)

	assert.Equal(t, items[3].SellIn, 0)
	assert.Equal(t, items[3].Quality, 80)

	assert.Equal(t, items[4].SellIn, -1)
	assert.Equal(t, items[4].Quality, 80)

	assert.Equal(t, items[5].SellIn, 14)
	assert.Equal(t, items[5].Quality, 21)

	assert.Equal(t, items[6].SellIn, 9)
	assert.Equal(t, items[6].Quality, 50)

	assert.Equal(t, items[7].SellIn, 4)
	assert.Equal(t, items[7].Quality, 50)

	assert.Equal(t, items[8].SellIn, 2)
	assert.Equal(t, items[8].Quality, 5)
}

func Test_Thirty_Days(t *testing.T) {

	var items = []*gildedrose.Item{
		{"+5 Dexterity Vest", 10, 20},
		{"Aged Brie", 2, 0},
		{"Elixir of the Mongoose", 5, 7},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
		{"Conjured Mana Cake", 3, 6}, // <-- :O
	}

	days := 30
	for day := 1; day < days; day++ {
		gildedrose.UpdateQuality(items)
	}

	assert.Equal(t, items[0].SellIn, -19)
	assert.Equal(t, items[0].Quality, 0)

	assert.Equal(t, items[1].SellIn, -27)
	assert.Equal(t, items[1].Quality, 50)

	assert.Equal(t, items[2].SellIn, -24)
	assert.Equal(t, items[2].Quality, 0)

	assert.Equal(t, items[3].SellIn, 0)
	assert.Equal(t, items[3].Quality, 80)

	assert.Equal(t, items[4].SellIn, -1)
	assert.Equal(t, items[4].Quality, 80)

	assert.Equal(t, items[5].SellIn, -14)
	assert.Equal(t, items[5].Quality, 0)

	assert.Equal(t, items[6].SellIn, -19)
	assert.Equal(t, items[6].Quality, 0)

	assert.Equal(t, items[7].SellIn, -24)
	assert.Equal(t, items[7].Quality, 0)

	assert.Equal(t, items[8].SellIn, -26)
	assert.Equal(t, items[8].Quality, 0)

}
