package libs

// PrepareInputData prepares the request
// to match the format that is programmed.
//
// This functions inserts all the measurements into an
// attribute field so they can be stored easily in the
// database
func PrepareInputData(body map[string]interface{}) map[string]interface{} {

	// Create a new body
	auxBody := make(map[string]interface{})

	// Get the id, type, descriptio, cipher fields
	ngsiID := body["id"]
	ngsiType := body["type"]
	desc := body["description"]
	cipher := body["cipher"]

	// Delete these  fields from the body
	delete(body, "id")
	delete(body, "type")
	delete(body, "description")
	delete(body, "cipher")

	// Group the rest of the fields into a new field called
	// attributes.
	//
	// Create an attributes interface
	attr := make(map[string]interface{})
	attr["type"] = "Object"
	attr["value"] = body
	auxBody["attributes"] = attr

	// Add the previous fields
	auxBody["id"] = ngsiID
	auxBody["type"] = ngsiType
	auxBody["description"] = desc
	auxBody["cipher"] = cipher

	return auxBody
}
