const net = require("net")

const server = net.createServer(
    socket => {
        socket.write("Hello.")
        socket.on("data", data => {
            console.log(data.toString())
        })
    }
)
const port = 8080
console.log("Server is running on port", port)
server.listen(port)