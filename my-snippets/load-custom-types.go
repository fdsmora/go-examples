func loadCustomType(filePath string) *model.CustomType{
	obj := loadObject(filePath, &model.CustomType{})
	request, ok := obj.(*model.CustomType)
	if !ok {
		log.Fatalf("error type asserting '%T' from file '%s'", &model.CustomType{}, filePath)
	}
	return request
}

func loadObject(filePath string, target interface{}) interface{} {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("error loading object from file: %s. err: %s", filePath, err)
	}
	if err = json.Unmarshal(fileBytes, target); err != nil {
		log.Fatalf("error parsing JSON object from file: %s, err: %s", filePath, err)
	}
	return target
}

func main () {
    var c *model.CustomType = loadCustomType("some-file-path-to-json-file/custom-type.json")
    // use c
}
