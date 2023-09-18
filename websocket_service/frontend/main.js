
let socket = null;
let o = document.getElementById("output");
let user = document.getElementById("username");
let message = document.getElementById("message");
// Letting server when leaving
window.onbeforeunload = function () {
  console.log("Leaving");
  let jsonData = {};
  jsonData["type"] = "left";
  socket.send(JSON.stringify(jsonData));
};
document.addEventListener("DOMContentLoaded", function () {
  const name = prompt("What is you name?")
  if (name==="" || name === undefined || name === null){
    alert("Please enter a name to connect to WebSocket")
    window.onload();
  }
  
  socket = new ReconnectingWebSocket(`ws://localhost:8081/ws?name=${name}`, null, {
    debug: true,
    recnnectInterval: 3000,
  });

  const offline = `<span class="badge bg-danger text-3">offline</span>`;
  const online = `<span class="badge bg-success">online</span>`;
  const onFire = `<span class="badge bg-danger">Fire is going on</span>`;
  const offFire = `<span class="badge bg-success">Normal</span>`;

  let status = document.getElementById("status");
  
  let fire = document.getElementById("fire")

console.log(fire)

  socket.onopen = () => {
    console.log("Successfully connected to WebSocket");
    status.innerHTML = online;
  };
  socket.onclose = () => {
    console.log("Failed to connect to WebSocket");
    status.innerHTML = offline;
  };

  socket.onerror = (error) => {
    console.log(error);
  };

  socket.onmessage = (message) => {
    console.log(message.data)
    let data = JSON.parse(message.data);
    switch (data.type) {
      case "list_users":
        let ul = document.getElementById("online");
        while (ul.firstChild) ul.removeChild(ul.firstChild);
        if (data.users.length > 0) {
          data.users.forEach(function (user) {
            let li = document.createElement("li");
            li.appendChild(document.createTextNode(user));
            ul.appendChild(li);
          });
        }
        break;
      case "broadcast":
        o.innerHTML = o.innerHTML + " " + data.message + "<br />";
        break;
    }
  };

  user.addEventListener("change", function () {
    if (user.value.length > 0) {
      let jsonData = {};
      jsonData["action"] = "username";
      jsonData["username"] = this.value;

      socket.send(JSON.stringify(jsonData));
    } else console.log("User is empty");
  });

  message.addEventListener("keydown", function (event) {
    if (event.code === "Enter") {
      if (!socket) {
        console.log("WebSocket is disconnected");
        return false;
      }
      event.preventDefault();
      event.stopPropagation();
      sendMessage();
    }
  });
  document
    .getElementById("sendMessage")
    .addEventListener("click", function (event) {
      if ( message.value !== "") {
        if (!socket) {
          console.log("WebSocket is disconnected");
          return false;
        }
        event.preventDefault();
        event.stopPropagation();
        sendMessage();
      } else alert("Message is empty");
    });
});


// sendMessage
function sendMessage() {
  if ( message.value !== "") {
    let jsonData = {};
    jsonData["type"] = "broadcast"
    jsonData["user_id"] = 1;
    jsonData["chat_id"] = 1;
    jsonData["message"] = message.value;
    socket.send(JSON.stringify(jsonData));
    console.log(jsonData)
    document.getElementById("message").value = "";
  } else alert("Username and message are empty");
}