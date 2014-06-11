$(document).ready(function() {
	var ws = new WebSocket("ws://localhost:8044/websocket");

	ws.onmessage = function(e) {
		$("#onmessage").append($("<li>").text(JSON.stringify(e)));

		var msg = $.parseJSON(e.data);
		msg.body = "server"
		console.log("send: " + JSON.stringify(msg))
		ws.send(JSON.stringify(msg));
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
});