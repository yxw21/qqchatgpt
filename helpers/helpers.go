package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/tuotoo/qrcode"
	"log"
	"os"
	"qqgpt/config"
	"time"
)

func LoginWithPassword(uin int64, password string) (*client.QQClient, error) {
	qq := client.NewClient(uin, password)
	loginResponse, err := qq.Login()
	if err != nil {
		qq.Disconnect()
		return nil, err
	}
	if !loginResponse.Success {
		qq.Disconnect()
		return nil, errors.New(loginResponse.Error.String())
	}
	log.Println("使用qq和密码登录成功")
	return qq, nil
}

func LoginWithQRCode(saveToken bool) (*client.QQClient, error) {
	var loginResponse *client.LoginResponse
	qq := client.NewClientEmpty()
	qrCodeLoginResponse, err := qq.FetchQRCode()
	if err != nil {
		qq.Disconnect()
		return nil, err
	}
	matrix, err := qrcode.Decode(bytes.NewReader(qrCodeLoginResponse.ImageData))
	if err != nil {
		qq.Disconnect()
		return nil, err
	}
	log.Println("扫描下面二维码登录")
	qrcodeTerminal.New().Get(matrix.Content).Print()
	for {
		time.Sleep(time.Second)
		newQRCodeLoginResponse, _ := qq.QueryQRCodeStatus(qrCodeLoginResponse.Sig)
		if newQRCodeLoginResponse == nil {
			continue
		}
		if newQRCodeLoginResponse.State == client.QRCodeCanceled {
			log.Println("扫码被用户取消，重新生成二维码")
			return LoginWithQRCode(saveToken)
		} else if newQRCodeLoginResponse.State == client.QRCodeTimeout {
			log.Println("二维码过期，重新生成二维码")
			return LoginWithQRCode(saveToken)
		} else if newQRCodeLoginResponse.State == client.QRCodeWaitingForConfirm {
			log.Println("扫码成功, 请在手机端确认登录.")
		} else if newQRCodeLoginResponse.State == client.QRCodeConfirmed {
			loginResponse, err = qq.QRCodeLogin(newQRCodeLoginResponse.LoginInfo)
			if err != nil {
				qq.Disconnect()
				return nil, err
			}
			if !loginResponse.Success {
				qq.Disconnect()
				return nil, errors.New(loginResponse.Error.String())
			}
			log.Println("使用二维码扫码登录成功")
			if saveToken {
				_ = os.WriteFile(config.Instance.TokenFile, qq.GenToken(), 0777)
			}
			break
		}
	}
	return qq, nil
}

func LoginWithToken() (*client.QQClient, error) {
	qq := client.NewClientEmpty()
	token, err := os.ReadFile(config.Instance.TokenFile)
	if err != nil {
		qq.Disconnect()
		return nil, err
	}
	if err = qq.TokenLogin(token); err != nil {
		_, _ = os.ReadFile(config.Instance.TokenFile)
		qq.Disconnect()
		return nil, err
	}
	log.Println("恢复登录状态成功")
	return qq, nil
}

func GetTextElementFromElements(uin int64, elements []message.IMessageElement) (bool, string) {
	var (
		isAT    bool
		content string
	)
	for _, element := range elements {
		if textElement, isTextElement := element.(*message.TextElement); isTextElement {
			content += textElement.Content
		} else if faceElement, isFaceElement := element.(*message.FaceElement); isFaceElement {
			content += faceElement.Name
		} else if animatedSticker, isAnimatedSticker := element.(*message.AnimatedSticker); isAnimatedSticker {
			content += animatedSticker.Name
		} else if tatElement, isAtElement := element.(*message.AtElement); isAtElement {
			if tatElement.Target == uin {
				isAT = true
			}
		}
	}
	return isAT, content
}

func AutoLoadDevice() error {
	fileBytes, err := os.ReadFile(config.Instance.DeviceFile)
	if err != nil {
		fileBytes = client.SystemDeviceInfo.ToJson()
		_ = os.WriteFile(config.Instance.DeviceFile, fileBytes, 0777)
	}
	if err = client.SystemDeviceInfo.ReadJson(fileBytes); err != nil {
		return err
	}
	return nil
}

func PrintJSONMarshal[T any](event T) {
	bs, err := json.Marshal(event)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bs))
}
