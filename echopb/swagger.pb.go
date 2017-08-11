package echopb 

const (
swagger = `{
  "swagger": "2.0",
  "info": {
    "title": "service.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/echo": {
      "post": {
        "operationId": "Echo",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/echopbEchoMessage"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/echopbEchoMessage"
            }
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    },
    "/v1/echo2": {
      "post": {
        "operationId": "Echo",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/echopbEchoMessage"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/echopbEchoMessage"
            }
          }
        ],
        "tags": [
          "Echo2Service"
        ]
      }
    },
    "/v1/echo3": {
      "post": {
        "operationId": "Echo",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/echopbEchoMessage"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/echopbEchoMessage"
            }
          }
        ],
        "tags": [
          "Echo3Service"
        ]
      }
    }
  },
  "definitions": {
    "echopbEchoMessage": {
      "type": "object",
      "properties": {
        "value": {
          "type": "string"
        }
      }
    }
  }
}
`
)
