#consumer="http://172.17.0.2:5050/notify"
#consumer="http://192.168.163.108:5050/notify"
#"192.168.163.115"#

curl -v localhost:1026/v2/subscriptions -s -S --header 'Content-type: application/json' --header 'Accept: application/json' --header 'Fiware-Service:smartsantander' --header 'Fiware-ServicePath:/trafficflowobserved' -d @- <<EOF
{
  "description": "A subscription to get the temperature info about Room1",
  "subject": {
    "entities": [
      {
        "id": "urn:ngsi-ld:TrafficFlowObserved:santander:traffic:flow:1022",
        "type": "TrafficFlowObserved"
      }
    ],
    "condition": {
      "attrs": ["attributes"]
    }
  },
  "notification": {
    "http": {
      "url": "http://localhost:5050/notify"
    },
    "attrs" : ["attributes"],
    "metadata": ["hash"]
  },
  "expires": "2040-01-01T14:00:00.00Z",
  "throttling": 1
}
EOF
