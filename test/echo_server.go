package main

import (
	"fmt"
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/cracker"
	"github.com/cbusbey/quickfixgo/fix"
	"github.com/cbusbey/quickfixgo/fix40"
	"github.com/cbusbey/quickfixgo/fix41"
	"github.com/cbusbey/quickfixgo/fix42"
	"github.com/cbusbey/quickfixgo/fix43"
	"github.com/cbusbey/quickfixgo/fix44"
	"github.com/cbusbey/quickfixgo/fix50"
	"github.com/cbusbey/quickfixgo/fix50sp1"
	"github.com/cbusbey/quickfixgo/fix50sp2"
	"github.com/cbusbey/quickfixgo/log"
	"github.com/cbusbey/quickfixgo/message"
	"github.com/cbusbey/quickfixgo/reject"
	"github.com/cbusbey/quickfixgo/settings"
	"github.com/cbusbey/quickfixgo/tag"
	"os"
	"os/signal"
)

type EchoApplication struct {
	cracker.MessageCracker
	OrderIds map[string]bool
}

func (e EchoApplication) OnCreate(sessionID quickfixgo.SessionID) {
	fmt.Printf("OnCreate %v\n", sessionID.String())
}
func (e *EchoApplication) OnLogon(sessionID quickfixgo.SessionID) {
	fmt.Printf("OnLogon %v\n", sessionID.String())
	e.OrderIds = make(map[string]bool)
}
func (e *EchoApplication) OnLogout(sessionID quickfixgo.SessionID) {
	fmt.Printf("OnLogout %v\n", sessionID.String())
}
func (e EchoApplication) ToAdmin(msgBuilder message.Builder, sessionID quickfixgo.SessionID) {}

func (e EchoApplication) ToApp(msgBuilder message.Builder, sessionID quickfixgo.SessionID) (err error) {
	return
}

func (e EchoApplication) FromAdmin(msg message.Message, sessionID quickfixgo.SessionID) (reject reject.MessageReject) {
	return
}

func (e *EchoApplication) FromApp(msg message.Message, sessionID quickfixgo.SessionID) (reject reject.MessageReject) {
	fmt.Println("Got Message ", msg)
	return cracker.Crack(msg, sessionID, e)
}

func (e *EchoApplication) processMsg(msg message.Message, sessionID quickfixgo.SessionID) (reject reject.MessageReject) {
	OrderId, ok := msg.Body.StringValue(tag.ClOrdID)
	if !ok {
		return
	}

	reply := message.NewMessage()
	sessionOrderId := sessionID.String() + OrderId
	if PossResend, err := msg.Header.BooleanValue(tag.PossResend); err == nil && PossResend {
		if e.OrderIds[sessionOrderId] {
			return
		}

		reply.Header.SetField(message.NewBooleanField(tag.PossResend, PossResend))
	}

	e.OrderIds[sessionOrderId] = true

	msgType, _ := msg.Header.Field(tag.MsgType)
	reply.Header.SetField(msgType)

	for _, tag := range msg.Body.Tags() {
		if field, ok := msg.Body.Field(tag); ok {
			reply.Body.SetField(field)
		}
	}

	quickfixgo.SendToTarget(reply, sessionID)

	return
}

func (e EchoApplication) OnFIX40NewOrderSingle(msg fix40.NewOrderSingle, sessionID quickfixgo.SessionID) reject.MessageReject {
	return e.processMsg(msg.Message, sessionID)
}

func (e EchoApplication) OnFIX41NewOrderSingle(msg fix41.NewOrderSingle, sessionID quickfixgo.SessionID) reject.MessageReject {
	return e.processMsg(msg.Message, sessionID)
}

func (e EchoApplication) OnFIX42NewOrderSingle(msg fix42.NewOrderSingle, sessionID quickfixgo.SessionID) reject.MessageReject {
	return e.processMsg(msg.Message, sessionID)
}

func (e EchoApplication) OnFIX43NewOrderSingle(msg fix43.NewOrderSingle, sessionID quickfixgo.SessionID) reject.MessageReject {
	return e.processMsg(msg.Message, sessionID)
}

func (e EchoApplication) OnFIX44NewOrderSingle(msg fix44.NewOrderSingle, sessionID quickfixgo.SessionID) reject.MessageReject {
	return e.processMsg(msg.Message, sessionID)
}

func (e EchoApplication) OnFIX50NewOrderSingle(msg fix50.NewOrderSingle, sessionID quickfixgo.SessionID) reject.MessageReject {
	return e.processMsg(msg.Message, sessionID)
}

func (e EchoApplication) OnFIX50SP1NewOrderSingle(msg fix50sp1.NewOrderSingle, sessionID quickfixgo.SessionID) reject.MessageReject {
	return e.processMsg(msg.Message, sessionID)
}

func (e EchoApplication) OnFIX50SP2NewOrderSingle(msg fix50sp2.NewOrderSingle, sessionID quickfixgo.SessionID) reject.MessageReject {
	return e.processMsg(msg.Message, sessionID)
}

func main() {
	app := new(EchoApplication)

	globalSettings := settings.NewDictionary()
	globalSettings.SetInt(settings.SocketAcceptPort, 5001)
	globalSettings.SetString(settings.SenderCompID, "ISLD")
	globalSettings.SetString(settings.TargetCompID, "TW")

	appSettings := settings.NewApplicationSettings(globalSettings)
	appSettings.AddSession("FIX40", settings.NewDictionary().SetString(settings.BeginString, fix.BeginString_FIX40))
	appSettings.AddSession("FIX41", settings.NewDictionary().SetString(settings.BeginString, fix.BeginString_FIX41))
	appSettings.AddSession("FIX42", settings.NewDictionary().SetString(settings.BeginString, fix.BeginString_FIX42))
	appSettings.AddSession("FIX43", settings.NewDictionary().SetString(settings.BeginString, fix.BeginString_FIX43))
	appSettings.AddSession("FIX44", settings.NewDictionary().SetString(settings.BeginString, fix.BeginString_FIX44))
	appSettings.AddSession("FIX50", settings.NewDictionary().SetString(settings.BeginString, fix.BeginString_FIXT11).
		SetString(settings.DefaultApplVerID, fix.ApplVerID_FIX50))
	appSettings.AddSession("FIX50SP1", settings.NewDictionary().SetString(settings.BeginString, fix.BeginString_FIXT11).
		SetString(settings.DefaultApplVerID, fix.ApplVerID_FIX50SP1))
	appSettings.AddSession("FIX50SP2", settings.NewDictionary().SetString(settings.BeginString, fix.BeginString_FIXT11).
		SetString(settings.DefaultApplVerID, fix.ApplVerID_FIX50SP2))

	acceptor, err := quickfixgo.NewAcceptor(app, appSettings, log.ScreenLogFactory{})
	if err != nil {
		fmt.Printf("Unable to create Acceptor: ", err)
		return
	}

	if err = acceptor.Start(); err != nil {
		fmt.Printf("Unable to start Acceptor: ", err)
		return
	}

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt)
	<-interrupt

	acceptor.Stop()
}
