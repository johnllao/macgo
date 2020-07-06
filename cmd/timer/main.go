package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"time"


	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type Sound struct {
	buff *bytes.Buffer
}

func (s Sound) Read(p []byte) (int, error) {
	return s.buff.Read(p)
}

func (s Sound) Close() error {
	return nil
}

var (
	d string
)

func init() {
	flag.StringVar(&d, "d", "10s", "duration of the timer")
	flag.Parse()
}

func main() {

	var err error
	var dd time.Duration

	dd, err = time.ParseDuration(d)
	if err != nil {
		fmt.Println("ERR: ", err)
		os.Exit(99)
	}

	fmt.Println("timer started")
	var t = time.NewTimer(dd)

	<-t.C
	fmt.Println("Time is up!")
	play()

}

func play() {
	var err error

	var soundb []byte
	soundb, err = base64.StdEncoding.DecodeString(sound)
	if err != nil {
		fmt.Println("ERR", err)
		os.Exit(99)
	}

	var s = Sound{
		buff: bytes.NewBuffer(soundb),
	}

	var strm beep.StreamSeekCloser
	var formt beep.Format
	strm, formt, err = mp3.Decode(s)
	if err != nil {
		fmt.Println("ERR: ", err)
		os.Exit(99)
	}
	defer strm.Close()

	err = speaker.Init(formt.SampleRate, formt.SampleRate.N(time.Second / 10))
	if err != nil {
		fmt.Println("ERR: ", err)
		os.Exit(99)
	}
	var done = make(chan bool)
	speaker.Play(beep.Seq(strm, beep.Callback(func(){
		done <- true
	})))
	<-done
}