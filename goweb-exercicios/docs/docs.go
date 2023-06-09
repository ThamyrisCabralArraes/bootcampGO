// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones",
        "contact": {
            "name": "API Support",
            "url": "https://developers.mercadolibre.com.ar/support"
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
        "/produtos/getAll": {
            "get": {
                "description": "get produtos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Produtos"
                ],
                "summary": "List produtos",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            }
        },
        "/produtos/post": {
            "post": {
                "description": "salvar Produtos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Produtos"
                ],
                "summary": "Salvar produtos",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Produto para Salvar",
                        "name": "Produto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.request"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            }
        },
        "/produtos/updateNome/{id}": {
            "patch": {
                "description": "Atualiza o nome de um produto",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Produtos"
                ],
                "summary": "UpdateNome produtos",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID do produto a ser atualizado",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Nome do produto a ser atualizado",
                        "name": "nome",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            }
        },
        "/produtos/updatePreco/{id}": {
            "patch": {
                "description": "UpdatePreco Produtos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Produtos"
                ],
                "summary": "UpdatePreco produtos",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID do produto a ser atualizado",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Preco do produto a ser atualizado",
                        "name": "preco",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            }
        },
        "/produtos/{id}": {
            "put": {
                "description": "Atualiza os detalhes de um produto",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Produtos"
                ],
                "summary": "Update produtos",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID do produto a ser atualizado",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Detalhes atualizados do produto",
                        "name": "Produto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.request"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete Produtos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Produtos"
                ],
                "summary": "Delete produtos",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ID do produto a ser excluído",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/web.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.request": {
            "type": "object",
            "properties": {
                "codigo": {
                    "type": "string"
                },
                "cor": {
                    "type": "string"
                },
                "dataCriacao": {
                    "type": "string"
                },
                "estoque": {
                    "type": "integer"
                },
                "nome": {
                    "type": "string"
                },
                "preco": {
                    "type": "number"
                },
                "publicacao": {
                    "type": "boolean"
                }
            }
        },
        "web.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "data": {},
                "error": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "MELI Bootcamp API - Produtos",
	Description:      "This API Handle MELI Products.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
