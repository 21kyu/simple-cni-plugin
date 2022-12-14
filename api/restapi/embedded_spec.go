// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Serve REST API Endpoints for Simple CNI plugin",
    "title": "Simple CNI Server",
    "version": "0.1.0"
  },
  "paths": {
    "/ipam": {
      "post": {
        "tags": [
          "ipam"
        ],
        "summary": "Allocate IP address",
        "parameters": [
          {
            "$ref": "#/parameters/ipam-family"
          },
          {
            "$ref": "#/parameters/ipam-owner"
          }
        ],
        "responses": {
          "201": {
            "description": "IP address allocated",
            "schema": {
              "$ref": "#/definitions/ipam-response"
            }
          }
        }
      }
    },
    "/ipam/{ip}": {
      "delete": {
        "tags": [
          "ipam"
        ],
        "summary": "Release an allocated IP address",
        "parameters": [
          {
            "$ref": "#/parameters/ipam-ip"
          }
        ],
        "responses": {
          "200": {
            "description": "IP address released"
          },
          "400": {
            "description": "Invalid IP address"
          },
          "404": {
            "description": "IP address not found"
          },
          "500": {
            "description": "IP address release failed"
          }
        }
      }
    }
  },
  "definitions": {
    "ipam-address": {
      "description": "Addressing information",
      "type": "object",
      "required": [
        "ip",
        "gateway"
      ],
      "properties": {
        "gateway": {
          "description": "Gateway address",
          "type": "string"
        },
        "ip": {
          "description": "IP address",
          "type": "string"
        }
      }
    },
    "ipam-response": {
      "description": "IPAM configuration",
      "type": "object",
      "required": [
        "address"
      ],
      "properties": {
        "address": {
          "$ref": "#/definitions/ipam-address"
        }
      }
    }
  },
  "parameters": {
    "ipam-family": {
      "enum": [
        "ipv4"
      ],
      "type": "string",
      "description": "IP family",
      "name": "family",
      "in": "query"
    },
    "ipam-ip": {
      "type": "string",
      "description": "IP address",
      "name": "ip",
      "in": "path",
      "required": true
    },
    "ipam-owner": {
      "type": "string",
      "description": "Owner of the IP address",
      "name": "owner",
      "in": "query"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Serve REST API Endpoints for Simple CNI plugin",
    "title": "Simple CNI Server",
    "version": "0.1.0"
  },
  "paths": {
    "/ipam": {
      "post": {
        "tags": [
          "ipam"
        ],
        "summary": "Allocate IP address",
        "parameters": [
          {
            "enum": [
              "ipv4"
            ],
            "type": "string",
            "description": "IP family",
            "name": "family",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Owner of the IP address",
            "name": "owner",
            "in": "query"
          }
        ],
        "responses": {
          "201": {
            "description": "IP address allocated",
            "schema": {
              "$ref": "#/definitions/ipam-response"
            }
          }
        }
      }
    },
    "/ipam/{ip}": {
      "delete": {
        "tags": [
          "ipam"
        ],
        "summary": "Release an allocated IP address",
        "parameters": [
          {
            "type": "string",
            "description": "IP address",
            "name": "ip",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "IP address released"
          },
          "400": {
            "description": "Invalid IP address"
          },
          "404": {
            "description": "IP address not found"
          },
          "500": {
            "description": "IP address release failed"
          }
        }
      }
    }
  },
  "definitions": {
    "ipam-address": {
      "description": "Addressing information",
      "type": "object",
      "required": [
        "ip",
        "gateway"
      ],
      "properties": {
        "gateway": {
          "description": "Gateway address",
          "type": "string"
        },
        "ip": {
          "description": "IP address",
          "type": "string"
        }
      }
    },
    "ipam-response": {
      "description": "IPAM configuration",
      "type": "object",
      "required": [
        "address"
      ],
      "properties": {
        "address": {
          "$ref": "#/definitions/ipam-address"
        }
      }
    }
  },
  "parameters": {
    "ipam-family": {
      "enum": [
        "ipv4"
      ],
      "type": "string",
      "description": "IP family",
      "name": "family",
      "in": "query"
    },
    "ipam-ip": {
      "type": "string",
      "description": "IP address",
      "name": "ip",
      "in": "path",
      "required": true
    },
    "ipam-owner": {
      "type": "string",
      "description": "Owner of the IP address",
      "name": "owner",
      "in": "query"
    }
  }
}`))
}
