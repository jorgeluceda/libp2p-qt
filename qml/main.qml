import QtQuick 2.7			//Item
import QtQuick.Controls 2.1	//Dialog
import QtQuick.Layouts 1.13 //GridLayout
import QtGraphicalEffects 1.12
import QtQuick.Controls.Material 2.12 //Material Style imports

Column {
	id: mainColumn
	Pane {
		id: mainPane
		Material.accent: Material.Teal
		Material.theme: Material.Light
		anchors.horizontalCenter: mainColumn.horizontalCenter
		anchors.verticalCenter: mainColumn.verticalCenter

		Pane {
			Material.elevation: 6
			implicitWidth: 600
			implicitHeight: 400
			Label {
				id: topologyCard
				Column {
					id: topologyCardContents
					Text {
						// anchors.horizontalCenter: topologyCardContents.horizontalCenter
						text: qsTr("<b>Network Topology Structure<b>")
						font.family: Roboto
						font.pointSize: 18
					}

					Text {
						// anchors.horizontalCenter: topologyCardContents.horizontalCenter
						text: qsTr("Current amount of nodes <b> {5} </b>")
						font.family: Roboto
					}


					Slider {
						from: 1
						value: 5
						to: 10
						stepSize: 1
						snapMode: SnapAlways
						// tickmarksEnabled: true
					}

					//A good separator to keep away the rats 
					MenuSeparator {
						padding: 0
						// topPadding: 12
						// bottomPadding: 12
						anchors.horizontalCenter: topologyCardContents.horizontalCenter
						contentItem: Rectangle {
							implicitWidth: 575
							implicitHeight: 1.5
							color: "#1E000000"
						}
					}

					Grid {
						padding: 20
						spacing: 20
						// ...

						Repeater {
							model: 5
							delegate: Button {
								height: 70
								text: qsTr("Received!")
								checkable: true
							}
						}
					}

					//A good separator to keep away the rats 
					MenuSeparator {
						padding: 0
						// topPadding: 12
						// bottomPadding: 12
						anchors.horizontalCenter: topologyCardContents.horizontalCenter
						contentItem: Rectangle {
							implicitWidth: 575
							implicitHeight: 1.5
							color: "#1E000000"
						}
					}
				}

			}
		}
		
	}
}


// ColumnLayout {
// 	id: column

// 	Item {
// 		width: 700
// 		height: 550
// 		anchors.horizontalCenter: column.horizontalCenter
// 		Item {
// 			id: container
// 			anchors.centerIn: parent
// 			width:  rect.width  + (2 * rectShadow.radius)
// 			height: rect.height + (2 * rectShadow.radius)
// 			visible: false

// 			Rectangle {
// 				id: rect
// 				width: 650
// 				height: 525
// 				color: "#f2f2f2"
// 				radius: 7
// 				antialiasing: true
// 				anchors.centerIn: parent

// 				border {
// 					width: 2
// 					color: "#f2f2f2"
// 				}

// 				Button {
// 					text: "<b>Network Topology Established</i>"
// 				}
// 			}
// 		}

// 		DropShadow {
// 			id: rectShadow
// 			anchors.fill: source
// 			cached: true
// 			horizontalOffset: 1
// 			verticalOffset: 1
// 			radius: 8.0
// 			samples: 16
// 			color: "#f2f2f2"
// 			smooth: true
// 			source: container
// 		}
// 	}

// 	Item {
// 		width: 700
// 		height: 250
// 		anchors.horizontalCenter: column.horizontalCenter
// 		Item {
// 			id: container2
// 			anchors.centerIn: parent
// 			width:  rect2.width  + (2 * rectShadow.radius)
// 			height: rect2.height + (2 * rectShadow.radius)
// 			visible: false

// 			Rectangle {
// 				id: rect2
// 				width: 650
// 				height: 200
// 				color: "#dbdbdb"
// 				radius: 7
// 				antialiasing: true
// 				border {
// 					width: 2
// 					color: "#dbdbdb"
// 				}
// 				anchors.centerIn: parent
// 			}
// 		}

// 		DropShadow {
// 			id: rectShadow2
// 			anchors.fill: source
// 			cached: true
// 			horizontalOffset: 1
// 			verticalOffset: 1
// 			radius: 8.0
// 			samples: 16
// 			color: "#dbdbdb"
// 			smooth: true
// 			source: container2
// 		}
// 	}

// }

