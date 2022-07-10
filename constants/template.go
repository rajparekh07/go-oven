package constants

type TemplateString struct {
	IP string
}

var templateStrings = TemplateString{
	IP: "<IP_ADDRESS>",
}

func Get() TemplateString {
	return templateStrings
}
