// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package mailer_template

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "fmt"
import "braces.dev/errtrace"

type Coupon struct {
	Persentage     int
	Code           string
	ActionBtnURL   string
	ActionBtnLabel string
}

func CouponDashed(coupon Coupon) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<table cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\" width=\"100%\"><tr><td style=\"padding: 0 24px;\"><table cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\" width=\"100%\"><tr><td class=\"col\" align=\"center\" width=\"100%\" style=\"padding: 0 8px;\"><table cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\" width=\"100%\"><tr><td class=\"spacer py-sm-16\" height=\"32\"></td></tr><tr><td class=\"px-sm-8\" align=\"center\" width=\"100%\" style=\"padding: 32px; border: 4px dashed #CCCCCC; color: #000000;\"><table align=\"center\" cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\" style=\"margin: 0 auto;\"><tr><th style=\"font-size: 96px; line-height: 100%; word-break: break-all;\">")
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ_7745c5c3_Err)
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d", coupon.Persentage))
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ.Error{Err: templ_7745c5c3_Err, FileName: `template/coupon.templ`, Line: 26, Col: 121})
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ_7745c5c3_Err)
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</th><th style=\"vertical-align: middle;\"><div style=\"font-size: 48px; line-height: 48px;\">%</div><div style=\"font-size: 32px; line-height: 32px; mso-line-height-rule: exactly; mso-text-raise: 2px;\">OFF</div></th></tr></table><div style=\"color: #999999;\">With coupon <span style=\"border: 1px solid #EA4B35; border-radius: 3px; color: #EA4B35; display: inline-block; font-size: 90%; padding: 1px 5px;\">")
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ_7745c5c3_Err)
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(coupon.Code)
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ.Error{Err: templ_7745c5c3_Err, FileName: `template/coupon.templ`, Line: 34, Col: 170})
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ_7745c5c3_Err)
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></div><div class=\"spacer py-sm-16\" style=\"line-height: 32px;\">\u200c</div><table cellpadding=\"0\" cellspacing=\"0\" role=\"presentation\"><tr><th bgcolor=\"#000000\" style=\"mso-padding-alt: 6px 32px 12px;\"><a href=\"")
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ_7745c5c3_Err)
		}
		var templ_7745c5c3_Var4 templ.SafeURL = templ.URL(coupon.ActionBtnURL)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var4)))
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ_7745c5c3_Err)
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" style=\"color: #FFFFFF; display: inline-block; font-size: 13px; line-height: 100%; padding: 12px 32px; text-decoration: none;\">")
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ_7745c5c3_Err)
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(coupon.ActionBtnLabel)
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ.Error{Err: templ_7745c5c3_Err, FileName: `template/coupon.templ`, Line: 40, Col: 205})
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ_7745c5c3_Err)
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" →</a></th></tr></table><div class=\"spacer py-sm-8\" style=\"line-height: 16px;\">\u200c</div></td></tr><tr><td class=\"spacer py-sm-16\" height=\"32\"></td></tr></table></td></tr></table></td></tr></table>")
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ_7745c5c3_Err)
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return errtrace.Wrap(templ_7745c5c3_Err)
	})
}