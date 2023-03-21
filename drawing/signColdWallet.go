package drawing

import (
	"crypto/ed25519"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type ColdWalletSigner struct {
}

func (m ColdWalletSigner) GetContent(myWindow fyne.Window) (content *fyne.Container) {
	// Create the first tab content

	address := widget.NewEntry()
	address.SetPlaceHolder("Enter Cold Wallet Address here...")

	lblPrivKey := widget.NewLabel("Private key")
	privKey := widget.NewMultiLineEntry()
	privKey.SetPlaceHolder("Enter PRIVATE KEY PEM here...")
	privKey.SetMinRowsVisible(5)

	lblpubKey := widget.NewLabel("Public key")
	pubKey := widget.NewMultiLineEntry()
	pubKey.SetPlaceHolder("Enter PUBLIC KEY PEM here...")
	pubKey.SetMinRowsVisible(5)

	envs := []string{"DEV", "SIT", "UAT/PROD/SANDBOX"}
	dropdown := widget.NewSelect(envs, func(selected string) {
		var exists = false
		for _, ks := range envs {
			if ks == selected {
				exists = true
				break
			}
		}
		if !exists {
			dialog.ShowError(errors.New("not supported"), myWindow)
		}

		m.doSelected(lblPrivKey, lblpubKey, privKey, pubKey, selected)
	})

	dropdown.SetSelectedIndex(0)

	signature := widget.NewMultiLineEntry()
	signature.SetPlaceHolder("Your signature will displays here...")

	signButton := widget.NewButton("SIGN!!", func() {
		m.sign(myWindow, signature, strings.TrimSpace(privKey.Text), strings.TrimSpace(pubKey.Text), strings.TrimSpace(dropdown.Selected), strings.TrimSpace(address.Text))
	})

	content = container.NewVBox(
		widget.NewLabel("If you have your own key, you may paste it below: (SIT & DEV cannot use own key)"),
		widget.NewLabel("Choose the environment:"),
		dropdown,
		widget.NewLabel("Cold wallet address:"),
		address,
		lblPrivKey,
		privKey,
		lblpubKey,
		pubKey,
		signButton,
		widget.NewSeparator(),
		widget.NewLabel("Address Signature:"),
		signature,
	)

	return
}

func (m ColdWalletSigner) doSelected(lblPriv, lblPub *widget.Label, privKey, pubKey *widget.Entry, selected string) {
	switch selected {
	case "DEV", "SIT":
		lblPriv.Hidden = true
		lblPub.Hidden = true
		privKey.Hidden = true
		pubKey.Hidden = true
	default:
		lblPriv.Hidden = false
		lblPub.Hidden = false
		privKey.Hidden = false
		pubKey.Hidden = false
	}
}

func (m ColdWalletSigner) sign(myWindow fyne.Window, s *widget.Entry, privKey64, pubKey64, selected, addr string) {
	var (
		pubKeyB, privKeyB []byte
		err               error
	)

	switch selected {
	case "DEV":
		dialog.ShowError(errors.New("not supporting dev. please talk to dev"), myWindow)
		return
	case "SIT":
		privKeyB, err = base64.StdEncoding.DecodeString("JRB2iRzIhH54YIDqVi04PmdQyFaF2zz4NGC4D4tTkuHc5F+kB3NgI36ng+7k1OGRv2YEP2Q7qqM+LFfTKI5bag==")
		if err != nil {
			dialog.ShowError(errors.New(fmt.Sprintf("Private Key Invalid. Please make sure it is base64. Error:%s", err)), myWindow)
			return
		}
		pubKeyB, err = base64.StdEncoding.DecodeString("3ORfpAdzYCN+p4Pu5NThkb9mBD9kO6qjPixX0yiOW2o=")
		if err != nil {
			dialog.ShowError(errors.New(fmt.Sprintf("Public Key Invalid. Please make sure it is base64. Error:%s", err)), myWindow)
			return
		}
	default:
		if privKey64 == "" || pubKey64 == "" {
			dialog.ShowError(errors.New("invalid keys."), myWindow)
			return
		}

		privKeyB, err = base64.StdEncoding.DecodeString(privKey64)
		if err != nil {
			dialog.ShowError(errors.New(fmt.Sprintf("Private Key Invalid. Please make sure it is base64. Error:%s", err)), myWindow)
			return
		}
		pubKeyB, err = base64.StdEncoding.DecodeString(pubKey64)
		if err != nil {
			dialog.ShowError(errors.New(fmt.Sprintf("Public Key Invalid. Please make sure it is base64. Error:%s", err)), myWindow)
			return
		}
	}

	addrByte := []byte(addr)

	// Sign the message
	signature := ed25519.Sign(privKeyB, addrByte)

	// Encode the signature to a base64 string
	signatureB64 := base64.StdEncoding.EncodeToString(signature)
	s.SetText(signatureB64)

	if ed25519.Verify(pubKeyB, addrByte, signature) {
		fmt.Println("Signature is valid.")
	} else {
		fmt.Println("Signature is not valid.")
	}
}
