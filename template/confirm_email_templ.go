// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package template

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "github.com/dmitrymomot/mailer/template/components"
import "github.com/dmitrymomot/mailer/template/utils"
import "time"
import "fmt"
import "braces.dev/errtrace"

type ConfirmEmailPayload struct {
	Lang         string
	Subject      string
	EmailPreview string
	LogoURL      string
	CompanyName  string
	CompanyURL   string
	Greeting     string
	SupportEmail string
	SupportURL   string
	ButtonText   string
	ButtonURL    string
	ExpiresIn    string
	Footer       string
}

func (p ConfirmEmailPayload) GetSupportMailToLink() templ.SafeURL {
	return templ.SafeURL("mailto:" + p.SupportEmail)
}

func ConfirmEmail(payload ConfirmEmailPayload) templ.Component {
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
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			templ_7745c5c3_Var3 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
				templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
				if !templ_7745c5c3_IsBuffer {
					templ_7745c5c3_Buffer = templ.GetBuffer()
					defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
				}
				templ_7745c5c3_Err = components.Spacer().Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return errtrace.Wrap(templ_7745c5c3_Err)
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
				if templ_7745c5c3_Err != nil {
					return errtrace.Wrap(templ_7745c5c3_Err)
				}
				templ_7745c5c3_Err = components.LogoImage(payload.LogoURL, payload.CompanyName, payload.CompanyURL, "200px").Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return errtrace.Wrap(templ_7745c5c3_Err)
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
				if templ_7745c5c3_Err != nil {
					return errtrace.Wrap(templ_7745c5c3_Err)
				}
				templ_7745c5c3_Err = components.Spacer().Render(ctx, templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return errtrace.Wrap(templ_7745c5c3_Err)
				}
				if !templ_7745c5c3_IsBuffer {
					_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
				}
				return errtrace.Wrap(templ_7745c5c3_Err)
			})
			templ_7745c5c3_Err = components.Wrapper().Render(templ.WithChildren(ctx, templ_7745c5c3_Var3), templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return errtrace.Wrap(templ_7745c5c3_Err)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
			if templ_7745c5c3_Err != nil {
				return errtrace.Wrap(templ_7745c5c3_Err)
			}
			templ_7745c5c3_Var4 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
				templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
				if !templ_7745c5c3_IsBuffer {
					templ_7745c5c3_Buffer = templ.GetBuffer()
					defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
				}
				templ_7745c5c3_Var5 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
					templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
					if !templ_7745c5c3_IsBuffer {
						templ_7745c5c3_Buffer = templ.GetBuffer()
						defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
					}
					templ_7745c5c3_Var6 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
						templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
						if !templ_7745c5c3_IsBuffer {
							templ_7745c5c3_Buffer = templ.GetBuffer()
							defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
						}
						templ_7745c5c3_Var7 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
							templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
							if !templ_7745c5c3_IsBuffer {
								templ_7745c5c3_Buffer = templ.GetBuffer()
								defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
							}
							templ_7745c5c3_Err = components.H2(payload.Greeting).Render(ctx, templ_7745c5c3_Buffer)
							if templ_7745c5c3_Err != nil {
								return errtrace.Wrap(templ_7745c5c3_Err)
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
							if templ_7745c5c3_Err != nil {
								return errtrace.Wrap(templ_7745c5c3_Err)
							}
							templ_7745c5c3_Var8 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
								templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
								if !templ_7745c5c3_IsBuffer {
									templ_7745c5c3_Buffer = templ.GetBuffer()
									defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
								}
								_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("Thank you for taking the next step with us! We're just one click away from getting everything set up for you. Please confirm your email address by clicking the button below:")
								if templ_7745c5c3_Err != nil {
									return errtrace.Wrap(templ_7745c5c3_Err)
								}
								if !templ_7745c5c3_IsBuffer {
									_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
								}
								return errtrace.Wrap(templ_7745c5c3_Err)
							})
							templ_7745c5c3_Err = components.P().Render(templ.WithChildren(ctx, templ_7745c5c3_Var8), templ_7745c5c3_Buffer)
							if templ_7745c5c3_Err != nil {
								return errtrace.Wrap(templ_7745c5c3_Err)
							}
							if !templ_7745c5c3_IsBuffer {
								_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
							}
							return errtrace.Wrap(templ_7745c5c3_Err)
						})
						templ_7745c5c3_Err = components.ColFull(utils.AlignLeft).Render(templ.WithChildren(ctx, templ_7745c5c3_Var7), templ_7745c5c3_Buffer)
						if templ_7745c5c3_Err != nil {
							return errtrace.Wrap(templ_7745c5c3_Err)
						}
						if !templ_7745c5c3_IsBuffer {
							_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
						}
						return errtrace.Wrap(templ_7745c5c3_Err)
					})
					templ_7745c5c3_Err = components.Row().Render(templ.WithChildren(ctx, templ_7745c5c3_Var6), templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err != nil {
						return errtrace.Wrap(templ_7745c5c3_Err)
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
					if templ_7745c5c3_Err != nil {
						return errtrace.Wrap(templ_7745c5c3_Err)
					}
					templ_7745c5c3_Var9 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
						templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
						if !templ_7745c5c3_IsBuffer {
							templ_7745c5c3_Buffer = templ.GetBuffer()
							defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
						}
						templ_7745c5c3_Var10 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
							templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
							if !templ_7745c5c3_IsBuffer {
								templ_7745c5c3_Buffer = templ.GetBuffer()
								defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
							}
							templ_7745c5c3_Err = components.ButtonFilled(utils.AlignCenter, utils.Primary, payload.ButtonText, payload.ButtonURL).Render(ctx, templ_7745c5c3_Buffer)
							if templ_7745c5c3_Err != nil {
								return errtrace.Wrap(templ_7745c5c3_Err)
							}
							if !templ_7745c5c3_IsBuffer {
								_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
							}
							return errtrace.Wrap(templ_7745c5c3_Err)
						})
						templ_7745c5c3_Err = components.ColFull(utils.AlignCenter).Render(templ.WithChildren(ctx, templ_7745c5c3_Var10), templ_7745c5c3_Buffer)
						if templ_7745c5c3_Err != nil {
							return errtrace.Wrap(templ_7745c5c3_Err)
						}
						if !templ_7745c5c3_IsBuffer {
							_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
						}
						return errtrace.Wrap(templ_7745c5c3_Err)
					})
					templ_7745c5c3_Err = components.Row().Render(templ.WithChildren(ctx, templ_7745c5c3_Var9), templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err != nil {
						return errtrace.Wrap(templ_7745c5c3_Err)
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
					if templ_7745c5c3_Err != nil {
						return errtrace.Wrap(templ_7745c5c3_Err)
					}
					templ_7745c5c3_Var11 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
						templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
						if !templ_7745c5c3_IsBuffer {
							templ_7745c5c3_Buffer = templ.GetBuffer()
							defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
						}
						templ_7745c5c3_Var12 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
							templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
							if !templ_7745c5c3_IsBuffer {
								templ_7745c5c3_Buffer = templ.GetBuffer()
								defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
							}
							if payload.ExpiresIn != "" {
								templ_7745c5c3_Var13 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
									templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
									if !templ_7745c5c3_IsBuffer {
										templ_7745c5c3_Buffer = templ.GetBuffer()
										defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
									}
									_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<b>Please note:</b> The confirmation link is valid for ")
									if templ_7745c5c3_Err != nil {
										return errtrace.Wrap(templ_7745c5c3_Err)
									}
									var templ_7745c5c3_Var14 string
									templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(payload.ExpiresIn)
									if templ_7745c5c3_Err != nil {
										return errtrace.Wrap(templ.Error{Err: templ_7745c5c3_Err, FileName: `template/confirm_email.templ`, Line: 53, Col: 82})
									}
									_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
									if templ_7745c5c3_Err != nil {
										return errtrace.Wrap(templ_7745c5c3_Err)
									}
									_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(". After this period, it will expire, and you will need to request a new confirmation if you haven't completed the process.")
									if templ_7745c5c3_Err != nil {
										return errtrace.Wrap(templ_7745c5c3_Err)
									}
									if !templ_7745c5c3_IsBuffer {
										_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
									}
									return errtrace.Wrap(templ_7745c5c3_Err)
								})
								templ_7745c5c3_Err = components.P().Render(templ.WithChildren(ctx, templ_7745c5c3_Var13), templ_7745c5c3_Buffer)
								if templ_7745c5c3_Err != nil {
									return errtrace.Wrap(templ_7745c5c3_Err)
								}
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
							if templ_7745c5c3_Err != nil {
								return errtrace.Wrap(templ_7745c5c3_Err)
							}
							if payload.SupportEmail != "" && payload.SupportURL != "" {
								templ_7745c5c3_Var15 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
									templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
									if !templ_7745c5c3_IsBuffer {
										templ_7745c5c3_Buffer = templ.GetBuffer()
										defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
									}
									_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("If you have any questions, please don't hesitate to contact us at <a href=\"")
									if templ_7745c5c3_Err != nil {
										return errtrace.Wrap(templ_7745c5c3_Err)
									}
									var templ_7745c5c3_Var16 templ.SafeURL = payload.GetSupportMailToLink()
									_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var16)))
									if templ_7745c5c3_Err != nil {
										return errtrace.Wrap(templ_7745c5c3_Err)
									}
									_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
									if templ_7745c5c3_Err != nil {
										return errtrace.Wrap(templ_7745c5c3_Err)
									}
									var templ_7745c5c3_Var17 string
									templ_7745c5c3_Var17, templ_7745c5c3_Err = templ.JoinStringErrs(payload.SupportEmail)
									if templ_7745c5c3_Err != nil {
										return errtrace.Wrap(templ.Error{Err: templ_7745c5c3_Err, FileName: `template/confirm_email.templ`, Line: 58, Col: 139})
									}
									_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var17))
									if templ_7745c5c3_Err != nil {
										return errtrace.Wrap(templ_7745c5c3_Err)
									}
									_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a> or visit our <a href=\"")
									if templ_7745c5c3_Err != nil {
										return errtrace.Wrap(templ_7745c5c3_Err)
									}
									var templ_7745c5c3_Var18 templ.SafeURL = templ.URL(payload.SupportURL)
									_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var18)))
									if templ_7745c5c3_Err != nil {
										return errtrace.Wrap(templ_7745c5c3_Err)
									}
									_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">support page</a>.")
									if templ_7745c5c3_Err != nil {
										return errtrace.Wrap(templ_7745c5c3_Err)
									}
									if !templ_7745c5c3_IsBuffer {
										_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
									}
									return errtrace.Wrap(templ_7745c5c3_Err)
								})
								templ_7745c5c3_Err = components.P().Render(templ.WithChildren(ctx, templ_7745c5c3_Var15), templ_7745c5c3_Buffer)
								if templ_7745c5c3_Err != nil {
									return errtrace.Wrap(templ_7745c5c3_Err)
								}
							}
							if !templ_7745c5c3_IsBuffer {
								_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
							}
							return errtrace.Wrap(templ_7745c5c3_Err)
						})
						templ_7745c5c3_Err = components.ColFull(utils.AlignLeft).Render(templ.WithChildren(ctx, templ_7745c5c3_Var12), templ_7745c5c3_Buffer)
						if templ_7745c5c3_Err != nil {
							return errtrace.Wrap(templ_7745c5c3_Err)
						}
						if !templ_7745c5c3_IsBuffer {
							_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
						}
						return errtrace.Wrap(templ_7745c5c3_Err)
					})
					templ_7745c5c3_Err = components.Row().Render(templ.WithChildren(ctx, templ_7745c5c3_Var11), templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err != nil {
						return errtrace.Wrap(templ_7745c5c3_Err)
					}
					if !templ_7745c5c3_IsBuffer {
						_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
					}
					return errtrace.Wrap(templ_7745c5c3_Err)
				})
				templ_7745c5c3_Err = components.Container(utils.BgWhite).Render(templ.WithChildren(ctx, templ_7745c5c3_Var5), templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return errtrace.Wrap(templ_7745c5c3_Err)
				}
				if !templ_7745c5c3_IsBuffer {
					_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
				}
				return errtrace.Wrap(templ_7745c5c3_Err)
			})
			templ_7745c5c3_Err = components.Wrapper().Render(templ.WithChildren(ctx, templ_7745c5c3_Var4), templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return errtrace.Wrap(templ_7745c5c3_Err)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
			if templ_7745c5c3_Err != nil {
				return errtrace.Wrap(templ_7745c5c3_Err)
			}
			templ_7745c5c3_Var19 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
				templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
				if !templ_7745c5c3_IsBuffer {
					templ_7745c5c3_Buffer = templ.GetBuffer()
					defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
				}
				templ_7745c5c3_Var20 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
					templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
					if !templ_7745c5c3_IsBuffer {
						templ_7745c5c3_Buffer = templ.GetBuffer()
						defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
					}
					templ_7745c5c3_Var21 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
						templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
						if !templ_7745c5c3_IsBuffer {
							templ_7745c5c3_Buffer = templ.GetBuffer()
							defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
						}
						templ_7745c5c3_Var22 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
							templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
							if !templ_7745c5c3_IsBuffer {
								templ_7745c5c3_Buffer = templ.GetBuffer()
								defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
							}
							templ_7745c5c3_Var23 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
								templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
								if !templ_7745c5c3_IsBuffer {
									templ_7745c5c3_Buffer = templ.GetBuffer()
									defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
								}
								if payload.Footer != "" {
									templ_7745c5c3_Err = templ.Raw(payload.Footer).Render(ctx, templ_7745c5c3_Buffer)
									if templ_7745c5c3_Err != nil {
										return errtrace.Wrap(templ_7745c5c3_Err)
									}
								} else {
									var templ_7745c5c3_Var24 string
									templ_7745c5c3_Var24, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d ", time.Now().Year()))
									if templ_7745c5c3_Err != nil {
										return errtrace.Wrap(templ.Error{Err: templ_7745c5c3_Err, FileName: `template/confirm_email.templ`, Line: 73, Col: 47})
									}
									_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var24))
									if templ_7745c5c3_Err != nil {
										return errtrace.Wrap(templ_7745c5c3_Err)
									}
									templ_7745c5c3_Err = components.Link(payload.CompanyName, payload.CompanyURL).Render(ctx, templ_7745c5c3_Buffer)
									if templ_7745c5c3_Err != nil {
										return errtrace.Wrap(templ_7745c5c3_Err)
									}
									_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" All rights reserved.")
									if templ_7745c5c3_Err != nil {
										return errtrace.Wrap(templ_7745c5c3_Err)
									}
								}
								if !templ_7745c5c3_IsBuffer {
									_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
								}
								return errtrace.Wrap(templ_7745c5c3_Err)
							})
							templ_7745c5c3_Err = components.SecondaryText().Render(templ.WithChildren(ctx, templ_7745c5c3_Var23), templ_7745c5c3_Buffer)
							if templ_7745c5c3_Err != nil {
								return errtrace.Wrap(templ_7745c5c3_Err)
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" ")
							if templ_7745c5c3_Err != nil {
								return errtrace.Wrap(templ_7745c5c3_Err)
							}
							templ_7745c5c3_Var25 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
								templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
								if !templ_7745c5c3_IsBuffer {
									templ_7745c5c3_Buffer = templ.GetBuffer()
									defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
								}
								_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("Aliquip irure fugiat fugiat sit est id laboris commodo laboris do. Culpa ad mollit eu id ex aliquip elit esse dolor id esse amet dolor. Esse culpa laborum occaecat sint aliqua. Elit dolore culpa voluptate laborum officia sit laborum laborum aliquip est commodo qui. Commodo anim et commodo elit qui consectetur eiusmod ipsum culpa consequat veniam. Qui ullamco nostrud cupidatat eu duis Lorem velit amet est ex culpa sint ad qui. Excepteur duis nostrud duis consectetur dolore elit minim sint do et commodo.")
								if templ_7745c5c3_Err != nil {
									return errtrace.Wrap(templ_7745c5c3_Err)
								}
								if !templ_7745c5c3_IsBuffer {
									_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
								}
								return errtrace.Wrap(templ_7745c5c3_Err)
							})
							templ_7745c5c3_Err = components.SecondaryText().Render(templ.WithChildren(ctx, templ_7745c5c3_Var25), templ_7745c5c3_Buffer)
							if templ_7745c5c3_Err != nil {
								return errtrace.Wrap(templ_7745c5c3_Err)
							}
							if !templ_7745c5c3_IsBuffer {
								_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
							}
							return errtrace.Wrap(templ_7745c5c3_Err)
						})
						templ_7745c5c3_Err = components.ColFull(utils.AlignCenter).Render(templ.WithChildren(ctx, templ_7745c5c3_Var22), templ_7745c5c3_Buffer)
						if templ_7745c5c3_Err != nil {
							return errtrace.Wrap(templ_7745c5c3_Err)
						}
						if !templ_7745c5c3_IsBuffer {
							_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
						}
						return errtrace.Wrap(templ_7745c5c3_Err)
					})
					templ_7745c5c3_Err = components.Row().Render(templ.WithChildren(ctx, templ_7745c5c3_Var21), templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err != nil {
						return errtrace.Wrap(templ_7745c5c3_Err)
					}
					if !templ_7745c5c3_IsBuffer {
						_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
					}
					return errtrace.Wrap(templ_7745c5c3_Err)
				})
				templ_7745c5c3_Err = components.Container(utils.Transparent).Render(templ.WithChildren(ctx, templ_7745c5c3_Var20), templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err != nil {
					return errtrace.Wrap(templ_7745c5c3_Err)
				}
				if !templ_7745c5c3_IsBuffer {
					_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
				}
				return errtrace.Wrap(templ_7745c5c3_Err)
			})
			templ_7745c5c3_Err = components.Wrapper().Render(templ.WithChildren(ctx, templ_7745c5c3_Var19), templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return errtrace.Wrap(templ_7745c5c3_Err)
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return errtrace.Wrap(templ_7745c5c3_Err)
		})
		templ_7745c5c3_Err = components.Layout(payload.Lang, payload.Subject, payload.EmailPreview).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return errtrace.Wrap(templ_7745c5c3_Err)
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return errtrace.Wrap(templ_7745c5c3_Err)
	})
}
