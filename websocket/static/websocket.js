var ws = new WebSocket("ws://localhost:8044/websocket");

ws.onmessage = function(e) {
	$("#onmessage").append($("<li>").text(JSON.stringify(e)));
};
ws.onopen = function(e) {
	$("#onopen").append($("<li>").text(JSON.stringify(e)));
};
ws.onerror = function(e) {
	$("#onerror").append($("<li>").text(JSON.stringify(e)));
};
ws.onclose = function(e) {
	$("#onclose").append($("<li>").text(JSON.stringify(e)));
};

var controller = {
	sendMessage: function() {
		console.log("send: " + JSON.stringify(ws))
		if (ws.readyState != WebSocket.OPEN){
		console.log("ws.readyState: " + JSON.stringify(ws.readyState))
			return;
		}
		var msg = {
			head: "msg",
			body: "hello",
		};
		console.log("send: " + JSON.stringify(msg))
		ws.send(JSON.stringify(msg));
	},

	setPingButton: function() {
		$('#ping-button').click(controller.sendMessage());
	},

};

$(document).ready(function() {
	controller.setPingButton();
});