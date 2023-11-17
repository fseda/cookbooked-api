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
                "tags": [
                    "Users"
                ],
                "summary": "Login user into the app",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "user-credentials",
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
                        "headers": {
                            "Authorization": {
                                "type": "string",
                                "description": "Bearer \u003ctoken\u003e"
                            }
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
                "tags": [
                    "Users"
                ],
                "summary": "Register user in the app",
                "parameters": [
                    {
                        "description": "New user credentials",
                        "name": "user-info",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.registerUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "headers": {
                            "Authorization": {
                                "type": "string",
                                "description": "Bearer \u003ctoken\u003e"
                            }
                        }
                    }
                }
            }
        },
        "/auth/validate": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Validate the JWT provided in the Authorization header",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Validate JWT",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "401": {
                        "description": "Unauthorized"
                    }
                }
            }
        },
        "/ingredients": {
            "get": {
                "description": "Get all ingredients",
                "tags": [
                    "Ingredients"
                ],
                "summary": "Get all ingredients",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.ingredientsResponse"
                        }
                    }
                }
            }
        },
        "/me": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieve the profile of the currently authenticated user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get logged in user's profile",
                "operationId": "get-user-profile",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.userProfileResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/httpstatus.GlobalErrorHandlerResp"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/httpstatus.GlobalErrorHandlerResp"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/httpstatus.GlobalErrorHandlerResp"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete the account of the authenticated user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Delete logged-in user's account",
                "operationId": "delete-user-by-id",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Invalid id. Should be a positive integer.",
                        "schema": {
                            "$ref": "#/definitions/httpstatus.GlobalErrorHandlerResp"
                        }
                    },
                    "500": {
                        "description": "Could not delete user.",
                        "schema": {
                            "$ref": "#/definitions/httpstatus.GlobalErrorHandlerResp"
                        }
                    }
                }
            }
        },
        "/recipes": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all recipes from a user, by user id",
                "tags": [
                    "Recipes"
                ],
                "summary": "Get all recipes from a user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.getAllRecipesResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/recipes/new": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a new recipe with the given input data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Recipes"
                ],
                "summary": "Create a new recipe",
                "parameters": [
                    {
                        "description": "Recipe creation data",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.createRecipeRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/controllers.createRecipeRequest"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httpstatus.GlobalErrorHandlerResp"
                        }
                    }
                }
            }
        },
        "/recipes/{recipe_id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get a recipe details, by recipe id",
                "tags": [
                    "Recipes"
                ],
                "summary": "Get a recipe details",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Recipe ID",
                        "name": "recipe_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.getRecipeDetailsResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete a recipe, by recipe id",
                "tags": [
                    "Recipes"
                ],
                "summary": "Delete a recipe",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Recipe ID",
                        "name": "recipe_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httpstatus.GlobalErrorHandlerResp"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/httpstatus.GlobalErrorHandlerResp"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update recipe details, by recipe id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Recipes"
                ],
                "summary": "Update recipe details",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Recipe ID",
                        "name": "recipe_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Recipe update data",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.updateRecipeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.updateRecipeRequest"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/httpstatus.GlobalErrorHandlerResp"
                        }
                    }
                }
            }
        },
        "/recipes/{recipe_id}/ingredients": {
            "patch": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Add multiple ingredients to a recipe, if it exists in the recipe update",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Recipes"
                ],
                "summary": "Add multiple ingredients to a recipe",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Recipe ID",
                        "name": "recipe_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Recipe ingredients data",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.recipeIngredientsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/recipes/{recipe_id}/ingredients/{recipe_ingredient_id}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Remove an ingredient from a recipe",
                "tags": [
                    "Recipes"
                ],
                "summary": "Remove an ingredient from a recipe",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Recipe ID",
                        "name": "recipe_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Recipe ingredient ID",
                        "name": "recipe_ingredient_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/units": {
            "get": {
                "description": "Get all units",
                "tags": [
                    "Units"
                ],
                "summary": "Get all units",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.unitsResponse"
                        }
                    }
                }
            }
        },
        "/users/exists": {
            "get": {
                "description": "Check if a user exists by their username or email.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Check if user exists",
                "operationId": "check-user-exists-by-username-or-email",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/users/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieve detailed information of a user based on their ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get user by ID",
                "operationId": "get-user-by-id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.userProfileResponse"
                        }
                    },
                    "400": {
                        "description": "User not Found",
                        "schema": {
                            "$ref": "#/definitions/httpstatus.GlobalErrorHandlerResp"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/httpstatus.GlobalErrorHandlerResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/httpstatus.GlobalErrorHandlerResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.createRecipeRequest": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "recipe_ingredients": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controllers.recipeIngredientRequest"
                    }
                },
                "tag_ids": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "title": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 3
                }
            }
        },
        "controllers.getAllRecipesResponse": {
            "type": "object",
            "properties": {
                "recipes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controllers.getRecipeResponse"
                    }
                }
            }
        },
        "controllers.getRecipeDetailsResponse": {
            "type": "object",
            "properties": {
                "body": {
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
                "title": {
                    "type": "string"
                }
            }
        },
        "controllers.getRecipeResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "controllers.ingredientResponse": {
            "type": "object",
            "properties": {
                "category": {
                    "$ref": "#/definitions/models.IngredientsCategory"
                },
                "category_id": {
                    "type": "integer"
                },
                "icon": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "controllers.ingredientsResponse": {
            "type": "object",
            "properties": {
                "ingredients": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controllers.ingredientResponse"
                    }
                }
            }
        },
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
        "controllers.recipeIngredientRequest": {
            "type": "object",
            "properties": {
                "ingredient_id": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "number"
                },
                "unit_id": {
                    "type": "integer"
                }
            }
        },
        "controllers.recipeIngredientsRequest": {
            "type": "object",
            "properties": {
                "recipe_ingredients": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controllers.recipeIngredientRequest"
                    }
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
                    "minLength": 4
                },
                "username": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 3
                }
            }
        },
        "controllers.unitResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "symbol": {
                    "type": "string"
                },
                "type": {
                    "$ref": "#/definitions/models.Type"
                }
            }
        },
        "controllers.unitsResponse": {
            "type": "object",
            "properties": {
                "units": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/controllers.unitResponse"
                    }
                }
            }
        },
        "controllers.updateRecipeRequest": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "title": {
                    "type": "string",
                    "maxLength": 255,
                    "minLength": 3
                }
            }
        },
        "controllers.userProfileResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "httpstatus.GlobalErrorHandlerResp": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Error message"
                },
                "success": {
                    "type": "boolean",
                    "example": false
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
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
        "models.RecipeIngredient": {
            "type": "object",
            "properties": {
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
                }
            }
        },
        "models.System": {
            "type": "string",
            "enum": [
                "metric",
                "farenheit"
            ],
            "x-enum-varnames": [
                "METRIC",
                "FARENHEIT"
            ]
        },
        "models.Type": {
            "type": "string",
            "enum": [
                "weight",
                "volume",
                "temperature"
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
                "user_id": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
