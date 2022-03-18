package statedb

import (
	"encoding/binary"
	"errors"

	"github.com/dfuse-io/fluxdb"
	zsw "github.com/zhongshuwen/zswchain-go"
)

var bigEndian = binary.BigEndian

var baseEntry = fluxdb.NewBaseSingletEntry
var baseRow = fluxdb.NewBaseTabletRow

var errABIUnmarshal = errors.New("unmarshal abi")

func bytesToName(bytes []byte) string {
	return zsw.NameToString(bigEndian.Uint64(bytes))
}

func bytesToName2(bytes []byte) (string, string) {
	return zsw.NameToString(bigEndian.Uint64(bytes)), zsw.NameToString(bigEndian.Uint64(bytes[8:]))
}

func bytesToName3(bytes []byte) (string, string, string) {
	return zsw.NameToString(bigEndian.Uint64(bytes)),
		zsw.NameToString(bigEndian.Uint64(bytes[8:])),
		zsw.NameToString(bigEndian.Uint64(bytes[16:]))
}

func bytesToJoinedName2(bytes []byte) string {
	return zsw.NameToString(bigEndian.Uint64(bytes)) + ":" + zsw.NameToString(bigEndian.Uint64(bytes[8:]))
}

func bytesToJoinedName3(bytes []byte) string {
	return zsw.NameToString(bigEndian.Uint64(bytes)) +
		":" + zsw.NameToString(bigEndian.Uint64(bytes[8:])) +
		":" + zsw.NameToString(bigEndian.Uint64(bytes[16:]))
}

var standardNameConverter = zsw.MustStringToName
var extendedNameConverter = mustExtendedStringToName

func standardNameToBytes(names ...string) (out []byte) {
	return nameToBytes(standardNameConverter, names)
}

func extendedNameToBytes(names ...string) (out []byte) {
	return nameToBytes(extendedNameConverter, names)
}

func nameToBytes(converter func(name string) uint64, names []string) (out []byte) {
	out = make([]byte, 8*len(names))
	moving := out
	for _, name := range names {
		bigEndian.PutUint64(moving, converter(name))
		moving = moving[8:]
	}

	return
}

func nameaToBytes(name zsw.AccountName) (out []byte) {
	out = make([]byte, 8)
	bigEndian.PutUint64(out, zsw.MustStringToName(string(name)))
	return
}

func namenToBytes(name zsw.Name) (out []byte) {
	out = make([]byte, 8)
	bigEndian.PutUint64(out, zsw.MustStringToName(string(name)))
	return
}

func mustExtendedStringToName(name string) uint64 {
	val, err := zsw.ExtendedStringToName(name)
	if err != nil {
		panic(err)
	}

	return val
}
