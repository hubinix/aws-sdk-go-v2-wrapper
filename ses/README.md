aws-sdk-go-v2-wrapper | SES
----


# Quick Usage

```go
import (
	"context"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/config"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/ses"
)

func main() {
	svc, err := ses.New(config.Config{
		AccessKey: "<...>",
		SecretKey: "<...>",
	})
	if err != nil {
		panic(err)
	}
	ctx := context.Background()


	// send email with csv file
	err = svc.XSendRawEmail(ctx, ses.XSendRawEmailRequest{
		From:     "from@example.com",
		To:       []string{"to1@example.com", "to2@example.com"},
		Subject:  "nice subject",
		TextBody: "nice body",
		HTMLBody: "<!DOCTYPE html><html lang="en"><body><p style="text-align: center;">I love you!</p></body></html>",
		Attachments: []ses.XAttachment{
			{
				Data:                    []byte(base64.StdEncoding.EncodeToString([]byte(`date,count\n2020-01-01,100`))),
				Filename:                "foo.csv",
				MIMEType:                "text/csv",
				ContentTransferEncoding: "base64",
			},
		},
	})
	if err != nil {
		panic(err)
	}
	// ...
}
```

# X API

| Name | Description |
|:--|:--|
| `XSendRawEmail` | sends email with easy option for attachment files. |
