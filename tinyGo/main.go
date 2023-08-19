package main

import (
	"machine"
	"time"

	"github.com/DrMagPie/tinyExperiment/tinyGo/drivers/hd44780i2c"
	// "tinygo.org/x/drivers/hd44780i2c"
)

func main() {
	machine.I2C0.Configure(machine.I2CConfig{})

	lcd := hd44780i2c.New(machine.I2C0, 0x27)
	lcd.Configure(hd44780i2c.Config{Width: 20, Height: 4})
	lcd.BacklightOn(true)
	lcd.ClearDisplay()

	lcd.CreateCharacter(0, []byte{0b00111, 0b01111, 0b11111, 0b11111, 0b11111, 0b11111, 0b11111, 0b11111})
	lcd.CreateCharacter(1, []byte{0b11111, 0b11111, 0b11111, 0b00000, 0b00000, 0b00000, 0b00000, 0b00000})
	lcd.CreateCharacter(2, []byte{0b11100, 0b11110, 0b11111, 0b11111, 0b11111, 0b11111, 0b11111, 0b11111})
	lcd.CreateCharacter(3, []byte{0b11111, 0b11111, 0b11111, 0b11111, 0b11111, 0b11111, 0b01111, 0b00111})
	lcd.CreateCharacter(4, []byte{0b00000, 0b00000, 0b00000, 0b00000, 0b00000, 0b11111, 0b11111, 0b11111})
	lcd.CreateCharacter(5, []byte{0b11111, 0b11111, 0b11111, 0b11111, 0b11111, 0b11111, 0b11110, 0b11100})
	lcd.CreateCharacter(6, []byte{0b11111, 0b11111, 0b11111, 0b00000, 0b00000, 0b00000, 0b11111, 0b11111})
	lcd.CreateCharacter(7, []byte{0b11111, 0b00000, 0b00000, 0b00000, 0b00000, 0b11111, 0b11111, 0b11111})

	var x, y uint8 = 3, 1

	lcd.SetCursor(x+1, y)
	lcd.Write(1)
	lcd.Write(2)
	lcd.SetCursor(x+2, y+1)
	lcd.Write(5)

	lcd.SetCursor(x+4, y)
	lcd.Write(0)
	lcd.Write(6)
	lcd.Write(6)
	lcd.SetCursor(x+4, y+1)
	lcd.Write(7)
	lcd.Write(7)
	lcd.Write(5)

	lcd.SetCursor(x+8, y)
	lcd.Write(3)
	lcd.Write(4)
	lcd.Write(2)
	lcd.SetCursor(x+8+2, y+1)
	lcd.Write(5)

	lcd.SetCursor(x+12, y)
	lcd.Write(6)
	lcd.Write(6)
	lcd.Write(2)
	lcd.SetCursor(x+12, y+1)
	lcd.Write(7)
	lcd.Write(7)
	lcd.Write(5)

	for {
		lcd.SetCursor(x+7, y)
		lcd.Write(165)
		lcd.SetCursor(x+7, y+1)
		lcd.Write(165)
		time.Sleep(time.Second / 2)

		lcd.SetCursor(x+7, y)
		lcd.Write(32)
		lcd.SetCursor(x+7, y+1)
		lcd.Write(32)
		time.Sleep(time.Second / 2)
	}
}
