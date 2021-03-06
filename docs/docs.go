// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Matthew Croft",
            "url": "https://www.linkedin.com/in/matthew-croft-44a5a5b3/"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/votecard": {
            "post": {
                "description": "Creates a VoteCard that can be used in a survey",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create VoteCard",
                "parameters": [
                    {
                        "description": "Create VoteCard request body",
                        "name": "createVoteCardRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.CreateVoteCardRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.VoteCard"
                        }
                    },
                    "400": {
                        "description": "Incorrect request body",
                        "schema": {
                            "$ref": "#/definitions/main.HttpErrorMessage"
                        }
                    }
                }
            }
        },
        "/votecard/{id}": {
            "get": {
                "description": "Returns a VoteCard",
                "produces": [
                    "application/json"
                ],
                "summary": "Get VoteCard",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "VoteCard ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.VoteCard"
                        }
                    },
                    "404": {
                        "description": "VoteCard not found",
                        "schema": {
                            "$ref": "#/definitions/main.HttpErrorMessage"
                        }
                    }
                }
            },
            "put": {
                "description": "Updates count for a certain option in the VoteCard",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update count on a VoteCard",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "VoteCard ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Option to update vote for",
                        "name": "option",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.VoteCard"
                        }
                    },
                    "400": {
                        "description": "Not a valid option",
                        "schema": {
                            "$ref": "#/definitions/main.HttpErrorMessage"
                        }
                    },
                    "404": {
                        "description": "VoteCard not found",
                        "schema": {
                            "$ref": "#/definitions/main.HttpErrorMessage"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.CreateVoteCardRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "options": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "main.HttpErrorMessage": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "main.VoteCard": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "options": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "votes": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "integer"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Survey Voting API",
	Description:      "This is a Survey Voting API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
