package main

type launchRequest struct {
	Request request `json:"request"`
}

type request struct {
	Type string `json:"type"`
}

type prevresp struct {
	Text string `json:"text"`
}

// alexa response

type backendResponse struct {
	Version           string `json:"version"`
	SessionAttributes struct {
		Key string `json:"key,omitempty"`
	} `json:"sessionAttributes,omitempty"`
	Response struct {
		OutputSpeech     outputSpeech  `json:"outputSpeech,omitempty"`
		Card             *card         `json:"card,omitempty"`
		Reprompt         *reprompt     `json:"reprompt,omitempty"`
		Directives       *[]directives `json:"directives,omitempty"`
		ShouldEndSession bool          `json:"shoulEndSession,omitempty"`
	} `json:"response"`
	Request struct {
		Type        string  `json:"type,omitempty"`
		RequestID   string  `json:"requestId,omitempty"`
		Timestamp   string  `json:"timestamp,omitempty"`
		DialogState string  `json:"dialogState,omitempty"`
		Locale      string  `json:"locale,omitempty"`
		Intent      *intent `json:"intent,omitempty"`
		Reason      string  `json:"reason,omitempty"`
		Error       *Error  `json:"error,omitempty"`
	} `json:"request"`
}

type outputSpeech struct {
	Type         string `json:"type"`
	Text         string `json:"text"`
	PlayBehavior string `json:"playBehavior,omitempty"`
}

type reprompt struct {
	OutputSpeech outputSpeech `json:"outputSpeech,omitempty"`
}

type directives struct {
	Type string `json:"type,omitempty"`
}

type card struct {
	Type  string `json:"type"`
	Title string `json:"title,omitempty"`
	Text  string `json:"text,omitempty"`
	Image struct {
		SmallImageURL string `json:"smallImageUrl,omitempty"`
		LargeImageURL string `json:"largeImageUrl,omitempty"`
	} `json:"image,omitempty"`
}

type intent struct {
	Name               string                 `json:"name"`
	ConfirmationStatus string                 `json:"confirmationStatus,omitempty"`
	Slots              map[string]interface{} `json:"slots,omitempty"`
}

// Error ...
type Error struct {
	Type    string `json:"type,omitempty"`
	Message string `json:"message,omitempty"`
}
