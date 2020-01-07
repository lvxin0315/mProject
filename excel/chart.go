package excel

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

func ChartStructToString(ChartStruct interface{}) string {
	chartBytes, err := json.Marshal(&ChartStruct)
	if err != nil {
		logrus.Error("ChartStructToString:", err)
		return ""
	}
	return string(chartBytes)
}

func InitFormatChart() FormatChart {
	format := FormatChart{
		Dimension: FormatChartDimension{
			Width:  480,
			Height: 290,
		},
		Format: FormatPicture{
			FPrintsWithSheet: true,
			FLocksWithSheet:  false,
			NoChangeAspect:   false,
			OffsetX:          0,
			OffsetY:          0,
			XScale:           1.0,
			YScale:           1.0,
		},
		Legend: FormatChartLegend{
			Position:      "bottom",
			ShowLegendKey: false,
		},
		Title: FormatChartTitle{
			Name: " ",
		},
		ShowBlanksAs: "gap",
	}
	return format
}

type FormatChart struct {
	Type      string               `json:"type"`
	Series    []FormatChartSeries  `json:"series"`
	Format    FormatPicture        `json:"format"`
	Dimension FormatChartDimension `json:"dimension"`
	Legend    FormatChartLegend    `json:"legend"`
	Title     FormatChartTitle     `json:"title"`
	XAxis     FormatChartAxis      `json:"x_axis"`
	YAxis     FormatChartAxis      `json:"y_axis"`
	Chartarea struct {
		Border struct {
			None bool `json:"none"`
		} `json:"border"`
		Fill struct {
			Color string `json:"color"`
		} `json:"fill"`
		Pattern struct {
			Pattern string `json:"pattern"`
			FgColor string `json:"fg_color"`
			BgColor string `json:"bg_color"`
		} `json:"pattern"`
	} `json:"chartarea"`
	Plotarea struct {
		ShowBubbleSize  bool `json:"show_bubble_size"`
		ShowCatName     bool `json:"show_cat_name"`
		ShowLeaderLines bool `json:"show_leader_lines"`
		ShowPercent     bool `json:"show_percent"`
		ShowSerName     bool `json:"show_series_name"`
		ShowVal         bool `json:"show_val"`
		Gradient        struct {
			Colors []string `json:"colors"`
		} `json:"gradient"`
		Border struct {
			Color    string `json:"color"`
			Width    int    `json:"width"`
			DashType string `json:"dash_type"`
		} `json:"border"`
		Fill struct {
			Color string `json:"color"`
		} `json:"fill"`
		Layout FormatLayout `json:"layout"`
	} `json:"plotarea"`
	ShowBlanksAs   string `json:"show_blanks_as"`
	ShowHiddenData bool   `json:"show_hidden_data"`
	SetRotation    int    `json:"set_rotation"`
	SetHoleSize    int    `json:"set_hole_size"`
}

// FormatChartLegend directly maps the Format settings of the chart legend.
type FormatChartLegend struct {
	None            bool         `json:"none"`
	DeleteSeries    []int        `json:"delete_series"`
	Font            FormatFont   `json:"font"`
	Layout          FormatLayout `json:"layout"`
	Position        string       `json:"position"`
	ShowLegendEntry bool         `json:"show_legend_entry"`
	ShowLegendKey   bool         `json:"show_legend_key"`
}

// FormatChartSeries directly maps the Format settings of the chart series.
type FormatChartSeries struct {
	Name       string `json:"name"`
	Categories string `json:"categories"`
	Values     string `json:"values"`
	Line       struct {
		None  bool   `json:"none"`
		Color string `json:"color"`
	} `json:"line"`
	Marker struct {
		Type   string  `json:"type"`
		Size   int     `json:"size"`
		Width  float64 `json:"width"`
		Border struct {
			Color string `json:"color"`
			None  bool   `json:"none"`
		} `json:"border"`
		Fill struct {
			Color string `json:"color"`
			None  bool   `json:"none"`
		} `json:"fill"`
	} `json:"marker"`
}

// FormatChartTitle directly maps the Format settings of the chart title.
type FormatChartTitle struct {
	None    bool         `json:"none"`
	Name    string       `json:"name"`
	Overlay bool         `json:"overlay"`
	Layout  FormatLayout `json:"layout"`
}

// FormatLayout directly maps the Format settings of the element layout.
type FormatLayout struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

// FormatFont directly maps the styles settings of the fonts.
type FormatFont struct {
	Bold      bool    `json:"bold"`
	Italic    bool    `json:"italic"`
	Underline string  `json:"underline"`
	Family    string  `json:"family"`
	Size      float64 `json:"size"`
	Strike    bool    `json:"strike"`
	Color     string  `json:"color"`
}

// FormatPicture directly maps the Format settings of the picture.
type FormatPicture struct {
	FPrintsWithSheet bool    `json:"print_obj"`
	FLocksWithSheet  bool    `json:"locked"`
	NoChangeAspect   bool    `json:"lock_aspect_ratio"`
	OffsetX          int     `json:"x_offset"`
	OffsetY          int     `json:"y_offset"`
	XScale           float64 `json:"x_scale"`
	YScale           float64 `json:"y_scale"`
	Hyperlink        string  `json:"hyperlink"`
	HyperlinkType    string  `json:"hyperlink_type"`
	Positioning      string  `json:"positioning"`
}

// FormatShape directly maps the Format settings of the shape.
type FormatShape struct {
	Type      string                 `json:"type"`
	Width     int                    `json:"width"`
	Height    int                    `json:"height"`
	Format    FormatPicture          `json:"format"`
	Color     FormatShapeColor       `json:"color"`
	Paragraph []FormatShapeParagraph `json:"paragraph"`
}

// FormatShapeParagraph directly maps the Format settings of the paragraph in
// the shape.
type FormatShapeParagraph struct {
	Font FormatFont `json:"font"`
	Text string     `json:"text"`
}

// FormatShapeColor directly maps the color settings of the shape.
type FormatShapeColor struct {
	Line   string `json:"line"`
	Fill   string `json:"fill"`
	Effect string `json:"effect"`
}

type FormatChartDimension struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

// FormatChartAxis directly maps the Format settings of the chart axis.
type FormatChartAxis struct {
	Crossing            string  `json:"crossing"`
	MajorGridlines      bool    `json:"major_grid_lines"`
	MajorTickMark       string  `json:"major_tick_mark"`
	MinorTickMark       string  `json:"minor_tick_mark"`
	MinorUnitType       string  `json:"minor_unit_type"`
	MajorUnit           int     `json:"major_unit"`
	MajorUnitType       string  `json:"major_unit_type"`
	DisplayUnits        string  `json:"display_units"`
	DisplayUnitsVisible bool    `json:"display_units_visible"`
	DateAxis            bool    `json:"date_axis"`
	ReverseOrder        bool    `json:"reverse_order"`
	Maximum             float64 `json:"maximum"`
	Minimum             float64 `json:"minimum"`
	NumFormat           string  `json:"num_Format"`
	NumFont             struct {
		Color     string `json:"color"`
		Bold      bool   `json:"bold"`
		Italic    bool   `json:"italic"`
		Underline bool   `json:"underline"`
	} `json:"num_font"`
	NameLayout FormatLayout `json:"name_layout"`
}
