package main

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "moc.h"
import "C"
import (
	"runtime"
	"strings"
	"unsafe"

	"github.com/therecipe/qt"
	std_core "github.com/therecipe/qt/core"
)

func cGoUnpackString(s C.struct_Moc_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_Moc_PackedString) []byte {
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return *(*[]byte)(unsafe.Pointer(&gs))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

type CtxObject_ITF interface {
	std_core.QObject_ITF
	CtxObject_PTR() *CtxObject
}

func (ptr *CtxObject) CtxObject_PTR() *CtxObject {
	return ptr
}

func (ptr *CtxObject) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *CtxObject) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromCtxObject(ptr CtxObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.CtxObject_PTR().Pointer()
	}
	return nil
}

func NewCtxObjectFromPointer(ptr unsafe.Pointer) (n *CtxObject) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(CtxObject)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *CtxObject:
			n = deduced

		case *std_core.QObject:
			n = &CtxObject{QObject: *deduced}

		default:
			n = new(CtxObject)
			n.SetPointer(ptr)
		}
	}
	return
}
func (this *CtxObject) Init() { this.init() }

//export callbackCtxObject8ed989_Constructor
func callbackCtxObject8ed989_Constructor(ptr unsafe.Pointer) {
	this := NewCtxObjectFromPointer(ptr)
	qt.Register(ptr, this)
	this.ConnectIncreased(this.increased)
	this.ConnectDecreased(this.decreased)
	this.init()
}

//export callbackCtxObject8ed989_Increased
func callbackCtxObject8ed989_Increased(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "increased"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *CtxObject) ConnectIncreased(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "increased") {
			C.CtxObject8ed989_ConnectIncreased(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "increased")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "increased"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "increased", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "increased", unsafe.Pointer(&f))
		}
	}
}

func (ptr *CtxObject) DisconnectIncreased() {
	if ptr.Pointer() != nil {
		C.CtxObject8ed989_DisconnectIncreased(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "increased")
	}
}

func (ptr *CtxObject) Increased() {
	if ptr.Pointer() != nil {
		C.CtxObject8ed989_Increased(ptr.Pointer())
	}
}

//export callbackCtxObject8ed989_Decreased
func callbackCtxObject8ed989_Decreased(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "decreased"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *CtxObject) ConnectDecreased(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "decreased") {
			C.CtxObject8ed989_ConnectDecreased(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "decreased")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "decreased"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "decreased", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "decreased", unsafe.Pointer(&f))
		}
	}
}

func (ptr *CtxObject) DisconnectDecreased() {
	if ptr.Pointer() != nil {
		C.CtxObject8ed989_DisconnectDecreased(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "decreased")
	}
}

func (ptr *CtxObject) Decreased() {
	if ptr.Pointer() != nil {
		C.CtxObject8ed989_Decreased(ptr.Pointer())
	}
}

//export callbackCtxObject8ed989_NodeAmount
func callbackCtxObject8ed989_NodeAmount(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "nodeAmount"); signal != nil {
		return C.longlong((*(*func() int64)(signal))())
	}

	return C.longlong(NewCtxObjectFromPointer(ptr).NodeAmountDefault())
}

func (ptr *CtxObject) ConnectNodeAmount(f func() int64) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "nodeAmount"); signal != nil {
			f := func() int64 {
				(*(*func() int64)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "nodeAmount", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "nodeAmount", unsafe.Pointer(&f))
		}
	}
}

func (ptr *CtxObject) DisconnectNodeAmount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "nodeAmount")
	}
}

func (ptr *CtxObject) NodeAmount() int64 {
	if ptr.Pointer() != nil {
		return int64(C.CtxObject8ed989_NodeAmount(ptr.Pointer()))
	}
	return 0
}

func (ptr *CtxObject) NodeAmountDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.CtxObject8ed989_NodeAmountDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackCtxObject8ed989_SetNodeAmount
func callbackCtxObject8ed989_SetNodeAmount(ptr unsafe.Pointer, nodeAmount C.longlong) {
	if signal := qt.GetSignal(ptr, "setNodeAmount"); signal != nil {
		(*(*func(int64))(signal))(int64(nodeAmount))
	} else {
		NewCtxObjectFromPointer(ptr).SetNodeAmountDefault(int64(nodeAmount))
	}
}

func (ptr *CtxObject) ConnectSetNodeAmount(f func(nodeAmount int64)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setNodeAmount"); signal != nil {
			f := func(nodeAmount int64) {
				(*(*func(int64))(signal))(nodeAmount)
				f(nodeAmount)
			}
			qt.ConnectSignal(ptr.Pointer(), "setNodeAmount", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setNodeAmount", unsafe.Pointer(&f))
		}
	}
}

func (ptr *CtxObject) DisconnectSetNodeAmount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setNodeAmount")
	}
}

func (ptr *CtxObject) SetNodeAmount(nodeAmount int64) {
	if ptr.Pointer() != nil {
		C.CtxObject8ed989_SetNodeAmount(ptr.Pointer(), C.longlong(nodeAmount))
	}
}

func (ptr *CtxObject) SetNodeAmountDefault(nodeAmount int64) {
	if ptr.Pointer() != nil {
		C.CtxObject8ed989_SetNodeAmountDefault(ptr.Pointer(), C.longlong(nodeAmount))
	}
}

//export callbackCtxObject8ed989_NodeAmountChanged
func callbackCtxObject8ed989_NodeAmountChanged(ptr unsafe.Pointer, nodeAmount C.longlong) {
	if signal := qt.GetSignal(ptr, "nodeAmountChanged"); signal != nil {
		(*(*func(int64))(signal))(int64(nodeAmount))
	}

}

func (ptr *CtxObject) ConnectNodeAmountChanged(f func(nodeAmount int64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "nodeAmountChanged") {
			C.CtxObject8ed989_ConnectNodeAmountChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "nodeAmountChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "nodeAmountChanged"); signal != nil {
			f := func(nodeAmount int64) {
				(*(*func(int64))(signal))(nodeAmount)
				f(nodeAmount)
			}
			qt.ConnectSignal(ptr.Pointer(), "nodeAmountChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "nodeAmountChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *CtxObject) DisconnectNodeAmountChanged() {
	if ptr.Pointer() != nil {
		C.CtxObject8ed989_DisconnectNodeAmountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "nodeAmountChanged")
	}
}

func (ptr *CtxObject) NodeAmountChanged(nodeAmount int64) {
	if ptr.Pointer() != nil {
		C.CtxObject8ed989_NodeAmountChanged(ptr.Pointer(), C.longlong(nodeAmount))
	}
}

func CtxObject_QRegisterMetaType() int {
	return int(int32(C.CtxObject8ed989_CtxObject8ed989_QRegisterMetaType()))
}

func (ptr *CtxObject) QRegisterMetaType() int {
	return int(int32(C.CtxObject8ed989_CtxObject8ed989_QRegisterMetaType()))
}

func CtxObject_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.CtxObject8ed989_CtxObject8ed989_QRegisterMetaType2(typeNameC)))
}

func (ptr *CtxObject) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.CtxObject8ed989_CtxObject8ed989_QRegisterMetaType2(typeNameC)))
}

func CtxObject_QmlRegisterType() int {
	return int(int32(C.CtxObject8ed989_CtxObject8ed989_QmlRegisterType()))
}

func (ptr *CtxObject) QmlRegisterType() int {
	return int(int32(C.CtxObject8ed989_CtxObject8ed989_QmlRegisterType()))
}

func CtxObject_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.CtxObject8ed989_CtxObject8ed989_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *CtxObject) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.CtxObject8ed989_CtxObject8ed989_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *CtxObject) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CtxObject8ed989___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CtxObject) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CtxObject8ed989___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CtxObject) __children_newList() unsafe.Pointer {
	return C.CtxObject8ed989___children_newList(ptr.Pointer())
}

func (ptr *CtxObject) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.CtxObject8ed989___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *CtxObject) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.CtxObject8ed989___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *CtxObject) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.CtxObject8ed989___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *CtxObject) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CtxObject8ed989___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CtxObject) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CtxObject8ed989___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CtxObject) __findChildren_newList() unsafe.Pointer {
	return C.CtxObject8ed989___findChildren_newList(ptr.Pointer())
}

func (ptr *CtxObject) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CtxObject8ed989___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CtxObject) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CtxObject8ed989___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CtxObject) __findChildren_newList3() unsafe.Pointer {
	return C.CtxObject8ed989___findChildren_newList3(ptr.Pointer())
}

func (ptr *CtxObject) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.CtxObject8ed989___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *CtxObject) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.CtxObject8ed989___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *CtxObject) __qFindChildren_newList2() unsafe.Pointer {
	return C.CtxObject8ed989___qFindChildren_newList2(ptr.Pointer())
}

func NewCtxObject(parent std_core.QObject_ITF) *CtxObject {
	CtxObject_QRegisterMetaType()
	tmpValue := NewCtxObjectFromPointer(C.CtxObject8ed989_NewCtxObject(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackCtxObject8ed989_DestroyCtxObject
func callbackCtxObject8ed989_DestroyCtxObject(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~CtxObject"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewCtxObjectFromPointer(ptr).DestroyCtxObjectDefault()
	}
}

func (ptr *CtxObject) ConnectDestroyCtxObject(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~CtxObject"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~CtxObject", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~CtxObject", unsafe.Pointer(&f))
		}
	}
}

func (ptr *CtxObject) DisconnectDestroyCtxObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~CtxObject")
	}
}

func (ptr *CtxObject) DestroyCtxObject() {
	if ptr.Pointer() != nil {
		C.CtxObject8ed989_DestroyCtxObject(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *CtxObject) DestroyCtxObjectDefault() {
	if ptr.Pointer() != nil {
		C.CtxObject8ed989_DestroyCtxObjectDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackCtxObject8ed989_ChildEvent
func callbackCtxObject8ed989_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewCtxObjectFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *CtxObject) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CtxObject8ed989_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackCtxObject8ed989_ConnectNotify
func callbackCtxObject8ed989_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewCtxObjectFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *CtxObject) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.CtxObject8ed989_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackCtxObject8ed989_CustomEvent
func callbackCtxObject8ed989_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewCtxObjectFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *CtxObject) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CtxObject8ed989_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackCtxObject8ed989_DeleteLater
func callbackCtxObject8ed989_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewCtxObjectFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *CtxObject) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.CtxObject8ed989_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackCtxObject8ed989_Destroyed
func callbackCtxObject8ed989_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackCtxObject8ed989_DisconnectNotify
func callbackCtxObject8ed989_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewCtxObjectFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *CtxObject) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.CtxObject8ed989_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackCtxObject8ed989_Event
func callbackCtxObject8ed989_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCtxObjectFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *CtxObject) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CtxObject8ed989_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackCtxObject8ed989_EventFilter
func callbackCtxObject8ed989_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewCtxObjectFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *CtxObject) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.CtxObject8ed989_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackCtxObject8ed989_ObjectNameChanged
func callbackCtxObject8ed989_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackCtxObject8ed989_TimerEvent
func callbackCtxObject8ed989_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewCtxObjectFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *CtxObject) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.CtxObject8ed989_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}

type TopologyObject_ITF interface {
	std_core.QObject_ITF
	TopologyObject_PTR() *TopologyObject
}

func (ptr *TopologyObject) TopologyObject_PTR() *TopologyObject {
	return ptr
}

func (ptr *TopologyObject) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *TopologyObject) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
	}
}

func PointerFromTopologyObject(ptr TopologyObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.TopologyObject_PTR().Pointer()
	}
	return nil
}

func NewTopologyObjectFromPointer(ptr unsafe.Pointer) (n *TopologyObject) {
	if gPtr, ok := qt.Receive(ptr); !ok {
		n = new(TopologyObject)
		n.SetPointer(ptr)
	} else {
		switch deduced := gPtr.(type) {
		case *TopologyObject:
			n = deduced

		case *std_core.QObject:
			n = &TopologyObject{QObject: *deduced}

		default:
			n = new(TopologyObject)
			n.SetPointer(ptr)
		}
	}
	return
}

//export callbackTopologyObject8ed989_Constructor
func callbackTopologyObject8ed989_Constructor(ptr unsafe.Pointer) {
	this := NewTopologyObjectFromPointer(ptr)
	qt.Register(ptr, this)
}

//export callbackTopologyObject8ed989_NodeAmount
func callbackTopologyObject8ed989_NodeAmount(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "nodeAmount"); signal != nil {
		return C.longlong((*(*func() int64)(signal))())
	}

	return C.longlong(NewTopologyObjectFromPointer(ptr).NodeAmountDefault())
}

func (ptr *TopologyObject) ConnectNodeAmount(f func() int64) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "nodeAmount"); signal != nil {
			f := func() int64 {
				(*(*func() int64)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "nodeAmount", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "nodeAmount", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TopologyObject) DisconnectNodeAmount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "nodeAmount")
	}
}

func (ptr *TopologyObject) NodeAmount() int64 {
	if ptr.Pointer() != nil {
		return int64(C.TopologyObject8ed989_NodeAmount(ptr.Pointer()))
	}
	return 0
}

func (ptr *TopologyObject) NodeAmountDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.TopologyObject8ed989_NodeAmountDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackTopologyObject8ed989_SetNodeAmount
func callbackTopologyObject8ed989_SetNodeAmount(ptr unsafe.Pointer, nodeAmount C.longlong) {
	if signal := qt.GetSignal(ptr, "setNodeAmount"); signal != nil {
		(*(*func(int64))(signal))(int64(nodeAmount))
	} else {
		NewTopologyObjectFromPointer(ptr).SetNodeAmountDefault(int64(nodeAmount))
	}
}

func (ptr *TopologyObject) ConnectSetNodeAmount(f func(nodeAmount int64)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setNodeAmount"); signal != nil {
			f := func(nodeAmount int64) {
				(*(*func(int64))(signal))(nodeAmount)
				f(nodeAmount)
			}
			qt.ConnectSignal(ptr.Pointer(), "setNodeAmount", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setNodeAmount", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TopologyObject) DisconnectSetNodeAmount() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setNodeAmount")
	}
}

func (ptr *TopologyObject) SetNodeAmount(nodeAmount int64) {
	if ptr.Pointer() != nil {
		C.TopologyObject8ed989_SetNodeAmount(ptr.Pointer(), C.longlong(nodeAmount))
	}
}

func (ptr *TopologyObject) SetNodeAmountDefault(nodeAmount int64) {
	if ptr.Pointer() != nil {
		C.TopologyObject8ed989_SetNodeAmountDefault(ptr.Pointer(), C.longlong(nodeAmount))
	}
}

//export callbackTopologyObject8ed989_NodeAmountChanged
func callbackTopologyObject8ed989_NodeAmountChanged(ptr unsafe.Pointer, nodeAmount C.longlong) {
	if signal := qt.GetSignal(ptr, "nodeAmountChanged"); signal != nil {
		(*(*func(int64))(signal))(int64(nodeAmount))
	}

}

func (ptr *TopologyObject) ConnectNodeAmountChanged(f func(nodeAmount int64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "nodeAmountChanged") {
			C.TopologyObject8ed989_ConnectNodeAmountChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "nodeAmountChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "nodeAmountChanged"); signal != nil {
			f := func(nodeAmount int64) {
				(*(*func(int64))(signal))(nodeAmount)
				f(nodeAmount)
			}
			qt.ConnectSignal(ptr.Pointer(), "nodeAmountChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "nodeAmountChanged", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TopologyObject) DisconnectNodeAmountChanged() {
	if ptr.Pointer() != nil {
		C.TopologyObject8ed989_DisconnectNodeAmountChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "nodeAmountChanged")
	}
}

func (ptr *TopologyObject) NodeAmountChanged(nodeAmount int64) {
	if ptr.Pointer() != nil {
		C.TopologyObject8ed989_NodeAmountChanged(ptr.Pointer(), C.longlong(nodeAmount))
	}
}

func TopologyObject_QRegisterMetaType() int {
	return int(int32(C.TopologyObject8ed989_TopologyObject8ed989_QRegisterMetaType()))
}

func (ptr *TopologyObject) QRegisterMetaType() int {
	return int(int32(C.TopologyObject8ed989_TopologyObject8ed989_QRegisterMetaType()))
}

func TopologyObject_QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.TopologyObject8ed989_TopologyObject8ed989_QRegisterMetaType2(typeNameC)))
}

func (ptr *TopologyObject) QRegisterMetaType2(typeName string) int {
	var typeNameC *C.char
	if typeName != "" {
		typeNameC = C.CString(typeName)
		defer C.free(unsafe.Pointer(typeNameC))
	}
	return int(int32(C.TopologyObject8ed989_TopologyObject8ed989_QRegisterMetaType2(typeNameC)))
}

func TopologyObject_QmlRegisterType() int {
	return int(int32(C.TopologyObject8ed989_TopologyObject8ed989_QmlRegisterType()))
}

func (ptr *TopologyObject) QmlRegisterType() int {
	return int(int32(C.TopologyObject8ed989_TopologyObject8ed989_QmlRegisterType()))
}

func TopologyObject_QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.TopologyObject8ed989_TopologyObject8ed989_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *TopologyObject) QmlRegisterType2(uri string, versionMajor int, versionMinor int, qmlName string) int {
	var uriC *C.char
	if uri != "" {
		uriC = C.CString(uri)
		defer C.free(unsafe.Pointer(uriC))
	}
	var qmlNameC *C.char
	if qmlName != "" {
		qmlNameC = C.CString(qmlName)
		defer C.free(unsafe.Pointer(qmlNameC))
	}
	return int(int32(C.TopologyObject8ed989_TopologyObject8ed989_QmlRegisterType2(uriC, C.int(int32(versionMajor)), C.int(int32(versionMinor)), qmlNameC)))
}

func (ptr *TopologyObject) __children_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TopologyObject8ed989___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TopologyObject) __children_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TopologyObject8ed989___children_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TopologyObject) __children_newList() unsafe.Pointer {
	return C.TopologyObject8ed989___children_newList(ptr.Pointer())
}

func (ptr *TopologyObject) __dynamicPropertyNames_atList(i int) *std_core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQByteArrayFromPointer(C.TopologyObject8ed989___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*std_core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *TopologyObject) __dynamicPropertyNames_setList(i std_core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.TopologyObject8ed989___dynamicPropertyNames_setList(ptr.Pointer(), std_core.PointerFromQByteArray(i))
	}
}

func (ptr *TopologyObject) __dynamicPropertyNames_newList() unsafe.Pointer {
	return C.TopologyObject8ed989___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *TopologyObject) __findChildren_atList(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TopologyObject8ed989___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TopologyObject) __findChildren_setList(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TopologyObject8ed989___findChildren_setList(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TopologyObject) __findChildren_newList() unsafe.Pointer {
	return C.TopologyObject8ed989___findChildren_newList(ptr.Pointer())
}

func (ptr *TopologyObject) __findChildren_atList3(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TopologyObject8ed989___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TopologyObject) __findChildren_setList3(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TopologyObject8ed989___findChildren_setList3(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TopologyObject) __findChildren_newList3() unsafe.Pointer {
	return C.TopologyObject8ed989___findChildren_newList3(ptr.Pointer())
}

func (ptr *TopologyObject) __qFindChildren_atList2(i int) *std_core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := std_core.NewQObjectFromPointer(C.TopologyObject8ed989___qFindChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *TopologyObject) __qFindChildren_setList2(i std_core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.TopologyObject8ed989___qFindChildren_setList2(ptr.Pointer(), std_core.PointerFromQObject(i))
	}
}

func (ptr *TopologyObject) __qFindChildren_newList2() unsafe.Pointer {
	return C.TopologyObject8ed989___qFindChildren_newList2(ptr.Pointer())
}

func NewTopologyObject(parent std_core.QObject_ITF) *TopologyObject {
	TopologyObject_QRegisterMetaType()
	tmpValue := NewTopologyObjectFromPointer(C.TopologyObject8ed989_NewTopologyObject(std_core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*std_core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackTopologyObject8ed989_DestroyTopologyObject
func callbackTopologyObject8ed989_DestroyTopologyObject(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~TopologyObject"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewTopologyObjectFromPointer(ptr).DestroyTopologyObjectDefault()
	}
}

func (ptr *TopologyObject) ConnectDestroyTopologyObject(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~TopologyObject"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~TopologyObject", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~TopologyObject", unsafe.Pointer(&f))
		}
	}
}

func (ptr *TopologyObject) DisconnectDestroyTopologyObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~TopologyObject")
	}
}

func (ptr *TopologyObject) DestroyTopologyObject() {
	if ptr.Pointer() != nil {
		C.TopologyObject8ed989_DestroyTopologyObject(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *TopologyObject) DestroyTopologyObjectDefault() {
	if ptr.Pointer() != nil {
		C.TopologyObject8ed989_DestroyTopologyObjectDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackTopologyObject8ed989_ChildEvent
func callbackTopologyObject8ed989_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*std_core.QChildEvent))(signal))(std_core.NewQChildEventFromPointer(event))
	} else {
		NewTopologyObjectFromPointer(ptr).ChildEventDefault(std_core.NewQChildEventFromPointer(event))
	}
}

func (ptr *TopologyObject) ChildEventDefault(event std_core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TopologyObject8ed989_ChildEventDefault(ptr.Pointer(), std_core.PointerFromQChildEvent(event))
	}
}

//export callbackTopologyObject8ed989_ConnectNotify
func callbackTopologyObject8ed989_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewTopologyObjectFromPointer(ptr).ConnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *TopologyObject) ConnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.TopologyObject8ed989_ConnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackTopologyObject8ed989_CustomEvent
func callbackTopologyObject8ed989_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		(*(*func(*std_core.QEvent))(signal))(std_core.NewQEventFromPointer(event))
	} else {
		NewTopologyObjectFromPointer(ptr).CustomEventDefault(std_core.NewQEventFromPointer(event))
	}
}

func (ptr *TopologyObject) CustomEventDefault(event std_core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TopologyObject8ed989_CustomEventDefault(ptr.Pointer(), std_core.PointerFromQEvent(event))
	}
}

//export callbackTopologyObject8ed989_DeleteLater
func callbackTopologyObject8ed989_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewTopologyObjectFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *TopologyObject) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.TopologyObject8ed989_DeleteLaterDefault(ptr.Pointer())
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackTopologyObject8ed989_Destroyed
func callbackTopologyObject8ed989_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*std_core.QObject))(signal))(std_core.NewQObjectFromPointer(obj))
	}
	qt.Unregister(ptr)

}

//export callbackTopologyObject8ed989_DisconnectNotify
func callbackTopologyObject8ed989_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*std_core.QMetaMethod))(signal))(std_core.NewQMetaMethodFromPointer(sign))
	} else {
		NewTopologyObjectFromPointer(ptr).DisconnectNotifyDefault(std_core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *TopologyObject) DisconnectNotifyDefault(sign std_core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.TopologyObject8ed989_DisconnectNotifyDefault(ptr.Pointer(), std_core.PointerFromQMetaMethod(sign))
	}
}

//export callbackTopologyObject8ed989_Event
func callbackTopologyObject8ed989_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QEvent) bool)(signal))(std_core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTopologyObjectFromPointer(ptr).EventDefault(std_core.NewQEventFromPointer(e)))))
}

func (ptr *TopologyObject) EventDefault(e std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TopologyObject8ed989_EventDefault(ptr.Pointer(), std_core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackTopologyObject8ed989_EventFilter
func callbackTopologyObject8ed989_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*std_core.QObject, *std_core.QEvent) bool)(signal))(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewTopologyObjectFromPointer(ptr).EventFilterDefault(std_core.NewQObjectFromPointer(watched), std_core.NewQEventFromPointer(event)))))
}

func (ptr *TopologyObject) EventFilterDefault(watched std_core.QObject_ITF, event std_core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.TopologyObject8ed989_EventFilterDefault(ptr.Pointer(), std_core.PointerFromQObject(watched), std_core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackTopologyObject8ed989_ObjectNameChanged
func callbackTopologyObject8ed989_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_Moc_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackTopologyObject8ed989_TimerEvent
func callbackTopologyObject8ed989_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*std_core.QTimerEvent))(signal))(std_core.NewQTimerEventFromPointer(event))
	} else {
		NewTopologyObjectFromPointer(ptr).TimerEventDefault(std_core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *TopologyObject) TimerEventDefault(event std_core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.TopologyObject8ed989_TimerEventDefault(ptr.Pointer(), std_core.PointerFromQTimerEvent(event))
	}
}
