package core

var webPage = []byte(`
<!DOCTYPE html>
<html>
<head>
    <title>Wireless Mouse Pad</title>
    <style>
        body {
            height: 2000px; /* Add enough content to enable scrolling */
            display: flex;
            align-items: center;
            justify-content: center;
        }
        
        .button-container {
            display: flex;
            flex-direction: column;
            align-items: center;
        }
        
        .scroll-button {
            width: 700px;
            height: 500px;
            margin-bottom: 250px;
            font-size: 75px;
            background-color: green;
            color: white;
        }
    </style>
</head>
<body>
    <script>
        // Prevent zoom on double-tap
        var lastTouchEnd = 0;
        document.addEventListener('touchend', function(event) {
            var now = new Date().getTime();
            if (now - lastTouchEnd <= 300) {
                event.preventDefault();
            }
            lastTouchEnd = now;
        }, false);

        if ('WebSocket' in window) {
            var protocol = 'ws://';
            var ws_request_path = '/sakthimahendran/wireless/mouse/pad';
            var address = protocol + window.location.host + ws_request_path;
            var socket = new WebSocket(address);
        } else {
            window.alert("Your browser doesn't support Wireless Mouse Pad");
            console.log("Your browser doesn't support Wireless Mouse Pad");
        }

        var scrollInterval;

        function scrollUp() {
            var message = {
                deltaX: 0,
                deltaY: 1  // Negative value indicates scrolling up
            };
            var jsonMessage = JSON.stringify(message);
            socket.send(jsonMessage);
        }

        function scrollDown() {
            var message = {
                deltaX: 0,
                deltaY: -1  // Positive value indicates scrolling down
            };
            var jsonMessage = JSON.stringify(message);
            socket.send(jsonMessage);
        }

        function startScrollingUp() {
            scrollInterval = setInterval(scrollUp, 200); // Adjust the scroll speed as desired
        }

        function startScrollingDown() {
            scrollInterval = setInterval(scrollDown, 200); // Adjust the scroll speed as desired
        }

        function stopScrolling() {
            clearInterval(scrollInterval);
        }
    </script>
    <div class="button-container">
        <button class="scroll-button" onmousedown="startScrollingUp()" onmouseup="stopScrolling()" ontouchstart="startScrollingUp()" ontouchend="stopScrolling()">Up</button>
        <button class="scroll-button" onmousedown="startScrollingDown()" onmouseup="stopScrolling()" ontouchstart="startScrollingDown()" ontouchend="stopScrolling()">Down</button>
    </div>d
</body>
</html>
`)
