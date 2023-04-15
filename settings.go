package main

import (
	kubewarden "github.com/kubewarden/policy-sdk-go"
	kubewarden_protocol "github.com/kubewarden/policy-sdk-go/protocol"
	easyjson "github.com/mailru/easyjson"
)

// Checks if given string is a palindrome
// O(N)time and O(N)space
func (s *Settings) isPalindrome(stringValue string) bool {
	result := []byte{}
	for i := len(stringValue) - 1; i >= 0; i-- {
		result = append(result, stringValue[i])
	}

	return stringValue == string(result)
}

func NewSettingsFromValidationReq(validationReq *kubewarden_protocol.ValidationRequest) (Settings, error) {
	settings := Settings{}
	err := easyjson.Unmarshal(validationReq.Settings, &settings)
	return settings, err
}

func validateSettings(payload []byte) ([]byte, error) {
	logger.Info("validating settings")
	return kubewarden.AcceptSettings()
}
