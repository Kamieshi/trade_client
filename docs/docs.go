// Package docs GENERATED BY SWAG; DO NOT EDIT
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
            "url": "https://github.com/Kamieshi"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/position/close/{userID}/{positionID}": {
            "get": {
                "tags": [
                    "position"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "position ID",
                        "name": "positionID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "userID ID",
                        "name": "userID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Position"
                        }
                    },
                    "400": {
                        "description": "Err",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/position/user/{userID}": {
            "get": {
                "tags": [
                    "position"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "user ID",
                        "name": "userID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Position"
                            }
                        }
                    },
                    "400": {
                        "description": "Err",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/position/{companyID}/open": {
            "post": {
                "tags": [
                    "position"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "company ID",
                        "name": "companyID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "postition for open",
                        "name": "position",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Position"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Position"
                        }
                    },
                    "400": {
                        "description": "Err",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/position/{positionID}": {
            "get": {
                "tags": [
                    "position"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "position ID",
                        "name": "positionID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Position"
                        }
                    },
                    "400": {
                        "description": "Err",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/price": {
            "get": {
                "tags": [
                    "price"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Price"
                            }
                        }
                    },
                    "400": {
                        "description": "bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "tags": [
                    "user"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.User"
                            }
                        }
                    },
                    "400": {
                        "description": "Err",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "description": "user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "400": {
                        "description": "bad Request",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/{userID}/updateBalance/{difference}": {
            "get": {
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "difference",
                        "name": "difference",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "userID",
                        "name": "userID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Not found User",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/{userName}": {
            "get": {
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "userName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "404": {
                        "description": "Not found User",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Position": {
            "type": "object",
            "properties": {
                "count_buy_position": {
                    "type": "integer"
                },
                "id": {
                    "type": "string",
                    "readOnly": true
                },
                "is_fixes": {
                    "description": "user limit or not",
                    "type": "boolean"
                },
                "is_opened": {
                    "type": "boolean",
                    "readOnly": true
                },
                "is_sales": {
                    "description": "true/false : sale/buy",
                    "type": "boolean"
                },
                "max_current_cost": {
                    "type": "integer"
                },
                "min_current_cost": {
                    "type": "integer"
                },
                "price": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.Price"
                        }
                    ],
                    "readOnly": true
                },
                "profit": {
                    "type": "integer",
                    "readOnly": true
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "model.Price": {
            "type": "object",
            "properties": {
                "Ask": {
                    "type": "integer"
                },
                "Bid": {
                    "type": "integer"
                },
                "CompanyID": {
                    "type": "string"
                },
                "Name": {
                    "type": "string"
                },
                "Time": {
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "integer"
                },
                "id": {
                    "type": "string",
                    "readOnly": true
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Swagger Client Trade service",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}