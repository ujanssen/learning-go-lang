$(document).ready(function() {
	var ws = new WebSocket("ws://localhost:8044/websocket");

	ws.onmessage = function(e) {
		alert(e.data);
	};
/*
	ws.send($.toJSON({"head" : "greeting", "body": "hello"}));
	ws.send({"head" : "greeting", "body": "hello"});
	ws.send("hello");
*/	
});