package odf

type RGB struct {
	r, g, b uint8
}

const (
	//
	ODF_TEMPLATES_DIR = "templates"

	// Mimetypes and their file extension
	ODF_TEXT                  = "application/vnd.oasis.opendocument.text"
	ODF_TEXT_TEMPLATE         = "application/vnd.oasis.opendocument.text-template"
	ODF_SPREADSHEET           = "application/vnd.oasis.opendocument.spreadsheet"
	ODF_SPREADSHEET_TEMPLATE  = "application/vnd.oasis.opendocument.spreadsheet-template"
	ODF_PRESENTATION          = "application/vnd.oasis.opendocument.presentation"
	ODF_PRESENTATION_TEMPLATE = "application/vnd.oasis.opendocument.presentation-template"
	ODF_DRAWING               = "application/vnd.oasis.opendocument.graphics"
	ODF_DRAWING_TEMPLATE      = "application/vnd.oasis.opendocument.graphics-template"
	ODF_CHART                 = "application/vnd.oasis.opendocument.chart"
	ODF_CHART_TEMPLATE        = "application/vnd.oasis.opendocument.chart-template"
	ODF_IMAGE                 = "application/vnd.oasis.opendocument.image"
	ODF_IMAGE_TEMPLATE        = "application/vnd.oasis.opendocument.image-template"
	ODF_FORMULA               = "application/vnd.oasis.opendocument.formula"
	ODF_FORMULA_TEMPLATE      = "application/vnd.oasis.opendocument.formula-template"
	ODF_MASTER                = "application/vnd.oasis.opendocument.text-master"
	ODF_WEB                   = "application/vnd.oasis.opendocument.text-web"

	// Paths of standard parts
	ODF_CONTENT  = "content.xml"
	ODF_META     = "meta.xml"
	ODF_SETTINGS = "settings.xml"
	ODF_STYLES   = "styles.xml"
	ODF_MANIFEST = "META-INF/manifest.xml"
)

var (

	// File extensions and their mimetype
	ODF_EXTENSIONS = map[string]string{
		"odt": ODF_TEXT,
		"ott": ODF_TEXT_TEMPLATE,
		"ods": ODF_SPREADSHEET,
		"ots": ODF_SPREADSHEET_TEMPLATE,
		"odp": ODF_PRESENTATION,
		"otp": ODF_PRESENTATION_TEMPLATE,
		"odg": ODF_DRAWING,
		"otg": ODF_DRAWING_TEMPLATE,
		"odc": ODF_CHART,
		"otc": ODF_CHART_TEMPLATE,
		"odi": ODF_IMAGE,
		"oti": ODF_IMAGE_TEMPLATE,
		"odf": ODF_FORMULA,
		"otf": ODF_FORMULA_TEMPLATE,
		"odm": ODF_MASTER,
		"oth": ODF_WEB,
	}

	// Mimetypes and their file extension
	ODF_MIMETYPES = map[string]string{
		ODF_TEXT:                  "odt",
		ODF_TEXT_TEMPLATE:         "ott",
		ODF_SPREADSHEET:           "ods",
		ODF_SPREADSHEET_TEMPLATE:  "ots",
		ODF_PRESENTATION:          "odp",
		ODF_PRESENTATION_TEMPLATE: "otp",
		ODF_DRAWING:               "odg",
		ODF_DRAWING_TEMPLATE:      "otg",
		ODF_CHART:                 "odc",
		ODF_CHART_TEMPLATE:        "otc",
		ODF_IMAGE:                 "odi",
		ODF_IMAGE_TEMPLATE:        "oti",
		ODF_FORMULA:               "odf",
		ODF_FORMULA_TEMPLATE:      "otf",
		ODF_MASTER:                "odm",
		ODF_WEB:                   "oth",
	}

	// Standard parts in the container (other are regular paths)
	ODF_PARTS = map[string]interface{}{
		"content":  nil,
		"meta":     nil,
		"settings": nil,
		"styles":   nil,
		"manifest": nil,
	}

	// Presentation classes (for layout)
	ODF_CLASSES = map[string]interface{}{
		"title":    nil,
		"outline":  nil,
		"subtitle": nil,
		"text":     nil,
		"graphic":  nil,
		"object":   nil,
		"chart":    nil,
		"table":    nil,
		"orgchart": nil,
		"page":     nil,
		"notes":    nil,
		"handout":  nil,
	}

	ODF_PROPERTIES = map[string]interface{}{
		"chart:angle-offset":                      nil,
		"chart:auto-position":                     nil,
		"chart:auto-size":                         nil,
		"chart:axis-label-position":               nil,
		"chart:axis-position":                     nil,
		"chart:connect-bars":                      nil,
		"chart:data-label-number":                 nil,
		"chart:data-label-symbol":                 nil,
		"chart:data-label-text":                   nil,
		"chart:deep":                              nil,
		"chart:display-label":                     nil,
		"chart:error-category":                    nil,
		"chart:error-lower-indicator":             nil,
		"chart:error-lower-limit":                 nil,
		"chart:error-lower-range":                 nil,
		"chart:error-margin":                      nil,
		"chart:error-percentage":                  nil,
		"chart:error-upper-indicator":             nil,
		"chart:error-upper-limit":                 nil,
		"chart:error-upper-range":                 nil,
		"chart:gap-width":                         nil,
		"chart:group-bars-per-axis":               nil,
		"chart:hole-size":                         nil,
		"chart:include-hidden-cells":              nil,
		"chart:interpolation":                     nil,
		"chart:interval-major":                    nil,
		"chart:interval-minor-divisor":            nil,
		"chart:japanese-candle-stick":             nil,
		"chart:label-arrangement":                 nil,
		"chart:label-position":                    nil,
		"chart:label-position-negative":           nil,
		"chart:lines":                             nil,
		"chart:link-data-style-to-source":         nil,
		"chart:logarithmic":                       nil,
		"chart:maximum":                           nil,
		"chart:mean-value":                        nil,
		"chart:minimum":                           nil,
		"chart:origin":                            nil,
		"chart:overlap":                           nil,
		"chart:percentage":                        nil,
		"chart:pie-offset":                        nil,
		"chart:regression-type":                   nil,
		"chart:reverse-direction":                 nil,
		"chart:right-angled-axes":                 nil,
		"chart:scale-text":                        nil,
		"chart:series-source":                     nil,
		"chart:solid-type":                        nil,
		"chart:sort-by-x-values":                  nil,
		"chart:spline-order":                      nil,
		"chart:spline-resolution":                 nil,
		"chart:stacked":                           nil,
		"chart:symbol-height":                     nil,
		"chart:symbol-name":                       nil,
		"chart:symbol-type":                       nil,
		"chart:symbol-width":                      nil,
		"chart:text-overlap":                      nil,
		"chart:three-dimensional":                 nil,
		"chart:tick-mark-position":                nil,
		"chart:tick-marks-major-inner":            nil,
		"chart:tick-marks-major-outer":            nil,
		"chart:tick-marks-minor-inner":            nil,
		"chart:tick-marks-minor-outer":            nil,
		"chart:treat-empty-cells":                 nil,
		"chart:vertical":                          nil,
		"chart:visible":                           nil,
		"dr3d:ambient-color":                      nil,
		"dr3d:back-scale":                         nil,
		"dr3d:backface-culling":                   nil,
		"dr3d:close-back":                         nil,
		"dr3d:close-front":                        nil,
		"dr3d:depth":                              nil,
		"dr3d:diffuse-color":                      nil,
		"dr3d:edge-rounding":                      nil,
		"dr3d:edge-rounding-mode":                 nil,
		"dr3d:emissive-color":                     nil,
		"dr3d:end-angle":                          nil,
		"dr3d:horizontal-segments":                nil,
		"dr3d:lighting-mode":                      nil,
		"dr3d:normals-direction":                  nil,
		"dr3d:normals-kind":                       nil,
		"dr3d:shadow":                             nil,
		"dr3d:shininess":                          nil,
		"dr3d:specular-color":                     nil,
		"dr3d:texture-filter":                     nil,
		"dr3d:texture-generation-mode-x":          nil,
		"dr3d:texture-generation-mode-y":          nil,
		"dr3d:texture-kind":                       nil,
		"dr3d:texture-mode":                       nil,
		"dr3d:vertical-segments":                  nil,
		"draw:auto-grow-height":                   nil,
		"draw:auto-grow-width":                    nil,
		"draw:background-size":                    nil,
		"draw:blue":                               nil,
		"draw:caption-angle":                      nil,
		"draw:caption-angle-type":                 nil,
		"draw:caption-escape":                     nil,
		"draw:caption-escape-direction":           nil,
		"draw:caption-fit-line-length":            nil,
		"draw:caption-gap":                        nil,
		"draw:caption-line-length":                nil,
		"draw:caption-type":                       nil,
		"draw:color-inversion":                    nil,
		"draw:color-mode":                         nil,
		"draw:contrast":                           nil,
		"draw:decimal-places":                     nil,
		"draw:draw-aspect":                        nil,
		"draw:end-guide":                          nil,
		"draw:end-line-spacing-horizontal":        nil,
		"draw:end-line-spacing-vertical":          nil,
		"draw:fill":                               nil,
		"draw:fill-color":                         nil,
		"draw:fill-gradient-name":                 nil,
		"draw:fill-hatch-name":                    nil,
		"draw:fill-hatch-solid":                   nil,
		"draw:fill-image-height":                  nil,
		"draw:fill-image-name":                    nil,
		"draw:fill-image-ref-point":               nil,
		"draw:fill-image-ref-point-x":             nil,
		"draw:fill-image-ref-point-y":             nil,
		"draw:fill-image-width":                   nil,
		"draw:fit-to-contour":                     nil,
		"draw:fit-to-size":                        nil,
		"draw:frame-display-border":               nil,
		"draw:frame-display-scrollbar":            nil,
		"draw:frame-margin-horizontal":            nil,
		"draw:frame-margin-vertical":              nil,
		"draw:gamma":                              nil,
		"draw:gradient-step-count":                nil,
		"draw:green":                              nil,
		"draw:guide-distance":                     nil,
		"draw:guide-overhang":                     nil,
		"draw:image-opacity":                      nil,
		"draw:line-distance":                      nil,
		"draw:luminance":                          nil,
		"draw:marker-end":                         nil,
		"draw:marker-end-center":                  nil,
		"draw:marker-end-width":                   nil,
		"draw:marker-start":                       nil,
		"draw:marker-start-center":                nil,
		"draw:marker-start-width":                 nil,
		"draw:measure-align":                      nil,
		"draw:measure-vertical-align":             nil,
		"draw:ole-draw-aspect":                    nil,
		"draw:opacity":                            nil,
		"draw:opacity-name":                       nil,
		"draw:parallel":                           nil,
		"draw:placing":                            nil,
		"draw:red":                                nil,
		"draw:secondary-fill-color":               nil,
		"draw:shadow":                             nil,
		"draw:shadow-color":                       nil,
		"draw:shadow-offset-x":                    nil,
		"draw:shadow-offset-y":                    nil,
		"draw:shadow-opacity":                     nil,
		"draw:show-unit":                          nil,
		"draw:start-guide":                        nil,
		"draw:start-line-spacing-horizontal":      nil,
		"draw:start-line-spacing-vertical":        nil,
		"draw:stroke":                             nil,
		"draw:stroke-dash":                        nil,
		"draw:stroke-dash-names":                  nil,
		"draw:stroke-linejoin":                    nil,
		"draw:symbol-color":                       nil,
		"draw:textarea-horizontal-align":          nil,
		"draw:textarea-vertical-align":            nil,
		"draw:tile-repeat-offset":                 nil,
		"draw:unit":                               nil,
		"draw:visible-area-height":                nil,
		"draw:visible-area-left":                  nil,
		"draw:visible-area-top":                   nil,
		"draw:visible-area-width":                 nil,
		"draw:wrap-influence-on-position":         nil,
		"fo:background-color":                     nil,
		"fo:border":                               nil,
		"fo:border-bottom":                        nil,
		"fo:border-left":                          nil,
		"fo:border-right":                         nil,
		"fo:border-top":                           nil,
		"fo:break-after":                          nil,
		"fo:break-before":                         nil,
		"fo:clip":                                 nil,
		"fo:color":                                nil,
		"fo:country":                              nil,
		"fo:font-family":                          nil,
		"fo:font-size":                            nil,
		"fo:font-style":                           nil,
		"fo:font-variant":                         nil,
		"fo:font-weight":                          nil,
		"fo:height":                               nil,
		"fo:hyphenate":                            nil,
		"fo:hyphenation-keep":                     nil,
		"fo:hyphenation-ladder-count":             nil,
		"fo:hyphenation-push-char-count":          nil,
		"fo:hyphenation-remain-char-count":        nil,
		"fo:keep-together":                        nil,
		"fo:keep-with-next":                       nil,
		"fo:language":                             nil,
		"fo:letter-spacing":                       nil,
		"fo:line-height":                          nil,
		"fo:margin":                               nil,
		"fo:margin-bottom":                        nil,
		"fo:margin-left":                          nil,
		"fo:margin-right":                         nil,
		"fo:margin-top":                           nil,
		"fo:max-height":                           nil,
		"fo:max-width":                            nil,
		"fo:min-height":                           nil,
		"fo:min-width":                            nil,
		"fo:orphans":                              nil,
		"fo:padding":                              nil,
		"fo:padding-bottom":                       nil,
		"fo:padding-left":                         nil,
		"fo:padding-right":                        nil,
		"fo:padding-top":                          nil,
		"fo:page-height":                          nil,
		"fo:page-width":                           nil,
		"fo:script":                               nil,
		"fo:text-align":                           nil,
		"fo:text-align-last":                      nil,
		"fo:text-indent":                          nil,
		"fo:text-shadow":                          nil,
		"fo:text-transform":                       nil,
		"fo:widows":                               nil,
		"fo:width":                                nil,
		"fo:wrap-option":                          nil,
		"presentation:background-objects-visible": nil,
		"presentation:background-visible":         nil,
		"presentation:display-date-time":          nil,
		"presentation:display-footer":             nil,
		"presentation:display-header":             nil,
		"presentation:display-page-number":        nil,
		"presentation:duration":                   nil,
		"presentation:transition-speed":           nil,
		"presentation:transition-style":           nil,
		"presentation:transition-type":            nil,
		"presentation:visibility":                 nil,
		"smil:direction":                          nil,
		"smil:fadeColor":                          nil,
		"smil:subtype":                            nil,
		"smil:type":                               nil,
		"style:auto-text-indent":                  nil,
		"style:background-transparency":           nil,
		"style:border-line-width":                 nil,
		"style:border-line-width-bottom":          nil,
		"style:border-line-width-left":            nil,
		"style:border-line-width-right":           nil,
		"style:border-line-width-top":             nil,
		"style:cell-protect":                      nil,
		"style:column-width":                      nil,
		"style:country-asian":                     nil,
		"style:country-complex":                   nil,
		"style:decimal-places":                    nil,
		"style:diagonal-bl-tr":                    nil,
		"style:diagonal-bl-tr-widths":             nil,
		"style:diagonal-tl-br":                    nil,
		"style:diagonal-tl-br-widths":             nil,
		"style:direction":                         nil,
		"style:dynamic-spacing":                   nil,
		"style:editable":                          nil,
		"style:first-page-number":                 nil,
		"style:flow-with-text":                    nil,
		"style:font-charset":                      nil,
		"style:font-charset-asian":                nil,
		"style:font-charset-complex":              nil,
		"style:font-family-asian":                 nil,
		"style:font-family-complex":               nil,
		"style:font-family-generic":               nil,
		"style:font-family-generic-asian":         nil,
		"style:font-family-generic-complex":       nil,
		"style:font-independent-line-spacing":     nil,
		"style:font-name":                         nil,
		"style:font-name-asian":                   nil,
		"style:font-name-complex":                 nil,
		"style:font-pitch":                        nil,
		"style:font-pitch-asian":                  nil,
		"style:font-pitch-complex":                nil,
		"style:font-relief":                       nil,
		"style:font-size-asian":                   nil,
		"style:font-size-complex":                 nil,
		"style:font-size-rel":                     nil,
		"style:font-size-rel-asian":               nil,
		"style:font-size-rel-complex":             nil,
		"style:font-style-asian":                  nil,
		"style:font-style-complex":                nil,
		"style:font-style-name":                   nil,
		"style:font-style-name-asian":             nil,
		"style:font-style-name-complex":           nil,
		"style:font-weight-asian":                 nil,
		"style:font-weight-complex":               nil,
		"style:footnote-max-height":               nil,
		"style:glyph-orientation-vertical":        nil,
		"style:horizontal-pos":                    nil,
		"style:horizontal-rel":                    nil,
		"style:join-border":                       nil,
		"style:justify-single-word":               nil,
		"style:language-asian":                    nil,
		"style:language-complex":                  nil,
		"style:layout-grid-base-height":           nil,
		"style:layout-grid-base-width":            nil,
		"style:layout-grid-color":                 nil,
		"style:layout-grid-display":               nil,
		"style:layout-grid-lines":                 nil,
		"style:layout-grid-mode":                  nil,
		"style:layout-grid-print":                 nil,
		"style:layout-grid-ruby-below":            nil,
		"style:layout-grid-ruby-height":           nil,
		"style:layout-grid-snap-to":               nil,
		"style:layout-grid-standard-mode":         nil,
		"style:letter-kerning":                    nil,
		"style:line-break":                        nil,
		"style:line-height-at-least":              nil,
		"style:line-spacing":                      nil,
		"style:may-break-between-rows":            nil,
		"style:min-row-height":                    nil,
		"style:mirror":                            nil,
		"style:num-format":                        nil,
		"style:num-letter-sync":                   nil,
		"style:num-prefix":                        nil,
		"style:num-suffix":                        nil,
		"style:number-wrapped-paragraphs":         nil,
		"style:overflow-behavior":                 nil,
		"style:page-number":                       nil,
		"style:paper-tray-name":                   nil,
		"style:print":                             nil,
		"style:print-content":                     nil,
		"style:print-orientation":                 nil,
		"style:print-page-order":                  nil,
		"style:protect":                           nil,
		"style:punctuation-wrap":                  nil,
		"style:register-true":                     nil,
		"style:register-truth-ref-style-name":     nil,
		"style:rel-column-width":                  nil,
		"style:rel-height":                        nil,
		"style:rel-width":                         nil,
		"style:repeat":                            nil,
		"style:repeat-content":                    nil,
		"style:rfc-language-tag":                  nil,
		"style:rfc-language-tag-asian":            nil,
		"style:rfc-language-tag-complex":          nil,
		"style:rotation-align":                    nil,
		"style:rotation-angle":                    nil,
		"style:row-height":                        nil,
		"style:ruby-align":                        nil,
		"style:ruby-position":                     nil,
		"style:run-through":                       nil,
		"style:scale-to":                          nil,
		"style:scale-to-pages":                    nil,
		"style:script-asian":                      nil,
		"style:script-complex":                    nil,
		"style:script-type":                       nil,
		"style:shadow":                            nil,
		"style:shrink-to-fit":                     nil,
		"style:snap-to-layout-grid":               nil,
		"style:tab-stop-distance":                 nil,
		"style:table-centering":                   nil,
		"style:text-align-source":                 nil,
		"style:text-autospace":                    nil,
		"style:text-blinking":                     nil,
		"style:text-combine":                      nil,
		"style:text-combine-end-char":             nil,
		"style:text-combine-start-char":           nil,
		"style:text-emphasize":                    nil,
		"style:text-line-through-color":           nil,
		"style:text-line-through-mode":            nil,
		"style:text-line-through-style":           nil,
		"style:text-line-through-text":            nil,
		"style:text-line-through-text-style":      nil,
		"style:text-line-through-type":            nil,
		"style:text-line-through-width":           nil,
		"style:text-outline":                      nil,
		"style:text-overline-color":               nil,
		"style:text-overline-mode":                nil,
		"style:text-overline-style":               nil,
		"style:text-overline-type":                nil,
		"style:text-overline-width":               nil,
		"style:text-position":                     nil,
		"style:text-rotation-angle":               nil,
		"style:text-rotation-scale":               nil,
		"style:text-scale":                        nil,
		"style:text-underline-color":              nil,
		"style:text-underline-mode":               nil,
		"style:text-underline-style":              nil,
		"style:text-underline-type":               nil,
		"style:text-underline-width":              nil,
		"style:use-optimal-column-width":          nil,
		"style:use-optimal-row-height":            nil,
		"style:use-window-font-color":             nil,
		"style:vertical-align":                    nil,
		"style:vertical-pos":                      nil,
		"style:vertical-rel":                      nil,
		"style:width":                             nil,
		"style:wrap":                              nil,
		"style:wrap-contour":                      nil,
		"style:wrap-contour-mode":                 nil,
		"style:wrap-dynamic-threshold":            nil,
		"style:writing-mode":                      nil,
		"style:writing-mode-automatic":            nil,
		"svg:fill-rule":                           nil,
		"svg:height":                              nil,
		"svg:stroke-color":                        nil,
		"svg:stroke-linecap":                      nil,
		"svg:stroke-opacity":                      nil,
		"svg:stroke-width":                        nil,
		"svg:width":                               nil,
		"svg:x":                                   nil,
		"svg:y":                                   nil,
		"table:align":                             nil,
		"table:border-model":                      nil,
		"table:display":                           nil,
		"text:anchor-page-number":                 nil,
		"text:anchor-type":                        nil,
		"text:animation":                          nil,
		"text:animation-delay":                    nil,
		"text:animation-direction":                nil,
		"text:animation-repeat":                   nil,
		"text:animation-start-inside":             nil,
		"text:animation-steps":                    nil,
		"text:animation-stop-inside":              nil,
		"text:condition":                          nil,
		"text:display":                            nil,
		"text:dont-balance-text-columns":          nil,
		"text:line-break":                         nil,
		"text:line-number":                        nil,
		"text:list-level-position-and-space-mode": nil,
		"text:min-label-distance":                 nil,
		"text:min-label-width":                    nil,
		"text:number-lines":                       nil,
		"text:space-before":                       nil,
	}

	// from CSS3 color map
	CSS3_COLORMAP = map[string]RGB{
		"indigo":               {75, 0, 130},
		"gold":                 {255, 215, 0},
		"firebrick":            {178, 34, 34},
		"indianred":            {205, 92, 92},
		"yellow":               {255, 255, 0},
		"darkolivegreen":       {85, 107, 47},
		"darkseagreen":         {143, 188, 143},
		"slategrey":            {112, 128, 144},
		"darkslategrey":        {47, 79, 79},
		"mediumvioletred":      {199, 21, 133},
		"mediumorchid":         {186, 85, 211},
		"chartreuse":           {127, 255, 0},
		"mediumslateblue":      {123, 104, 238},
		"black":                {0, 0, 0},
		"springgreen":          {0, 255, 127},
		"crimson":              {220, 20, 60},
		"lightsalmon":          {255, 160, 122},
		"brown":                {165, 42, 42},
		"turquoise":            {64, 224, 208},
		"olivedrab":            {107, 142, 35},
		"lightcyan":            {224, 255, 255},
		"cyan":                 {0, 255, 255},
		"silver":               {192, 192, 192},
		"skyblue":              {135, 206, 235},
		"gray":                 {128, 128, 128},
		"darkturquoise":        {0, 206, 209},
		"goldenrod":            {218, 165, 32},
		"darkgreen":            {0, 100, 0},
		"darkviolet":           {148, 0, 211},
		"darkgray":             {169, 169, 169},
		"lightpink":            {255, 182, 193},
		"teal":                 {0, 128, 128},
		"darkmagenta":          {139, 0, 139},
		"lightgoldenrodyellow": {250, 250, 210},
		"lavender":             {230, 230, 250},
		"yellowgreen":          {154, 205, 50},
		"thistle":              {216, 191, 216},
		"violet":               {238, 130, 238},
		"navy":                 {0, 0, 128},
		"dimgrey":              {105, 105, 105},
		"orchid":               {218, 112, 214},
		"blue":                 {0, 0, 255},
		"ghostwhite":           {248, 248, 255},
		"honeydew":             {240, 255, 240},
		"cornflowerblue":       {100, 149, 237},
		"darkblue":             {0, 0, 139},
		"darkkhaki":            {189, 183, 107},
		"mediumpurple":         {147, 112, 216},
		"cornsilk":             {255, 248, 220},
		"red":                  {255, 0, 0},
		"bisque":               {255, 228, 196},
		"slategray":            {112, 128, 144},
		"darkcyan":             {0, 139, 139},
		"khaki":                {240, 230, 140},
		"wheat":                {245, 222, 179},
		"deepskyblue":          {0, 191, 255},
		"darkred":              {139, 0, 0},
		"steelblue":            {70, 130, 180},
		"aliceblue":            {240, 248, 255},
		"lightslategrey":       {119, 136, 153},
		"gainsboro":            {220, 220, 220},
		"mediumturquoise":      {72, 209, 204},
		"floralwhite":          {255, 250, 240},
		"plum":                 {221, 160, 221},
		"purple":               {128, 0, 128},
		"lightgrey":            {211, 211, 211},
		"burlywood":            {222, 184, 135},
		"darksalmon":           {233, 150, 122},
		"beige":                {245, 245, 220},
		"azure":                {240, 255, 255},
		"lightsteelblue":       {176, 196, 222},
		"oldlace":              {253, 245, 230},
		"greenyellow":          {173, 255, 47},
		"royalblue":            {65, 105, 225},
		"lightseagreen":        {32, 178, 170},
		"sienna":               {160, 82, 45},
		"lightcoral":           {240, 128, 128},
		"orangered":            {255, 69, 0},
		"navajowhite":          {255, 222, 173},
		"lime":                 {0, 255, 0},
		"palegreen":            {152, 251, 152},
		"mistyrose":            {255, 228, 225},
		"seashell":             {255, 245, 238},
		"mediumspringgreen":    {0, 250, 154},
		"fuchsia":              {255, 0, 255},
		"papayawhip":           {255, 239, 213},
		"blanchedalmond":       {255, 235, 205},
		"peru":                 {205, 133, 63},
		"aquamarine":           {127, 255, 212},
		"white":                {255, 255, 255},
		"darkslategray":        {47, 79, 79},
		"lightgray":            {211, 211, 211},
		"ivory":                {255, 255, 240},
		"dodgerblue":           {30, 144, 255},
		"lawngreen":            {124, 252, 0},
		"chocolate":            {210, 105, 30},
		"orange":               {255, 165, 0},
		"forestgreen":          {34, 139, 34},
		"darkgrey":             {169, 169, 169},
		"olive":                {128, 128, 0},
		"mintcream":            {245, 255, 250},
		"antiquewhite":         {250, 235, 215},
		"darkorange":           {255, 140, 0},
		"cadetblue":            {95, 158, 160},
		"moccasin":             {255, 228, 181},
		"limegreen":            {50, 205, 50},
		"saddlebrown":          {139, 69, 19},
		"grey":                 {128, 128, 128},
		"darkslateblue":        {72, 61, 139},
		"lightskyblue":         {135, 206, 250},
		"deeppink":             {255, 20, 147},
		"coral":                {255, 127, 80},
		"aqua":                 {0, 255, 255},
		"darkgoldenrod":        {184, 134, 11},
		"maroon":               {128, 0, 0},
		"sandybrown":           {244, 164, 96},
		"tan":                  {210, 180, 140},
		"magenta":              {255, 0, 255},
		"rosybrown":            {188, 143, 143},
		"pink":                 {255, 192, 203},
		"lightblue":            {173, 216, 230},
		"palevioletred":        {216, 112, 147},
		"mediumseagreen":       {60, 179, 113},
		"slateblue":            {106, 90, 205},
		"dimgray":              {105, 105, 105},
		"powderblue":           {176, 224, 230},
		"seagreen":             {46, 139, 87},
		"snow":                 {255, 250, 250},
		"mediumblue":           {0, 0, 205},
		"midnightblue":         {25, 25, 112},
		"paleturquoise":        {175, 238, 238},
		"palegoldenrod":        {238, 232, 170},
		"whitesmoke":           {245, 245, 245},
		"darkorchid":           {153, 50, 204},
		"salmon":               {250, 128, 114},
		"lightslategray":       {119, 136, 153},
		"lemonchiffon":         {255, 250, 205},
		"lightgreen":           {144, 238, 144},
		"tomato":               {255, 99, 71},
		"hotpink":              {255, 105, 180},
		"lightyellow":          {255, 255, 224},
		"lavenderblush":        {255, 240, 245},
		"linen":                {250, 240, 230},
		"mediumaquamarine":     {102, 205, 170},
		"green":                {0, 128, 0},
		"blueviolet":           {138, 43, 226},
		"peachpuff":            {255, 218, 185},
	}
)
