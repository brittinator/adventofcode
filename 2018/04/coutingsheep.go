package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	f, err := os.Open("./input.txt")
	defer f.Close()
	if err != nil {
		log.Fatal("can't open file ", err)
	}

	var allLogs entries

	scanner := bufio.NewScanner(f)
	// put logs into memory
	for scanner.Scan() {
		l := scanner.Text()
		separated := strings.Split(l, "] ")

		date := separated[0]
		message := separated[1]

		newLog := newLogEntry(date, message)
		fmt.Println(l, newLog.date.Format(layout), newLog.msg)
		allLogs = append(allLogs, newLog)
	}

	// first, order the logs by date and time
	sort.Sort(allLogs)
	for _, log := range allLogs {
		fmt.Println(log.date.Format(layout), log.msg)
	}

	strategyOne(allLogs)

}

type logEntry struct {
	date time.Time
	msg  string
}

const layout = "[2006-01-02 15:04"

func newLogEntry(timeStamp, msg string) logEntry {
	// parse the date [year-month-day hour:minute] [1518-11-03 00:36]
	date, err := time.Parse(layout, timeStamp)
	if err != nil {
		log.Fatal("error converting time", err)
	}
	return logEntry{
		date: date,
		msg:  msg,
	}
}

func strategyOne(allLogs []logEntry) {
	// find the guard who is asleep the most
	type sleepTime struct {
		start, end int
	}

	type guard struct {
		id            string
		asleepRange   []sleepTime
		minutesAsleep int
	}

	allGuards := make(map[string]guard)
	var currentGuard guard
	var asleepTime sleepTime

	for _, l := range allLogs {
		fmt.Println("current guard: ", currentGuard.id)
		fmt.Println(l.msg, asleepTime)
		// fmt.Println(l, l.date.Format(layout), l.msg)

		if strings.Contains(l.msg, "Guard") {
			// zero out sleep time
			asleepTime = sleepTime{}
			contents := strings.Split(l.msg, " ")
			var id string
			for _, text := range contents {
				if strings.Contains(text, "#") {
					id = text
					break
				}
			}

			// find guard by id
			if g, ok := allGuards[id]; ok {
				fmt.Println("guard already found: ", g, id)
				currentGuard = g
				continue
			} else {
				// start a new guard entry
				currentGuard = guard{
					id:          id,
					asleepRange: []sleepTime{},
				}
				fmt.Println("new guard> ", id, l.msg, len(allGuards))
				// add to list
				allGuards[id] = currentGuard
			}
		} else {
			fmt.Println("in else, currentGuard: ", currentGuard, l.date.Format(layout), l.msg)
			// continue with current entry
			if l.msg == "falls asleep" {
				asleepTime.start = l.date.Minute()
				fmt.Println("falls asleep")
				// currentGuard.asleepRange = append(
				// 	currentGuard.asleepRange, sleepTime{start: l.date.Minute()},
				// )
			}
			if l.msg == "wakes up" {
				fmt.Println("wakes up")
				// stop the sleep time
				asleepTime.end = l.date.Minute()
				// fmt.Println("last sleep: ", asleepTime)
				currentGuard.asleepRange = append(currentGuard.asleepRange, asleepTime)
				fmt.Println(currentGuard.asleepRange)
				allGuards[currentGuard.id] = currentGuard
			}
		}
	}

	for i, g := range allGuards {
		fmt.Println("guard ", g.id)
		var totalZZZs int
		for _, r := range g.asleepRange {
			fmt.Printf("	%v --> %v\n", r.start, r.end)
			totalZZZs = totalZZZs + r.end - r.start
		}
		// fmt.Println("total amount slept: ", totalZZZs)
		g.minutesAsleep = totalZZZs
		allGuards[i] = g
	}

	var highScore int
	var guardID string
	for _, g := range allGuards {
		// fmt.Println("guard: ", g)
		if g.minutesAsleep > highScore {
			highScore = g.minutesAsleep
			guardID = g.id
		}
	}

	fmt.Println("guard that slept the most: ", guardID, highScore)
	// find the minute that guard is asleep the most
	sleepFrequency := make(map[int]int)
	chosenGuard := allGuards[guardID]
	fmt.Println(chosenGuard)

	for _, sleepTimes := range chosenGuard.asleepRange {
		for i := sleepTimes.start; i < sleepTimes.end; i++ {
			sleepFrequency[i]++
		}
	}
	fmt.Println("sleep frequency: ", sleepFrequency)
	var highestProbabilityForSnooze int
	var value int
	for k, newVal := range sleepFrequency {
		if newVal > value {
			value = newVal
			highestProbabilityForSnooze = k
		}
	}

	fmt.Println("best time is: ", highestProbabilityForSnooze, value)

	fmt.Printf("answer %v id x %v minute\n", guardID, highestProbabilityForSnooze)
	// capture time windows

	// [1518-11-21 23:50 Guard #2953 begins shift
	// [1518-11-22 00:04 falls asleep
	// [1518-11-22 00:37 wakes up
	// [4...37]

	// [1518-11-22 00:51 falls asleep
	// [1518-11-22 00:53 wakes up
	// [51..53]

	// so this day would be [4..37, 51..53] (non-inclusive of the last minute)
	// (37-4) + (53-51)
	// for a total of 35 minutes asleep

	// [1518-09-30 00:00 Guard #2693 begins shift
	// [1518-09-30 00:33 falls asleep
	// [1518-09-30 00:40 wakes up
	// [33..40]

	// [1518-09-30 00:46 falls asleep
	// [1518-09-30 00:58 wakes up
	// [46..58]

	// so this day would be [33..40, 46..58]
	// (40-33) + (58-46) = 7+ 12 = 19 min asleep

	// the overlap would be
	// [4..37, 33..40, 46..58, 51..53]
	// 51,52 minutes the guard is asleep both days
}

// entries sorts the logs by date.
type entries []logEntry

func (l entries) Len() int {
	return len(l)
}

func (l entries) Less(i, j int) bool {
	iDate := l[i].date
	jDate := l[j].date
	return iDate.Sub(jDate) < 0
}

func (l entries) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
