$(document).ready(function() {
	var ws = new WebSocket("ws://localhost:8044/websocket");

	ws.onmessage = function(e) {
		var msg = $.parseJSON(e.data);

		$("#list").append($("<li>").text(msg.head));
		$("#list").append($("<li>").text(msg.body));
	};
/*
	ws.send($.toJSON({"head" : "greeting", "body": "hello"}));
	ws.send({"head" : "greeting", "body": "hello"});
	ws.send("hello");
*/	
});