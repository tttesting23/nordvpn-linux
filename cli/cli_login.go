package cli

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/url"

	"github.com/NordSecurity/nordvpn-linux/client"
	"github.com/NordSecurity/nordvpn-linux/daemon/pb"
	"github.com/NordSecurity/nordvpn-linux/internal"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

const (
	// LoginUsageText is shown next to login command by nordvpn --help
	LoginUsageText = "Logs you in"

	// LoginFlagUsernameUsageText is shown next to username flag by nordvpn login --help
	LoginFlagUsernameUsageText = "Specify a user account"

	// LoginFlagPasswordUsageText is shown next to password flag by nordvpn login --help
	LoginFlagPasswordUsageText = "Specify the password for the user specified in --username"
	// LoginFlagLegacyUsageText is shown next to legacy flag by nordvpn login --help
	LoginFlagLegacyUsageText = "Use legacy login method. Does not support MFA."

	// LoginFlagTokenUsageText is shown next to token flag by nordvpn login --help
	LoginFlagTokenUsageText = "Use token login method. Does not support MFA." // #nosec

	// LoginCallbackUsageText is shown next to callback flag by nordvpn login --help
	LoginCallbackUsageText = "Usually used by the browser to finish Nord Account login flow. Also useful in headless setups."
)

func (c *cmd) Login(ctx *cli.Context) error {
	if ctx.IsSet(flagLoginCallback) {
		return c.oauth2(ctx)
	}

	if ctx.IsSet(flagUsername) && !ctx.IsSet(flagPassword) || !ctx.IsSet(flagUsername) && ctx.IsSet(flagPassword) {
		return formatError(argsCountError(ctx))
	}

	if ctx.IsSet(flagLegacy) || ctx.IsSet(flagUsername) || ctx.IsSet(flagPassword) {
		resp, err := c.client.IsLoggedIn(context.Background(), &pb.Empty{})
		if err != nil || resp.GetValue() {
			return formatError(internal.ErrAlreadyLoggedIn)
		}

		err = c.loginDefault(ctx)
		if err != nil {
			return formatError(err)
		}

		return nil
	}

	if ctx.IsSet(flagToken) {
		resp, err := c.client.IsLoggedIn(context.Background(), &pb.Empty{})
		if err != nil || resp.GetValue() {
			return formatError(internal.ErrAlreadyLoggedIn)
		}

		err = c.loginWithToken(ctx)
		if err != nil {
			return formatError(err)
		}

		return nil
	}

	cl, err := c.client.LoginOAuth2(
		context.Background(),
		&pb.Empty{},
	)
	if err != nil {
		return formatError(err)
	}

	for {
		resp, err := cl.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return formatError(err)
		}
		if url := resp.GetData(); url != "" {
			color.Green("Continue in the browser: %s", url)
		}
	}

	return nil
}

func (c *cmd) loginDefault(ctx *cli.Context) error {
	color.Yellow(MsgLoginLegacyDeprecated)
	username, password := ctx.String(flagUsername), ctx.String(flagPassword)
	var err error
	if isStdInAvailable() {
		username, password, err = ReadCredentialsFromStdIn()
		if err != nil {
			return err
		}
	}
	// if flags are not provided, read from terminal
	if username == "" && password == "" {
		// show intro message
		color.Green(LoginStart)
		err = c.LoginWithTerminal(ctx, MaxLoginAttempts)
		if err != nil {
			return formatError(err)
		}
		return nil
	}
	err = c.LoginWithFlags(ctx, username, password)
	if err != nil {
		return formatError(err)
	}
	return nil
}

func (c *cmd) loginWithToken(ctx *cli.Context) error {
	// nordvpn login --token b50fc06c2bf6331522c1ef5f1d449ca99b818a16ef10253d67b4a4804d9x0xd6
	token := ctx.Args().First()
	if token == "" {
		return formatError(errors.New(client.TokenLoginFailure))
	}

	resp, err := c.client.LoginWithToken(context.Background(), &pb.LoginWithTokenRequest{
		Token: token,
	})
	if err != nil {
		return formatError(err)
	}
	return LoginRespHandler(ctx, resp)
}

func (c *cmd) LoginWithTerminal(ctx *cli.Context, attempt int) error {
	if attempt == 0 {
		return formatError(fmt.Errorf(LoginTooManyAttempts, ctx.App.Name))
	} else if attempt != MaxLoginAttempts {
		color.Yellow(client.LegacyLoginFailure)
		color.Yellow(LoginAttempt, MaxLoginAttempts-attempt+1, MaxLoginAttempts)
	}

	username, password, err := ReadCredentialsFromTerminal()
	if err != nil {
		return formatError(err)
	}

	resp, err := c.client.Login(context.Background(), &pb.LoginRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		return formatError(err)
	}

	switch resp.Type {
	case internal.CodeUnauthorized:
		return c.LoginWithTerminal(ctx, attempt-1)
	}
	return LoginRespHandler(ctx, resp)
}

func (c *cmd) LoginWithFlags(ctx *cli.Context, username, password string) error {
	resp, err := c.client.Login(context.Background(), &pb.LoginRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		return formatError(err)
	}
	return LoginRespHandler(ctx, resp)
}

func LoginRespHandler(ctx *cli.Context, resp *pb.LoginResponse) error {
	switch resp.Type {
	case internal.CodeGatewayError:
		return formatError(errors.New(client.ConnectTimeoutError))
	case internal.CodeUnauthorized:
		return formatError(errors.New(client.LegacyLoginFailure))
	case internal.CodeBadRequest:
		return formatError(errors.New(client.LoginFailure))
	case internal.CodeTokenLoginFailure:
		return formatError(errors.New(client.TokenLoginFailure))
	case internal.CodeTokenInvalid:
		return formatError(errors.New(client.TokenInvalid))
	case internal.CodeSuccess:
		color.Green(LoginSuccess, ctx.App.Name)
	}
	return nil
}

// oauth2 is called by the browser during login via OAuth2.
func (c *cmd) oauth2(ctx *cli.Context) error {
	if ctx.NArg() != 1 {
		return formatError(errors.New("expected a url"))
	}

	url, err := url.Parse(ctx.Args().First())
	if err != nil {
		return formatError(err)
	}

	if url.Scheme != "nordvpn" {
		return formatError(errors.New("expected a url with nordvpn scheme"))
	}

	_, err = c.client.LoginOAuth2Callback(context.Background(), &pb.String{
		Data: url.Query().Get("exchange_token"),
	})
	if err != nil {
		return formatError(err)
	}

	color.Green(LoginSuccess, ctx.App.Name)
	return nil
}
