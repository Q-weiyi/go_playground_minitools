package drawing

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/Q-weiyi/go_playground_minitools/helper"
)

type BodySigner struct {
	PrivKey *widget.Entry
}

func (m BodySigner) GetContent(myWindow fyne.Window) (content *container.Split) {
	// Create the first tab content
	privKey := widget.NewMultiLineEntry()
	privKey.SetPlaceHolder("Enter PRIVATE KEY PEM here...")
	privKey.SetMinRowsVisible(10)

	pubKey := widget.NewMultiLineEntry()
	pubKey.SetPlaceHolder("Enter PUBLIC KEY PEM here...")
	pubKey.SetMinRowsVisible(10)

	keySizes := []string{"1024", "2048", "4096"}
	dropdown := widget.NewSelect(keySizes, func(selected string) {
		var exists = false
		for _, ks := range keySizes {
			if ks == selected {
				exists = true
				break
			}
		}
		if !exists {
			dialog.ShowError(errors.New("no_keysize_selected"), myWindow)
		}
	})

	genKeyButton := widget.NewButton("2. Generate KEYS!!", func() {
		m.generateKeys(myWindow, privKey, pubKey, dropdown.Selected)
	})

	c1 := container.NewVBox(
		widget.NewLabel("STEP 1: GET YOUR KEYS READY!"),
		widget.NewLabel("Generate new key if you dont have one:"),
		widget.NewLabel("1. Select RSA key size:"),
		dropdown,
		genKeyButton,
		widget.NewSeparator(),
		widget.NewLabel("If you have your own key, you may paste it below:"),
		widget.NewLabel("Private key"),
		privKey,
		widget.NewLabel("Public key"),
		pubKey,
	)

	body := widget.NewMultiLineEntry()
	body.SetPlaceHolder("Enter body text here...")
	body.SetMinRowsVisible(10)

	query := widget.NewMultiLineEntry()
	query.SetPlaceHolder("Enter query string here...")

	signature := widget.NewMultiLineEntry()
	signature.SetPlaceHolder("Your signature will displays here...")

	signButton := widget.NewButton("SIGN!!", func() {
		m.sign(myWindow, signature, strings.TrimSpace(privKey.Text), strings.TrimSpace(pubKey.Text), strings.TrimSpace(body.Text), strings.TrimSpace(query.Text))
	})

	c2 := container.NewVBox(
		widget.NewLabel("STEP 2: SIGN!!"),
		widget.NewLabel("3. Body"),
		body,
		widget.NewLabel("4. Query String"),
		query,
		signButton,
		widget.NewSeparator(),
		widget.NewLabel("5. Signature"),
		signature,
	)

	content = container.NewHSplit(c1, c2)

	return
}

func (m BodySigner) generateKeys(myWindow fyne.Window, privKey *widget.Entry, pubKey *widget.Entry, ks string) {
	var seleted helper.KeySize

	switch ks {
	case fmt.Sprint(helper.KeySize1024):
		seleted = helper.KeySize1024
	case fmt.Sprint(helper.KeySize2048):
		seleted = helper.KeySize2048
	case fmt.Sprint(helper.KeySize4096):
		seleted = helper.KeySize4096
	}

	priv, pub, err := helper.GenerateRSAKeys(seleted)
	if err != nil {
		dialog.ShowError(err, myWindow)
		return
	}

	privKey.SetText(priv)
	pubKey.SetText(pub)
}

func (m BodySigner) sign(myWindow fyne.Window, signTA *widget.Entry, privKey, pubKey, body, queryString string) {
	data, err := helper.MinifyJSON(body)
	if err != nil {
		dialog.ShowError(err, myWindow)
	}

	qs := url.QueryEscape(queryString)

	if queryString != "" {
		data = append(data, []byte(qs)...)
	}

	signature, err := helper.RSASignPSS(data, privKey)
	if err != nil {
		dialog.ShowError(err, myWindow)
		return
	}

	err = helper.RSAVerifyPSS([]byte(data), signature, pubKey)
	if err != nil {
		dialog.ShowError(err, myWindow)
		return
	}

	signTA.SetText(signature)
}
