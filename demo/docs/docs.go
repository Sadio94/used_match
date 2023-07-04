// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/api/v1/:id/user": {
            "get": {
                "description": "用户详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/rest.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/rest.UserInfo"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/user/register": {
            "post": {
                "description": "用户注册",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "description": "用户注册基本信息",
                        "name": "registerInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/rest.RegisterInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/rest.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/v1/users": {
            "get": {
                "description": "用户列表",
                "parameters": [
                    {
                        "maximum": 100,
                        "minimum": 1,
                        "type": "integer",
                        "description": "分页数量(默认15)",
                        "name": "count",
                        "in": "query"
                    },
                    {
                        "maxLength": 30,
                        "minLength": 1,
                        "type": "string",
                        "description": "非必传 但是传了之后需要满足tag对应的规则",
                        "name": "filterName",
                        "in": "query"
                    },
                    {
                        "minimum": 0,
                        "type": "integer",
                        "description": "分页起始位置(默认0)",
                        "name": "index",
                        "in": "query"
                    },
                    {
                        "enum": [
                            1,
                            2
                        ],
                        "type": "integer",
                        "name": "userStatus",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/rest.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/rest.Paging"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "rest.Paging": {
            "type": "object",
            "properties": {
                "metaInfo": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/rest.UserInfo"
                    }
                },
                "totalNum": {
                    "type": "integer"
                }
            }
        },
        "rest.RegisterInfo": {
            "type": "object",
            "required": [
                "email",
                "nickName",
                "password",
                "registerType",
                "repeatPassword"
            ],
            "properties": {
                "email": {
                    "description": "必传 符合邮箱格式",
                    "type": "string"
                },
                "nickName": {
                    "description": "最长不超过30",
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 1
                },
                "password": {
                    "type": "string"
                },
                "registerType": {
                    "description": "之一",
                    "type": "integer",
                    "enum": [
                        1,
                        2,
                        3,
                        4
                    ]
                },
                "repeatPassword": {
                    "description": "需要和Password字段一样",
                    "type": "string"
                }
            }
        },
        "rest.Result": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {}
            }
        },
        "rest.UserInfo": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "nickName": {
                    "type": "string"
                },
                "registerTime": {
                    "type": "string"
                },
                "updateTime": {
                    "type": "string"
                },
                "userStatus": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
