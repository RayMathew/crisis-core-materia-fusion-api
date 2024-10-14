// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Ray Mathew",
            "url": "https://github.com/RayMathew/",
            "email": "ray10mathew@gmail.com"
        },
        "license": {
            "name": "GNU GENERAL PUBLIC LICENSE",
            "url": "https://www.gnu.org/licenses/gpl-3.0.en.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/fusion": {
            "post": {
                "description": "Simulate the fusion of two materia and get the resulting materia.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Fuse two materia",
                "parameters": [
                    {
                        "description": "Fusion Request Data",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.MateriaFusionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Fused Materia Response",
                        "schema": {
                            "$ref": "#/definitions/main.MateriaDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponseDTO"
                        }
                    },
                    "415": {
                        "description": "Unsupported Media Type",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponseDTO"
                        }
                    },
                    "422": {
                        "description": "Failed Validation",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponseDTO"
                        }
                    },
                    "429": {
                        "description": "Too Many Requests",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Error response",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponseDTO"
                        }
                    },
                    "504": {
                        "description": "Request Timed Out",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponseDTO"
                        }
                    }
                }
            }
        },
        "/materia": {
            "get": {
                "description": "Get list of all materia used in the game.",
                "consumes": [
                    "application/json"
                ],
                "summary": "List of all materia",
                "responses": {
                    "200": {
                        "description": "List of all materia",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.MateriaDTO"
                            }
                        }
                    },
                    "429": {
                        "description": "Too Many Requests",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Error response",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponseDTO"
                        }
                    },
                    "504": {
                        "description": "Request Timed Out",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponseDTO"
                        }
                    }
                }
            }
        },
        "/status": {
            "get": {
                "description": "Use this endpoint to check that the server is up and responsive.",
                "consumes": [
                    "application/json"
                ],
                "summary": "Health Check",
                "responses": {
                    "200": {
                        "description": "Successful Response",
                        "schema": {
                            "$ref": "#/definitions/main.StatusDTO"
                        }
                    },
                    "429": {
                        "description": "You have reached maximum request limit.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Error response",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponseDTO"
                        }
                    },
                    "504": {
                        "description": "Request Timed Out",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponseDTO"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.ErrorResponseDTO": {
            "type": "object",
            "properties": {
                "Error": {
                    "type": "string",
                    "example": "The server encountered a problem and could not process your request"
                }
            }
        },
        "main.MateriaDTO": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "Shoots lightning forward dealing thunder damage."
                },
                "name": {
                    "type": "string",
                    "example": "Thunder"
                },
                "type": {
                    "type": "string",
                    "example": "Magic"
                }
            }
        },
        "main.MateriaFusionRequest": {
            "type": "object",
            "properties": {
                "materia1mastered": {
                    "type": "boolean",
                    "example": true
                },
                "materia1name": {
                    "type": "string",
                    "example": "Fire"
                },
                "materia2mastered": {
                    "type": "boolean",
                    "example": false
                },
                "materia2name": {
                    "type": "string",
                    "example": "Blizzard"
                }
            }
        },
        "main.StatusDTO": {
            "type": "object",
            "properties": {
                "Status": {
                    "type": "string",
                    "example": "OK"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "crisis-core-materia-fusion-api-546461677134.us-central1.run.app",
	BasePath:         "",
	Schemes:          []string{"https"},
	Title:            "Crisi Core Materia Fusion API",
	Description:      "A server for simulating Materia Fusion outputs in the game Crisis Core: Final Fantasy VII",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
