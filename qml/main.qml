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
			implicitWidth: 650
			implicitHeight: 500
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
						text: qsTr("Current number of nodes: <b>" + ctxObject.nodeAmount + " </b>")
						font.family: Roboto
					}

					Row {
						spacing: 20
						Button {
							// anchors.horizontalCenter: topologyCardContents.horizontalCenter
							text: qsTr("Decrease Node Amount")
							font.family: Roboto
							onClicked: ctxObject.decreased()
						}
						Button {
							// anchors.horizontalCenter: topologyCardContents.horizontalCenter
							text: qsTr("Increase Node Amount")
							font.family: Roboto
							onClicked: ctxObject.increased()
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

					Grid {
						// padding: 20
						spacing: 20
						// ...
						anchors.horizontalCenter: topologyCardContents.horizontalCenter

						Repeater {
							model: ctxObject.nodeAmount
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