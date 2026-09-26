import pygame
import sys
import os

SCREEN_WIDTH = 800
SCREEN_HEIGHT = 600
TILE_SIZE = 40 

COLOR_WHITE = (255, 255, 255)
COLOR_GOLD = (255, 215, 0)
COLOR_RED = (200, 0, 0)

#путь к папке с материалами
try:
    ASSET_DIR = os.path.dirname(os.path.abspath(__file__))
except NameError:
    ASSET_DIR = os.path.abspath('.')

# функция для загрузки изображений
def load_image(filename):
    path = os.path.join(ASSET_DIR, filename)
    try:
        image = pygame.image.load(path).convert_alpha()
        return pygame.transform.scale(image, (TILE_SIZE, TILE_SIZE))
    except pygame.error:
        print(f"Ошибка: Не удалось загрузить изображение '{filename}' по пути '{path}'")
        sys.exit()

class Player(pygame.sprite.Sprite):
    def __init__(self, x, y):
        super().__init__()
        self.image = load_image('player.png')
        self.rect = self.image.get_rect(topleft=(x, y))

class Wall(pygame.sprite.Sprite):
    def __init__(self, x, y):
        super().__init__()
        self.image = load_image('wall.png')
        self.rect = self.image.get_rect(topleft=(x, y))

class Collectible(pygame.sprite.Sprite):
    def __init__(self, x, y):
        super().__init__()
        self.image = load_image('collectible.png') 
        self.rect = self.image.get_rect(topleft=(x, y))

class Exit(pygame.sprite.Sprite):
    def __init__(self, x, y):
        super().__init__()
        self.image = load_image('exit.png') 
        self.rect = self.image.get_rect(topleft=(x, y))

class Enemy(pygame.sprite.Sprite):
    def __init__(self, x, y):
        super().__init__()
        self.image = load_image('enemy.png') 
        self.rect = self.image.get_rect(topleft=(x, y))




class Game:
    def __init__(self):
        pygame.init()
        self.screen = pygame.display.set_mode((SCREEN_WIDTH, SCREEN_HEIGHT))
        pygame.display.set_caption("Моя 2D Игра")
        self.clock = pygame.time.Clock()
        
        # загружаем фоновое изображение
        bg_path = os.path.join(ASSET_DIR, 'background.png') 
        try:
            background_image_raw = pygame.image.load(bg_path).convert()
            self.background_image = pygame.transform.scale(background_image_raw, (SCREEN_WIDTH, SCREEN_HEIGHT))
        except pygame.error:
            print(f"Ошибка: Не удалось загрузить фон '{bg_path}'")
            sys.exit()

        self.font = pygame.font.Font(None, 50)
        
        self.game_over = False
        self.win = False
        self.moves = 0
        self.collectibles_count = 0
        self.load_map("map.txt")

    def load_map(self, filename):
        self.all_sprites = pygame.sprite.Group()
        self.walls = pygame.sprite.Group()
        self.collectibles = pygame.sprite.Group()
        self.exits = pygame.sprite.Group()
        self.enemies = pygame.sprite.Group()

        map_path = os.path.join(ASSET_DIR, filename)
        try:
            with open(map_path, 'r') as f:
                for y, line in enumerate(f):
                    for x, char in enumerate(line.strip()):
                        if char == '1':
                            self.all_sprites.add(Wall(x * TILE_SIZE, y * TILE_SIZE))
                            self.walls.add(self.all_sprites.sprites()[-1])
                        elif char == 'P':
                            self.player = Player(x * TILE_SIZE, y * TILE_SIZE)
                            self.all_sprites.add(self.player)
                        elif char == 'C':
                            self.all_sprites.add(Collectible(x * TILE_SIZE, y * TILE_SIZE))
                            self.collectibles.add(self.all_sprites.sprites()[-1])
                            self.collectibles_count += 1
                        elif char == 'E':
                            self.all_sprites.add(Exit(x * TILE_SIZE, y * TILE_SIZE))
                            self.exits.add(self.all_sprites.sprites()[-1])
                        elif char == 'X':
                            self.all_sprites.add(Enemy(x * TILE_SIZE, y * TILE_SIZE))
                            self.enemies.add(self.all_sprites.sprites()[-1])
        except FileNotFoundError:
            print(f"Ошибка: Файл карты '{filename}' не найден в папке проекта!")
            sys.exit()


    def run(self):
        while not self.game_over:
            self.events()
            self.update()
            self.draw()
            self.clock.tick(60)
        self.show_final_screen()

    def events(self):
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                self.game_over = True
            if event.type == pygame.KEYDOWN:
                if event.key == pygame.K_ESCAPE:
                    self.game_over = True

                new_x, new_y = self.player.rect.x, self.player.rect.y
                moved = False
                if event.key == pygame.K_w:
                    new_y -= TILE_SIZE
                    moved = True
                elif event.key == pygame.K_s:
                    new_y += TILE_SIZE
                    moved = True
                elif event.key == pygame.K_a:
                    new_x -= TILE_SIZE
                    moved = True
                elif event.key == pygame.K_d:
                    new_x += TILE_SIZE
                    moved = True

                if moved:
                    if not self.check_collision(new_x, new_y):
                        self.player.rect.x = new_x
                        self.player.rect.y = new_y
                        self.moves += 1

    def check_collision(self, x, y):
        temp_rect = pygame.Rect(x, y, TILE_SIZE, TILE_SIZE)
        for wall in self.walls:
            if temp_rect.colliderect(wall.rect):
                return True
        return False

    def update(self):
        collected = pygame.sprite.spritecollide(self.player, self.collectibles, True)
        if collected:
            self.collectibles_count -= len(collected)

        if self.collectibles_count == 0:
            if pygame.sprite.spritecollide(self.player, self.exits, False):
                self.win = True
                self.game_over = True

        if pygame.sprite.spritecollide(self.player, self.enemies, False):
            self.win = False
            self.game_over = True

    def draw_text(self, text, color, x, y):
        text_surface = self.font.render(text, True, color)
        text_rect = text_surface.get_rect(topleft=(x, y))
        self.screen.blit(text_surface, text_rect)
    def draw(self):
        self.screen.blit(self.background_image, (0, 0))
        self.all_sprites.draw(self.screen)
        self.draw_text(f"MOVES: {self.moves}", COLOR_WHITE, 20, 10)
        pygame.display.flip()

    def show_final_screen(self):
        final_font = pygame.font.Font(None, 100)
        if self.win:
            text = final_font.render("YOU WIN!", True, COLOR_GOLD)
        else:
            text = final_font.render("GAME OVER", True, COLOR_RED)
        
        text_rect = text.get_rect(center=(SCREEN_WIDTH / 2, SCREEN_HEIGHT / 2))
        self.screen.blit(text, text_rect)
        pygame.display.flip()
        
        waiting = True
        while waiting:
            for event in pygame.event.get():
                if event.type == pygame.QUIT or event.type == pygame.KEYDOWN:
                    waiting = False

if __name__ == '__main__':
    game = Game()
    game.run()
    pygame.quit()
    sys.exit()