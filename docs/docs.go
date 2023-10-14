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
            "name": "Felipe Seda"
        },
        "license": {
            "name": "MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "description": "Logs an existing user into the app",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Login into the app",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.loginUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.loginUserResponse"
                        }
                    }
                }
            }
        },
        "/auth/signup": {
            "post": {
                "description": "Registers a new user in the app",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Register user in the app",
                "parameters": [
                    {
                        "description": "New user credentials",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.registerUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.registerUserResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.loginUserRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "controllers.loginUserResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "controllers.registerUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "maxLength": 72,
                    "minLength": 6
                },
                "username": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 3
                }
            }
        },
        "controllers.registerUserResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "ingredients": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Ingredient"
                    }
                },
                "recipes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Recipe"
                    }
                },
                "role": {
                    "$ref": "#/definitions/models.Role"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Tag"
                    }
                },
                "units": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Unit"
                    }
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.Ingredient": {
            "type": "object",
            "properties": {
                "category": {
                    "$ref": "#/definitions/models.IngredientsCategory"
                },
                "category_id": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "icon": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_system_ingredient": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "recipe_ingredients": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.RecipeIngredient"
                    }
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.IngredientsCategory": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Recipe": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "link": {
                    "type": "string"
                },
                "recipe_ingredients": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.RecipeIngredient"
                    }
                },
                "recipe_tags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.RecipeTag"
                    }
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.RecipeIngredient": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "ingredient": {
                    "$ref": "#/definitions/models.Ingredient"
                },
                "ingredient_id": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "number"
                },
                "recipe_id": {
                    "type": "integer"
                },
                "unit": {
                    "$ref": "#/definitions/models.Unit"
                },
                "unit_id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.RecipeTag": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "recipe": {
                    "$ref": "#/definitions/models.Recipe"
                },
                "recipe_id": {
                    "type": "integer"
                },
                "tag": {
                    "$ref": "#/definitions/models.Tag"
                },
                "tag_id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Role": {
            "type": "string",
            "enum": [
                "admin",
                "user"
            ],
            "x-enum-varnames": [
                "ADMIN",
                "USER"
            ]
        },
        "models.System": {
            "type": "string",
            "enum": [
                "METRIC",
                "FARENHEIT"
            ],
            "x-enum-varnames": [
                "METRIC",
                "FARENHEIT"
            ]
        },
        "models.Tag": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_system_tag": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "recipeTags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.RecipeTag"
                    }
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.Type": {
            "type": "string",
            "enum": [
                "WEIGHT",
                "VOLUME",
                "TEMPERATURE"
            ],
            "x-enum-varnames": [
                "WEIGHT",
                "VOLUME",
                "TEMPERATURE"
            ]
        },
        "models.Unit": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_system_unit": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "recipeIngredients": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.RecipeIngredient"
                    }
                },
                "symbol": {
                    "type": "string"
                },
                "system": {
                    "$ref": "#/definitions/models.System"
                },
                "type": {
                    "$ref": "#/definitions/models.Type"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "CookBooked API",
	Description:      "API for CookBooked, a recipe management app.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}