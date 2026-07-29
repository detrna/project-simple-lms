package infrastructure

import (
	"context"
	"main/internal/config"
	"main/internal/pkg"

	gomail "github.com/wneessen/go-mail"
)

type GoMailer struct {
	client *gomail.Client
	from   string
}

func NewGoMailer(cfg *config.MailConfig, appCfg *config.AppConfig) (pkg.Mailer, error) {
	var auth gomail.SMTPAuthType
	var tls gomail.TLSPolicy

	switch appCfg.Mode {
	case "DEV":
		auth = gomail.SMTPAuthNoAuth
		tls = gomail.NoTLS

	case "DEMO":
		auth = gomail.SMTPAuthPlain
		tls = gomail.DefaultTLSPolicy

	}

	client, err := gomail.NewClient(
		cfg.Host,
		gomail.WithPort(cfg.Port),
		gomail.WithUsername(cfg.Username),
		gomail.WithPassword(cfg.Password),
		gomail.WithSMTPAuth(auth),
		gomail.WithTLSPolicy(tls),
	)

	if err != nil {
		return nil, err
	}

	return &GoMailer{
		client: client,
		from:   cfg.From,
	}, nil
}

func (gm *GoMailer) Send(ctx context.Context, to, subject, html string) error {
	msg := gomail.NewMsg()

	if err := msg.From(gm.from); err != nil {
		return err
	}

	if err := msg.To(to); err != nil {
		return err
	}

	msg.Subject(subject)
	msg.SetBodyString(gomail.TypeTextHTML, html)

	return gm.client.DialAndSendWithContext(ctx, msg)
}
