//
// Tencent is pleased to support the open source community by making trpc-agent-go available.
//
// Copyright (C) 2025 Tencent.  All rights reserved.
//
// trpc-agent-go is licensed under the Apache License Version 2.0.
//
//

package codeinterpreter

// ChartType represents the kind of chart returned by the server.
type ChartType string

const (
	// ChartTypeLine represents a line chart.
	ChartTypeLine ChartType = "line"
	// ChartTypeScatter represents a scatter plot.
	ChartTypeScatter ChartType = "scatter"
	// ChartTypeBar represents a bar chart.
	ChartTypeBar ChartType = "bar"
	// ChartTypePie represents a pie chart.
	ChartTypePie ChartType = "pie"
	// ChartTypeBoxAndWhisker represents a box and whisker plot.
	ChartTypeBoxAndWhisker ChartType = "box_and_whisker"
	// ChartTypeSuperChart represents a super chart.
	ChartTypeSuperChart ChartType = "superchart"
	// ChartTypeUnknown represents an unknown chart type.
	ChartTypeUnknown ChartType = "unknown"
)

// ScaleType represents an axis scale type (linear, log, etc.)
type ScaleType string

const (
	// ScaleTypeLinear represents a linear scale.
	ScaleTypeLinear ScaleType = "linear"
	// ScaleTypeDatetime represents a datetime scale.
	ScaleTypeDatetime ScaleType = "datetime"
	// ScaleTypeCategorical represents a categorical scale.
	ScaleTypeCategorical ScaleType = "categorical"
	// ScaleTypeLog represents a log scale.
	ScaleTypeLog ScaleType = "log"
	// ScaleTypeSymlog represents a symlog scale.
	ScaleTypeSymlog ScaleType = "symlog"
	// ScaleTypeLogit represents a logit scale.
	ScaleTypeLogit ScaleType = "logit"
	// ScaleTypeFunction represents a function scale.
	ScaleTypeFunction ScaleType = "function"
	// ScaleTypeFunctionLog represents a function log scale.
	ScaleTypeFunctionLog ScaleType = "functionlog"
	// ScaleTypeAsinh represents an asinh scale.
	ScaleTypeAsinh ScaleType = "asinh"
	// ScaleTypeUnknown represents an unknown scale.
	ScaleTypeUnknown ScaleType = "unknown"
)

// Chart is the common interface implemented by all concrete chart types.
// Use a type switch on the concrete types to inspect specialized fields, e.g.
//
//	switch c := result.Chart.(type) {
//	case *LineChart: ...
//	case *BarChart: ...
//	}
type Chart interface {
	ChartType() ChartType
	ChartTitle() string
	// ToDict returns the raw JSON representation of the chart.
	ToJSON() map[string]any
}

// BaseChart contains the fields shared by every chart type.
type BaseChart struct {
	Type     ChartType      `json:"type"`
	Title    string         `json:"title"`
	Elements []any          `json:"elements"`
	raw      map[string]any `json:"-"`
}

// ChartType returns the type of the chart.
func (c *BaseChart) ChartType() ChartType {
	_ = "STUB: not implemented"

	// ChartTitle returns the title of the chart.
	return *new(ChartType)
}

func (c *BaseChart) ChartTitle() string {
	_ = "STUB: not implemented"

	// ToJSON returns the raw JSON representation of the chart.
	return ""
}

func (c *BaseChart) ToJSON() map[string]any {
	_ = "STUB: not implemented"

	// Chart2D is the base for charts that live on a 2D plane.
	return nil
}

type Chart2D struct {
	BaseChart
	XLabel string `json:"x_label,omitempty"`
	YLabel string `json:"y_label,omitempty"`
	XUnit  string `json:"x_unit,omitempty"`
	YUnit  string `json:"y_unit,omitempty"`
}

// PointData is one series in a point based chart (line/scatter).
type PointData struct {
	Label  string   `json:"label"`
	Points [][2]any `json:"points"`
}

// PointChart is the base for line/scatter.
type PointChart struct {
	Chart2D
	XTicks      []any       `json:"x_ticks"`
	XTickLabels []string    `json:"x_tick_labels"`
	XScale      ScaleType   `json:"x_scale"`
	YTicks      []any       `json:"y_ticks"`
	YTickLabels []string    `json:"y_tick_labels"`
	YScale      ScaleType   `json:"y_scale"`
	Points      []PointData `json:"-"`
}

// LineChart represents a line chart.
type LineChart struct {
	PointChart
}

// ScatterChart represents a scatter chart.
type ScatterChart struct {
	PointChart
}

// BarData represents a single bar in a bar chart.
type BarData struct {
	Label string `json:"label"`
	Value string `json:"value"`
	Group string `json:"group"`
}

// BarChart represents a bar chart.
type BarChart struct {
	Chart2D
	Bars []BarData `json:"-"`
}

// PieData represents a slice of a pie chart.
type PieData struct {
	Label  string  `json:"label"`
	Angle  float64 `json:"angle"`
	Radius float64 `json:"radius"`
}

// PieChart represents a pie chart.
type PieChart struct {
	BaseChart
	Slices []PieData `json:"-"`
}

// BoxAndWhiskerData represents one box-and-whisker series.
type BoxAndWhiskerData struct {
	Label         string    `json:"label"`
	Min           float64   `json:"min"`
	FirstQuartile float64   `json:"first_quartile"`
	Median        float64   `json:"median"`
	ThirdQuartile float64   `json:"third_quartile"`
	Max           float64   `json:"max"`
	Outliers      []float64 `json:"outliers"`
}

// BoxAndWhiskerChart represents a box-and-whisker chart.
type BoxAndWhiskerChart struct {
	Chart2D
	Boxes []BoxAndWhiskerData `json:"-"`
}

// SuperChart is a composite chart containing multiple sub-charts.
type SuperChart struct {
	BaseChart
	Charts []Chart `json:"-"`
}

// UnknownChart is used when the server returns a chart type that the SDK does
// not yet understand; all data is still accessible through ToDict().
type UnknownChart struct {
	BaseChart
}

// deserializeChart converts the raw JSON payload coming from the server into
// the matching Chart implementation.
func deserializeChart(data map[string]any) Chart { _ = "STUB: not implemented"; return *new(Chart) }

func buildBarChart(base BaseChart, data map[string]any) *BarChart {
	_ = "STUB: not implemented"
	return nil
}

func buildPointChart(base BaseChart, data map[string]any) PointChart {
	_ = "STUB: not implemented"
	return *new(PointChart)
}

func getString(m map[string]any, key string) string { _ = "STUB: not implemented"; return "" }

func getFloat(m map[string]any, key string) float64 { _ = "STUB: not implemented"; return 0 }

func getFloatSlice(m map[string]any, key string) []float64 { _ = "STUB: not implemented"; return nil }

func getStringSlice(m map[string]any, key string) []string { _ = "STUB: not implemented"; return nil }

func getInterfaceSlice(m map[string]any, key string) []any { _ = "STUB: not implemented"; return nil }
