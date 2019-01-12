package api

import (
	"errors"
	"fmt"
	"log"
	"net/url"

	"hub/config"
)

type CreateSecretResponse struct {
	Id string
}

func CreateSecret(entityKind, selector, name, component, kind string, values map[string]string) {
	id := ""
	resource := ""
	var parameters []Parameter
	var qErr error

	switch entityKind {
	case "environment":
		env, err := environmentBy(selector)
		qErr = err
		if err == nil && env != nil {
			id = env.Id
			resource = environmentsResource
			parameters = env.Parameters
		}
	case "stackTemplate":
		template, err := templateBy(selector)
		qErr = err
		if err == nil && template != nil {
			id = template.Id
			resource = templatesResource
			parameters = template.Parameters
		}
	case "stackInstance":
		instance, err := stackInstanceBy(selector)
		qErr = err
		if err == nil && instance != nil {
			id = instance.Id
			resource = stackInstancesResource
			parameters = instance.Parameters
		}
	default:
		log.Fatalf("Unknown entity kind `%s`", entityKind)
	}
	if id == "" && qErr == nil {
		qErr = errors.New("Not Found")
	}
	if qErr != nil {
		log.Fatalf("Unable to query %s: %v", entityKind, qErr)
	}

	for _, existing := range parameters {
		if existing.Value != "" && existing.Name == name &&
			((existing.Component == "" && component == "") || existing.Component == component) {
			qname := name
			if component != "" {
				qname = fmt.Sprintf("%s|%s", name, component)
			}
			log.Fatalf("Parameter `%s` already exist in %s `%s` and is not empty", qname, entityKind, selector)
		}
	}

	secretId, err := createSecret(resource, id, name, component, kind, values)
	if err != nil {
		log.Fatalf("Unable to create %s secret: %v", entityKind, err)
	}
	if config.Verbose {
		log.Printf("Secret `%s` created in %s `%s` with id `%s`",
			name, entityKind, selector, secretId)
	}
	if config.Verbose {
		switch entityKind {
		case "environment":
			Environments(selector, false, false, false, false, false)
		case "stackTemplate":
			Templates(selector, false, false, false, false)
		case "stackInstance":
			StackInstances(selector, false, false)
		}
	}
}

func createSecret(resource, id, name, component, kind string, values map[string]string) (string, error) {
	values["name"] = name
	if component != "" {
		values["component"] = component
	}
	values["kind"] = kind
	path := fmt.Sprintf("%s/%s/secrets", resource, url.PathEscape(id))
	var jsResp CreateSecretResponse
	code, err := post(hubApi, path, values, &jsResp)
	if err != nil {
		return "", fmt.Errorf("Error creating Hub Service `%s` Secret `%s`: %v",
			id, name, err)
	}
	if code != 201 {
		return "", fmt.Errorf("Got %d HTTP creating Hub Service `%s` Secret `%s`, expected 201 HTTP",
			code, id, name)
	}
	return jsResp.Id, nil
}

func secret(resource, id string) (map[string]string, error) {
	path := fmt.Sprintf("%s/secrets/%s", resource, url.PathEscape(id))
	var jsResp map[string]string
	code, err := get(hubApi, path, &jsResp)
	if err != nil {
		return nil, fmt.Errorf("Error querying Hub Service `%s` Secret `%s`: %v",
			resource, id, err)
	}
	if code != 200 {
		return nil, fmt.Errorf("Got %d HTTP querying Hub Service `%s` Secret `%s`, expected 200 HTTP",
			code, resource, id)
	}
	return jsResp, nil
}

func formatSecret(s map[string]string) string {
	str := ""
	if kind, ok := s["kind"]; ok {
		switch kind {
		case "text", "password", "certificate", "sshKey", "privateKey", "loginToken":
			str = s[kind]
		case "gitAccessToken":
			str = s["loginToken"]
		case "usernamePassword":
			str = fmt.Sprintf("%s/%s", s["username"], s["password"])
		case "cloudAccessKeys":
			str = fmt.Sprintf("%s:%s", s["accessKey"], s["secretKey"])
		case "cloudAccount":
			str = fmt.Sprintf("%s/%s", s["roleArn"], s["externalId"])
		}
	}
	if str == "" {
		str = fmt.Sprintf("%+v", s)
	}
	return str
}
