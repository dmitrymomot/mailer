// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/dmitrymomot/mailer/template/utils"
import "fmt"
import "braces.dev/errtrace"

func ButtonFilled(align utils.Alignment, color utils.Color, buttonText, buttonLink string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<table cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\" style=\"margin: 0 auto;\" align=\"")
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ_7745c5c3_Err)
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(align.String()))
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ_7745c5c3_Err)
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><tr><th bgcolor=\"")
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ_7745c5c3_Err)
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(color.String()))
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ_7745c5c3_Err)
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" style=\"border-radius: 3px; mso-padding-alt: 6px 42px 12px;\"><a href=\"")
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ_7745c5c3_Err)
		}
		var templ_7745c5c3_Var2 templ.SafeURL = templ.URL(buttonLink)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var2)))
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ_7745c5c3_Err)
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" style=\"color: #FFFFFF; display: inline-block; font-size: 13px; line-height: 100%; padding: 12px 42px; text-decoration: none;\">")
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ_7745c5c3_Err)
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(buttonText)
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ.Error{Err: templ_7745c5c3_Err, FileName: `template/components/button.templ`, Line: 10, Col: 17})
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ_7745c5c3_Err)
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></th></tr></table>")
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ_7745c5c3_Err)
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return errtrace.Wrap(templ_7745c5c3_Err)
	})
}

func ButtonOutlined(align utils.Alignment, color utils.Color, buttonText, buttonLink string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var4 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var4 == nil {
			templ_7745c5c3_Var4 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<table cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\" style=\"margin: 0 auto;\" align=\"")
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ_7745c5c3_Err)
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(align.String()))
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ_7745c5c3_Err)
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><tr><th")
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ_7745c5c3_Err)
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, buttonOutlinedAttrs(color))
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ_7745c5c3_Err)
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("><a href=\"")
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ_7745c5c3_Err)
		}
		var templ_7745c5c3_Var5 templ.SafeURL = templ.URL(buttonLink)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var5)))
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ_7745c5c3_Err)
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ_7745c5c3_Err)
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, buttonOutlinedLinkAttrs(color))
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ_7745c5c3_Err)
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(">")
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ_7745c5c3_Err)
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(buttonText)
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ.Error{Err: templ_7745c5c3_Err, FileName: `template/components/button.templ`, Line: 22, Col: 17})
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ_7745c5c3_Err)
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></th></tr></table>")
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ_7745c5c3_Err)
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return errtrace.Wrap(templ_7745c5c3_Err)
	})
}

func buttonOutlinedAttrs(color utils.Color) templ.Attributes {
	return templ.Attributes{"style": fmt.Sprintf("border: 2px solid %s; border-radius: 3px; mso-padding-alt: 6px 42px 12px;", color.String())}
}

func buttonOutlinedLinkAttrs(color utils.Color) templ.Attributes {
	return templ.Attributes{"style": fmt.Sprintf("color: %s; display: inline-block; font-size: 13px; line-height: 100%%; padding: 12px 42px; text-decoration: none;", color.String())}
}
