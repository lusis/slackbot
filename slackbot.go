package slack

import (
	"github.com/gorilla/websocket"
	"io"
	"io/ioutil"
)

type SlackBot struct {
	allEventsHandlers           []func(Event, []byte)
	unknownEventHandlers        []func(Event, []byte)
	helloEventHandlers          []func(HelloEvent)
	presenceChangeEventHandlers []func(PresenceChangeEvent)
	messageEventHandlers        []func(MessageEvent)

	token string
	conn  *websocket.Conn

	started bool
}

func NewSlackBot() SlackBot {
	s := SlackBot{}
	s.started = false
	s.unknownEventHandlers = make([]func(Event, []byte), 0)
	s.helloEventHandlers = make([]func(HelloEvent), 0)
	s.presenceChangeEventHandlers = make([]func(PresenceChangeEvent), 0)
	s.messageEventHandlers = make([]func(MessageEvent), 0)

	return s
}

func (s SlackBot) Start() error {

	conn, err := GetSlackRtm(s.token)
	if err != nil {
		return err
	}
	s.conn = conn

	s.started = true

	s.runLoop()
	return err

}

func (s SlackBot) runLoop() error {
	c := s.conn
	for {
		_, r, err := c.NextReader()

		if err != nil {
			return err
		}

		err = s.parseEvent(r)
		if err != nil {
			return err
		}

	}
}

func (s SlackBot) Token() string {
	return s.token
}
func (s *SlackBot) SetToken(tok string) {
	s.token = tok
}

func (s SlackBot) triggerAllEvents(evt Event, evtstring []byte) {
	for _, handler := range s.allEventsHandlers {
		handler(evt, evtstring)
	}
}
func (s *SlackBot) OnAllEvents(handler func(Event, []byte)) {
	s.unknownEventHandlers = append(s.allEventsHandlers, handler)
}

func (s SlackBot) triggerUnknownEvents(evt Event, evtstring []byte) {
	for _, handler := range s.unknownEventHandlers {
		handler(evt, evtstring)
	}
}
func (s *SlackBot) OnUnknownEvents(handler func(Event, []byte)) {
	s.unknownEventHandlers = append(s.unknownEventHandlers, handler)
}

func (s SlackBot) triggerHelloEvents(evt HelloEvent) {
	for _, handler := range s.helloEventHandlers {
		handler(evt)
	}
}
func (s *SlackBot) OnHelloEvents(handler func(HelloEvent)) {
	s.helloEventHandlers = append(s.helloEventHandlers, handler)
}

func (s SlackBot) triggerPresenceChangeEvents(evt PresenceChangeEvent) {
	for _, handler := range s.presenceChangeEventHandlers {
		handler(evt)
	}
}
func (s *SlackBot) OnPresenceChangeEvents(handler func(PresenceChangeEvent)) {
	s.presenceChangeEventHandlers = append(s.presenceChangeEventHandlers, handler)
}

func (s SlackBot) triggerMessageEvents(evt MessageEvent) {
	for _, handler := range s.messageEventHandlers {
		handler(evt)
	}
}

func (s *SlackBot) OnMessageEvents(handler func(MessageEvent)) {
	s.messageEventHandlers = append(s.messageEventHandlers, handler)
}

func (s SlackBot) parseEvent(evtReader io.Reader) error {
	evtstring, _ := ioutil.ReadAll(evtReader)

	genericEvt := parseEvent(evtstring)

	s.triggerAllEvents(genericEvt, evtstring)

	switch genericEvt.Type {
	default:
		s.triggerUnknownEvents(genericEvt, evtstring)
	case "hello":
		evt := parseHelloEvent(evtstring)
		s.triggerHelloEvents(evt)
	case "presence_change":
		evt := parsePresenceChangeEvent(evtstring)
		s.triggerPresenceChangeEvents(evt)
	case "message":
		evt := parseMessageEvent(evtstring)
		s.triggerMessageEvents(evt)
	}

	return nil
}
