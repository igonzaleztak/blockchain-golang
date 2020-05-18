curl localhost:5051/notify -s -S --header 'Content-Type: application/json' --header 'Accept: application/json' --header 'Fiware-Service:RoomsControl' --header 'Fiware-ServicePath:/house1' -X POST -d @- <<EOF
{
  "id": "Room1",
  "type": "Room",
  "attributes": {
    "type": "Object",
    "value": {
      "temperature": {
        "value": 21,
        "type": "Number"
      },
      "sensorID": {
        "value": "sensor1",
        "type": "String"
      },
      "timestamp": {
        "value": 1588008799886,
        "type": "Number"
      }
    }
  },
  "description": {
    "value": "Measurement of the temperature in the living room of house 1",
    "type": "String"
  },
  "cipher": {
    "type": "String",
		"value": "0xfccae1c24e6f9d0363cdae2cb9ea9eb202ca0020b47164ce7501572e7394437665d1dde2f19f51144457cb5fed65b135b1a075800020f43883224a8079d66b2d31d1505063a33f21ad700c536a4
fd5556a7db416712262d796051684c5d8238f80db030ace2cd101ffa5c1a8ca2055b82c573c6ae65dd0e3b20d763cc5a8bef4a5c05c58fb1e"
  }
}
EOF
