import websocket

GUEST_IP = "192.168.100.193"
GUEST_PORT = "8080"

if __name__ == "__main__":
    websocket.enableTrace(True)
    ws = websocket.create_connection("ws://{0}:{1}/".format(GUEST_IP, GUEST_PORT))
    opts = ('cpu', 'mem', 'iface', 'disk', 'uptime')

    for opt in opts:
        ws.send(opt)
        result = ws.recv()
        print("Received: {}".format(result))
