/*
 * Copyright 2025 coze-dev Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package workflow

// ExportWorkflowRequest 工作流导出请求
type ExportWorkflowRequest struct {
	WorkflowID          int64  `json:"workflow_id" form:"workflow_id" query:"workflow_id"`                               // 工作流ID
	Version             string `json:"version,omitempty" form:"version" query:"version"`                                // 版本号，不传则导出草稿
	IncludeDependencies bool   `json:"include_dependencies,omitempty" form:"include_dependencies" query:"include_dependencies"` // 是否包含依赖资源
	ExportFormat        string `json:"export_format,omitempty" form:"export_format" query:"export_format"`              // 导出格式，默认json
}

// ExportWorkflowResponse 工作流导出响应
type ExportWorkflowResponse struct {
	Code int64                    `json:"code" form:"code" query:"code"`
	Msg  string                   `json:"msg" form:"msg" query:"msg"`
	Data *WorkflowExportData      `json:"data,omitempty" form:"data" query:"data"`
}

// WorkflowExportData 工作流导出数据
type WorkflowExportData struct {
	WorkflowExport *WorkflowExport `json:"workflow_export" form:"workflow_export" query:"workflow_export"`
}

// WorkflowExport 工作流导出结构
type WorkflowExport struct {
	ExportVersion string                 `json:"export_version" form:"export_version" query:"export_version"` // 导出版本
	ExportTime    string                 `json:"export_time" form:"export_time" query:"export_time"`          // 导出时间
	Meta          *WorkflowExportMeta    `json:"meta" form:"meta" query:"meta"`                              // 工作流元数据
	Canvas        *WorkflowExportCanvas  `json:"canvas" form:"canvas" query:"canvas"`                        // 工作流画布
	Dependencies  *WorkflowDependencies  `json:"dependencies,omitempty" form:"dependencies" query:"dependencies"` // 依赖资源
}

// WorkflowExportMeta 工作流导出元数据
type WorkflowExportMeta struct {
	ID          int64  `json:"id" form:"id" query:"id"`                               // 工作流ID
	Name        string `json:"name" form:"name" query:"name"`                         // 工作流名称
	Description string `json:"description" form:"description" query:"description"`   // 工作流描述
	IconURI     string `json:"icon_uri" form:"icon_uri" query:"icon_uri"`             // 图标URI
	CreatorID   int64  `json:"creator_id" form:"creator_id" query:"creator_id"`       // 创建者ID
	SpaceID     int64  `json:"space_id" form:"space_id" query:"space_id"`             // 空间ID
	CreatedAt   string `json:"created_at" form:"created_at" query:"created_at"`       // 创建时间
	UpdatedAt   string `json:"updated_at" form:"updated_at" query:"updated_at"`       // 更新时间
	Version     string `json:"version,omitempty" form:"version" query:"version"`      // 版本号
}

// WorkflowExportCanvas 工作流导出画布
type WorkflowExportCanvas struct {
	Nodes       []*Node `json:"nodes" form:"nodes" query:"nodes"`             // 节点列表
	Edges       []*Edge `json:"edges" form:"edges" query:"edges"`             // 边列表
	InputParams any     `json:"input_params,omitempty" form:"input_params" query:"input_params"` // 输入参数
	OutputParams any    `json:"output_params,omitempty" form:"output_params" query:"output_params"` // 输出参数
}

// WorkflowDependencies 工作流依赖资源
type WorkflowDependencies struct {
	Plugins       []*PluginDependency       `json:"plugins,omitempty" form:"plugins" query:"plugins"`             // 插件依赖
	KnowledgeBases []*KnowledgeDependency   `json:"knowledge_bases,omitempty" form:"knowledge_bases" query:"knowledge_bases"` // 知识库依赖
	Databases     []*DatabaseDependency     `json:"databases,omitempty" form:"databases" query:"databases"`       // 数据库依赖
	SubWorkflows  []*SubWorkflowDependency  `json:"sub_workflows,omitempty" form:"sub_workflows" query:"sub_workflows"` // 子工作流依赖
}

// PluginDependency 插件依赖
type PluginDependency struct {
	ID          int64  `json:"id" form:"id" query:"id"`                               // 插件ID
	Name        string `json:"name" form:"name" query:"name"`                         // 插件名称
	Version     string `json:"version" form:"version" query:"version"`                // 插件版本
	ServerURL   string `json:"server_url" form:"server_url" query:"server_url"`       // 服务器URL
	PluginType  int    `json:"plugin_type" form:"plugin_type" query:"plugin_type"`    // 插件类型
	Manifest    any    `json:"manifest,omitempty" form:"manifest" query:"manifest"`   // 插件清单
	OpenAPIDoc  any    `json:"openapi_doc,omitempty" form:"openapi_doc" query:"openapi_doc"` // OpenAPI文档
}

// KnowledgeDependency 知识库依赖
type KnowledgeDependency struct {
	ID          int64  `json:"id" form:"id" query:"id"`                               // 知识库ID
	Name        string `json:"name" form:"name" query:"name"`                         // 知识库名称
	Description string `json:"description" form:"description" query:"description"`   // 知识库描述
	FormatType  int    `json:"format_type" form:"format_type" query:"format_type"`    // 格式类型
	IconURI     string `json:"icon_uri" form:"icon_uri" query:"icon_uri"`             // 图标URI
}

// DatabaseDependency 数据库依赖
type DatabaseDependency struct {
	ID              int64  `json:"id" form:"id" query:"id"`                                           // 数据库ID
	TableName       string `json:"table_name" form:"table_name" query:"table_name"`                   // 表名
	TableDesc       string `json:"table_desc" form:"table_desc" query:"table_desc"`                   // 表描述
	TableField      string `json:"table_field" form:"table_field" query:"table_field"`               // 表字段信息
	IconURI         string `json:"icon_uri" form:"icon_uri" query:"icon_uri"`                         // 图标URI
	PhysicalTableName string `json:"physical_table_name" form:"physical_table_name" query:"physical_table_name"` // 物理表名
	RWMode          int64  `json:"rw_mode" form:"rw_mode" query:"rw_mode"`                            // 读写模式
}

// SubWorkflowDependency 子工作流依赖
type SubWorkflowDependency struct {
	ID          int64  `json:"id" form:"id" query:"id"`                               // 工作流ID
	Name        string `json:"name" form:"name" query:"name"`                         // 工作流名称
	Description string `json:"description" form:"description" query:"description"`   // 工作流描述
	Version     string `json:"version" form:"version" query:"version"`                // 版本号
	IconURI     string `json:"icon_uri" form:"icon_uri" query:"icon_uri"`             // 图标URI
} 