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
	// there are quite a few things to unpack here.

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

	// date is within []
	// [year-month-day hour:minute]
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

func strategyOne() {

	// find the guard who is asleep the most
	// find the minute that guard is asleep the most
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
