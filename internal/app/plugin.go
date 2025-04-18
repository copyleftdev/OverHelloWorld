package app

import "log"

type Plugin interface {
	OnHelloSaid(message string)
}

type ASCIIPlugin struct{}
func (p *ASCIIPlugin) OnHelloSaid(message string) {
	log.Printf("[ASCII Plugin] %s", message)
}

type TTSPlugin struct{}
func (p *TTSPlugin) OnHelloSaid(message string) {
	log.Printf("[TTS Plugin] (speaking): %s", message)
}

type LEDPlugin struct{}
func (p *LEDPlugin) OnHelloSaid(message string) {
	log.Printf("[LED Plugin] (flashing LEDs for): %s", message)
}
