

#pragma once

#ifndef GO_MOC_8ed989_H
#define GO_MOC_8ed989_H

#include <stdint.h>

#ifdef __cplusplus
class TopologyObject8ed989;
void TopologyObject8ed989_TopologyObject8ed989_QRegisterMetaTypes();
class CtxObject8ed989;
void CtxObject8ed989_CtxObject8ed989_QRegisterMetaTypes();
extern "C" {
#endif

struct Moc_PackedString { char* data; long long len; };
struct Moc_PackedList { void* data; long long len; };
void CtxObject8ed989_ConnectIncreased(void* ptr, long long t);
void CtxObject8ed989_DisconnectIncreased(void* ptr);
void CtxObject8ed989_Increased(void* ptr);
void CtxObject8ed989_ConnectDecreased(void* ptr, long long t);
void CtxObject8ed989_DisconnectDecreased(void* ptr);
void CtxObject8ed989_Decreased(void* ptr);
long long CtxObject8ed989_NodeAmount(void* ptr);
long long CtxObject8ed989_NodeAmountDefault(void* ptr);
void CtxObject8ed989_SetNodeAmount(void* ptr, long long nodeAmount);
void CtxObject8ed989_SetNodeAmountDefault(void* ptr, long long nodeAmount);
void CtxObject8ed989_ConnectNodeAmountChanged(void* ptr, long long t);
void CtxObject8ed989_DisconnectNodeAmountChanged(void* ptr);
void CtxObject8ed989_NodeAmountChanged(void* ptr, long long nodeAmount);
int CtxObject8ed989_CtxObject8ed989_QRegisterMetaType();
int CtxObject8ed989_CtxObject8ed989_QRegisterMetaType2(char* typeName);
int CtxObject8ed989_CtxObject8ed989_QmlRegisterType();
int CtxObject8ed989_CtxObject8ed989_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName);
void* CtxObject8ed989___children_atList(void* ptr, int i);
void CtxObject8ed989___children_setList(void* ptr, void* i);
void* CtxObject8ed989___children_newList(void* ptr);
void* CtxObject8ed989___dynamicPropertyNames_atList(void* ptr, int i);
void CtxObject8ed989___dynamicPropertyNames_setList(void* ptr, void* i);
void* CtxObject8ed989___dynamicPropertyNames_newList(void* ptr);
void* CtxObject8ed989___findChildren_atList(void* ptr, int i);
void CtxObject8ed989___findChildren_setList(void* ptr, void* i);
void* CtxObject8ed989___findChildren_newList(void* ptr);
void* CtxObject8ed989___findChildren_atList3(void* ptr, int i);
void CtxObject8ed989___findChildren_setList3(void* ptr, void* i);
void* CtxObject8ed989___findChildren_newList3(void* ptr);
void* CtxObject8ed989___qFindChildren_atList2(void* ptr, int i);
void CtxObject8ed989___qFindChildren_setList2(void* ptr, void* i);
void* CtxObject8ed989___qFindChildren_newList2(void* ptr);
void* CtxObject8ed989_NewCtxObject(void* parent);
void CtxObject8ed989_DestroyCtxObject(void* ptr);
void CtxObject8ed989_DestroyCtxObjectDefault(void* ptr);
void CtxObject8ed989_ChildEventDefault(void* ptr, void* event);
void CtxObject8ed989_ConnectNotifyDefault(void* ptr, void* sign);
void CtxObject8ed989_CustomEventDefault(void* ptr, void* event);
void CtxObject8ed989_DeleteLaterDefault(void* ptr);
void CtxObject8ed989_DisconnectNotifyDefault(void* ptr, void* sign);
char CtxObject8ed989_EventDefault(void* ptr, void* e);
char CtxObject8ed989_EventFilterDefault(void* ptr, void* watched, void* event);
;
void CtxObject8ed989_TimerEventDefault(void* ptr, void* event);
long long TopologyObject8ed989_NodeAmount(void* ptr);
long long TopologyObject8ed989_NodeAmountDefault(void* ptr);
void TopologyObject8ed989_SetNodeAmount(void* ptr, long long nodeAmount);
void TopologyObject8ed989_SetNodeAmountDefault(void* ptr, long long nodeAmount);
void TopologyObject8ed989_ConnectNodeAmountChanged(void* ptr, long long t);
void TopologyObject8ed989_DisconnectNodeAmountChanged(void* ptr);
void TopologyObject8ed989_NodeAmountChanged(void* ptr, long long nodeAmount);
int TopologyObject8ed989_TopologyObject8ed989_QRegisterMetaType();
int TopologyObject8ed989_TopologyObject8ed989_QRegisterMetaType2(char* typeName);
int TopologyObject8ed989_TopologyObject8ed989_QmlRegisterType();
int TopologyObject8ed989_TopologyObject8ed989_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName);
void* TopologyObject8ed989___children_atList(void* ptr, int i);
void TopologyObject8ed989___children_setList(void* ptr, void* i);
void* TopologyObject8ed989___children_newList(void* ptr);
void* TopologyObject8ed989___dynamicPropertyNames_atList(void* ptr, int i);
void TopologyObject8ed989___dynamicPropertyNames_setList(void* ptr, void* i);
void* TopologyObject8ed989___dynamicPropertyNames_newList(void* ptr);
void* TopologyObject8ed989___findChildren_atList(void* ptr, int i);
void TopologyObject8ed989___findChildren_setList(void* ptr, void* i);
void* TopologyObject8ed989___findChildren_newList(void* ptr);
void* TopologyObject8ed989___findChildren_atList3(void* ptr, int i);
void TopologyObject8ed989___findChildren_setList3(void* ptr, void* i);
void* TopologyObject8ed989___findChildren_newList3(void* ptr);
void* TopologyObject8ed989___qFindChildren_atList2(void* ptr, int i);
void TopologyObject8ed989___qFindChildren_setList2(void* ptr, void* i);
void* TopologyObject8ed989___qFindChildren_newList2(void* ptr);
void* TopologyObject8ed989_NewTopologyObject(void* parent);
void TopologyObject8ed989_DestroyTopologyObject(void* ptr);
void TopologyObject8ed989_DestroyTopologyObjectDefault(void* ptr);
void TopologyObject8ed989_ChildEventDefault(void* ptr, void* event);
void TopologyObject8ed989_ConnectNotifyDefault(void* ptr, void* sign);
void TopologyObject8ed989_CustomEventDefault(void* ptr, void* event);
void TopologyObject8ed989_DeleteLaterDefault(void* ptr);
void TopologyObject8ed989_DisconnectNotifyDefault(void* ptr, void* sign);
char TopologyObject8ed989_EventDefault(void* ptr, void* e);
char TopologyObject8ed989_EventFilterDefault(void* ptr, void* watched, void* event);
;
void TopologyObject8ed989_TimerEventDefault(void* ptr, void* event);

#ifdef __cplusplus
}
#endif

#endif