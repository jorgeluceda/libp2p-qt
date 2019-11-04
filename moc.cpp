

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QLayout>
#include <QMetaMethod>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QQuickItem>
#include <QString>
#include <QTimerEvent>
#include <QWidget>
#include <QWindow>

#ifdef QT_QML_LIB
	#include <QQmlEngine>
#endif


class CtxObject8ed989: public QObject
{
Q_OBJECT
Q_PROPERTY(qint64 nodeAmount READ nodeAmount WRITE setNodeAmount NOTIFY nodeAmountChanged)
public:
	CtxObject8ed989(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");CtxObject8ed989_CtxObject8ed989_QRegisterMetaType();CtxObject8ed989_CtxObject8ed989_QRegisterMetaTypes();callbackCtxObject8ed989_Constructor(this);};
	void Signal_Increased() { callbackCtxObject8ed989_Increased(this); };
	void Signal_Decreased() { callbackCtxObject8ed989_Decreased(this); };
	qint64 nodeAmount() { return callbackCtxObject8ed989_NodeAmount(this); };
	void setNodeAmount(qint64 nodeAmount) { callbackCtxObject8ed989_SetNodeAmount(this, nodeAmount); };
	void Signal_NodeAmountChanged(qint64 nodeAmount) { callbackCtxObject8ed989_NodeAmountChanged(this, nodeAmount); };
	 ~CtxObject8ed989() { callbackCtxObject8ed989_DestroyCtxObject(this); };
	void childEvent(QChildEvent * event) { callbackCtxObject8ed989_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackCtxObject8ed989_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackCtxObject8ed989_CustomEvent(this, event); };
	void deleteLater() { callbackCtxObject8ed989_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackCtxObject8ed989_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackCtxObject8ed989_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackCtxObject8ed989_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackCtxObject8ed989_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackCtxObject8ed989_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackCtxObject8ed989_TimerEvent(this, event); };
	qint64 nodeAmountDefault() { return _nodeAmount; };
	void setNodeAmountDefault(qint64 p) { if (p != _nodeAmount) { _nodeAmount = p; nodeAmountChanged(_nodeAmount); } };
signals:
	void increased();
	void decreased();
	void nodeAmountChanged(qint64 nodeAmount);
public slots:
private:
	qint64 _nodeAmount;
};

Q_DECLARE_METATYPE(CtxObject8ed989*)


void CtxObject8ed989_CtxObject8ed989_QRegisterMetaTypes() {
}

class TopologyObject8ed989: public QObject
{
Q_OBJECT
Q_PROPERTY(qint64 nodeAmount READ nodeAmount WRITE setNodeAmount NOTIFY nodeAmountChanged)
public:
	TopologyObject8ed989(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");TopologyObject8ed989_TopologyObject8ed989_QRegisterMetaType();TopologyObject8ed989_TopologyObject8ed989_QRegisterMetaTypes();callbackTopologyObject8ed989_Constructor(this);};
	qint64 nodeAmount() { return callbackTopologyObject8ed989_NodeAmount(this); };
	void setNodeAmount(qint64 nodeAmount) { callbackTopologyObject8ed989_SetNodeAmount(this, nodeAmount); };
	void Signal_NodeAmountChanged(qint64 nodeAmount) { callbackTopologyObject8ed989_NodeAmountChanged(this, nodeAmount); };
	 ~TopologyObject8ed989() { callbackTopologyObject8ed989_DestroyTopologyObject(this); };
	void childEvent(QChildEvent * event) { callbackTopologyObject8ed989_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackTopologyObject8ed989_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackTopologyObject8ed989_CustomEvent(this, event); };
	void deleteLater() { callbackTopologyObject8ed989_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackTopologyObject8ed989_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackTopologyObject8ed989_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackTopologyObject8ed989_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackTopologyObject8ed989_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackTopologyObject8ed989_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackTopologyObject8ed989_TimerEvent(this, event); };
	qint64 nodeAmountDefault() { return _nodeAmount; };
	void setNodeAmountDefault(qint64 p) { if (p != _nodeAmount) { _nodeAmount = p; nodeAmountChanged(_nodeAmount); } };
signals:
	void nodeAmountChanged(qint64 nodeAmount);
public slots:
private:
	qint64 _nodeAmount;
};

Q_DECLARE_METATYPE(TopologyObject8ed989*)


void TopologyObject8ed989_TopologyObject8ed989_QRegisterMetaTypes() {
}

void CtxObject8ed989_ConnectIncreased(void* ptr, long long t)
{
	QObject::connect(static_cast<CtxObject8ed989*>(ptr), static_cast<void (CtxObject8ed989::*)()>(&CtxObject8ed989::increased), static_cast<CtxObject8ed989*>(ptr), static_cast<void (CtxObject8ed989::*)()>(&CtxObject8ed989::Signal_Increased), static_cast<Qt::ConnectionType>(t));
}

void CtxObject8ed989_DisconnectIncreased(void* ptr)
{
	QObject::disconnect(static_cast<CtxObject8ed989*>(ptr), static_cast<void (CtxObject8ed989::*)()>(&CtxObject8ed989::increased), static_cast<CtxObject8ed989*>(ptr), static_cast<void (CtxObject8ed989::*)()>(&CtxObject8ed989::Signal_Increased));
}

void CtxObject8ed989_Increased(void* ptr)
{
	static_cast<CtxObject8ed989*>(ptr)->increased();
}

void CtxObject8ed989_ConnectDecreased(void* ptr, long long t)
{
	QObject::connect(static_cast<CtxObject8ed989*>(ptr), static_cast<void (CtxObject8ed989::*)()>(&CtxObject8ed989::decreased), static_cast<CtxObject8ed989*>(ptr), static_cast<void (CtxObject8ed989::*)()>(&CtxObject8ed989::Signal_Decreased), static_cast<Qt::ConnectionType>(t));
}

void CtxObject8ed989_DisconnectDecreased(void* ptr)
{
	QObject::disconnect(static_cast<CtxObject8ed989*>(ptr), static_cast<void (CtxObject8ed989::*)()>(&CtxObject8ed989::decreased), static_cast<CtxObject8ed989*>(ptr), static_cast<void (CtxObject8ed989::*)()>(&CtxObject8ed989::Signal_Decreased));
}

void CtxObject8ed989_Decreased(void* ptr)
{
	static_cast<CtxObject8ed989*>(ptr)->decreased();
}

long long CtxObject8ed989_NodeAmount(void* ptr)
{
	return static_cast<CtxObject8ed989*>(ptr)->nodeAmount();
}

long long CtxObject8ed989_NodeAmountDefault(void* ptr)
{
	return static_cast<CtxObject8ed989*>(ptr)->nodeAmountDefault();
}

void CtxObject8ed989_SetNodeAmount(void* ptr, long long nodeAmount)
{
	static_cast<CtxObject8ed989*>(ptr)->setNodeAmount(nodeAmount);
}

void CtxObject8ed989_SetNodeAmountDefault(void* ptr, long long nodeAmount)
{
	static_cast<CtxObject8ed989*>(ptr)->setNodeAmountDefault(nodeAmount);
}

void CtxObject8ed989_ConnectNodeAmountChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<CtxObject8ed989*>(ptr), static_cast<void (CtxObject8ed989::*)(qint64)>(&CtxObject8ed989::nodeAmountChanged), static_cast<CtxObject8ed989*>(ptr), static_cast<void (CtxObject8ed989::*)(qint64)>(&CtxObject8ed989::Signal_NodeAmountChanged), static_cast<Qt::ConnectionType>(t));
}

void CtxObject8ed989_DisconnectNodeAmountChanged(void* ptr)
{
	QObject::disconnect(static_cast<CtxObject8ed989*>(ptr), static_cast<void (CtxObject8ed989::*)(qint64)>(&CtxObject8ed989::nodeAmountChanged), static_cast<CtxObject8ed989*>(ptr), static_cast<void (CtxObject8ed989::*)(qint64)>(&CtxObject8ed989::Signal_NodeAmountChanged));
}

void CtxObject8ed989_NodeAmountChanged(void* ptr, long long nodeAmount)
{
	static_cast<CtxObject8ed989*>(ptr)->nodeAmountChanged(nodeAmount);
}

int CtxObject8ed989_CtxObject8ed989_QRegisterMetaType()
{
	return qRegisterMetaType<CtxObject8ed989*>();
}

int CtxObject8ed989_CtxObject8ed989_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<CtxObject8ed989*>(const_cast<const char*>(typeName));
}

int CtxObject8ed989_CtxObject8ed989_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<CtxObject8ed989>();
#else
	return 0;
#endif
}

int CtxObject8ed989_CtxObject8ed989_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<CtxObject8ed989>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* CtxObject8ed989___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CtxObject8ed989___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CtxObject8ed989___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* CtxObject8ed989___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CtxObject8ed989___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* CtxObject8ed989___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* CtxObject8ed989___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CtxObject8ed989___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CtxObject8ed989___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CtxObject8ed989___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CtxObject8ed989___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CtxObject8ed989___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CtxObject8ed989___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CtxObject8ed989___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CtxObject8ed989___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CtxObject8ed989_NewCtxObject(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new CtxObject8ed989(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new CtxObject8ed989(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new CtxObject8ed989(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new CtxObject8ed989(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new CtxObject8ed989(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new CtxObject8ed989(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new CtxObject8ed989(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new CtxObject8ed989(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new CtxObject8ed989(static_cast<QWindow*>(parent));
	} else {
		return new CtxObject8ed989(static_cast<QObject*>(parent));
	}
}

void CtxObject8ed989_DestroyCtxObject(void* ptr)
{
	static_cast<CtxObject8ed989*>(ptr)->~CtxObject8ed989();
}

void CtxObject8ed989_DestroyCtxObjectDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void CtxObject8ed989_ChildEventDefault(void* ptr, void* event)
{
	static_cast<CtxObject8ed989*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void CtxObject8ed989_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<CtxObject8ed989*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void CtxObject8ed989_CustomEventDefault(void* ptr, void* event)
{
	static_cast<CtxObject8ed989*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void CtxObject8ed989_DeleteLaterDefault(void* ptr)
{
	static_cast<CtxObject8ed989*>(ptr)->QObject::deleteLater();
}

void CtxObject8ed989_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<CtxObject8ed989*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char CtxObject8ed989_EventDefault(void* ptr, void* e)
{
	return static_cast<CtxObject8ed989*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char CtxObject8ed989_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<CtxObject8ed989*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}



void CtxObject8ed989_TimerEventDefault(void* ptr, void* event)
{
	static_cast<CtxObject8ed989*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

long long TopologyObject8ed989_NodeAmount(void* ptr)
{
	return static_cast<TopologyObject8ed989*>(ptr)->nodeAmount();
}

long long TopologyObject8ed989_NodeAmountDefault(void* ptr)
{
	return static_cast<TopologyObject8ed989*>(ptr)->nodeAmountDefault();
}

void TopologyObject8ed989_SetNodeAmount(void* ptr, long long nodeAmount)
{
	static_cast<TopologyObject8ed989*>(ptr)->setNodeAmount(nodeAmount);
}

void TopologyObject8ed989_SetNodeAmountDefault(void* ptr, long long nodeAmount)
{
	static_cast<TopologyObject8ed989*>(ptr)->setNodeAmountDefault(nodeAmount);
}

void TopologyObject8ed989_ConnectNodeAmountChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<TopologyObject8ed989*>(ptr), static_cast<void (TopologyObject8ed989::*)(qint64)>(&TopologyObject8ed989::nodeAmountChanged), static_cast<TopologyObject8ed989*>(ptr), static_cast<void (TopologyObject8ed989::*)(qint64)>(&TopologyObject8ed989::Signal_NodeAmountChanged), static_cast<Qt::ConnectionType>(t));
}

void TopologyObject8ed989_DisconnectNodeAmountChanged(void* ptr)
{
	QObject::disconnect(static_cast<TopologyObject8ed989*>(ptr), static_cast<void (TopologyObject8ed989::*)(qint64)>(&TopologyObject8ed989::nodeAmountChanged), static_cast<TopologyObject8ed989*>(ptr), static_cast<void (TopologyObject8ed989::*)(qint64)>(&TopologyObject8ed989::Signal_NodeAmountChanged));
}

void TopologyObject8ed989_NodeAmountChanged(void* ptr, long long nodeAmount)
{
	static_cast<TopologyObject8ed989*>(ptr)->nodeAmountChanged(nodeAmount);
}

int TopologyObject8ed989_TopologyObject8ed989_QRegisterMetaType()
{
	return qRegisterMetaType<TopologyObject8ed989*>();
}

int TopologyObject8ed989_TopologyObject8ed989_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<TopologyObject8ed989*>(const_cast<const char*>(typeName));
}

int TopologyObject8ed989_TopologyObject8ed989_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<TopologyObject8ed989>();
#else
	return 0;
#endif
}

int TopologyObject8ed989_TopologyObject8ed989_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<TopologyObject8ed989>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* TopologyObject8ed989___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TopologyObject8ed989___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TopologyObject8ed989___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* TopologyObject8ed989___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TopologyObject8ed989___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* TopologyObject8ed989___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* TopologyObject8ed989___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TopologyObject8ed989___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TopologyObject8ed989___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TopologyObject8ed989___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TopologyObject8ed989___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TopologyObject8ed989___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TopologyObject8ed989___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TopologyObject8ed989___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TopologyObject8ed989___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TopologyObject8ed989_NewTopologyObject(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new TopologyObject8ed989(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new TopologyObject8ed989(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new TopologyObject8ed989(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new TopologyObject8ed989(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new TopologyObject8ed989(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new TopologyObject8ed989(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new TopologyObject8ed989(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new TopologyObject8ed989(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new TopologyObject8ed989(static_cast<QWindow*>(parent));
	} else {
		return new TopologyObject8ed989(static_cast<QObject*>(parent));
	}
}

void TopologyObject8ed989_DestroyTopologyObject(void* ptr)
{
	static_cast<TopologyObject8ed989*>(ptr)->~TopologyObject8ed989();
}

void TopologyObject8ed989_DestroyTopologyObjectDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void TopologyObject8ed989_ChildEventDefault(void* ptr, void* event)
{
	static_cast<TopologyObject8ed989*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void TopologyObject8ed989_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<TopologyObject8ed989*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void TopologyObject8ed989_CustomEventDefault(void* ptr, void* event)
{
	static_cast<TopologyObject8ed989*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void TopologyObject8ed989_DeleteLaterDefault(void* ptr)
{
	static_cast<TopologyObject8ed989*>(ptr)->QObject::deleteLater();
}

void TopologyObject8ed989_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<TopologyObject8ed989*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char TopologyObject8ed989_EventDefault(void* ptr, void* e)
{
	return static_cast<TopologyObject8ed989*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char TopologyObject8ed989_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<TopologyObject8ed989*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}



void TopologyObject8ed989_TimerEventDefault(void* ptr, void* event)
{
	static_cast<TopologyObject8ed989*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
