// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package mailer_template

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"
import "braces.dev/errtrace"

func Main() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\" xmlns:v=\"urn:schemas-microsoft-com:vml\" xmlns:o=\"urn:schemas-microsoft-com:office:office\"><head><meta charset=\"utf-8\"><meta http-equiv=\"x-ua-compatible\" content=\"ie=edge\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1\"><meta name=\"x-apple-disable-message-reformatting\"><title>Acorn Email Framework</title><!--[if mso]>\n      <xml>\n        <o:OfficeDocumentSettings>\n        <o:PixelsPerInch>96</o:PixelsPerInch>\n        </o:OfficeDocumentSettings>\n      </xml>\n      <style>\n        table {border-collapse: collapse;}\n        .spacer,.divider {mso-line-height-rule: exactly;}\n        td,th,div,p,a {font-size: 16px; line-height: 25px;}\n        td,th,div,p,a,h1,h2,h3,h4,h5,h6 {font-family:\"Segoe UI\",Helvetica,Arial,sans-serif;}\n      </style>\n      <![endif]--><style type=\"text/css\">\n        @import url('https://fonts.googleapis.com/css?family=Merriweather|Open+Sans');\n\n        img {border: 0; line-height: 100%; vertical-align: middle;}\n        .col {font-size: 16px; line-height: 25px; vertical-align: top;}\n\n        @media screen {\n        .col, td, th, div, p {font-family: -apple-system,system-ui,BlinkMacSystemFont,\"Segoe UI\",\"Roboto\",\"Helvetica Neue\",Arial,sans-serif;}\n        .sans-serif {font-family: 'Open Sans', Arial, sans-serif;}\n        .serif {font-family: 'Merriweather', Georgia, serif;}\n        img {max-width: 100%;}\n        }\n\n        @media (max-width: 632px) {\n        .container {width: 100%!important;}\n        }\n\n        @media (max-width: 480px) {\n        .col {\n        display: inline-block!important;\n        line-height: 23px;\n        width: 100%!important;\n        }\n\n        .col-sm-1 {max-width: 25%;}\n        .col-sm-2 {max-width: 50%;}\n        .col-sm-3 {max-width: 75%;}\n        .col-sm-third {max-width: 33.33333%;}\n\n        .col-sm-push-1 {margin-left: 25%;}\n        .col-sm-push-2 {margin-left: 50%;}\n        .col-sm-push-3 {margin-left: 75%;}\n        .col-sm-push-third {margin-left: 33.33333%;}\n\n        .full-width-sm {display: table!important; width: 100%!important;}\n        .stack-sm-first {display: table-header-group!important;}\n        .stack-sm-last {display: table-footer-group!important;}\n        .stack-sm-top {display: table-caption!important; max-width: 100%; padding-left: 0!important;}\n\n        .toggle-content {\n        max-height: 0;\n        overflow: auto;\n        transition: max-height .4s linear;\n        -webkit-transition: max-height .4s linear;\n        }\n        .toggle-trigger:hover + .toggle-content,\n        .toggle-content:hover {max-height: 999px!important;}\n\n        .show-sm {\n        display: inherit!important;\n        font-size: inherit!important;\n        line-height: inherit!important;\n        max-height: none!important;\n        }\n        .hide-sm {display: none!important;}\n\n        .align-sm-center {\n        display: table!important;\n        float: none;\n        margin-left: auto!important;\n        margin-right: auto!important;\n        }\n        .align-sm-left {float: left;}\n        .align-sm-right {float: right;}\n\n        .text-sm-center {text-align: center!important;}\n        .text-sm-left {text-align: left!important;}\n        .text-sm-right {text-align: right!important;}\n\n        .borderless-sm {border: none!important;}\n        .nav-sm-vertical .nav-item {display: block;}\n        .nav-sm-vertical .nav-item a {display: inline-block; padding: 4px 0!important;}\n\n        .spacer {height: 0;}\n\n        .p-sm-0 {padding: 0!important;}\n        .p-sm-8 {padding: 8px!important;}\n        .p-sm-16 {padding: 16px!important;}\n        .p-sm-24 {padding: 24px!important;}\n        .pt-sm-0 {padding-top: 0!important;}\n        .pt-sm-8 {padding-top: 8px!important;}\n        .pt-sm-16 {padding-top: 16px!important;}\n        .pt-sm-24 {padding-top: 24px!important;}\n        .pr-sm-0 {padding-right: 0!important;}\n        .pr-sm-8 {padding-right: 8px!important;}\n        .pr-sm-16 {padding-right: 16px!important;}\n        .pr-sm-24 {padding-right: 24px!important;}\n        .pb-sm-0 {padding-bottom: 0!important;}\n        .pb-sm-8 {padding-bottom: 8px!important;}\n        .pb-sm-16 {padding-bottom: 16px!important;}\n        .pb-sm-24 {padding-bottom: 24px!important;}\n        .pl-sm-0 {padding-left: 0!important;}\n        .pl-sm-8 {padding-left: 8px!important;}\n        .pl-sm-16 {padding-left: 16px!important;}\n        .pl-sm-24 {padding-left: 24px!important;}\n        .px-sm-0 {padding-right: 0!important; padding-left: 0!important;}\n        .px-sm-8 {padding-right: 8px!important; padding-left: 8px!important;}\n        .px-sm-16 {padding-right: 16px!important; padding-left: 16px!important;}\n        .px-sm-24 {padding-right: 24px!important; padding-left: 24px!important;}\n        .py-sm-0 {padding-top: 0!important; padding-bottom: 0!important;}\n        .py-sm-8 {padding-top: 8px!important; padding-bottom: 8px!important;}\n        .py-sm-16 {padding-top: 16px!important; padding-bottom: 16px!important;}\n        .py-sm-24 {padding-top: 24px!important; padding-bottom: 24px!important;}\n        }\n      </style></head><body style=\"margin:0;padding:0;width:100%;word-break:break-word;-webkit-font-smoothing:antialiased;\"><div lang=\"en\" style=\"display:none;\"><!-- Add your preheader text here --></div><table lang=\"en\" bgcolor=\"#EEEEEE\" cellpadding=\"16\" cellspacing=\"0\" role=\"presentation\" width=\"100%\"><tr><td align=\"center\"><table class=\"container\" bgcolor=\"#FFFFFF\" cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\" width=\"600\"><tr><td align=\"left\"><!-- ADD ROWS HERE --></td></tr></table></td></tr></table></body></html>")
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ_7745c5c3_Err)
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return errtrace.Wrap(templ_7745c5c3_Err)
	})
}
