from djitellopy import Tello
import cv2
from threading import Thread


epicDrone = Tello()

epicDrone.connect()
print(f'Battery: {epicDrone()}%')

keepRecording = True
epicDrone.streamon()
frame_read = epicDrone.get_frame_read()

##def videoRecorder():
##    height, width, _ = frame_read.frame.shape
##    video = cv2.VideoWriter('video.avi', cv2 VideoWriter_fourrcc (*'XVID'), 30, (width, height))
##
##    while keepRecording:
##        video.write(frame_read.frame)
##        time.sleep(1/30)
##    
##    video.release()


epicDrone.takeoff()

epicDrone.move_up(100)

epicDrone.move_left(100)
epicDrone.flip_forward()







