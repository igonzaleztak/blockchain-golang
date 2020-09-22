curl localhost:5053/notify -s -S --header 'Content-Type: application/json' --header 'Accept: application/json' --header 'Fiware-ServicePath:/' -X POST -d @- <<EOF
{
	"data": [
	  {
		"id": "urn:ngsi-ld:TrafficFlowObserved:santander:traffic:flow:1001",
		"type": "TrafficFlowObserved",
		"dateModified": {
			"type": "ISO8601",
			"value": "2020-09-09T13:00:00.00Z",
			"metadata": {

			}
		},
		"dateObserved": {
			"type": "ISO8601",
			"value": "2020-09-09T13:00:00.00Z",
			"metadata": {

			}
		},
		"intensity": {
			"type": "Number",
			"value": 305,
			"metadata": {

			}
		},
		"laneId": {
			"type": "Number",
			"value": 0,
			"metadata": {

			}
		},
		"location": {
			"type": "geo:json",
			"value": {
				"type": "Point",
				"coordinates": [
					-3.8295937,
					43.4535859
				]
			},
			"metadata": {

			}
		},
		"occupancy": {
			"type": "Number",
			"value": 0.08,
			"metadata": {

			}
		},
		"roadLoad": {
			"type": "Number",
			"value": 17,
			"metadata": {

			}
		}
	}]
}
EOF