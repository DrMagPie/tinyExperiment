#include <LiquidCrystal_I2C.h>

#define DISPLAY_ADDR 0x27
LiquidCrystal_I2C lcd(DISPLAY_ADDR, 20, 4);

uint8_t x = 3;
uint8_t y = 1;

void setup()
{
  lcd.init();
  lcd.backlight();
  lcd.clear();

  uint8_t LT[8] = {0b00111, 0b01111, 0b11111, 0b11111, 0b11111, 0b11111, 0b11111, 0b11111};
  uint8_t UB[8] = {0b11111, 0b11111, 0b11111, 0b00000, 0b00000, 0b00000, 0b00000, 0b00000};
  uint8_t RT[8] = {0b11100, 0b11110, 0b11111, 0b11111, 0b11111, 0b11111, 0b11111, 0b11111};
  uint8_t LL[8] = {0b11111, 0b11111, 0b11111, 0b11111, 0b11111, 0b11111, 0b01111, 0b00111};
  uint8_t LB[8] = {0b00000, 0b00000, 0b00000, 0b00000, 0b00000, 0b11111, 0b11111, 0b11111};
  uint8_t LR[8] = {0b11111, 0b11111, 0b11111, 0b11111, 0b11111, 0b11111, 0b11110, 0b11100};
  uint8_t UMB[8] = {0b11111, 0b11111, 0b11111, 0b00000, 0b00000, 0b00000, 0b11111, 0b11111};
  uint8_t LMB[8] = {0b11111, 0b00000, 0b00000, 0b00000, 0b00000, 0b11111, 0b11111, 0b11111};

  lcd.createChar(0, LT);
  lcd.createChar(1, UB);
  lcd.createChar(2, RT);
  lcd.createChar(3, LL);
  lcd.createChar(4, LB);
  lcd.createChar(5, LR);
  lcd.createChar(6, UMB);
  lcd.createChar(7, LMB);

  lcd.setCursor(x + 1, y);
  lcd.write(1);
  lcd.write(2);
  lcd.setCursor(x + 2, y + 1);
  lcd.write(5);

  lcd.setCursor(x + 4, y);
  lcd.write(0);
  lcd.write(6);
  lcd.write(6);
  lcd.setCursor(x + 4, y + 1);
  lcd.write(7);
  lcd.write(7);
  lcd.write(5);

  lcd.setCursor(x + 8, y);
  lcd.write(3);
  lcd.write(4);
  lcd.write(2);
  lcd.setCursor(x + 8 + 2, y + 1);
  lcd.write(5);

  lcd.setCursor(x + 12, y);
  lcd.write(3);
  lcd.write(4);
  lcd.write(2);
  lcd.setCursor(x + 12 + 2, y + 1);
  lcd.write(5);
}

void loop()
{
  lcd.setCursor(x + 7, y);
  lcd.write(165);
  lcd.setCursor(x + 7, y + 1);
  lcd.write(165);
  delay(500);

  lcd.setCursor(x + 7, y);
  lcd.write(32);
  lcd.setCursor(x + 7, y + 1);
  lcd.write(32);
  delay(500);
}
