package main

import (
	"errors"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	player      *Player
	maps        *Maps
	items       []*Item
	enemy       []*Enemy
	playerSprite *ebiten.Image
	wallSprite   *ebiten.Image
	emptySprite  *ebiten.Image
	itemSprite   *ebiten.Image
	exitSprite   *ebiten.Image
	enemySprite  *ebiten.Image
}

func InitGame(player *Player, maps *Maps, items []*Item, enemy []*Enemy) *Game {
	return &Game{
		player: player,
		maps:   maps,
		items:  items,
		enemy:  enemy,
	}
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || ebiten.IsWindowBeingClosed() {
		//закрытие программы
		return errors.New("ты вышел из игры")
	}
	if g.player != nil {
		return g.player.Move(0, 0)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Отрисовка карты
	g.maps.Draw(screen, g.wallSprite, g.emptySprite, g.itemSprite, g.exitSprite, g.enemySprite)
	
	// Отрисовка игрока
	if g.player != nil {
		g.player.Draw(screen)
	}

	ebitenutil.DebugPrint(screen, "start")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	// Загрузка спрайтов один раз при запуске
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

	// Инициализация карты
	maps := InitMap()
	
	// Поиск начальной позиции игрока на карте
	playerX, playerY := -1, -1
	for y, row := range maps.Mmap {
		for x, cell := range row {
			if cell == 'P' {
				playerX, playerY = x, y
				break
			}
		}
		if playerX != -1 {
			break
		}
	}
	
	if playerX == -1 || playerY == -1 {
		log.Fatal("не найдена начальная позиция игрока на карте")
	}

	// Инициализация игрока
	player := InitPlayer(playerX, playerY, playerSprite, 0, 0, false, maps)
	
	// Заменяем символ 'P' на карте на пустую клетку '0', так как игрок отрисовывается отдельно
	maps.Mmap[playerY][playerX] = '0'

	game := &Game{
		player:       player,
		maps:         maps,
		items:        []*Item{},
		enemy:        []*Enemy{},
		playerSprite: playerSprite,
		wallSprite:   wallSprite,
		emptySprite:  emptySprite,
		itemSprite:   itemSprite,
		exitSprite:   exitSprite,
		enemySprite:  enemySprite,
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("servive in murino")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
