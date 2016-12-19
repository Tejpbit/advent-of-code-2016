package day10

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
	"strconv"
)

var betweenMatch = regexp.MustCompile("bot (\\d+) gives low to (output|bot) (\\d+) and high to (output|bot) (\\d+)")
var firstHandMatch = regexp.MustCompile("value (\\d+) goes to bot (\\d+)")

type Bot struct {
	input chan int //input channel for chip values
	id    string      // bot id
	low   chan int
	high  chan int
}

func createBot(id string) *Bot {
	return &Bot{make(chan int, 2), id, nil, nil}
}

var bots = make(map[string]*Bot)
var outputs = make(map[string]chan int)

func getBot(id string) *Bot {
	bot, ok := bots[id]
	if !ok {
		bot = createBot(id)
		bots[id] = bot
	}
	return bot
}

func (b *Bot) setupHandofChannels(lowType, highType, lowId, highId string) {
	if lowType == "output" {
		outputChan, ok := outputs[lowId]
		if !ok {
			outputChan = make(chan int, 1)
			outputs[lowId] = outputChan
		}
		b.low = outputChan
	} else if lowType == "bot" {
		receiverBot := getBot(lowId)
		b.low = receiverBot.input
	}

	if highType == "output" {
		outputChan, ok := outputs[highId]
		if !ok {
			outputChan = make(chan int, 1)
			outputs[highId] = outputChan
		}
		b.high = outputChan
	} else if highType == "bot" {
		receiverBot := getBot(highId)
		b.high = receiverBot.input
	}
}

func Run(input string, task int) {
	for _, line := range strings.Split(input, "\n") {
		res := betweenMatch.FindStringSubmatch(line)
		if len(res) != 0 {
			fromBotId := res[1]
			lowType, lowId := res[2], res[3]
			highType, highId := res[4], res[5]
			giverBot := getBot(fromBotId)
			giverBot.setupHandofChannels(lowType, highType, lowId, highId)
		}
		res = firstHandMatch.FindStringSubmatch(line)
		if len(res) != 0 {
			value, err := strconv.Atoi(res[1])
			if err != nil {
				panic(err)
			}
			toBotId := res[2]

			toBot, ok := bots[toBotId]
			if !ok {
				toBot = getBot(toBotId)
				bots[toBotId] = toBot
			}

			toBot.input <- value
		}
	}

	var wg sync.WaitGroup
	for _, e := range bots {
		wg.Add(1)
		go botRoutine(e, &wg)
	}
	wg.Add(1)
	go func () {
		val0 := <-outputs["0"]
		val1 := <-outputs["1"]
		val2 := <-outputs["2"]
		fmt.Printf("outpus 0+1+2: %d\n", val0 * val1 * val2)
	}()

	wg.Wait()

	fmt.Println("helo")

}

func botRoutine(b *Bot, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	v1 := <-b.input
	v2 := <-b.input
	if v1 == 61 && v2 == 17 || v1 == 17 && v2 == 61 {
		fmt.Printf("Bot %s comparing 17 and 61\n", b.id)
	}
	if v1 < v2 {
		b.low <- v1
		b.high <- v2
	} else {
		b.low <- v2
		b.high <- v1
	}
}
