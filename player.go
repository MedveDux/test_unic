package main

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	x, y        int
	sprite      *ebiten.Image
	moveCounter int // счётчик перемещений
	gainedItems int // собранные предметы
	isAtExit    bool
	gameMap     *Maps
}

func InitPlayer(x, y int, sprite *ebiten.Image, moveCounter int, gainedItems int, isAtExit bool, gameMap *Maps) *Player { //конструктор
	return &Player{
		x:           x,
		y:           y,
		sprite:      sprite,
		moveCounter: 0,
		gainedItems: 0,
		isAtExit:    false,
		gameMap:     gameMap,
	}
}

func (p *Player) Move(x, y int) error {
	xx, yy := p.x, p.y
	moved := false

	switch {
	case ebiten.IsKeyPressed(ebiten.KeyW): //вверх
		yy--
		moved = true
	case ebiten.IsKeyPressed(ebiten.KeyS): //вниз
		yy++
		moved = true
	case ebiten.IsKeyPressed(ebiten.KeyD): //вправо
		xx++
		moved = true
	case ebiten.IsKeyPressed(ebiten.KeyA): //влево
		xx--
		moved = true
	}

	// Если клавиша не нажата, не двигаемся
	if !moved {
		return nil
	}

	if xx >= 0 && xx < len(p.gameMap.Mmap[0]) && yy >= 0 && yy < len(p.gameMap.Mmap) { //проверка на границы карты
		cell := p.gameMap.Mmap[yy][xx]
		if cell != '1' { //если это не стена
			p.x, p.y = xx, yy
			p.moveCounter++
			log.Printf("сделано %d\n шагов", p.moveCounter)

			//проверка на предмет сбора
			if cell == 'C' {
				p.gainedItems++
				p.gameMap.Mmap[yy][xx] = '0' // Удаляем предмет с карты
				log.Printf("собрано %d\n предметов", p.gainedItems)
			}

			//проверка на внезапный выход из игры
			if cell == 'E' {
				p.isAtExit = true
				log.Printf("упс... игра прервана")
			}

			//проверка на столкновение с врагом
			if cell == 'X' {
				return errors.New("ты наткнулся на врага, игра завершена")
			}
		}
	}
	return nil
}

func (p *Player) Update() {}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.x)*32, float64(p.y)*32)
	screen.DrawImage(p.sprite, op)
}

func (p *Player) Reset(x, y int) { //сброс игры при попадании в начальную позицию
	xx, yy := x, y
	cell := p.gameMap.Mmap[yy][xx]
	if cell == 'P' {
		p.moveCounter = 0
		p.gainedItems = 0
	}
}

func (p *Player) ReturnMoveCounter() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("пройдено %d\n шагов", p.moveCounter))
	return sb.String()
}
