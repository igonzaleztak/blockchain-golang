curl localhost:5051/notify -s -S --header 'Content-Type: application/json' --header 'Accept: application/json' --header 'Fiware-Service:RoomsControl' --header 'Fiware-ServicePath:/house1' -X POST -d @- <<EOF
{
  "id": "Room1",
  "type": "Room",
  "attributes": {
    "type": "Object",
    "value": {
      "temperature": {
        "value": 27,
        "type": "Number"
      },
      "sensorID": {
        "value": "sensor1",
        "type": "String"
      },
      "timestamp": {
        "value": 1588004722686,
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
		"value": "0x04f889334300f5f63bbd2233aabc5d4b40f823604073ac96ccd5b79e03da25cddd15c6f0e01b0f293980eaf22402bdcdad77a0b4512fc2ffea58494dfbf61303feb74fe0109d350d0e5c64503e212e22c9f2cda68accd08c98dd70ab28010596e4d6"
  }
}
EOF