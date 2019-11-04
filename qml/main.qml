import QtQuick 2.7						//Item
import QtQuick.Controls 2.1				//Dialog
import QtQuick.Layouts 1.13 			//GridLayout
import QtGraphicalEffects 1.12			//DropShadow
import QtQuick.Controls.Material 2.12 	//Material Style imports
import QtQuick.Controls.Styles 1.4 //For styling Node buttons

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
							onClicked: {
								if(ctxObject.nodeAmount == 2) {

								} else {
									ctxObject.decreased()
								}
							}
						}
						Button {
							// anchors.horizontalCenter: topologyCardContents.horizontalCenter
							text: qsTr("Increase Node Amount")
							font.family: Roboto
							onClicked: {
								if(ctxObject.nodeAmount == 16) {

								} else {
									ctxObject.increased()
								}
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

					Grid {
						// padding: 20
						spacing: 20
						// ...
						anchors.horizontalCenter: topologyCardContents.horizontalCenter

						Repeater {
							id: nodeRepeater
							model: NodeModel
							delegate: Button {
								id: nodeButton
								height: 70
								width: 80
								text: nodeName
								// checked: true
								onClicked: nodeDialog.open()
								Dialog {
									id: nodeDialog
									Material.accent: Material.Teal
									Material.theme: Material.Light
									width: 300
									height: 250
									x: (mainPane.width - width) * 0.5
									y: (mainPane.height - height) * 0.5
									parent: mainPane

									contentWidth: mainPane.width * 0.5
									contentHeight: mainPane.height * 0.25
									standardButtons: Dialog.Ok | Dialog.Cancel

									contentItem: Label {
										Column {
											property string textFieldValue: ""
											id: nodeDialogContents
											Text { text: "<b>Node Name: </b>" + nodeName }
											Text { text: "<b>Node Message: </b>" + nodeMessage }
											//A good separator to keep away the rats 
											MenuSeparator {
												padding: 0
												// topPadding: 12
												// bottomPadding: 12
												anchors.horizontalCenter: nodeDialogContents.horizontalCenter
												contentItem: Rectangle {
													implicitWidth: 200
													implicitHeight: 1.5
													color: "#1E000000"
												}
											}
											Text { text: "<b>TRANSMIT MESSAGE</b>"}

											TextField {
												topPadding: 20
												bottomPadding: 20
												placeholderText: "..."
												cursorVisible: false
												// onAccepted: { placeholderText = "...";}
											}
											
										}

									}
									onAccepted: { 
										nodeMessage = textFieldValue 
										console.log("Ok clicked"); 
									}
    								onRejected: console.log("Cancel clicked")
								}

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


