package mdgwmsg


const LoginMsgType = "S001"
const LogoutMsgType = "S002"
const HBMsgType = "S003"

type MsgHeader struct {
    msgType [4]byte
    SendingTtime uint64
    MsgSeq uint64
    BodyLength uint32
}

type MsgTail struct {
    CheckSum uint32
}

//登录消息
type LoginMsg struct {
    Header MsgHeader
    SenderCompID    [32]byte
    TargetCompID    [32]byte
    HeartBtInt      uint16
    AppVerID        [8]byte
    Tail MsgTail
}

//注销消息
type LogoutMsg struct {
    Header MsgHeader
    SessionStatus uint32
    Text    [256]byte
    Tail MsgTail
}

//心跳消息
type HeartBtMsg struct {
    Header MsgHeader
    Tail MsgTail
}

//市场状态消息
type MktStatusMsg struct {
    Header MsgHeader
    SecurityType uint8
    TradSesMode uint8
    TradingSessionID [8]byte
    TotNoRelatedSym uint32
    Tail MsgTail
}

//行情快照
type SnapMsg struct {
    Header MsgHeader
    SecurityType uint8
    TradSesMode uint8
    TradeDate uint32
    LastUpdateTime  uint32
    MDStreamID [5]byte
    SecurityID [8]byte
    Symbol [8]byte
    PreClosePx uint64
    TotalVolumeTraded uint64
    NumTrades uint64
    TotalValueTraded uint64
    TradingPhaseCode [8]byte
}

//指数行情快照
//根据条目个数需要进行扩展
type IndexSnap struct {
    SnapData SnapMsg
    NoMDEntries uint16
    MDEntryType [2]byte
    MDEntryPx uint64
}

//竞价行情快照
type BidSnap struct {
    SnapData SnapMsg
    NoMDEntries uint16
    MDEntryType [2]byte
    MDEntryPx uint64
    MDEntrySize uint64
    MDEntryPositionNo uint8
}


