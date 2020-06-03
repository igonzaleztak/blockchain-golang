curl localhost:5051/notify -s -S --header 'Content-Type: application/json' --header 'Accept: application/json' --header 'Fiware-Service:RoomsControl' --header 'Fiware-ServicePath:/house1' -X POST -d @- <<EOF
{
  "id": "Room1",
  "type": "Room",
  "attributes": {
    "type": "Object",
    "value": {
      "temperature": {
        "value": 25,
        "type": "Number"
      },
      "sensorID": {
        "value": "sensor1",
        "type": "String"
      },
      "timestamp": {
        "value": 1588050810892,
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
		"value": "0x04596ed847c44a0d84d112f773fef57b1e71d0e2c410788b3554bd3aed2c432ff8e20545f41bfa7a540007fb20acfc8d3cc3d2068bbf5b901f0040612135f82a0742a049049209204701fcc7410c39c58faf88f5b5536239bf30c0f0a09f0f35ead1"
  }
}
EOF
