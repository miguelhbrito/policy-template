package main

import (
	"fmt"
	"strings"

	onelog "github.com/francoispqt/onelog"
	corev1 "github.com/kubewarden/k8s-objects/api/core/v1"
	kubewarden "github.com/kubewarden/policy-sdk-go"
	kubewarden_protocol "github.com/kubewarden/policy-sdk-go/protocol"
	"github.com/mailru/easyjson"
)

func validate(payload []byte) ([]byte, error) {
	// Create a ValidationRequest instance from the incoming payload
	validationRequest := kubewarden_protocol.ValidationRequest{}
	err := easyjson.Unmarshal(payload, &validationRequest)
	if err != nil {
		return kubewarden.RejectRequest(
			kubewarden.Message(err.Error()),
			kubewarden.Code(400))
	}

	// Create a Settings instance from the ValidationRequest object
	settings, err := NewSettingsFromValidationReq(&validationRequest)
	if err != nil {
		return kubewarden.RejectRequest(
			kubewarden.Message(err.Error()),
			kubewarden.Code(400))
	}

	// Access the **raw** JSON that describes the object
	podJSON := validationRequest.Request.Object

	// Try to create a Pod instance using the RAW JSON we got from the
	// ValidationRequest.
	pod := &corev1.Pod{}
	if err := easyjson.Unmarshal([]byte(podJSON), pod); err != nil {
		return kubewarden.RejectRequest(
			kubewarden.Message(
				fmt.Sprintf("Cannot decode Pod object: %s", err.Error())),
			kubewarden.Code(400))
	}

	logger.DebugWithFields("validating pod object", func(e onelog.Entry) {
		e.String("name", pod.Metadata.Name)
		e.String("namespace", pod.Metadata.Namespace)
	})

	if settings.IsNameDenied(pod.Metadata.Name) {
		logger.InfoWithFields("rejecting pod object", func(e onelog.Entry) {
			e.String("name", pod.Metadata.Name)
			e.String("denied_names", strings.Join(settings.DeniedNames, ","))
		})

		return kubewarden.RejectRequest(
			kubewarden.Message(
				fmt.Sprintf("The '%s' name is on the deny list", pod.Metadata.Name)),
			kubewarden.NoCode)
	}

	logger.Debug("validating labels keys")

	// Validates every key label to check if is a palindrome
	var keysLabels []string
	for key, _ := range pod.Metadata.Labels {
		if settings.IsPalindrome(key) {
			logger.InfoWithFields("rejecting label key", func(e onelog.Entry) {
				e.String("key label", key)
			})

			keysLabels = append(keysLabels, key)
		}
	}
	if len(keysLabels) > 0 {
		return kubewarden.RejectRequest(
			kubewarden.Message(
				fmt.Sprintf("The keys labels '%v' are palindromes", keysLabels)),
			kubewarden.NoCode)
	}

	return kubewarden.AcceptRequest()
}
