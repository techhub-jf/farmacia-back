// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/farmacia-tech/auth": {
            "post": {
                "description": "Authenticates a user and returns a JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "User Login",
                "parameters": [
                    {
                        "description": "Login Credentials",
                        "name": "loginRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Login successful",
                        "schema": {
                            "$ref": "#/definitions/schema.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/v1/farmacia-tech/clients": {
            "get": {
                "description": "Returns Clients",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Client"
                ],
                "summary": "Get Clients",
                "parameters": [
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "\"name\"",
                        "description": "Sort by field",
                        "name": "sort_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "\"asc\"",
                        "description": "Sort type (asc/desc)",
                        "name": "sort_type",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 10,
                        "description": "Limit of records per page",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of clients",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/schema.ClientResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/v1/farmacia-tech/deliveries": {
            "get": {
                "description": "Returns deliveries",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Delivery"
                ],
                "summary": "Get Deliveries",
                "parameters": [
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 10,
                        "description": "Number of items per page",
                        "name": "items_per_page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "\"name\"",
                        "description": "Field to sort by",
                        "name": "sort_by",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "\"asc\"",
                        "description": "Type of sorting (asc/desc)",
                        "name": "sort_type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of deliveries\"s",
                        "schema": {
                            "$ref": "#/definitions/schema.ListDeliveriesOutput"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "post": {
                "description": "Create a new delivery record",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Delivery"
                ],
                "summary": "Create Delivery",
                "parameters": [
                    {
                        "description": "Delivery data",
                        "name": "delivery",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.CreateDeliveryRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created delivery",
                        "schema": {
                            "$ref": "#/definitions/schema.CreateDeliveryResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/v1/farmacia-tech/deliveries/reference/{reference}": {
            "get": {
                "description": "Get details of a specific delivery using its reference",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Delivery"
                ],
                "summary": "Get Delivery by Reference",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Delivery Reference",
                        "name": "reference",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Delivery details",
                        "schema": {
                            "$ref": "#/definitions/schema.GetDeliveryByReferenceResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/v1/farmacia-tech/medicine": {
            "get": {
                "description": "Returns medicines",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Medicine"
                ],
                "summary": "Get Medicines",
                "responses": {
                    "200": {
                        "description": "List of medicines",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Medicine"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Medicine": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "email": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "medicine_id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "qty": {
                    "type": "integer"
                },
                "unit_id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "schema.ClientResponse": {
            "type": "object",
            "properties": {
                "cpf": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "phone": {
                    "type": "string"
                },
                "reference": {
                    "type": "string"
                },
                "rg": {
                    "type": "string"
                }
            }
        },
        "schema.CreateDeliveryRequest": {
            "type": "object",
            "properties": {
                "client_id": {
                    "type": "integer"
                },
                "medicine_id": {
                    "type": "integer"
                },
                "qty": {
                    "type": "integer"
                },
                "unit_id": {
                    "type": "integer"
                }
            }
        },
        "schema.CreateDeliveryResponse": {
            "type": "object",
            "properties": {
                "delivery": {
                    "$ref": "#/definitions/schema.CreatedDeliveryResponse"
                }
            }
        },
        "schema.CreatedDeliveryResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "reference": {
                    "type": "string"
                }
            }
        },
        "schema.DeliveryResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "qty": {
                    "type": "integer"
                },
                "reference": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "schema.GetDeliveryByReferenceResponse": {
            "type": "object",
            "properties": {
                "delivery": {
                    "$ref": "#/definitions/schema.DeliveryResponse"
                }
            }
        },
        "schema.ListDeliveriesOutput": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.ListDeliveriesResponse"
                    }
                },
                "metadata": {
                    "$ref": "#/definitions/schema.Meta"
                }
            }
        },
        "schema.ListDeliveriesResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "qty": {
                    "type": "integer"
                },
                "reference": {
                    "type": "string"
                }
            }
        },
        "schema.LoginRequest": {
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
        "schema.LoginResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "schema.Meta": {
            "type": "object",
            "properties": {
                "current_page": {
                    "type": "integer"
                },
                "items_per_page": {
                    "type": "integer"
                },
                "total_items": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Farmacia-back API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
