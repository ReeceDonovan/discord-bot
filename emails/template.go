package emails

import "fmt"

const template = `
<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width">
        <style>
                @import url('https://fonts.googleapis.com/css2?family=Roboto:wght@100;300;400;500;700;900&display=swap');

                html, body {
                        margin: 0;
                        padding: 0;
                        font-size: 14px;
                        font-weight: 400;
                        font-family: 'Roboto', sans-serif !important;
                        color: white;
                        overflow-x: hidden;
                }
        </style>
    </head>
    <body style="margin: 0 auto;padding: 0;font-size: 14px;font-weight: 400 !important;color: white;overflow-x: hidden;font-family: 'Roboto', sans-serif !important;">
        <div id="body" style="margin: 0 auto;background-color: rgb(33,33,33);height: 100%%;position: absolute;width: 100%%;">
            <div id="kc-header" class="" style="height: 64px;background-color: #2196F3;position: relative;z-index: 2;box-shadow: 0px 0px 8px rgba(0,0,0,0.35);margin: 0 auto;">
                <div id="kc-header-wrapper" class="" style="position: relative;height: 64px;display: block;text-align: center;">
                        <img src="https://raw.githubusercontent.com/UCCNetsoc/wiki/master/assets/logo-horizontal-inverted.png" style="height: 32px;margin: 16px auto;padding: 0;">
                </div>
            </div>
			<h1 style="color: white; font-size: 18px;font-weight: 200;text-align: center;padding: 8px 0;">
				%s
            </h1>
            <div style="max-width: 460px; margin: 1em auto">
                <p style="border-top: 1px solid rgb(55,55,55);border-bottom: 1px solid rgb(55,55,55);color: white;text-align: center;max-width: 500px;color: #fff;padding: 1em;">
                    %s
                </p>
				<div style="background-color: #111;max-width: max-content; margin: 10px auto 15px auto;padding: 10px; border-left: #2196F3 solid 4px;">%s</div>
            </div>
        </div>
    </body>
</html>
`

// FillTemplate for custom email.
func FillTemplate(header, paragraph, code string) string {
	return fmt.Sprintf(template, header, paragraph, code)
}
