# Import the necessary modules
import socket
import time
import os


class Swarm:
    def __init__(self, drone_range=range(1), rescan=False) -> None:

        if rescan:
            self.update_scan(drone_range)

        self.drone_ips = open('drone_ips.txt', 'r')
        # 8889 is the default port Tellos communicate on
        self.drone_ips = tuple((line.strip(), 8889) for line in self.drone_ips)

        # creates unique listening ports which correspond to the last two number's of the drones ips
        self.listening_ports = tuple(('', 9010 + int(drone[0][-2:]))
                                     for drone in self.drone_ips)

        self.socks = tuple(socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
                           for _ in range(len(self.drone_ips)))

        # binding socks
        for i, sock in enumerate(self.socks):
            sock.bind(self.listening_ports[i])

    @staticmethod
    def update_scan(drone_range):
        # scan the network and append only the ips to a file
        scan = os.popen('nmap -sn -T5 192.168.0.101-199').read()
        scan = scan.split("report")
        scan.pop(0)

        # that way you don't rescan every time unnecessarily when flying the same swarm
        try:
            ips = tuple(scan[ip][scan[ip].index(
                'for')+4:scan[ip].index('H')-1] for ip in drone_range)

            with open('drone_ips.txt', 'w') as drone_ips:
                for ip in ips:
                    drone_ips.write(ip+'\n')

        except IndexError:
            print("Not enough drones in range, exiting..")

    # How we receive Tello error messages and info
    def receive(self):
        cancel_flight = False
        # Continuously loop and listen for incoming messages
        while True:
            # Try to receive the message otherwise print the exception
            try:
                responses = []
                for sock in self.socks:
                    responses.append(sock.recvfrom(128)[0])

                for i, response in enumerate(responses):
                    print(
                        f"Received message: from Tello EDU #{i}: " + response.decode(encoding='utf-8'))
                    # since our port and the drones ips are the same I can use the computers port to find which drone is dying
                    # if response.decode(encoding='utf-8') < "30":
                    #     print(
                    #         f"Drone ip: 192.168.0.1{self.socks[0].getsockname()[1] - 9000} is low battery!")
                    #     cancel_flight = True
                break
            except Exception as e:
                # If there's an error close the socket and break out of the loop
                for sock in self.socks:
                    sock.close()
                print("Error receiving: " + str(e))
                break
        # return cancel_flight

    # Send the message to Tello and allow for a delay in seconds
    def send(self, message, delay):
        # Try to send the message otherwise print the exception
        try:
            for i, sock in enumerate(self.socks):
                sock.sendto(message.encode(), self.drone_ips[i])

            print("Sending message: " + message)
        except Exception as e:
            print("Error sending: " + str(e))

        # Delay for a user-defined period of time
        time.sleep(delay)


# Example flight path for large room 
# DO NOT USE in enclosed area 

fleet_1 = Swarm(range(1))
fleet_2 = Swarm(range(3), False, "drone_ips2.txt")
fleet_3 = Swarm(range(3), False, "drone_ips3.txt")


fleet_1.send('command', 5)
fleet_2.send('command', 5)
fleet_3.send('command', 5)

fleet_1.send('takeoff', 5)
fleet_2.send('takeoff', 5)
fleet_3.send('takeoff', 5)


# # fleet 3 flies up 400, 2 up 600, 1 up 800

fleet_1.send('up 800', 4)
fleet_2.send('up 600', 4)
fleet_3.send('up 400', 4)


# #  fleet 1 rotates 360 cw, 2 rotates 360 ccw, 3 rotates 360 cw
fleet_1.send('cw 360', 4)
fleet_2.send('ccw 3600', 4)
fleet_3.send('cw 360', 4)



# # fleet 1 down 400, 3 up 400
fleet_1.send('down 400', 3)
fleet_3.send('up 400', 3)



# up for debate

# #  flip right, left, forward
fleet_1.send('flip r', 4)
fleet_2.send('flip r', 4)
fleet_3.send('flip r', 4)

fleet_1.send('flip l', 4)
fleet_2.send('flip l', 4)
fleet_3.send('flip l', 4)

fleet_1.send('flip f', 3)
fleet_2.send('flip f', 3)
fleet_3.send('flip f', 3)


fleet_1.send('land', 5)
fleet_2.send('land', 5)
fleet_3.send('land', 5)