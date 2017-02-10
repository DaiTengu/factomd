package dbInfo

import (
	"encoding/gob"

	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
	"time"
)

type DirBlockInfo struct {
	// Serial hash for the directory block
	DBHash    interfaces.IHash
	DBHeight  uint32 //directory block height
	Timestamp int64  // time of this dir block info being created
	// BTCTxHash is the Tx hash returned from rpcclient.SendRawTransaction
	BTCTxHash interfaces.IHash // use string or *btcwire.ShaHash ???
	// BTCTxOffset is the index of the TX in this BTC block
	BTCTxOffset int32
	// BTCBlockHeight is the height of the block where this TX is stored in BTC
	BTCBlockHeight int32
	//BTCBlockHash is the hash of the block where this TX is stored in BTC
	BTCBlockHash interfaces.IHash // use string or *btcwire.ShaHash ???
	// DBMerkleRoot is the merkle root of the Directory Block
	// and is written into BTC as OP_RETURN data
	DBMerkleRoot interfaces.IHash
	// A flag to to show BTC anchor confirmation
	BTCConfirmed bool
}

var _ interfaces.Printable = (*DirBlockInfo)(nil)
var _ interfaces.BinaryMarshallableAndCopyable = (*DirBlockInfo)(nil)
var _ interfaces.DatabaseBatchable = (*DirBlockInfo)(nil)
var _ interfaces.IDirBlockInfo = (*DirBlockInfo)(nil)

func (e *DirBlockInfo) Init() {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddbInfoDirBlockInfoInit.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	if e.DBHash == nil {
		e.DBHash = primitives.NewZeroHash()
	}
	if e.BTCTxHash == nil {
		e.BTCTxHash = primitives.NewZeroHash()
	}
	if e.BTCBlockHash == nil {
		e.BTCBlockHash = primitives.NewZeroHash()
	}
	if e.DBMerkleRoot == nil {
		e.DBMerkleRoot = primitives.NewZeroHash()
	}
}

func NewDirBlockInfo() *DirBlockInfo {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddbInfoNewDirBlockInfo.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	dbi := new(DirBlockInfo)
	dbi.DBHash = primitives.NewZeroHash()
	dbi.BTCTxHash = primitives.NewZeroHash()
	dbi.BTCBlockHash = primitives.NewZeroHash()
	dbi.DBMerkleRoot = primitives.NewZeroHash()
	return dbi
}

func (e *DirBlockInfo) String() string {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddbInfoDirBlockInfoString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	str, _ := e.JSONString()
	return str
}

func (e *DirBlockInfo) JSONByte() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddbInfoDirBlockInfoJSONByte.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSON(e)
}

func (e *DirBlockInfo) JSONString() (string, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddbInfoDirBlockInfoJSONString.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.EncodeJSONString(e)
}

func (c *DirBlockInfo) New() interfaces.BinaryMarshallableAndCopyable {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddbInfoDirBlockInfoNew.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return NewDirBlockInfo()
}

func (c *DirBlockInfo) GetDatabaseHeight() uint32 {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddbInfoDirBlockInfoGetDatabaseHeight.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return c.DBHeight
}

func (c *DirBlockInfo) GetDBHeight() uint32 {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddbInfoDirBlockInfoGetDBHeight.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return c.DBHeight
}

func (c *DirBlockInfo) GetBTCConfirmed() bool {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddbInfoDirBlockInfoGetBTCConfirmed.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return c.BTCConfirmed
}

func (c *DirBlockInfo) GetChainID() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddbInfoDirBlockInfoGetChainID.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	id := make([]byte, 32)
	copy(id, []byte("DirBlockInfo"))
	return primitives.NewHash(id)
}

func (c *DirBlockInfo) DatabasePrimaryIndex() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddbInfoDirBlockInfoDatabasePrimaryIndex.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	c.Init()
	return c.DBMerkleRoot
}

func (c *DirBlockInfo) DatabaseSecondaryIndex() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddbInfoDirBlockInfoDatabaseSecondaryIndex.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	c.Init()
	return c.DBHash
}

func (e *DirBlockInfo) GetDBMerkleRoot() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddbInfoDirBlockInfoGetDBMerkleRoot.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	e.Init()
	return e.DBMerkleRoot
}

func (e *DirBlockInfo) GetBTCTxHash() interfaces.IHash {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddbInfoDirBlockInfoGetBTCTxHash.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	e.Init()
	return e.BTCTxHash
}

func (e *DirBlockInfo) GetTimestamp() interfaces.Timestamp {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddbInfoDirBlockInfoGetTimestamp.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return primitives.NewTimestampFromMilliseconds(uint64(e.Timestamp))
}

func (e *DirBlockInfo) GetBTCBlockHeight() int32 {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddbInfoDirBlockInfoGetBTCBlockHeight.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	return e.BTCBlockHeight
}

func (e *DirBlockInfo) MarshalBinary() ([]byte, error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddbInfoDirBlockInfoMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	e.Init()
	var data primitives.Buffer

	enc := gob.NewEncoder(&data)

	err := enc.Encode(newDirBlockInfoCopyFromDBI(e))
	if err != nil {
		return nil, err
	}
	return data.DeepCopyBytes(), nil
}

func (e *DirBlockInfo) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddbInfoDirBlockInfoUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	dec := gob.NewDecoder(primitives.NewBuffer(data))
	dbic := newDirBlockInfoCopy()
	err = dec.Decode(dbic)
	if err != nil {
		return nil, err
	}
	e.parseDirBlockInfoCopy(dbic)
	return nil, nil
}

func (e *DirBlockInfo) UnmarshalBinary(data []byte) (err error) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddbInfoDirBlockInfoUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	_, err = e.UnmarshalBinaryData(data)
	return
}

func (e *DirBlockInfo) SetTimestamp(timestamp interfaces.Timestamp) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddbInfoDirBlockInfoSetTimestamp.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	e.Timestamp = timestamp.GetTimeMilli()
}

type dirBlockInfoCopy struct {
	// Serial hash for the directory block
	DBHash    interfaces.IHash
	DBHeight  uint32 //directory block height
	Timestamp int64  // time of this dir block info being created
	// BTCTxHash is the Tx hash returned from rpcclient.SendRawTransaction
	BTCTxHash interfaces.IHash // use string or *btcwire.ShaHash ???
	// BTCTxOffset is the index of the TX in this BTC block
	BTCTxOffset int32
	// BTCBlockHeight is the height of the block where this TX is stored in BTC
	BTCBlockHeight int32
	//BTCBlockHash is the hash of the block where this TX is stored in BTC
	BTCBlockHash interfaces.IHash // use string or *btcwire.ShaHash ???
	// DBMerkleRoot is the merkle root of the Directory Block
	// and is written into BTC as OP_RETURN data
	DBMerkleRoot interfaces.IHash
	// A flag to to show BTC anchor confirmation
	BTCConfirmed bool
}

func newDirBlockInfoCopyFromDBI(dbi *DirBlockInfo) *dirBlockInfoCopy {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddbInfonewDirBlockInfoCopyFromDBI.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	dbic := new(dirBlockInfoCopy)
	dbic.DBHash = dbi.DBHash
	dbic.DBHeight = dbi.DBHeight
	dbic.Timestamp = dbi.Timestamp
	dbic.BTCTxHash = dbi.BTCTxHash
	dbic.BTCTxOffset = dbi.BTCTxOffset
	dbic.BTCBlockHeight = dbi.BTCBlockHeight
	dbic.BTCBlockHash = dbi.BTCBlockHash
	dbic.DBMerkleRoot = dbi.DBMerkleRoot
	dbic.BTCConfirmed = dbi.BTCConfirmed
	return dbic
}

func newDirBlockInfoCopy() *dirBlockInfoCopy {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddbInfonewDirBlockInfoCopy.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	dbi := new(dirBlockInfoCopy)
	dbi.DBHash = primitives.NewZeroHash()
	dbi.BTCTxHash = primitives.NewZeroHash()
	dbi.BTCBlockHash = primitives.NewZeroHash()
	dbi.DBMerkleRoot = primitives.NewZeroHash()
	return dbi
}

func (dbic *DirBlockInfo) parseDirBlockInfoCopy(dbi *dirBlockInfoCopy) {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddbInfoDirBlockInfoparseDirBlockInfoCopy.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	dbic.DBHash = dbi.DBHash
	dbic.DBHeight = dbi.DBHeight
	dbic.Timestamp = dbi.Timestamp
	dbic.BTCTxHash = dbi.BTCTxHash
	dbic.BTCTxOffset = dbi.BTCTxOffset
	dbic.BTCBlockHeight = dbi.BTCBlockHeight
	dbic.BTCBlockHash = dbi.BTCBlockHash
	dbic.DBMerkleRoot = dbi.DBMerkleRoot
	dbic.BTCConfirmed = dbi.BTCConfirmed
}

// NewDirBlockInfoFromDirBlock creates a DirDirBlockInfo from DirectoryBlock
func NewDirBlockInfoFromDirBlock(dirBlock interfaces.IDirectoryBlock) *DirBlockInfo {
	/////START PROMETHEUS/////
	callTime := time.Now().UnixNano()
	defer factomddbInfoNewDirBlockInfoFromDirBlock.Observe(float64(time.Now().UnixNano() - callTime))
	/////STOP PROMETHEUS/////

	dbi := new(DirBlockInfo)
	dbi.DBHash = dirBlock.GetHash()
	dbi.DBHeight = dirBlock.GetDatabaseHeight()
	dbi.DBMerkleRoot = dirBlock.GetKeyMR()
	dbi.SetTimestamp(dirBlock.GetHeader().GetTimestamp())
	dbi.BTCTxHash = primitives.NewZeroHash()
	dbi.BTCBlockHash = primitives.NewZeroHash()
	dbi.BTCConfirmed = false
	return dbi
}
