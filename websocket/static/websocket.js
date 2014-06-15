
var controller = {
	ws: {},

	sendMessage: function() {
		console.log("send: " + JSON.stringify(controller.ws))
		if (controller.ws.readyState != WebSocket.OPEN){
			console.log("ws.readyState: " + JSON.stringify(controller.ws.readyState))
			return;
		}
		var msg = {
			head: "ping",
			body: "ping",
		};
		console.log("send: " + JSON.stringify(msg))
		controller.ws.send(JSON.stringify(msg));
	},

	disconnectServer: function() {
		controller.ws.close();		
		$("#onmessage").empty();
		$("#onerror").empty();
		$("#onopen").empty();
		$("#onclose").empty();
	},
	checkConnection: function() {
		setTimeout(function() { 
			controller.checkConnection();
		}, 1000);

		if(typeof controller.ws.readyState == "undefined"){
			console.log("no ready state");
			controller.connectToServer();
		}
	},

	connectToServer: function() {
		console.log("connectToServer")
		controller.ws = new WebSocket("ws://localhost:8044/websocket");
		controller.ws.onmessage = function(e) {
			var msg = JSON.parse(e.data)
			$("#onmessage").append($("<li>").text(msg.head +" - "+ msg.body));
		};
		controller.ws.onopen = function(e) {
			$("#onopen").append($("<li>").text(JSON.stringify(e)));
		};
		controller.ws.onerror = function(e) {
			$("#onerror").append($("<li>").text(JSON.stringify(e)));
			controller.ws = {};
		};
		controller.ws.onclose = function(e) {
			$("#onclose").append($("<li>").text(JSON.stringify(e)));
			controller.ws = {};
		};
	},

	setPingButton: function() {
		$('#ping-button').click(function(){controller.sendMessage()});
	},

	setConnectButton: function() {
		$('#connect-button').click(function(){controller.connectToServer()});
	},

	setDisconnectButton: function() {
		$('#disconnect-button').click(function(){controller.disconnectServer()});
	},
};

$(document).ready(function() {
	controller.setDisconnectButton();
	controller.setConnectButton();
	controller.setPingButton();

	controller.checkConnection();

});