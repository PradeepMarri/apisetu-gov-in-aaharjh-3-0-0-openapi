package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// AcademicCertificateSchema represents the AcademicCertificateSchema schema from the OpenAPI specification
type AcademicCertificateSchema struct {
	Issuedate string `json:"issueDate"`
	Language string `json:"language"`
	Number int `json:"number"`
	TypeField string `json:"type"`
	Status string `json:"status"`
	Certificatedata map[string]interface{} `json:"CertificateData"`
	Issuedto map[string]interface{} `json:"IssuedTo"`
	Validfromdate string `json:"validFromDate"`
	Issuedat string `json:"issuedAt"`
	Name string `json:"name"`
	Issuedby map[string]interface{} `json:"IssuedBy"`
}

// ConsentArtifactSchema represents the ConsentArtifactSchema schema from the OpenAPI specification
type ConsentArtifactSchema struct {
	Consent map[string]interface{} `json:"consent"`
	Signature map[string]interface{} `json:"signature"`
}
