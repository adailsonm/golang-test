// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Adailson Moreira",
            "email": "adailson.moreira16@gmail.com"
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
        "/login": {
            "post": {
                "description": "Login with email and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login to the platform",
                "parameters": [
                    {
                        "description": "user email",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Handler.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Login successful",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/profile": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get the details of the logged-in user's profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get user profile",
                "responses": {
                    "200": {
                        "description": "User profile",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Register a new user with their details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "Data Form User",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Registration successful",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/slot/history": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get the history of previous game spins",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Game"
                ],
                "summary": "Get game history",
                "responses": {
                    "200": {
                        "description": "History of spins",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/slot/spin": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Make a bet on the slot machine",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Game"
                ],
                "summary": "Spin a slot machine",
                "parameters": [
                    {
                        "description": "Spint Reques",
                        "name": "spinGame",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Handler.SpinRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Spin result",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/wallet/deposit": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Deposit a specified amount into the wallet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wallet"
                ],
                "summary": "Deposit money into wallet",
                "parameters": [
                    {
                        "description": "Wallet Data",
                        "name": "wallet",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Models.Wallet"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Deposit successful",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/wallet/withdraw": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Withdraw a specified amount from the wallet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Wallet"
                ],
                "summary": "Withdraw money from wallet",
                "parameters": [
                    {
                        "description": "Wallet Data",
                        "name": "wallet",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Models.Wallet"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Withdraw successful",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "Handler.LoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "Handler.SpinRequest": {
            "type": "object",
            "properties": {
                "bet_amount": {
                    "type": "number"
                }
            }
        },
        "Models.User": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "number"
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "Models.Wallet": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "transaction": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/Models.User"
                },
                "userId": {
                    "type": "string"
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
	Title:            "NodeArt - Golang",
	Description:      "Created by Adailson",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
