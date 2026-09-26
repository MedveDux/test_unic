package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func Load() (*ebiten.Image, *ebiten.Image, *ebiten.Image, *ebiten.Image, *ebiten.Image, *ebiten.Image) {

	playerSprite, _, err := ebitenutil.NewImageFromFile("asserts\\images\\player.png")
	if err != nil {
		log.Fatal("не удалось загрузить изображение игрока", err)
	}

	itemSprite, _, err := ebitenutil.NewImageFromFile("asserts\\images\\item.png")
	if err != nil {
		log.Fatal("не удалось загрузить изображение предмета", err)
	}

	wallSprite, _, err := ebitenutil.NewImageFromFile("asserts\\images\\wall.png")
	if err != nil {
		log.Fatal("не удалось загрузить изображение стены", err)
	}

	exitSprite, _, err := ebitenutil.NewImageFromFile("asserts\\images\\exit.png")
	if err != nil {
		log.Fatal("не удалось загрузить изображение выхода", err)
	}

	emptySprite, _, err := ebitenutil.NewImageFromFile("asserts\\images\\empty.png")
	if err != nil {
		log.Fatal("не удалось загрузить изображение пустого пространства", err)
	}

	enemySprite, _, err := ebitenutil.NewImageFromFile("asserts\\images\\enemy.png")
	if err != nil {
		log.Fatal("не удалось загрузить изображение врага", err)
	}

	return playerSprite, itemSprite, wallSprite, exitSprite, emptySprite, enemySprite

}
