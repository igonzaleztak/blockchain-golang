#consumer="http://172.17.0.2:5050/notify"
#consumer="http://192.168.163.108:5050/notify"
#"192.168.163.115"#

curl -v localhost:1026/v2/subscriptions -s -S --header 'Content-type: application/json' --header 'Accept: application/json' --header 'Fiware-Service:RoomsControl' --header 'Fiware-ServicePath:/house1' -d @- <<EOF
{
  "description": "A subscription to get the temperature info about Room1",
  "subject": {
    "entities": [
      {
        "id": "Room1",
        "type": "Room"
      }
    ],
    "condition": {
      "attrs": [ "attributes"]
    }
  },
  "notification": {
    "http": {
      "url": "http://localhost:5050/notify"
    },
    "metadata": ["hash"]
  },
  "expires": "2040-01-01T14:00:00.00Z",
  "throttling": 5
}
EOF
