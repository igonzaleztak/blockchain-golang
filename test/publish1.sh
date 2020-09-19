curl localhost:5051/notify -s -S --header 'Content-Type: application/json' --header 'Accept: application/json' --header 'Fiware-Service:RoomsControl' --header 'Fiware-ServicePath:/house1' -X POST -d @- <<EOF
{
  "id": "Room1",
  "type": "Room",
  "attributes": {
    "type": "Object",
    "value": {
      "temperature": {
        "value": 15,
        "type": "Number"
      },
      "sensorID": {
        "value": "sensor1",
        "type": "String"
      },
      "timestamp": {
        "value": 1595412064,
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
		"value": "0x0435c015c0a1076894ba8f862db6c356dab54ddfdee61160daf8ba795b758180c18ee2ee014c8ddc31419df932e57eefa37f72193760210a479d345e7d9f9bb2d4e6decff266eeb7148aa7d8a31523e4ce74149e6e18eab5bde5462f12c1bda74901"
  }
}
EOF