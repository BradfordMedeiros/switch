#!/usr/bin/env bash

function publish-message(){
	TOPIC=$(echo $0 | awk '{ print $0 }')
	MESSAGE=$(echo $0 | awk '{ print $1 }')
	mqtt --publish -t "$TOPIC" -m "$MESSAGE"
}

./listen-mqtt-messages.sh | switch -i "
	(floor1 -> floor2) : up 
	(floor2 -> floor3) : up
	(floor3 -> floor4) : up
	(floor4 -> floor3) : down
	(floor3 -> floor2) : down
	(floor2 -> floor1) : down
	(floor1 -> off) : shutdown - power_state_off
	(off -> floor1) : turnon - power_state_on
	
	start as off
	when - power_state_off - /housing/room1/power on 
	when - power_state_on - /housing/room2/power off

" | publish-message



