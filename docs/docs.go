// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/movies": {
            "get": {
                "description": "Get all movies in order of their release date from earliest to newest in the cache or from swapi if the cache is empty",
                "produces": [
                    "application/json"
                ],
                "summary": "Get all movies",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Movie"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    }
                }
            }
        },
        "/api/v1/movies/{movie_id}/characters/{filter}/{sort_by}/{order}": {
            "get": {
                "description": "Get all characters for a movie by movie id use the sort parameter",
                "produces": [
                    "application/json"
                ],
                "summary": "Get characters",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Movie ID",
                        "name": "movie_id",
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
                                "$ref": "#/definitions/models.Character"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    }
                }
            }
        },
        "/api/v1/movies/{movie_id}/comments": {
            "get": {
                "description": "Get all comments for a movie by movie id",
                "produces": [
                    "application/json"
                ],
                "summary": "Get comments",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Movie ID",
                        "name": "movie_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    }
                }
            },
            "post": {
                "description": "Add comments to a movie by movie id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Add comments",
                "parameters": [
                    {
                        "description": "Comment",
                        "name": "comment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CommentRequest"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "MovieId",
                        "name": "movie_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Comment"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ApiError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ApiError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.Character": {
            "type": "object",
            "properties": {
                "birth_year": {
                    "type": "string"
                },
                "eye_color": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "hair_color": {
                    "type": "string"
                },
                "height": {
                    "type": "string"
                },
                "mass": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "skin_color": {
                    "type": "string"
                }
            }
        },
        "models.Comment": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "ip": {
                    "type": "string"
                },
                "movie_id": {
                    "type": "integer"
                }
            }
        },
        "models.CommentRequest": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                }
            }
        },
        "models.Movie": {
            "type": "object",
            "properties": {
                "comment_count": {
                    "type": "integer"
                },
                "episode_id": {
                    "type": "integer"
                },
                "opening_crawl": {
                    "type": "string"
                },
                "release_date": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "swapi-movie-api.herokuapp.com",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Go + Gin Movie API",
	Description: "This is a movie server. You can visit the GitHub repository at https://github.com/olusola/movie-api",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register("swagger", &s{})
}
