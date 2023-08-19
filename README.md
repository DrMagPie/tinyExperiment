# tiny Experiment

Hardware Board Arduino Nano (with old bootloader)

Arduino implementation uses this [LiquidCrystal_I2C](https://github.com/johnrickman/LiquidCrystal_I2C) library.

tinyGo implementation uses original [hd44780i2c](https://github.com/tinygo-org/drivers/tree/release/hd44780i2c) driver with slight modification. I have exposed `sendData` method and renamed it to `Write` in order to have a closer match with arduino lib.