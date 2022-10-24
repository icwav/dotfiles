from djitellopy import Tello
from time import sleep
drone = Tello()

drone.connect(False)
# set ssid name and password for individual tello

# drone.set_wifi_credentials("Tello-CBF", "12345678")

# connects drone to router

drone.connect_to_wifi("TP-Link_A66E", "80742982")
