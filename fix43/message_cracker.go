package fix43

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/message"
	"github.com/cbusbey/quickfixgo/reject"
	"github.com/cbusbey/quickfixgo/tag"
)

func Crack(msg message.Message, sessionID quickfixgo.SessionID, router MessageRouter) reject.MessageReject {
	switch msgType, _ := msg.Header.StringValue(tag.MsgType); msgType {
	case "0":
		return router.OnFIX43Heartbeat(Heartbeat{msg}, sessionID)
	case "1":
		return router.OnFIX43TestRequest(TestRequest{msg}, sessionID)
	case "2":
		return router.OnFIX43ResendRequest(ResendRequest{msg}, sessionID)
	case "3":
		return router.OnFIX43Reject(Reject{msg}, sessionID)
	case "4":
		return router.OnFIX43SequenceReset(SequenceReset{msg}, sessionID)
	case "5":
		return router.OnFIX43Logout(Logout{msg}, sessionID)
	case "6":
		return router.OnFIX43IOI(IOI{msg}, sessionID)
	case "7":
		return router.OnFIX43Advertisement(Advertisement{msg}, sessionID)
	case "8":
		return router.OnFIX43ExecutionReport(ExecutionReport{msg}, sessionID)
	case "9":
		return router.OnFIX43OrderCancelReject(OrderCancelReject{msg}, sessionID)
	case "A":
		return router.OnFIX43Logon(Logon{msg}, sessionID)
	case "B":
		return router.OnFIX43News(News{msg}, sessionID)
	case "C":
		return router.OnFIX43Email(Email{msg}, sessionID)
	case "D":
		return router.OnFIX43NewOrderSingle(NewOrderSingle{msg}, sessionID)
	case "E":
		return router.OnFIX43NewOrderList(NewOrderList{msg}, sessionID)
	case "F":
		return router.OnFIX43OrderCancelRequest(OrderCancelRequest{msg}, sessionID)
	case "G":
		return router.OnFIX43OrderCancelReplaceRequest(OrderCancelReplaceRequest{msg}, sessionID)
	case "H":
		return router.OnFIX43OrderStatusRequest(OrderStatusRequest{msg}, sessionID)
	case "J":
		return router.OnFIX43Allocation(Allocation{msg}, sessionID)
	case "K":
		return router.OnFIX43ListCancelRequest(ListCancelRequest{msg}, sessionID)
	case "L":
		return router.OnFIX43ListExecute(ListExecute{msg}, sessionID)
	case "M":
		return router.OnFIX43ListStatusRequest(ListStatusRequest{msg}, sessionID)
	case "N":
		return router.OnFIX43ListStatus(ListStatus{msg}, sessionID)
	case "P":
		return router.OnFIX43AllocationAck(AllocationAck{msg}, sessionID)
	case "Q":
		return router.OnFIX43DontKnowTrade(DontKnowTrade{msg}, sessionID)
	case "R":
		return router.OnFIX43QuoteRequest(QuoteRequest{msg}, sessionID)
	case "S":
		return router.OnFIX43Quote(Quote{msg}, sessionID)
	case "T":
		return router.OnFIX43SettlementInstructions(SettlementInstructions{msg}, sessionID)
	case "V":
		return router.OnFIX43MarketDataRequest(MarketDataRequest{msg}, sessionID)
	case "W":
		return router.OnFIX43MarketDataSnapshotFullRefresh(MarketDataSnapshotFullRefresh{msg}, sessionID)
	case "X":
		return router.OnFIX43MarketDataIncrementalRefresh(MarketDataIncrementalRefresh{msg}, sessionID)
	case "Y":
		return router.OnFIX43MarketDataRequestReject(MarketDataRequestReject{msg}, sessionID)
	case "Z":
		return router.OnFIX43QuoteCancel(QuoteCancel{msg}, sessionID)
	case "a":
		return router.OnFIX43QuoteStatusRequest(QuoteStatusRequest{msg}, sessionID)
	case "b":
		return router.OnFIX43MassQuoteAcknowledgement(MassQuoteAcknowledgement{msg}, sessionID)
	case "c":
		return router.OnFIX43SecurityDefinitionRequest(SecurityDefinitionRequest{msg}, sessionID)
	case "d":
		return router.OnFIX43SecurityDefinition(SecurityDefinition{msg}, sessionID)
	case "e":
		return router.OnFIX43SecurityStatusRequest(SecurityStatusRequest{msg}, sessionID)
	case "f":
		return router.OnFIX43SecurityStatus(SecurityStatus{msg}, sessionID)
	case "g":
		return router.OnFIX43TradingSessionStatusRequest(TradingSessionStatusRequest{msg}, sessionID)
	case "h":
		return router.OnFIX43TradingSessionStatus(TradingSessionStatus{msg}, sessionID)
	case "i":
		return router.OnFIX43MassQuote(MassQuote{msg}, sessionID)
	case "j":
		return router.OnFIX43BusinessMessageReject(BusinessMessageReject{msg}, sessionID)
	case "k":
		return router.OnFIX43BidRequest(BidRequest{msg}, sessionID)
	case "l":
		return router.OnFIX43BidResponse(BidResponse{msg}, sessionID)
	case "m":
		return router.OnFIX43ListStrikePrice(ListStrikePrice{msg}, sessionID)
	case "o":
		return router.OnFIX43RegistrationInstructions(RegistrationInstructions{msg}, sessionID)
	case "p":
		return router.OnFIX43RegistrationInstructionsResponse(RegistrationInstructionsResponse{msg}, sessionID)
	case "q":
		return router.OnFIX43OrderMassCancelRequest(OrderMassCancelRequest{msg}, sessionID)
	case "r":
		return router.OnFIX43OrderMassCancelReport(OrderMassCancelReport{msg}, sessionID)
	case "s":
		return router.OnFIX43NewOrderCross(NewOrderCross{msg}, sessionID)
	case "u":
		return router.OnFIX43CrossOrderCancelRequest(CrossOrderCancelRequest{msg}, sessionID)
	case "t":
		return router.OnFIX43CrossOrderCancelReplaceRequest(CrossOrderCancelReplaceRequest{msg}, sessionID)
	case "v":
		return router.OnFIX43SecurityTypeRequest(SecurityTypeRequest{msg}, sessionID)
	case "w":
		return router.OnFIX43SecurityTypes(SecurityTypes{msg}, sessionID)
	case "x":
		return router.OnFIX43SecurityListRequest(SecurityListRequest{msg}, sessionID)
	case "y":
		return router.OnFIX43SecurityList(SecurityList{msg}, sessionID)
	case "z":
		return router.OnFIX43DerivativeSecurityListRequest(DerivativeSecurityListRequest{msg}, sessionID)
	case "AA":
		return router.OnFIX43DerivativeSecurityList(DerivativeSecurityList{msg}, sessionID)
	case "AB":
		return router.OnFIX43NewOrderMultileg(NewOrderMultileg{msg}, sessionID)
	case "AC":
		return router.OnFIX43MultilegOrderCancelReplaceRequest(MultilegOrderCancelReplaceRequest{msg}, sessionID)
	case "AD":
		return router.OnFIX43TradeCaptureReportRequest(TradeCaptureReportRequest{msg}, sessionID)
	case "AE":
		return router.OnFIX43TradeCaptureReport(TradeCaptureReport{msg}, sessionID)
	case "AF":
		return router.OnFIX43OrderMassStatusRequest(OrderMassStatusRequest{msg}, sessionID)
	case "AG":
		return router.OnFIX43QuoteRequestReject(QuoteRequestReject{msg}, sessionID)
	case "AH":
		return router.OnFIX43RFQRequest(RFQRequest{msg}, sessionID)
	case "AI":
		return router.OnFIX43QuoteStatusReport(QuoteStatusReport{msg}, sessionID)
	}
	return reject.NewInvalidMessageType(msg)
}

type MessageRouter interface {
	OnFIX43Heartbeat(msg Heartbeat, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43TestRequest(msg TestRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43ResendRequest(msg ResendRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43Reject(msg Reject, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43SequenceReset(msg SequenceReset, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43Logout(msg Logout, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43IOI(msg IOI, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43Advertisement(msg Advertisement, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43ExecutionReport(msg ExecutionReport, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43OrderCancelReject(msg OrderCancelReject, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43Logon(msg Logon, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43News(msg News, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43Email(msg Email, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43NewOrderSingle(msg NewOrderSingle, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43NewOrderList(msg NewOrderList, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43OrderCancelRequest(msg OrderCancelRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43OrderStatusRequest(msg OrderStatusRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43Allocation(msg Allocation, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43ListCancelRequest(msg ListCancelRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43ListExecute(msg ListExecute, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43ListStatusRequest(msg ListStatusRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43ListStatus(msg ListStatus, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43AllocationAck(msg AllocationAck, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43DontKnowTrade(msg DontKnowTrade, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43QuoteRequest(msg QuoteRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43Quote(msg Quote, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43SettlementInstructions(msg SettlementInstructions, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43MarketDataRequest(msg MarketDataRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43MarketDataRequestReject(msg MarketDataRequestReject, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43QuoteCancel(msg QuoteCancel, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43QuoteStatusRequest(msg QuoteStatusRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43MassQuoteAcknowledgement(msg MassQuoteAcknowledgement, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43SecurityDefinition(msg SecurityDefinition, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43SecurityStatusRequest(msg SecurityStatusRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43SecurityStatus(msg SecurityStatus, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43TradingSessionStatus(msg TradingSessionStatus, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43MassQuote(msg MassQuote, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43BusinessMessageReject(msg BusinessMessageReject, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43BidRequest(msg BidRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43BidResponse(msg BidResponse, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43ListStrikePrice(msg ListStrikePrice, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43RegistrationInstructions(msg RegistrationInstructions, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43RegistrationInstructionsResponse(msg RegistrationInstructionsResponse, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43OrderMassCancelRequest(msg OrderMassCancelRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43OrderMassCancelReport(msg OrderMassCancelReport, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43NewOrderCross(msg NewOrderCross, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43CrossOrderCancelRequest(msg CrossOrderCancelRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43CrossOrderCancelReplaceRequest(msg CrossOrderCancelReplaceRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43SecurityTypeRequest(msg SecurityTypeRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43SecurityTypes(msg SecurityTypes, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43SecurityListRequest(msg SecurityListRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43SecurityList(msg SecurityList, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43DerivativeSecurityListRequest(msg DerivativeSecurityListRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43DerivativeSecurityList(msg DerivativeSecurityList, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43NewOrderMultileg(msg NewOrderMultileg, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43MultilegOrderCancelReplaceRequest(msg MultilegOrderCancelReplaceRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43TradeCaptureReportRequest(msg TradeCaptureReportRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43TradeCaptureReport(msg TradeCaptureReport, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43OrderMassStatusRequest(msg OrderMassStatusRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43QuoteRequestReject(msg QuoteRequestReject, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43RFQRequest(msg RFQRequest, sessionID quickfixgo.SessionID) reject.MessageReject
	OnFIX43QuoteStatusReport(msg QuoteStatusReport, sessionID quickfixgo.SessionID) reject.MessageReject
}
type FIX43MessageCracker struct{}

func (c *FIX43MessageCracker) OnFIX43Heartbeat(msg Heartbeat, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43TestRequest(msg TestRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43ResendRequest(msg ResendRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43Reject(msg Reject, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43SequenceReset(msg SequenceReset, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43Logout(msg Logout, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43IOI(msg IOI, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43Advertisement(msg Advertisement, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43ExecutionReport(msg ExecutionReport, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43OrderCancelReject(msg OrderCancelReject, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43Logon(msg Logon, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43News(msg News, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43Email(msg Email, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43NewOrderSingle(msg NewOrderSingle, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43NewOrderList(msg NewOrderList, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43OrderCancelRequest(msg OrderCancelRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43OrderCancelReplaceRequest(msg OrderCancelReplaceRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43OrderStatusRequest(msg OrderStatusRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43Allocation(msg Allocation, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43ListCancelRequest(msg ListCancelRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43ListExecute(msg ListExecute, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43ListStatusRequest(msg ListStatusRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43ListStatus(msg ListStatus, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43AllocationAck(msg AllocationAck, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43DontKnowTrade(msg DontKnowTrade, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43QuoteRequest(msg QuoteRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43Quote(msg Quote, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43SettlementInstructions(msg SettlementInstructions, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43MarketDataRequest(msg MarketDataRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43MarketDataSnapshotFullRefresh(msg MarketDataSnapshotFullRefresh, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43MarketDataIncrementalRefresh(msg MarketDataIncrementalRefresh, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43MarketDataRequestReject(msg MarketDataRequestReject, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43QuoteCancel(msg QuoteCancel, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43QuoteStatusRequest(msg QuoteStatusRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43MassQuoteAcknowledgement(msg MassQuoteAcknowledgement, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43SecurityDefinitionRequest(msg SecurityDefinitionRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43SecurityDefinition(msg SecurityDefinition, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43SecurityStatusRequest(msg SecurityStatusRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43SecurityStatus(msg SecurityStatus, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43TradingSessionStatusRequest(msg TradingSessionStatusRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43TradingSessionStatus(msg TradingSessionStatus, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43MassQuote(msg MassQuote, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43BusinessMessageReject(msg BusinessMessageReject, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43BidRequest(msg BidRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43BidResponse(msg BidResponse, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43ListStrikePrice(msg ListStrikePrice, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43RegistrationInstructions(msg RegistrationInstructions, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43RegistrationInstructionsResponse(msg RegistrationInstructionsResponse, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43OrderMassCancelRequest(msg OrderMassCancelRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43OrderMassCancelReport(msg OrderMassCancelReport, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43NewOrderCross(msg NewOrderCross, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43CrossOrderCancelRequest(msg CrossOrderCancelRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43CrossOrderCancelReplaceRequest(msg CrossOrderCancelReplaceRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43SecurityTypeRequest(msg SecurityTypeRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43SecurityTypes(msg SecurityTypes, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43SecurityListRequest(msg SecurityListRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43SecurityList(msg SecurityList, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43DerivativeSecurityListRequest(msg DerivativeSecurityListRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43DerivativeSecurityList(msg DerivativeSecurityList, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43NewOrderMultileg(msg NewOrderMultileg, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43MultilegOrderCancelReplaceRequest(msg MultilegOrderCancelReplaceRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43TradeCaptureReportRequest(msg TradeCaptureReportRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43TradeCaptureReport(msg TradeCaptureReport, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43OrderMassStatusRequest(msg OrderMassStatusRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43QuoteRequestReject(msg QuoteRequestReject, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43RFQRequest(msg RFQRequest, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
func (c *FIX43MessageCracker) OnFIX43QuoteStatusReport(msg QuoteStatusReport, sessionId quickfixgo.SessionID) reject.MessageReject {
	return reject.NewUnsupportedMessageType(msg.Message)
}
